package hahnair

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "postgres"
)

func saveCard(car Card) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic("could not connect to postgres")
	}

	fmt.Fprintln(os.Stdout, "Successfully connected to postgres!")

	// only a sample to disply concept
	sqlStatement := `
INSERT INTO cards_info (card_holder_name, card_number)
VALUES ('Tariq Damra', '123456789')`
	_, err = db.Exec(sqlStatement)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not execute psql command: %s\n", err)
	}
}
