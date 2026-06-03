package main

import (
	"context"
	"dbms-project/cmd"
	"dbms-project/database/db"
	"dbms-project/rest"
	"net/http"
)

func main() {
	dbPool := cmd.ConnectDB(context.Background())
	defer dbPool.Close()

	store := db.NewStore(dbPool)
	server := rest.NewServer(store)

	err := http.ListenAndServe(":8080", server.Router)
	if err != nil {
		panic(err)
	}
}
