package main

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	dsn := os.Getenv("DB_DSN")
	if dsn == "" {
		log.Fatal("DB_DSN environment variable is not set")
	}

	// 1. Connect with retry logic (handles DB container warmup delay)
	var db *sql.DB
	var err error

	for i := 1; i <= 10; i++ {
		db, err = sql.Open("mysql", dsn)
		if err == nil {
			err = db.Ping()
			if err == nil {
				fmt.Println("✅ Successfully connected to MariaDB!")
				break
			}
		}
		log.Printf("Waiting for database to be ready... (Attempt %d/10)", i)
		time.Sleep(2 * time.Second)
	}

	if err != nil {
		log.Fatalf("Could not connect to database: %v", err)
	}
	defer db.Close()

	// 2. Run Database Migrations
	runMigrations(db)

	// 3. Health Check Route
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Backend is healthy and connected to DB!"))
	})

	// 4. Start HTTP Server
	port := "8080"
	fmt.Printf("🚀 Server starting on port %s...\n", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatalf("Server crashed: %v", err)
	}
}

func runMigrations(db *sql.DB) {
	driver, err := mysql.WithInstance(db, &mysql.Config{})
	if err != nil {
		log.Fatalf("Could not create migration driver: %v", err)
	}

	// Loads migration files from database/migrations/
	m, err := migrate.NewWithDatabaseInstance(
		"file://database/migrations",
		"mysql",
		driver,
	)
	if err != nil {
		log.Fatalf("Migration initialization failed: %v", err)
	}

	// Apply migrations
	if err := m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		log.Fatalf("Failed to run migrations: %v", err)
	}

	log.Println("✅ Database migrations applied successfully!")
}
