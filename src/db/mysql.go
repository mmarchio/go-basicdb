package db

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
)

const CONN_STRING = "username:password@tcp(127.0.0.1:3306)/test"

type IModel interface{
	isModel()
	DBHydrate(*sql.Rows)
	ToMSI() map[string]interface{}
}

type MySQLDatabase struct {
	Id string `json:"id"`
	Instance *sql.DB `json:"instance"`
}

func (c MySQLDatabase) Connect() {
	db, err := sql.Open("mysql", CONN_STRING)
	if Error(err) {
		panic(err.Error)
	}
	c.Instance = db
}

func (c MySQLDatabase) Query(q string, model IModel) (*Result, error) {
	r, err := c.Instance.Query(q)
	if Error(err) {
		return nil, err
	}
	model.DBHydrate(r)
	numRows := 0
	for r.Next() {
		numRows += 1
	}
	result := Result{
		Id: uuid.NewString(),
		DataFrame: model.ToMSI(),
		NumRows: int64(numRows),
	}
	return &result, nil
}

func (c MySQLDatabase) QueryCount(q string) (int64, error) {
	count := int64(0)
	err := c.Instance.QueryRow(q).Scan(&count)
	return count, err
}

func (c MySQLDatabase) Insert(q string) (*Result, error) {
	result := Result{}
	res, err := c.Instance.Exec(q)
	if Error(err) {
		return nil, err
	}
	result.Id = uuid.NewString()
	result.NumRows, err = res.RowsAffected()
	if Error(err) {
		return nil, err
	}
	return &result, err
}

func Error(e error) bool {
	return e != nil
}

type Result struct {
	Id string `json:"id"`
	DataFrame map[string]interface{} `json:"dataFrame"`
	NumRows int64 `json:"numRows"`
}

