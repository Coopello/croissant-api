package infrastructure

import (
	"CoopeLunch-api/interfaces/database"
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type SqlHandler struct {
	Conn *sql.DB
}

type SqlRows struct {
	Rows *sql.Rows
}

func NewSqlHandler() database.SqlHandler {
	conn, err := sql.Open("mysql", fmt.Sprintf("%s:%s%s/%s", "go_test", "password", "@tcp(db:3306)", "go_database"))
	if err != nil {
		panic(err.Error())
	}
	sqlHandler := new(SqlHandler)
	sqlHandler.Conn = conn

	return sqlHandler
}

func (handler *SqlHandler) ErrNoRows() error {
	return sql.ErrNoRows
}

func (handler *SqlHandler) Query(statement string, args ...interface{}) (database.Rows, error) {
	rows, err := handler.Conn.Query(statement, args...)
	if err != nil {
		return new(SqlRows), err
	}
	row := new(SqlRows)
	row.Rows = rows
	return rows, nil
}

type SqlResult struct {
	Result sql.Result
}

func (handler *SqlHandler) Execute(statement string, args ...interface{}) (database.Result, error) {
	res := SqlResult{}
	stmt, err := handler.Conn.Prepare(statement)
	defer stmt.Close()
	if err != nil {
		return res, err
	}
	exe, err := stmt.Exec(args...)
	if err != nil {
		return res, err
	}
	res.Result = exe
	return res, nil
}

func (r SqlRows) Scan(dest ...interface{}) error {
	return r.Rows.Scan(dest...)
}

func (r SqlRows) Next() bool {
	return r.Rows.Next()
}

func (r SqlRows) Close() error {
	return r.Rows.Close()
}

func (r SqlResult) LastInsertId() (int64, error) {
	return r.Result.LastInsertId()
}

func (r SqlResult) RowsAffected() (int64, error) {
	return r.Result.RowsAffected()
}
