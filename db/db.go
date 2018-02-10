package db

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/go-gorp/gorp"
	_ "github.com/lib/pq" //import postgres
)

//DB ...
type DB struct {
	*sql.DB
}

const (
	host     = "localhost"
	port     = 5432
	user     = "ichabod"
	password = "A3fq7Hx3uCnu9iXLsv3C"
	dbname   = "ichabod"
)

var db *gorp.DbMap

//Init ...
func Init() {

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	var err error
	db, err = ConnectDB(psqlInfo)

	if err != nil {
		log.Fatal(err)
	}
}

//ConnectDB ...
func ConnectDB(dataSourceName string) (*gorp.DbMap, error) {
	db, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.PostgresDialect{}}
	//dbmap.TraceOn("[gorp]", log.New(os.Stdout, "golang-gin:", log.Lmicroseconds)) //Trace database requests
	return dbmap, nil
}

//GetDB ...
func GetDB() *gorp.DbMap {
	return db
}
