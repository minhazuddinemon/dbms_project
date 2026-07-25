package main

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	// Note: Replace "dbms_project" with your actual module name found in your go.mod file
	"dbms-project/internal/db"
	"dbms-project/internal/handler/rest"
	"dbms-project/internal/logger"
	"dbms-project/internal/middleware"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load("../.env")
	_ = godotenv.Load(".env")

	dsn := os.Getenv("DB_DSN")
	if dsn == "" {
		log.Fatal("DB_DSN environment variable is not set")
	}

	// 1. Connect with retry logic
	var dbConn *sql.DB
	var err error

	for i := 1; i <= 10; i++ {
		dbConn, err = sql.Open("mysql", dsn)
		if err == nil {
			err = dbConn.Ping()
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
	defer dbConn.Close()

	// 2. Run Database Migrations
	runMigrations(dbConn)

	// 3. Initialize sqlc queries
	queries := db.New(dbConn)
	h := rest.NewHandler(queries)

	// 4. Set up HTTP Router
	mux := http.NewServeMux()

	// Health Check Route
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Backend is healthy and connected to DB!"))
	})

	// Registration Route
	mux.HandleFunc("/register", h.HandleRegister)
	mux.HandleFunc("/login", h.HandleLogin)
	mux.HandleFunc("/profile", middleware.RequireAuth(h.HandleProfile))
	mux.HandleFunc("/programs", h.HandleListPrograms)
	mux.HandleFunc("/programs/detail", h.HandleGetProgramByID)
	// Protected Routes (Requires valid JWT)
	mux.HandleFunc("/profile", middleware.RequireAuth(h.HandleProfile))
	mux.HandleFunc("/programs/eligible", middleware.RequireAuth(h.HandleEligiblePrograms))

	// Wrap the entire router with the logger middleware
	loggedMux := logger.RequestLogger(mux)

	// 5. Start HTTP Server
	port := "8080"
	fmt.Printf("🚀 Server starting on port %s...\n", port)
	if err := http.ListenAndServe(":"+port, loggedMux); err != nil {
		log.Fatalf("Server crashed: %v", err)
	}
}

func runMigrations(dbConn *sql.DB) {
	driver, err := mysql.WithInstance(dbConn, &mysql.Config{})
	if err != nil {
		log.Fatalf("Could not create migration driver: %v", err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://database/migrations",
		"mysql",
		driver,
	)
	if err != nil {
		log.Fatalf("Migration initialization failed: %v", err)
	}

	if err := m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		log.Fatalf("Failed to run migrations: %v", err)
	}

	log.Println("✅ Database migrations applied successfully!")
}
