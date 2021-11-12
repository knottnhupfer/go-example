package main

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"service/handlers"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	"gopkg.in/yaml.v2"
)

var db *sql.DB

type Configuration struct {
	Service struct {
		BindAddress     string `yaml:"bindAddress"`
		DbConnectString string `yaml:"dbConnectString"`
	}
}

func main() {

	conf, _ := readConfiguration("config/application.yaml")

	connectToMySqlDatabase(*conf)
	connectToPostgresDatabase(*conf)

	log.Println("Starting server on address: ", conf.Service.BindAddress)
	handlers.RegisterHandlers()
	http.ListenAndServe(conf.Service.BindAddress, nil)
}

func connectToMySqlDatabase(conf Configuration) {
	db, err := sql.Open("mysql", "mysql:secret123@tcp(mysql:3306)/demomysql")

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var version string

	err2 := db.QueryRow("SELECT VERSION()").Scan(&version)

	if err2 != nil {
		log.Fatal(err2)
	}

	log.Println("MySQL version: ", version)
}

func connectToPostgresDatabase(conf Configuration) {

	log.Println("Configured db connection string is: ", conf.Service.DbConnectString)

	db, err := sql.Open("postgres", conf.Service.DbConnectString)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	log.Println("Database connected!")
}

func readConfiguration(filename string) (*Configuration, error) {

	buf, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Println("Configuration file not found, use default values.")
		configuration := Configuration{}
		configuration.Service.BindAddress = "0.0.0.0:7080"
		configuration.Service.DbConnectString = "user=postgres dbname=demopostgres password=secure host=0.0.0.0 sslmode=disable"
		return &configuration, err
	}

	c := &Configuration{}
	err = yaml.Unmarshal(buf, c)
	if err != nil {
		fmt.Println("Unable to unmarshal configuration, terminate application.")
		panic(err)
	}

	log.Println("Loaded configuration!")
	return c, nil
}
