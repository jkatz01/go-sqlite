package main

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3" // Import go-sqlite3 library
	"log"
	"os"
	"time"
)

type host_data struct {
	host_name string
	last_ping time.Time
}

func main() {
	//os.Remove("sqlite-database.db") // I delete the file to avoid duplicated records.
	// SQLite is a file based database.

	log.Println("Creating sqlite-database.db...")
	file, err := os.Create("sqlite-database.db") // Create SQLite file
	if err != nil {
		log.Fatal(err.Error())
	}
	file.Close()
	log.Println("sqlite-database.db created")

	sqliteDatabase, _ := sql.Open("sqlite3", "./sqlite-database.db") // Open the created SQLite File
	defer sqliteDatabase.Close()                                     // Defer Closing the database
	create_host_table(sqliteDatabase)                                // Create Database Tables

	// INSERT RECORDS
	//insertStudent(sqliteDatabase, "0001", "Liana Kim", "Bachelor")
	//insertStudent(sqliteDatabase, "0002", "Glen Rangel", "Bachelor")
	host_1 := host_data{host_name: "uhhum", last_ping: time.Now()}
	time.Sleep(200 * time.Millisecond)
	host_2 := host_data{host_name: "someone", last_ping: time.Now()}
	time.Sleep(200 * time.Millisecond)
	host_3 := host_data{host_name: "guh", last_ping: time.Now()}
	time.Sleep(200 * time.Millisecond)

	insert_host(sqliteDatabase, host_1)
	insert_host(sqliteDatabase, host_2)
	insert_host(sqliteDatabase, host_3)
	insert_host(sqliteDatabase, host_3)
	insert_host(sqliteDatabase, host_3)
	insert_host(sqliteDatabase, host_3)
	insert_host(sqliteDatabase, host_3)

	// DISPLAY INSERTED RECORDS
	displayStudents(sqliteDatabase)
}

func create_host_table(db *sql.DB) {
	create_host_table_SQL := `
		CREATE TABLE IF NOT EXISTS hosts (
		"id" integer NOT NULL PRIMARY KEY AUTOINCREMENT,		
		"host_name" TEXT UNIQUE,
		"last_ping" TEXT		
	  );` // SQL Statement for Create Table

	//log.Println("Create host table...")
	statement, err := db.Prepare(create_host_table_SQL) // Prepare SQL Statement
	if err != nil {
		log.Fatal(err.Error())
	}
	statement.Exec() // Execute SQL Statements
	log.Println("host table created")
}

// We are passing db reference connection from main to our method with other parameters
func insert_host(db *sql.DB, host host_data) {
	log.Println("Inserting host record ...")
	// TODO: change REPLACE in query to update + insert if fails
	// 	 because REPLACE changes the primary key id
	insert_host_SQL := `
		REPLACE INTO hosts(host_name, last_ping) 
		VALUES (?, ?)`
	statement, err := db.Prepare(insert_host_SQL) // Prepare statement.
	// This is good to avoid SQL injections
	if err != nil {
		log.Fatalln(err.Error())
	}
	_, err = statement.Exec(host.host_name, host.last_ping)
	if err != nil {
		log.Fatalln(err.Error())
	}
}

func displayStudents(db *sql.DB) {
	row, err := db.Query("SELECT * FROM hosts ORDER BY host_name")
	if err != nil {
		log.Fatal(err)
	}
	defer row.Close()
	for row.Next() { // Iterate and fetch the records from result cursor
		var id int
		var host_name string
		var last_ping string
		// TODO: use proper Time type
		row.Scan(&id, &host_name, &last_ping)
		log.Println("Host: ", host_name, " ", last_ping)
	}
}
