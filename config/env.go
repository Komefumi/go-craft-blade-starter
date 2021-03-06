package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"regexp"

	_ "github.com/lib/pq"
)

var envValueRegexp *regexp.Regexp = regexp.MustCompile("(\\w+)=(.+)")
var DB *sql.DB
var ServerPortString string
var JWTSecret string

func init() {
	setEnvironmentVariables()
	ServerPortString = os.Getenv("SERVER_PORT")
	JWTSecret = os.Getenv("JWT_SECRET")
	setupDB()
}

func setEnvironmentVariables() {
	res, err := os.ReadFile(".env")
	if err != nil {
		log.Fatal(err)
	}
	matches := envValueRegexp.FindAllStringSubmatch(string(res), -1)
	for _, v := range matches {
		os.Setenv(v[1], v[2])
	}
	// fmt.Println(os.Getenv("SERVER_PORT"))
	// fmt.Println(os.Getenv("DB_DSN"))
}

func setupDB() {
	// _, err := sql.Open("postgres", "dbname=craft-blade; sslmode=disable;")
	var err error
	DB, err = sql.Open(os.Getenv("DB_TYPE"), os.Getenv("DB_DSN"))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connection Established!")
	// fmt.Println(DB)
}

func CloseDB() {
	DB.Close()
}

func main() {
	var col string
	row := DB.QueryRow("SELECT 10;")
	err := row.Scan(&col)
	if err != nil {
		log.Fatal(err)
	}
}
