package service_user

import (
	"errors"

	"github.com/dijer/otus-highload/backend/internal/models"
	storage_user "github.com/dijer/otus-highload/backend/internal/storage/user"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	storage *storage_user.UserStorage
}

func New(storage *storage_user.UserStorage) *UserService {
	return &UserService{
		storage: storage,
	}
}

func (s *UserService) CreateUser(user models.UserWithPassword) error {
	hashedPassword, err := hashPassword(user.Password)
	if err != nil {
		return err
	}

	return s.storage.CreateUser(user.User, hashedPassword)
}

func (s *UserService) CheckUserPassword(user models.UserWithPassword) (int, error) {
	hashedPassword, userID, err := s.storage.GetHashedPassword(user.UserName)
	if err != nil {
		return 0, err
	}

	if !checkHashedPassword(user.Password, hashedPassword) {
		return 0, errors.New("not match")
	}

	return userID, nil
}

func (s *UserService) GetUser(userID int) (*models.User, error) {
	return s.storage.GetUser(userID)
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func checkHashedPassword(password, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}

func (s *UserService) GetUsers(firstname, lastname string) ([]models.User, error) {
	return s.storage.GetUsers(firstname, lastname)
}
