package main

import (
	"fmt"
	"net/http"
	"os"
    "log"
    "database/sql"

    _ "github.com/go-sql-driver/mysql"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		
		db, err := sql.Open(
			"mysql", 
			""+os.Getenv("MYSQL_USERNAME")+":"+os.Getenv("MYSQL_PASSWORD")+"@tcp("+os.Getenv("MYSQL_HOST")+":"+os.Getenv("MYSQL_PORT")+")/"+os.Getenv("MYSQL_DB"))
		defer db.Close()

		if err != nil {
			log.Fatal(err)
		}

		var version string

		err2 := db.QueryRow("SELECT VERSION()").Scan(&version)

		if err2 != nil {
			log.Fatal(err2)
		}

		fmt.Fprintf(w, "Config Status : "+os.Getenv("CONFIG_STATUS")+" | MySQL Version : "+version )
	})

	http.ListenAndServe(":8080", nil)
}

