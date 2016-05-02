package converters

import (
	"database/sql"
	"strconv"
)

type TableContent struct {
	TASKID   int
	TASK     string
	PRIORITY string
}

func ToStructObjects(rows *sql.Rows) ([]TableContent) {
	dbData := []TableContent{}
	if (rows != nil) {
		for rows.Next() {
			var r TableContent
			rows.Scan(&r.TASKID, &r.TASK, &r.PRIORITY)
			dbData = append(dbData, r)
		}
	}
	return dbData
}

func ToArrayOfString(rows *sql.Rows) ([][]string) {
	dbData := [][]string{}
	data := ToStructObjects(rows)
	for i := 0; i < len(data); i++ {
		r := []string{strconv.Itoa(data[i].TASKID),data[i].TASK,data[i].PRIORITY}
		dbData = append(dbData,r)
	}
	return dbData
}