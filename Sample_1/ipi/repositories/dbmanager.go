package repositories

import (
	"Sample_1/ipi/models"
	"fmt"
	"log"
	"strings"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Dbmanager struct {
	*gorm.DB
}

type UserRepositories interface {
	GetUserbyToken(tocken string) (*models.User, error)
	GetTocken(tocken string) (bool, error)
}

func (m *Dbmanager) GetTocken(tocken string) (bool, error) {
	auth := models.Auth_token{}
	tocken1 := strings.Split(tocken, " ")
	s := ""
	for i := 1; i < len(tocken1); i++ {
		if i == 1 {
			s = s + tocken1[i]
		} else {
			s = s + " " + tocken1[i]
		}
	}
	if err := m.Where(&models.Auth_token{Tocken: s}).First(&auth).Error; err != nil {
		return false, err
	}
	return true, nil
}
func NewDbmanager() (UserRepositories, error) {
	db, err := gorm.Open(postgres.Open("host=localhost user=postgres password=1234 dbname=Hieu port=5432 sslmode=disable TimeZone=Asia/Ho_Chi_Minh"))
	if err != nil {
		log.Fatal("error conecting dattabase", err)
	}
	db = db.Debug()
	err = db.AutoMigrate(&models.User{}, &models.Auth_token{})
	if err != nil {
		log.Fatal("error create dattabase", err)
	}
	return &Dbmanager{db}, nil
}

func (m *Dbmanager) GetUserbyToken(tocken string) (*models.User, error) {
	user := models.User{}
	auth := models.Auth_token{}
	tocken1 := strings.Split(tocken, " ")
	s := ""
	for i := 1; i < len(tocken1); i++ {
		if i == 1 {
			s = s + tocken1[i]
		} else {
			s = s + " " + tocken1[i]
		}
	}
	fmt.Println(s)
	if err := m.Where(&models.Auth_token{Tocken: s}).First(&auth).Error; err != nil {
		return nil, err
	}
	if err := m.Where(&models.User{Id: auth.UserID}).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
