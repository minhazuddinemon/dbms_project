package cmd

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/jackc/pgx/v5/stdlib"
)

func ConnectDB(ctx context.Context) *pgxpool.Pool {
	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	dst := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", dbHost, dbUser, dbPassword, dbName)

	var db *pgxpool.Pool
	var err error
	for i := 1; i <= 10; i++ {
		fmt.Printf("connecting to the database attempt %d/10 ......\n", i)
		db, err = pgxpool.New(ctx, dst)
		if err == nil {
			err = db.Ping(context.Background())
			if err == nil {
				fmt.Println("database connected successfully. ")
				break
			}
		}
		time.Sleep(2 * time.Second)
	}

	if err != nil {
		fmt.Println("could not connect to the database", err)
		os.Exit(1)
	}
	return db
}
