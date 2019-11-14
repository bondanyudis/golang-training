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
// 	user := User{}
// 	c.Bind(&user)

// 	return c.JSON(http.StatusOK, map[string]interface{}{
// 		"message": "Success create user",
// 		"user":    user,
// 	})
// }

// func main() {
// 	e := echo.New()
// 	e.GET("/user", GetUser)
// 	e.Start(":3004")
// }
