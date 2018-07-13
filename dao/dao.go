package dao

import (
	"database/sql"
	"reflect"
)

func getRows(rows *sql.Rows, varType interface{}) []interface{} {

	result := make([]interface{}, 0)
	s := reflect.ValueOf(varType).Elem()
	leng := s.NumField()
	onerow := make([]interface{}, leng)
	for i := 0; i < leng; i++ {
		onerow[i] = s.Field(i).Addr().Interface()
	}
	for rows.Next() {
		rows.Scan(onerow...)
		result = append(result, s.Interface())
	}
	return result
}

func Get(query string, varType interface{}, cond ...interface{}) []interface{} {

	db, err := sql.Open("mysql", dataSource)
	if err != nil {
		return nil
	}
	defer db.Close()

	stmt, err := db.Prepare(query)
	if err != nil {
		return nil
	}
	defer stmt.Close()
	rows, err := stmt.Query(cond...)
	if err != nil {
		return nil
	}
	defer rows.Close()
	return getRows(rows, varType)
}

func Modify(modify string, cond ...interface{}) (int, error) {
	db, err := sql.Open("mysql", dataSource)
	if err != nil {
		return 0, err
	}
	defer db.Close()

	stmt, err := db.Prepare(modify)
	if err != nil {
		return 0, err
	}

	res, err := stmt.Exec(cond...)
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(id), nil

}
