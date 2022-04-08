package data

import (
	"fmt"

	_ "github.com/lib/pq"
	"xorm.io/xorm"
)

var Engine *xorm.Engine

type User struct {
	Id       int64
	Name     string
	Email    string
	Password string `json:"-"`
}

var (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "password"
	dbname   = "authserver"
)

func CreateDBEngine() (*xorm.Engine, error) {
	connectionInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	var err error
	Engine, err = xorm.NewEngine("postgres", connectionInfo)

	if err != nil {
		return nil, err
	}

	if err := Engine.Ping(); err != nil {
		return nil, err
	}

	if err := Engine.Sync(new(User)); err != nil {
		return nil, err
	}

	return Engine, nil

}
