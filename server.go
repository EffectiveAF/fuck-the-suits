package main

import (
	"database/sql"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/cryptag/gosecure/canary"
	"github.com/cryptag/gosecure/content"
	"github.com/cryptag/gosecure/csp"
	"github.com/cryptag/gosecure/frame"
	"github.com/cryptag/gosecure/hsts"
	"github.com/cryptag/gosecure/referrer"
	"github.com/cryptag/gosecure/xss"
	"github.com/gorilla/mux"
	"github.com/justinas/alice"
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/acme/autocert"
)

var (
	DEFAULT_POSTGREST_BASE_URL = "http://127.0.0.1:4999/"
	// POSTGREST_BASE_URL         = os.Getenv("INTERNAL_POSTGREST_BASE_URL")

	pgDB *sql.DB
)

func init() {
	// if POSTGREST_BASE_URL == "" {
	// 	POSTGREST_BASE_URL = DEFAULT_POSTGREST_BASE_URL
	// }

	// Create generic Postgres connection
	db, err := sql.Open("postgres", POSTGRES_CONNECT)
	if err != nil {
		log.Fatal("Error connecting to Postgres: " + err.Error())
	}

	// Setting global var
	pgDB = db
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !strings.HasPrefix(r.RequestURI, "/build") &&
			!strings.HasPrefix(r.RequestURI, "/img") &&
			!strings.HasPrefix(r.RequestURI, "/global.css") &&
			!strings.HasPrefix(r.RequestURI, "/favicon") {

			log.Infof("New request to %s\n", r.RequestURI)
		}
		next.ServeHTTP(w, r)
	})
}

func NewRouter() *mux.Router {
	r := mux.NewRouter()
	r.Use(loggingMiddleware)

	r.HandleFunc("/tos", GetIndex).Methods("GET")
	r.HandleFunc("/tos/", GetIndex).Methods("GET")

	r.HandleFunc("/methodology", GetIndex).Methods("GET")
	r.HandleFunc("/methodology/", GetIndex).Methods("GET")

	r.PathPrefix("/chart").HandlerFunc(GetIndex)

	// postgrestAPI, _ := url.Parse(POSTGREST_BASE_URL)
	// handlePostgrest := http.StripPrefix("/postgrest",
	// 	httputil.NewSingleHostReverseProxy(postgrestAPI))

	handleBuildDir := http.FileServer(http.Dir("./public"))

	// // TODO(elimisteve): Remove?
	// r.PathPrefix("/postgrest").Handler(handlePostgrest)

	// r.HandleFunc("/api/ws/all", WSAllHandler).Methods("GET")

	r.PathPrefix("/").Handler(gzipHandler(handleBuildDir)).Methods("GET")

	http.Handle("/", r)
	return r
}

func NewServer(httpAddr string) *http.Server {
	r := NewRouter()

	return &http.Server{
		Addr:         httpAddr,
		ReadTimeout:  1000 * time.Second,
		WriteTimeout: 1000 * time.Second,
		IdleTimeout:  120 * time.Second,
		Handler:      r,
	}
}

func ProductionServer(srv *http.Server, httpsAddr string, domain string, manager *autocert.Manager) {
	gotWarrant := false
	middleware := alice.New(canary.GetHandler(&gotWarrant),
		csp.GetCustomHandlerStyleUnsafeInline(domain, "api."+domain),
		hsts.PreloadHandler, frame.GetHandler("api."+domain), content.GetHandler,
		xss.GetHandler, referrer.NoHandler)

	srv.Handler = middleware.Then(manager.HTTPHandler(srv.Handler))

	srv.Addr = httpsAddr
	srv.TLSConfig = manager.TLSConfig()
}

func GetIndex(w http.ResponseWriter, req *http.Request) {
	contents, err := ioutil.ReadFile("public/index.html")
	if err != nil {
		log.Errorf("Error serving index.html: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error: couldn't serve you index.html!"))
		return
	}
	w.Write(contents)
}

func redirectToHTTPS(httpAddr, httpsPort string, manager *autocert.Manager) {
	srv := &http.Server{
		Addr:         httpAddr,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
		Handler: http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			w.Header().Set("Connection", "close")
			domain := strings.SplitN(req.Host, ":", 2)[0]
			url := "https://" + domain + ":" + httpsPort + req.URL.String()
			http.Redirect(w, req, url, http.StatusFound)
		}),
	}
	log.Infof("Listening on %v", httpAddr)
	log.Fatal(srv.ListenAndServe())
}

func getAutocertManager(domain string) *autocert.Manager {
	return &autocert.Manager{
		Prompt:     autocert.AcceptTOS,
		HostPolicy: autocert.HostWhitelist(domain),
		Cache:      autocert.DirCache("./" + domain),
	}
}
