package main

import (
	// "github.com/couchbase/gocb/v2"
	_ "github.com/joho/godotenv/autoload"
	"prima-integrasi.com/fendiya/syncronizer/internal/server"
)

func main() {
	// gocb.SetLogger(gocb.VerboseStdioLogger())
	server.Run()

}
