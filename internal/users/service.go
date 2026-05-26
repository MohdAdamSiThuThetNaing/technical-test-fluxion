package users

import (
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"

	"github.com/MohdAdamSiThuThetNaing/technical-test-fluxion/internal/audit"
	"github.com/MohdAdamSiThuThetNaing/technical-test-fluxion/internal/db"
	"github.com/MohdAdamSiThuThetNaing/technical-test-fluxion/internal/queue"
)

type UserServiceReason string

const (

	SYSTEM_ACTION UserServiceReason = "SYSTEM_ACTION"
)

type Service struct{}

func (s *Service) CreateUser(input CreateUserDTO) error {

	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(input.Password),
		bcrypt.DefaultCost,
	)

	if err != nil {
		return err
	}

	user := User{
		ID:       uuid.New(),
		Name:     input.Name,
		Email:    input.Email,
		Password: string(hashedPassword),
	}

	err = db.DB.Create(&user).Error

	if err != nil {
		return err
	}

	queue.Publish(map[string]interface{}{
		"user_id":      user.ID,
		"event":        string(audit.USER_CREATED),
		"data":         user,
		"action_by":    string(SYSTEM_ACTION),
		"created_at":   time.Now(),
	})

	return nil
}

func (s *Service) GetUsers() ([]User, error) {

	var users []User

	err := db.DB.Find(&users).Error

	if err != nil {
		return nil, err
	}

	return users, nil
}

func (s *Service) GetUserByID(id string) (*User, error) {

	var user User

	err := db.DB.First(
		&user,
		"id = ?",
		id,
	).Error

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (s *Service) UpdateUser(id string, input UpdateUserDTO) error {

	var user User

	err := db.DB.First(
		&user,
		"id = ?",
		id,
	).Error

	if err != nil {
		return err
	}

	user.Name = input.Name
	user.Email = input.Email
	err = db.DB.Save(&user).Error

	if err != nil {
		return err
	}

	queue.Publish(map[string]interface{}{
		"user_id":      user.ID,
		"event":        string(audit.USER_UPDATED),
		"data":         user,
		"action_by":    string(SYSTEM_ACTION),
		"created_at":   time.Now(),
	})

	return nil
}

func (s *Service) DeleteUser(id string) error {

	var user User

	err := db.DB.First(
		&user,
		"id = ?",
		id,
	).Error

	if err != nil {
		return err
	}

	err = db.DB.Delete(&user).Error

	if err != nil {
		return err
	}

	queue.Publish(map[string]interface{}{
		"user_id":      user.ID,
		"event":        string(audit.USER_DELETED),
		"data":         user,
		"action_by":    string(SYSTEM_ACTION),
		"created_at":   time.Now(),
	})

	return nil
}