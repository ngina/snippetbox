package main

import (
	"database/sql"
	"flag"
	"log"
	"net/http"
	"os"
	
	"snippetbox.ngina.com/internal/models"

	_ "github.com/go-sql-driver/mysql"

)

func main() {
	addr := flag.String("addr", ":4000", "HTTP network address")
	dsn := flag.String("dsn", "web:webapp@/snippetbox?parseTime=true", "MySQL data source name")
	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	db, err := openDB(*dsn)
	if err != nil { 
		errorLog.Fatal(err)
	}
	
	// We also defer a call to db.Close(), so that the connection pool is closed 
	// before the main() function exits.
	defer db.Close()

	app := &application{
		infoLog:  infoLog,
		errorLog: errorLog,
		snippets: &models.SnippetModel{DB: db},
	}

	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  app.routes(),
	}

	infoLog.Printf("Starting server on : %s", *addr)
	err = srv.ListenAndServe()
	errorLog.Fatal(err)
}

// The openDB() function wraps sql.Open() and returns a sql.DB connection pool
// for a given DSN.
func openDB(dsn string) (*sql.DB, error) { 
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err 
	}
	if err = db.Ping(); err != nil { 
		return nil, err
	}
		return db, nil 
	}
