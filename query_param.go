package main

// import (
// 	"net/http"

// 	"github.com/labstack/echo"
// )

// type User struct {
// 	Id    int
// 	Name  string
// 	Email string
// }

// func GetUser(c echo.Context) error {
// 	match := c.QueryParam("match")
// 	return c.JSON(http.StatusOK, map[string]interface{}{
// 		"match":  match,
// 		"result": []string{"adi", "aan", "yudis"},
// 	})
// }

// func main() {
// 	e := echo.New()
// 	e.GET("/user", GetUser)
// 	e.Start(":3004")
// }

// // func HelloController(c echo.Context) error {
// // 	return c.String(http.StatusOK, "Hello World")
// // }
