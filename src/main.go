package main

import (
	"net/http"

	"github.com/GG_Backend_tech_challenge/src/handler"
	"github.com/GG_Backend_tech_challenge/src/repository"
)

func main() {
	db := repository.GetDataBaseConnectionWithTablesAndData("root", "turing221997", "localhost", 3306, "event")
	defer db.Close()
	eventHandler := handler.EventHandler{
		DB: db,
	}

	router := handler.NewRouter(eventHandler)

	http.ListenAndServe(":8001", router)

}
