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
		_, err = db.NamedExec(`INSERT INTO request (external_id, name, type, task_text, time) 
VALUES (:external_id, :name, :type, :task_text, :time)`, &request)
		if err != nil {
			panic(err)
		}
		//if count, _ := result.RowsAffected(); count == 0 {
		//	ctx.JSON(map[string]string{
		//		"error" : "require fields",
		//	})
		//	return
		//}

		ctx.JSON(request)

	})
	app.Run(iris.Addr(":8081"))
}
