package db

import (
	"github.com/couchbase/gocb/v2"
	"github.com/jmoiron/sqlx"
)

type Connection interface {
	OpenConn() *sqlx.DB
	CloseConn(*sqlx.DB)
}

type ConnectionCouchbaseSDK interface {
	OpenConn() *gocb.Cluster
}

type DbProperties struct {
	Hostname string
	Port     string
	Dbname   string
	Username string
	Password string
}
