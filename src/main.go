package main

import (
	"fmt"

	"github.com/GG_Backend_tech_challenge/src/repository"
)

func main() {
	db := repository.GetDataBaseConnectionWithTablesAndData("root", "turing221997", "localhost", 3306, "event")
	fmt.Println("here")
	db.Close()
}
