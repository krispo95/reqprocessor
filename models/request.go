package models

type Request struct {
	Id         int64  `db:"id"`
	ExternalId int    `db:"external_id"`
	Name       string `db:"name"  validate:"required"`
	Type       string `db:"type"`
	Text       string `db:"task_text" validate:"required"`
	WorkTime   int    `db:"work_time"  validate:"required"`
	Status     string `db:"status"`
	Result     string `db:"result"`
}
