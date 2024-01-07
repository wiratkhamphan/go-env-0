package database

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/wiratkhamphan/go-env-0/fileenv"
)

type MySQLStore struct {
	db *sql.DB
}

func DBConnection() (*sql.DB, error) {
	env, err := fileenv.NewEnv()
	if err != nil {
		return nil, err
	}
	db, err := sql.Open(env.DBdriver, fmt.Sprintf("%s:%s@/%s", env.DBuser, env.DBpass, env.DBname))
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	// defer db.Close()
	return db, err

}
func Newdb() {

}
