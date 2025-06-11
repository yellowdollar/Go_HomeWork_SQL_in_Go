package iternals

import (
	"home_work_sql_gin/iternals/repositories"

	"github.com/jmoiron/sqlx"
)

var Db *sqlx.DB

func InitDB() {
	Db = repositories.ConnectToDatabase()
}
