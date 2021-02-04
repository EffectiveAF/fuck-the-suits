package main

import (
	"flag"
	"strings"

	"github.com/justinas/alice"
	"github.com/rs/cors"
	log "github.com/sirupsen/logrus"
)

var (
	THIS_DOMAIN_BASE_URL string

	IS_PROD = false
)

func main() {
	httpAddr := flag.String("http", "127.0.0.1:5001",
		"Address to listen on HTTP")
	httpsAddr := flag.String("https", "127.0.0.1:8443",
		"Address to listen on HTTPS")
	domain := flag.String("domain", "", "Domain of this service")
	prod := flag.Bool("prod", false, "Run in Production mode")
	flag.Parse()

	// Used in server.go; must be set before NewServer() is called
	// below
	IS_PROD = *prod

	srv := NewServer(*httpAddr)

	if *prod {
		log.SetLevel(log.DebugLevel)

		if *domain == "" {
			log.Fatal("You must specify a -domain when using the -prod flag.")
		}

		THIS_DOMAIN_BASE_URL = "https://" + *domain

		manager := getAutocertManager(*domain)

		// Setup http->https redirection
		httpsPort := strings.SplitN(*httpsAddr, ":", 2)[1]
		go redirectToHTTPS(*httpAddr, httpsPort, manager)

		// Production modifications to `srv`
		ProductionServer(srv, *httpsAddr, *domain, manager)

		// TODO(elimisteve): Add graceful shutdown
		log.Infof("Listening on %v", *httpsAddr)
		log.Fatal(srv.ListenAndServeTLS("", ""))
	} else {
		log.SetLevel(log.DebugLevel)

		// Because the JavaScript ecosystem is shit. (Rollup can't
		// proxy to local APIs and Svelte doesn't work reliably with
		// WebPack, contra `template-webpack`. So, enabling CORS when
		// in dev mode so that "cross-origin" AJAX calls work.)
		srv.Handler = alice.New(cors.AllowAll().Handler).Then(srv.Handler)

		THIS_DOMAIN_BASE_URL = "http://" + *httpAddr
		log.Infof("Listening on %v", *httpAddr)
		log.Fatal(srv.ListenAndServe())
	}
}
