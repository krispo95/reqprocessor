package main

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/kataras/iris"
	_ "github.com/lib/pq"
	"gopkg.in/go-playground/validator.v9"
	"log"
	"reqprocessor/models"
)

func main() {
	var requestsCh = make(chan models.Request, 20)
	db, err := sqlx.Connect("postgres", "user=postgres dbname=reqprocessor sslmode=disable")
	if err != nil {
		log.Fatalln(err)
	}
	var validate *validator.Validate
	validate = validator.New()

	app := iris.Default()
	app.Post("/task", func(ctx iris.Context) {
		var request models.Request
		if err := ctx.ReadJSON(&request); err != nil {
			panic(err)
		}
		err := validate.Struct(request)
		if err != nil {

			// This check is only needed when your code could produce
			// an invalid value for validation such as interface with nil
			// value most including myself do not usually have code like this.

			ctx.StatusCode(iris.StatusInternalServerError)
			ctx.WriteString(err.Error())
			return
		}
		fmt.Printf("%+v\n", request)
		request.Status = "В очереди"
		result, err := db.NamedExec(`INSERT INTO request (external_id, name, type, task_text, work_time, status) 
VALUES (:external_id, :name, :type, :task_text, :work_time, :status)`, &request)
		if err != nil {
			panic(err)
		}
		id, err := result.LastInsertId()
		if err != nil {
			panic(err)
		}
		request.Id = id
		requestsCh <- request
		ctx.JSON(request)

	})
	app.Run(iris.Addr(":8081"))
}
