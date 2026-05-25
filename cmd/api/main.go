package main

import (
	"context"
	"net/http"
	"os"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"github.com/MohdAdamSiThuThetNaing/technical-test-fluxion/internal/ai"
	"github.com/MohdAdamSiThuThetNaing/technical-test-fluxion/internal/db"
	"github.com/MohdAdamSiThuThetNaing/technical-test-fluxion/internal/guard"
	"github.com/MohdAdamSiThuThetNaing/technical-test-fluxion/internal/logs"
	"github.com/MohdAdamSiThuThetNaing/technical-test-fluxion/internal/queue"
	"github.com/MohdAdamSiThuThetNaing/technical-test-fluxion/internal/users"

	"golang.org/x/crypto/bcrypt"
)

func main() {

	db.ConnectPostgres()
	db.ConnectMongo()

	db.DB.AutoMigrate(&users.User{})

	r := gin.Default()

	store := cookie.NewStore([]byte("super-secret-session"))
	r.Use(sessions.Sessions("fluxion_session", store))
	r.LoadHTMLGlob("templates/*")

	r.GET("/login", func(c *gin.Context) {

		c.HTML(http.StatusOK, "login.html", nil)
	})

	r.POST("/login", func(c *gin.Context) {

		email := c.PostForm("email")
		password := c.PostForm("password")
		adminEmail := os.Getenv("ADMIN_EMAIL")
		adminPassword := os.Getenv("ADMIN_PASSWORD")

		if email != adminEmail || password != adminPassword {

			c.HTML(http.StatusUnauthorized, "login.html", gin.H{
				"Error": "Invalid credentials",
			})
			return
		}

		session := sessions.Default(c)
		session.Set("admin_logged_in", true)
		session.Save()

		c.Redirect(http.StatusFound, "/")
	})

	r.GET("/logout", func(c *gin.Context) {

		session := sessions.Default(c)
		session.Clear()
		session.Save()
		c.Redirect(http.StatusFound, "/login")
	})

	r.GET("/", guard.AuthMiddleware(), func(c *gin.Context) {

		var totalUsers int64
		db.DB.Model(&users.User{}).Count(&totalUsers)

		c.HTML(http.StatusOK, "dashboard.html", gin.H{
			"TotalUsers": totalUsers,
		})
	})

	r.GET("/ai/test", guard.AuthMiddleware(), func(c *gin.Context) {

		collection := db.MongoClient.Database("fluxion_logs").Collection("logs")
		cursor, err := collection.Find(context.Background(), map[string]interface{}{})

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		defer cursor.Close(context.Background())
		var logsText string

		for cursor.Next(context.Background()) {
			var logData logs.Log
			cursor.Decode(&logData)
			logsText += "Event: " + logData.Event + ", Email: " + logData.UserEmail + ", Name: " + logData.UpdatedName + "\n"
		}
		result, err := ai.GenerateSummary(logsText)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"result": result,
		})
	})

	// LOGS
	r.GET("/logs", guard.AuthMiddleware(), func(c *gin.Context) {

		collection := db.MongoClient.
			Database("fluxion_logs").
			Collection("logs")

		cursor, err := collection.Find(
			context.Background(),
			map[string]interface{}{},
		)

		if err != nil {
			c.JSON(
				http.StatusInternalServerError,
				gin.H{
					"error": err.Error(),
				},
			)
			return
		}

		defer cursor.Close(
			context.Background(),
		)

		var logsList []logs.Log
		for cursor.Next(
			context.Background(),
		) {
			var logData logs.Log
			cursor.Decode(&logData)
			logsList = append(
				logsList,
				logData,
			)
		}

		c.HTML(
			http.StatusOK,
			"logs.html",
			gin.H{
				"Logs": logsList,
			},
		)
	})

	r.GET("/users", guard.AuthMiddleware(), func(c *gin.Context) {

		var usersList []users.User
		db.DB.Find(&usersList)

		c.HTML(http.StatusOK, "users.html", gin.H{
			"Users": usersList,
		})
	})

	r.GET("/users/create", guard.AuthMiddleware(), func(c *gin.Context) {

		c.HTML(http.StatusOK, "create-user.html", nil)
	})

	r.POST("/users/create", guard.AuthMiddleware(), func(c *gin.Context) {

		name := c.PostForm("name")
		email := c.PostForm("email")
		password := c.PostForm("password")
		hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

		user := users.User{
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
	})

	r.GET("/users/edit/:id", guard.AuthMiddleware(), func(c *gin.Context) {

		id := c.Param("id")
		var user users.User
		result := db.DB.Where("id = ?", id).First(&user)
		if result.Error != nil {
			c.String(http.StatusNotFound, "User not found")
			return
		}

		c.HTML(http.StatusOK, "edit-user.html", gin.H{
			"User": user,
		})
	})

	r.POST("/users/edit/:id", guard.AuthMiddleware(), func(c *gin.Context) {

		id := c.Param("id")

		var user users.User
		result := db.DB.Where("id = ?", id).First(&user)
		if result.Error != nil {
			c.String(http.StatusNotFound, "User not found")
			return
		}

		user.Name = c.PostForm("name")
		user.Email = c.PostForm("email")
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
	})

	r.POST("/users/delete/:id", guard.AuthMiddleware(), func(c *gin.Context) {

		id := c.Param("id")

		var user users.User
		db.DB.Where("id = ?", id).First(&user)
		db.DB.Delete(&users.User{}, "id = ?", id)

		queue.Publish(gin.H{
			"event":        "USER_DELETED",
			"user_id":      user.ID.String(),
			"user_email":   user.Email,
			"updated_name": user.Name,
			"action_by":    "admin",
			"created_at":   time.Now(),
		})
		c.Redirect(http.StatusFound, "/users")
	})

	r.Run(":8080")
}