package repository

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

const (
	blogs = "Blogs"
)

type Config struct {
	User   string
	Passwd string
	Net    string
	Addr   string
	DBName string
}

func NewMysqlDB(cfg Config) (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:root@/ItBlog")
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
