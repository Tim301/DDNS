package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/mux"
)

func redirect(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	vars := mux.Vars(r)
	location := vars["location"]
	key := vars["key"]
	redirection := getDDNS(SqliteDatabase, location, key)
	w.Header().Set("pragma", "no-cache")
	w.Header().Set("cache-control", "no-cache")

	if redirection != "Not found" {
		http.Redirect(w, r, redirection, 301)
		fmt.Println("Redirection succed")
	} else {
		fmt.Fprintf(w, "Not found")
		fmt.Println("Not found")
	}

	elapsed := time.Since(start)
	log.Printf("redirect took %s", elapsed)
}

func updateIP(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	location := vars["location"]
	key := vars["key"]
	updatedIP := "http://" + getIP(r)
	fmt.Println(updatedIP)
	setDDNS(SqliteDatabase, location, key, updatedIP)
	fmt.Fprintf(w, "Updated")
}

func getIP(r *http.Request) string {
	if r.RemoteAddr != "" {
		end := strings.Index(r.RemoteAddr, ":")
		return r.RemoteAddr[:end]
	}
	return r.RemoteAddr
}

func main() {
	SqliteDatabase, _ = sql.Open("sqlite3", "./sqlite-database.db") // Open the created SQLite File
	defer SqliteDatabase.Close()                                    // Defer Closing the database

	//fmt.Println(getDDNS(SqliteDatabase, "Studiotech", "96a3f2da"))
	r := mux.NewRouter()
	r.HandleFunc("/access/{location}/{key}", redirect)
	r.HandleFunc("/update/{location}/{key}", updateIP)

	err := http.ListenAndServe(":80", r)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
