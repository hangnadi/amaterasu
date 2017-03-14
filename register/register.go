package register

import (
	"fmt"
	x "github.com/hangnadi/amaterasu/common"
	sql "github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
	"net/http"
)

var Database *sql.DB

const (
	DB_USER           = "postgres"
	DB_PASSWORD       = "postgres"
	DB_NAME           = "ProjectIcowDB"
	FORMAT_DATASOURCE = "user=%s password=%s dbname=%s sslmode=disable"
)

func init() {
	driverName := "postgres"
	dataSourceName := fmt.Sprintf(FORMAT_DATASOURCE,
		DB_USER, DB_PASSWORD, DB_NAME)

	// Open db connection
	db, err := sql.Open(driverName, dataSourceName)

	if err != nil {
		// Cannot open database
		log.Fatal(err.Error())
	}

	if err := db.Ping(); err != nil {
		// Cannot establish db connection
		log.Fatal(err.Error())
	}

	Database = db
}

func main() {

}

func ProvideRegister() http.HandlerFunc {
	return func(rw http.ResponseWriter, req *http.Request) {
		tx, err := Database.Begin()
		if err != nil {
			// Cannot begin transaction
			log.Fatal(err.Error())
		}

		username := req.PostFormValue("username")
		fullname := req.PostFormValue("fullname")
		email := req.PostFormValue("email")
		password := req.PostFormValue("password")
		passwordHash := x.GeneratePasswordHash(password)

		query := `INSERT INTO tb_user 
		(username, fullname, email, password_hash, is_verified)
		VALUES ($1, $2, $3, $4, $5)`

		_, err = tx.Exec(query, username, fullname, email, passwordHash, false)

		if err != nil {
			tx.Rollback()
			log.Fatal(err.Error())
		}

		err = tx.Commit()
		if err != nil {
			tx.Rollback()
			log.Fatal(err.Error())
		}

	}
}
