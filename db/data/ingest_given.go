// Steve Phillips / elimisteve
// 2021.02.01

package main

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/lib/pq"
	log "github.com/sirupsen/logrus"
)

const (
	POSTGRES_CONNECT   = "postgres://superuser:superuser@127.0.0.1:5432/fuckthesuits?sslmode=disable"
	NUM_COLUMNS_WANTED = 6
)

var (
	USAGE = fmt.Sprintf("Usage:\n\n$ go run %s.go raw/* [ CNMSshvol2021mmdd.txt ]\n", filepath.Base(os.Args[0]))
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println(USAGE)
		os.Exit(1)
	}

	// Log everything
	log.SetLevel(log.DebugLevel)

	filenames := os.Args[1:]

	db, err := sql.Open("postgres", POSTGRES_CONNECT)
	if err != nil {
		log.Fatalf("Error connecting to Postgres: %s", err)
	}

	for _, filename := range filenames {
		log.Infof("Ingesting file `%s` ...", filename)

		contents, err := ioutil.ReadFile(filename)
		if err != nil {
			log.Errorf("Error reading %s: %v\n", filename, err)
			continue
		}

		rows := strings.Split(string(contents), "\r\n")
		if strings.Count(rows[0], "|") != NUM_COLUMNS_WANTED-1 {
			log.Errorf("First row isn't %d columns wide! Row: `%s`\n",
				NUM_COLUMNS_WANTED, rows[0])
			continue
		}

		rows = rows[1:]
		inserted := 0

		for i, rowStr := range rows {
			row := strings.Split(rowStr, "|")

			if len(row) != NUM_COLUMNS_WANTED {
				if i == len(rows)-1 {
					continue
				}
				fmt.Printf("Skipping invalid row #%d: %q\n  Row as slice: %#v\n", i, rowStr, row)
				continue
			}

			date := row[0]
			sym := row[1]

			shortVolume, err := strconv.Atoi(row[2])
			if err != nil {
				log.Errorf("Error converting shortVolume string `%s` from row %d to int\n", i, row[2])
			}

			shortExemptVolume, err := strconv.Atoi(row[3])
			if err != nil {
				log.Errorf("Error converting shortExemptVolume string `%s` from row %d to int\n", row[3], i)
				continue
			}

			totalVolume, err := strconv.Atoi(row[4])
			if err != nil {
				log.Errorf("Error converting shortExemptVolume string `%s` from row %d to int\n", row[4], i)
				continue
			}

			market := strings.Split(row[5], ",")

			_, err = db.Exec(`INSERT INTO daily_short_volume
			(date_orig, date, symbol, short_volume, short_exempt_volume, total_volume, market)
				VALUES
			($1, $2, $3, $4, $5, $6, $7)`,
				date,
				date[:4]+"-"+date[4:6]+"-"+date[6:8],
				sym,
				shortVolume,
				shortExemptVolume,
				totalVolume,
				pq.Array(market),
			)
			if err != nil {
				log.Errorf("Error inserting row into DB: %v\n", err)
				// if strings.Contains(err.Error(), "duplicate") {
				// 	break
				// }
				continue
			}

			inserted++
		}

		fmt.Printf("Successfully inserted %d rows from file %s\n", inserted, filename)
	}
}
