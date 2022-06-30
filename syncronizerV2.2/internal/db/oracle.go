package db

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/sijms/go-ora/v2"
)

type Oracle struct {
	DbProperties
}

func (p Oracle) OpenConn() *sqlx.DB {
	url := "oracle://" + p.Username + ":" + p.Password + "@" + p.Hostname + ":" + p.Port + "/" + p.Dbname
	log.Println("connect to db using url :" + url)
	db, err := sqlx.Open("oracle", url)
	if err != nil {
		panic(fmt.Errorf("error in sql.Open: %v", err))
	}

	if err := db.Ping(); err != nil {
		panic(fmt.Errorf("error in db.Ping: %v", err))
	}
	fmt.Println("Connection Open to Oracle Success!")
	return db
}

func (p Oracle) CloseConn(c *sqlx.DB) {
	err := c.Close()
	if err != nil {
		fmt.Println("Can't close connection: ", err)
	}
	fmt.Println("Connection to Oracle Closed!")
}
