package main

// import (
// 	"net/http"

// 	"github.com/labstack/echo"
// )

// type User struct {
// 	Name  string `json:"name" form:"name"`
// 	Email string `json:"email" form:"email"`
// }

// func GetUser(c echo.Context) error {
// 	name := c.FormValue("name")
// 	email := c.FormValue("email")

// 	var user User

// 	user.Name = name
// 	user.Email = email

// 	return c.JSON(http.StatusOK, map[string]interface{}{
// 		"message": "Successcreate User",
// 		"user":    user,
// 	})
// }

// func main() {
// 	e := echo.New()
// 	e.GET("/user", GetUser)
// 	e.Start(":3004")
// }
