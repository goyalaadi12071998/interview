package db

import (
	"fmt"
	"interview/cmd/migrations"
	"log"
	"strconv"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DbClient *gorm.DB

type Config struct {
	Name     string
	Host     string
	Port     int
	Username string
	Password string
}

func GetDb() *gorm.DB {
	return DbClient
}

func InitDB(config Config) error {
	db, err := connectDB(config)
	if err != nil {
		return err
	}
	DbClient = db
	return nil
}

func connectDB(config Config) (*gorm.DB, error) {
	dbConnectionUrl := getDbConnectionUrl(config)
	db, err := gorm.Open(mysql.Open(dbConnectionUrl), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	fmt.Println("Db connection successful")

	res := db.Exec(migrations.CREATE_USER_MODEL)
	if res.Error != nil {
		log.Fatal(err)
	}

	fmt.Println("USER table initialize successfully")

	return db, nil
}

func getDbConnectionUrl(config Config) string {
	return config.Username + ":" + config.Password + "@tcp" + "(" + config.Host + ":" + strconv.Itoa(config.Port) + ")/" + config.Name + "?" + "parseTime=true&loc=Local"
}
