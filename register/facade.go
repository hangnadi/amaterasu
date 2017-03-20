package register

import (
	"fmt"
	x "github.com/hangnadi/amaterasu/common"
	sql "github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
)

var Database *sql.DB

const (
	DB_USER           = "postgres"
	DB_PASSWORD       = "postgres"
	DB_NAME           = "ProjectIcowDB"
	DB_HOST           = "localhost"
	DB_PORT           = 5432
	FORMAT_DATASOURCE = "host=%s port=%d user=%s password=%s dbname=%s sslmode=disable"
)

func init() {
	driverName := "postgres"
	dataSourceName := fmt.Sprintf(FORMAT_DATASOURCE,
		DB_HOST, DB_PORT, DB_USER, DB_PASSWORD, DB_NAME)

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

func NewAccounts(model *RequestModel) int {

	passwordHash := x.GeneratePasswordHash(model.password)

	query := `INSERT INTO tb_user 
		(username, fullname, email, password_hash, is_verified)
		VALUES ($1, $2, $3, $4, $5)`

	_, err := Database.Exec(query, model.username, model.fullname, model.email, passwordHash, false)

	if err != nil {
		return 0
	}

	return 1
}
