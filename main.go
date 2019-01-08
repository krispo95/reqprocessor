package main

import (
	"github.com/jmoiron/sqlx"
	"github.com/kataras/iris"
	_ "github.com/lib/pq"
	"gopkg.in/go-playground/validator.v9"
	"log"
	"reqprocessor/handler"
	"reqprocessor/models"
	"reqprocessor/processor"
)

func main() {
	var workersCount = 4
	var requestsCh = make(chan models.Task, 20)
	db, err := sqlx.Connect("postgres", "user=postgres password=postgres dbname=reqprocessor sslmode=disable")
	if err != nil {
		log.Fatalln(err)
	}

	app := iris.Default()

	api := handler.NewHandler(validator.New(), db, requestsCh)

	app.Post("/putTask", api.PutTask)
	app.Get("/getResult", api.GetResult)
	for i := 0; i < workersCount; i++ {
		go processor.ProcessRequest(requestsCh, db)
	}

	app.Run(iris.Addr(":8081"))
}
