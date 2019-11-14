package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

type User struct {
	Id       int    `json:"id" form:"id"`
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

var users []User

func GetUsersController(c echo.Context) error {
	// users = []User{{Id: 1, Name: "yudis", Email: "email", Password: "1111"}, {Id: 2, Name: "haloo", Email: "haloo", Password: "haloo"}}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all users",
		"users":   users,
	})
}
func GetUserController(c echo.Context) error {
	iddata, _ := strconv.Atoi(c.Param("id"))
	for _, value := range users {
		if value.Id == iddata {
			return c.JSON(http.StatusOK, map[string]interface{}{
				"message": "success get data",
				"users":   value,
			})
		}
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "data not found",
	})
}
func DeleteUserController(c echo.Context) error {
	// users[1] = users[len(users)-1]
	// users[len(users)-1] = User{}
	// users = users[:len(users)-1]
	iddata, _ := strconv.Atoi(c.Param("id"))
	for index, value := range users {
		if value.Id == iddata {
			users[index] = users[len(users)-1]
			users[len(users)-1] = User{}
			users = users[:len(users)-1]
			return c.JSON(http.StatusOK, map[string]interface{}{
				"message": "data successfully delete,",
				"users":   value,
			})
		}
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "data not found",
	})
}
func UpdateUserController(c echo.Context) error {
	iddata, _ := strconv.Atoi(c.Param("id"))
	user := User{}
	c.Bind(&user)
	fmt.Println("datanya : ", user)
	if user.Email == "" || user.Name == "" || user.Password == "" {
		return c.JSON(http.StatusNotAcceptable, map[string]interface{}{
			"message": "406 invalida data request",
		})
	}
	for index, value := range users {
		if value.Id == iddata {
			users[index] = user
			users[index].Id = iddata
			return c.JSON(http.StatusOK, map[string]interface{}{
				"message": "success get all users",
				"users":   users,
			})
		}
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "data not found",
	})
}
func CreateUserController(c echo.Context) error {
	user := User{}
	c.Bind(&user)

	if len(users) == 0 {
		user.Id = 1
	} else {
		newId := users[len(users)-1].Id + 1
		user.Id = newId
	}
	users = append(users, user)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create user",
		"user":    user,
	})
}

func main() {
	e := echo.New()
	e.GET("/users", GetUsersController)
	e.GET("/users/:id", GetUserController)
	e.PUT("/users/:id", UpdateUserController)
	e.DELETE("/users/:id", DeleteUserController)
	e.POST("/users", CreateUserController)

	e.Logger.Fatal(e.Start(":3008"))
}
