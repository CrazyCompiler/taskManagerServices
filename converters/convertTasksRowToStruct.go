package converters

import (
	"database/sql"
)

type TableContent struct {
	TASKID   int
	TASK     string
	PRIORITY string
}

func ConvertRowsToStructObjects(rows *sql.Rows) ([]TableContent) {
	dbData := []TableContent{}
	if(rows != nil) {
		for rows.Next() {
			var r TableContent
			rows.Scan(&r.TASKID, &r.TASK, &r.PRIORITY)
			dbData = append(dbData, r)
		}
	}
	return dbData
}
