package main

// import (
// 	"net/http"
// 	"strconv"

// 	"github.com/labstack/echo"
// )

// type User struct {
// 	Id    int
// 	Name  string
// 	Email string
// }

// func GetUser(c echo.Context) error {
// 	//
// 	id, _ := strconv.Atoi(c.Param("id"))
// 	user := User{Id: id, Name: "Ismail", Email: "ismail@alterra.id"}
// 	return c.JSON(http.StatusOK, map[string]interface{}{
// 		"user": user,
// 	})
// }

// func main() {
// 	e := echo.New()
// 	e.GET("/user/:id", GetUser)
// 	e.Start(":3004")
// }

// // func HelloController(c echo.Context) error {
// // 	return c.String(http.StatusOK, "Hello World")
// // }
