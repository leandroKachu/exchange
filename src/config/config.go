package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	Port             = 0
	ConnectionString = ""
	ApiKey           = ""
)

// InitBaseConfig:

func InitBaseConfig() {
	var err error
	if err = godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	ApiKey = os.Getenv("APIKEY")

	Port, err = strconv.Atoi(os.Getenv("API_PORT"))
	if err != nil {
		Port = 9000
	}
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")

	// Fazendo a conexão do banco usando a string 'sprintf' e passando as variasveis
	ConnectionString = fmt.Sprintf("%s:%s@tcp(localhost:3306)/%s?charset=utf8&parseTime=True&loc=Local",
		dbUser,
		dbPass,
		dbName,
	)

}
