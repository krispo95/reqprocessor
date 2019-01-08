package handler

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/kataras/iris"
	"gopkg.in/go-playground/validator.v9"
	"reqprocessor/models"
)

type Handler struct {
	validate   *validator.Validate
	db         *sqlx.DB
	requestsCh chan<- models.Task
}

func NewHandler(validate *validator.Validate, db *sqlx.DB, requestsCh chan<- models.Task) *Handler {
	return &Handler{
		validate:   validate,
		db:         db,
		requestsCh: requestsCh,
	}
}

func (h *Handler) PutTask(ctx iris.Context) {
	var request models.Task
	if err := ctx.ReadJSON(&request); err != nil {
		panic(err)
	}
	err := h.validate.Struct(request)
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
	rows, err := h.db.NamedQuery(`INSERT INTO requests (external_id, name, type, task_text, work_time, status) 
VALUES (:external_id, :name, :type, :task_text, :work_time, :status) RETURNING id`, &request)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	var id int64
	for rows.Next() {
		err := rows.Scan(&id)
		if err != nil {
			panic(err)
		}
	}
	if err != nil {
		panic(err)
	}
	request.Id = id
	h.requestsCh <- request
	if _, err := ctx.JSON(request); err != nil {
		panic(err)
	}
}
func (h *Handler) GetResult(ctx iris.Context) {
	var response []models.Task

	id, err := ctx.URLParamInt64("id")
	if err != nil {
		panic(err)
	}

	err = h.db.Select(&response, `SELECT id, result, status FROM requests WHERE id = $1 `, id)
	if err != nil {
		panic(err)
	}
	if _, err := ctx.JSON(response[0]); err != nil {
		panic(err)
	}

}
