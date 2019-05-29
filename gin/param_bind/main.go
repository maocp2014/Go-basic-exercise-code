package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	Username string `form:"username" json:"username" binding:"required"`
	Passwd   string `form:"passwd" json:"passwd" bdinding:"required"`
	Age      int    `form:"age" json:"age"`
}

// 1、方式1
/*
func main() {
	engine := gin.Default()

	engine.POST("/login", func(c *gin.Context) {
		var user User
		var err error
		contentType := c.Request.Header.Get("Content-Type")

		switch contentType {
			case "application/json":
				err = c.BindJSON(&user)
			case "application/x-www-form-urlencoded":
				err = c.BindWith(&user, binding.Form)
		}

		if err != nil {
			fmt.Println(err)
			log.Fatal(err)
		}

		c.JSON(http.StatusOK, gin.H{
			"user":   user.Username,
			"passwd": user.Passwd,
			"age":    user.Age,
		})
	})

    engine.Run()
}
*/

// 2、方式2
func main() {
	engine := gin.Default()

	engine.POST("/login", func(c *gin.Context) {
		var user User

		err := c.Bind(&user)
		if err != nil {
			fmt.Println(err)
			log.Fatal(err)
		}

		c.JSON(http.StatusOK, gin.H{
			"username": user.Username,
			"passwd":   user.Passwd,
			"age":      user.Age,
		})
	})

	engine.Run()
}
