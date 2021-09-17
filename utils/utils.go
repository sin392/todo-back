package utils

import (
	"database/sql"
	"encoding/json"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func GetConnection() *sql.DB {
	// about param: https://budougumi0617.github.io/2019/03/31/go_db_unsupported_scan_storing_driver_value_type_uint8_into_type_time_time/
	db, err := sql.Open("mysql", "user:password@/todo_db?parseTime=true")
	if err != nil {
		panic(err.Error())
	}
	return db
}

func DecodeBody(r *http.Request, out interface{}) error {
	defer r.Body.Close()
	decoder := json.NewDecoder(r.Body)
	return decoder.Decode(out)
}

// func EncodeBody(w *http.ResponseWriter, out interface{}) error {
// 	w.Header().Set("content-type", "application/json")
// 	var res
// 	defer w.Body.Close()
// 	encoder := json.NewEncoder(res)
// 	return encoder.Decode(out)
// }
