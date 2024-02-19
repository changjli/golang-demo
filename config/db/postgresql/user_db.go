package user_db_postgresql

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

// func name(parameter input) (parameter output, parameter output)

// connect ke db
func OpenConnection() *sql.DB {
	// var name string
	// name = "audric"
	// name := "audric"
	// fmt.Println(name)

	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	psqlInfo := fmt.Sprintf(`
		host=%s port=%s user=%s password=%s dbname=%s sslmode=disable`,
		os.Getenv("DB_HOST_POSTGRES"),
		os.Getenv("DB_PORT_POSTGRES"),
		os.Getenv("DB_USER_POSTGRES"),
		os.Getenv("DB_PW_POSTGRES"),
		os.Getenv("DB_NAME_POSTGRES"),
	)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		//langsung matiin os
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	return db
}
