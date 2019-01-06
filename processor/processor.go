package processor

import (
	"github.com/jmoiron/sqlx"
	"reqprocessor/models"
	"strings"
	"time"
)

func ProcessRequest(ch <-chan models.Request, db *sqlx.DB) {
	for request := range ch {
		request.Status = "В обработке"
		_, err := db.NamedExec(`UPDATE requests SET status=:status WHERE id=:id`, &request)
		if err != nil {
			panic(err)
		}
		time.Sleep(time.Duration(request.WorkTime) * time.Millisecond)
		request.Result = strings.ToUpper(request.Text)
		request.Status = "Выполнено"
		_, err = db.NamedExec(`UPDATE requests SET result=:result, status=:status WHERE id=:id`, &request)
		if err != nil {
			panic(err)
		}
	}

}
