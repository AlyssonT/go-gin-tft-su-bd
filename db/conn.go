package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

// const (
// 	host     = "tft_su_bd_db"
// 	port     = 5432
// 	user     = "postgres"
// 	password = "1234"
// 	dbname   = "postgres"
// )

const urldb = "postgresql://dpg-ctc98fd2ng1s73c0j0mg-a.ohio-postgres.render.com/tft_su_bd_db?user=tft_su_bd_db_user&password=Fo7hUnm3J1iLkzvAFvzdqk5LG2gkApgm"

func ConnectDB() (*sql.DB, error) {
	// psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
	// 	"password=%s dbname=%s sslmode=disable",
	// 	host, port, user, password, dbname)
	db, err := sql.Open("postgres", urldb)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Connected to " + urldb)

	return db, nil
}
