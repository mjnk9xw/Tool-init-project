package formatter

import (
	"time"

	"github.com/go-sql-driver/mysql"
)

func ConvertMySQLTimeToUnix(value mysql.NullTime) int64 {
	if !value.Valid {
		return 0
	}

	return value.Time.Unix()
}

func ConvertUnixTimestampToMySQLTime(value int64) mysql.NullTime {
	if value > 0 {
		return mysql.NullTime{
			Time:  time.Unix(value, 0),
			Valid: true,
		}
	}

	return mysql.NullTime{
		Time:  time.Time{},
		Valid: false,
	}
}
