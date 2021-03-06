package main

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3" // Import go-sqlite3 library
)

// SqliteDatabase is Gobal var use on main.go
var SqliteDatabase *sql.DB

func init() {
	_, err := os.Stat("sqlite-database.db")
	if os.IsNotExist(err) {

		log.Println("Creating sqlite-database.db...")
		file, err := os.Create("sqlite-database.db") // Create SQLite file
		if err != nil {
			log.Fatal(err.Error())
		}
		file.Close()
		log.Println("sqlite-database.db created")
		sqliteDatabaseInit, _ := sql.Open("sqlite3", "./sqlite-database.db")                                // Open the created SQLite File
		createTable(sqliteDatabaseInit)                                                                     // Create Database Tables
		insertDDNS(sqliteDatabaseInit, "example", "96a3f2da", "351c5ze2", "https://github.com/Tim301/DDNS") // Example
		sqliteDatabaseInit.Close()                                                                          // Closing the database
	}
}

func createTable(db *sql.DB) {
	createStudentTableSQL := `CREATE TABLE DDNS (
		"id" integer NOT NULL PRIMARY KEY AUTOINCREMENT,		
		"location" TEXT,
		"accesskey" TEXT,
		"updatekey" TEXT,
		"ip" TEXT
	  );` // SQL Statement for Create Table

	log.Println("Create DDNS table...")
	statement, err := db.Prepare(createStudentTableSQL) // Prepare SQL Statement
	if err != nil {
		log.Fatal(err.Error())
	}
	statement.Exec() // Execute SQL Statements
	log.Println("DDNS table created")
}

// We are passing db reference connection from main to our method with other parameters
func insertDDNS(db *sql.DB, location string, accesskey string, updatekey string, ip string) {
	log.Println("Inserting DDNS record ...")
	insertDDNSSQL := `INSERT INTO DDNS(location, accesskey, updatekey, ip) VALUES (?, ?, ?, ?)`
	statement, err := db.Prepare(insertDDNSSQL) // Prepare statement.
	if err != nil {
		log.Fatalln(err.Error())
	}
	_, err = statement.Exec(location, accesskey, updatekey, ip)
	if err != nil {
		log.Fatalln(err.Error())
	}
}

func getDDNS(db *sql.DB, target string, key string) (redirection string) {
	row, err := db.Query("SELECT * FROM DDNS")
	if err != nil {
		log.Fatal(err)
	}
	defer row.Close()
	for row.Next() { // Iterate and fetch the records from result cursor
		var id int
		var location string
		var accesskey string
		var updatekey string
		var ip string
		redirection = "Not found"
		row.Scan(&id, &location, &accesskey, &updatekey, &ip)
		log.Println("DDNS: ", location, " ", accesskey, " ", ip)
		if target == location && key == accesskey {
			redirection = ip
			break
		}
	}
	return
}

func setDDNS(db *sql.DB, location string, key string, ip string) {
	statement, err := db.Prepare("update DDNS set ip=? where location=?") // Prepare SQL Statement
	if err != nil {
		log.Fatal(err.Error())
	}
	statement.Exec(ip, location) // Execute SQL Statements
}
