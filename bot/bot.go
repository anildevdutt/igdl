package bot

import (
	"database/sql"
	"log"

	"github.com/anildevdutt/igdl/igapi"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB
var dbfile = "db/igdl.db"

func chkerr(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func SaveUserData(userdata igapi.UserData) {
	_, err := db.Exec("")
	chkerr(err)
}

func Start() {
	var err error
	db, err = sql.Open("sqlite3", dbfile)
	chkerr(err)

	igapi.GetUserNameInfo("rem_ram_zeros")
}
