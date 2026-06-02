package cmd

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"
)

func Serve() {
	db := servedb()
	defer db.Close()
	http.HandleFunc("/", root)
	http.ListenAndServe(":8080", nil)
}

func root(w http.ResponseWriter, r *http.Request) {
	str := "database connected successfully"
	strb := []byte(str)
	w.Write(strb)
}

func servedb() *sql.DB {
	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	dst := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", dbHost, dbUser, dbPassword, dbName)

	var db *sql.DB
	var err error
	for i := 1; i <= 10; i++ {
		fmt.Printf("connecting to the database attempt %d/10 ......\n", i)
		db, err = sql.Open("postgres", dst)
		if err == nil {
			err = db.Ping()
			if err == nil {
				fmt.Println("database connected successfully.")
				break
			}
		}
	}

	if err != nil {
		fmt.Println("could not connect to the database", err)
		os.Exit(1)
	}
	return db
}
