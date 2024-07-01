package main

import (
	"github.com/yardan-org/final-task-pbi-rakamin-fullstack-yardan/app/database"
	"github.com/yardan-org/final-task-pbi-rakamin-fullstack-yardan/app/router"
)

func main() {
	database.Connect()

	r := router.SetupRouter()
	r.Run(":8080")
}
