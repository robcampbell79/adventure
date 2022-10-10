package Connect

import(
	"fmt"
	"log"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func Conn() (db *sql.DB, err error) {
	db, err = sql.Open("mysql", "root:toor@tcp(127.0.0.1:3306)/adventure")
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println("Connected to db")

	return db, err
}