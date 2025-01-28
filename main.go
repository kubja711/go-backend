package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"time"

	"gihub.com/kubja711/GO_backend/internal/models"
	_ "github.com/go-sql-driver/mysql"
)

type application struct {
	errorLog    *log.Logger
	infoLog     *log.Logger
	links       *models.LinkModel
	temperature *models.TempImageModel
	meteo       *models.MeteoModel
}

func main() {
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	cmlDb, err := openCmlDB()
	if err != nil {
		errorLog.Fatal(err)
	}
	defer cmlDb.Close()

	telcoOutDb, err := openTelcoOutDB()
	if err != nil {
		errorLog.Fatal(err)
	}
	defer telcoOutDb.Close()

	app := &application{
		errorLog:    errorLog,
		infoLog:     infoLog,
		links:       &models.LinkModel{DB: cmlDb},
		temperature: &models.TempImageModel{DB: telcoOutDb},
		meteo:       &models.MeteoModel{DB: cmlDb},
	}

	server := &http.Server{
		Addr:     ":8081",
		ErrorLog: errorLog,
		Handler:  app.routes(),
	}
	infoLog.Printf("Starting server on %s", server.Addr)
	err = server.ListenAndServe()
	errorLog.Fatal(err)
}

func openCmlDB() (*sql.DB, error) {
	db, err := sql.Open("mysql", "LOGIN@tcp(192.168.64.168)/cml_metadata")
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	return db, nil
}

func openTelcoOutDB() (*sql.DB, error) {
	db, err := sql.Open("mysql", "LOGIN@tcp(192.168.64.168)/telcorain_output")
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	return db, nil
}
