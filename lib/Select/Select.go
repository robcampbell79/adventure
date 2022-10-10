package Select

import(
	"log"
	"adventure/lib/Connect"
)

type Tester struct {
	Id		int
	Name	string
	Num		int
	Rando	string
}

func SelAll() Tester {
	db, err := Connect.Conn()
	if err != nil {
		log.Fatal(err.Error())
	}

	sqlParse, err := db.Query("SELECT * FROM test")
	if err != nil {
		log.Fatal(err.Error())
	}

	defer sqlParse.Close()
	defer db.Close()

	res := Tester{}

	for sqlParse.Next() {
		err := sqlParse.Scan(&res.Id, &res.Name, &res.Num, &res.Rando)
		if err != nil {
			log.Fatal(err.Error())
		}
	}

	return res
}