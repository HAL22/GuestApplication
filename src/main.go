package main

import (
	"net/http"
	"os"

	"github.com/GG_Backend_tech_challenge/src/handler"
	"github.com/GG_Backend_tech_challenge/src/repository"
)

func main() {
	os.Setenv("USER", "root")
	os.Setenv("PASSWORD", "turing221997")
	os.Setenv("IP_ADDRESS", "localhost")
	os.Setenv("DB_NAME", "event")
	os.Setenv("PORT_NUM", "3306")
	os.Setenv("HTTP_PORT", ":8001")

	db := repository.GetDataBaseConnectionWithTablesAndData(os.Getenv("USER"), os.Getenv("PASSWORD"), os.Getenv("IP_ADDRESS"), os.Getenv("PORT_NUM"), os.Getenv("DB_NAME"))
	defer db.Close()
	eventHandler := handler.EventHandler{
		DB: db,
	}

	router := handler.NewRouter(eventHandler)

	http.ListenAndServe(os.Getenv("HTTP_PORT"), router)

}
