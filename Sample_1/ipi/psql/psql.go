package psql

import (
	"Sample_1/ipi/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Dbmanager struct {
	*gorm.DB
}

func NewDbmanager() (*Dbmanager, error) {
	db, err := gorm.Open(postgres.Open("host=localhost user=postgres password=1234 dbname=Hieu port=5432 sslmode=disable TimeZone=Asia/Ho_Chi_Minh"))
	if err != nil {
		log.Fatal("error conecting dattabase", err)
	}
	db = db.Debug()
	err = db.AutoMigrate(&models.User{}, &models.Auth_token{})
	if err != nil {
		log.Fatal("error create dattabase", err)
	}
	return &Dbmanager{
		DB: db,
	}, nil
}
