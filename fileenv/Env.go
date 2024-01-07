package fileenv

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Env struct {
	DBdriver string
	DBuser   string
	DBpass   string
	DBname   string
}

func NewEnv() (*Env, error) {
	err := godotenv.Load("fileenv/mysql.env")

	if err != nil {
		log.Fatal("Error loadinh .env file: ", err)
	}

	env := &Env{
		DBdriver: os.Getenv("DB_Driver"),
		DBuser:   os.Getenv("DB_User"),
		DBpass:   os.Getenv("DB_Pass"),
		DBname:   os.Getenv("DB_name"),
	}
	// fmt.Println(env)
	return env, nil
}
