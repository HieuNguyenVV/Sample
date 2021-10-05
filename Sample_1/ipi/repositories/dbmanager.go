package repositories

import (
	"Sample_1/ipi/models"
	"Sample_1/ipi/psql"
	"strings"
)

type IUserRepository interface {
	GetUserbyID(id int64) (*models.User, error)
	GetTocken(tocken string) (*models.Auth_token, error)
}
type UserRepository struct {
	db *psql.Dbmanager
}

func NewUserRepository(db *psql.Dbmanager) IUserRepository {
	return &UserRepository{
		db: db,
	}
}
func (m *UserRepository) GetTocken(tocken string) (*models.Auth_token, error) {
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
	if err := m.db.Where(&models.Auth_token{Tocken: s}).First(&auth).Error; err != nil {
		return nil, err
	}
	return &auth, nil
}

func (m *UserRepository) GetUserbyID(id int64) (*models.User, error) {
	user := models.User{}
	if err := m.db.Where(&models.User{Id: id}).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
