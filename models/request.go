package models

type Request struct {
	Id         int    `db:"id"`
	ExternalId int    `db:"external_id"`
	Name       string `db:"name"  validate:"required"`
	Type       string `db:"type"`
	Text       string `db:"task_text" validate:"required"`
	Time       int    `db:"time"  validate:"required"`
}
