package users

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"

	"github.com/MohdAdamSiThuThetNaing/technical-test-fluxion/internal/db"
	"github.com/MohdAdamSiThuThetNaing/technical-test-fluxion/internal/queue"
)

func ListUsers(c *gin.Context) {

	var usersList []User
	db.DB.Find(&usersList)
	c.HTML(http.StatusOK, "users.html", gin.H{
		"Users": usersList,
	})
}

func ShowCreateUser(c *gin.Context) {

	c.HTML(http.StatusOK, "create-user.html", nil)
}

func CreateUser(c *gin.Context) {

	name := c.PostForm("name")
	email := c.PostForm("email")
	password := c.PostForm("password")

	var existingUser User
	if err := db.DB.
		Where("email = ?", email).
		First(&existingUser).Error; err == nil {

		c.HTML(http.StatusBadRequest, "create-user.html", gin.H{
			"Error": "Email already exists",
		})
		return
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword(
		[]byte(password),
		bcrypt.DefaultCost,
	)

	user := User{
		ID:       uuid.New(),
		Name:     name,
		Email:    email,
		Password: string(hashedPassword),
	}

	db.DB.Create(&user)

	queue.Publish(gin.H{
		"event":        "USER_CREATED",
		"user_id":      user.ID.String(),
		"user_email":   user.Email,
		"updated_name": user.Name,
		"action_by":    "admin",
		"created_at":   time.Now(),
	})

	c.Redirect(http.StatusFound, "/users")
}

func ShowEditUser(c *gin.Context) {

	id := c.Param("id")
	var user User

	result := db.DB.
		Where("id = ?", id).
		First(&user)

	if result.Error != nil {

		c.String(http.StatusNotFound, "User not found")
		return
	}

	c.HTML(http.StatusOK, "edit-user.html", gin.H{
		"User": user,
	})
}

func UpdateUser(c *gin.Context) {

	id := c.Param("id")
	var user User

	result := db.DB.
		Where("id = ?", id).
		First(&user)

	if result.Error != nil {

		c.String(http.StatusNotFound, "User not found")
		return
	}

	name := c.PostForm("name")
	email := c.PostForm("email")
	var existingUser User

	if err := db.DB.
		Where("email = ? AND id != ?", email, id).
		First(&existingUser).Error; err == nil {

		c.HTML(http.StatusBadRequest, "edit-user.html", gin.H{
			"User":  user,
			"Error": "Email already exists",
		})
		return
	}

	user.Name = name
	user.Email = email
	db.DB.Save(&user)

	queue.Publish(gin.H{
		"event":        "USER_UPDATED",
		"user_id":      user.ID.String(),
		"user_email":   user.Email,
		"updated_name": user.Name,
		"action_by":    "admin",
		"created_at":   time.Now(),
	})

	c.Redirect(http.StatusFound, "/users")
}

func DeleteUser(c *gin.Context) {

	id := c.Param("id")
	var user User
	db.DB.Where("id = ?", id).First(&user)
	db.DB.Delete(&User{}, "id = ?", id)

	queue.Publish(gin.H{
		"event":        "USER_DELETED",
		"user_id":      user.ID.String(),
		"user_email":   user.Email,
		"updated_name": user.Name,
		"action_by":    "admin",
		"created_at":   time.Now(),
	})

	c.Redirect(http.StatusFound, "/users")
}