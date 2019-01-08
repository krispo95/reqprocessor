package models

type Task struct {
	Id         int64   `json:"id,omitempty" db:"id"`
	ExternalId int     `json:"external_id,omitempty" db:"external_id"`
	Name       string  `json:"name,omitempty" db:"name"  validate:"required"`
	Type       string  `json:"type,omitempty" db:"type"`
	TaskText   string  `json:"task_text,omitempty" db:"task_text" validate:"required"`
	WorkTime   int     `json:"work_time,omitempty" db:"work_time"  validate:"required"`
	Status     string  `json:"status,omitempty" db:"status"`
	Result     *string `json:"result,omitempty" db:"result"`
}
type ReqResult struct {
	Id int64 `json:"id" db:"id" validate:"required"`
}

//type Result struct {
//	Id         int64  `json:"id,omitempty" db:"id"`
//	Status     string `json:"status,omitempty" db:"status"`
//	Result     sql.NullString `json:"result,omitempty" db:"result"`
//}
