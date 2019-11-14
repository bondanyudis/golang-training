package main

// import (
// 	"encoding/json"
// 	"fmt"
// 	"io/ioutil"
// 	"net/http"
// )

// type StarWarsPeople struct {
// 	Name      string   `json:"name"`
// 	Height    string   `json:"height"`
// 	Mass      string   `json:"mass"`
// 	HairColor string   `json:"hair_color"`
// 	SkinColor string   `json:"skin_color"`
// 	EyeColor  string   `json:"eye_color"`
// 	BirthYear string   `json:"birth_year"`
// 	Gender    string   `json:"gender"`
// 	Films     []string `json:"films"`
// 	Homeworld string   `json:"homeworld"`
// }

// func main() {
// 	response, _ := http.Get("https://swapi.co/api/people/1")
// 	responseData, _ := ioutil.ReadAll(response.Body)
// 	defer response.Body.Close()

// 	var LukeSkywalker StarWarsPeople
// 	json.Unmarshal(responseData, &LukeSkywalker)
// 	fmt.Println(responseData)
// 	fmt.Println(LukeSkywalker)

// 	fmt.Println("------print field-------")
// 	fmt.Println(LukeSkywalker)
// 	fmt.Println(LukeSkywalker.Name)
// 	fmt.Println(LukeSkywalker.Height)
// 	fmt.Println(LukeSkywalker.Mass)
// 	fmt.Println(LukeSkywalker.HairColor)
// 	fmt.Println(LukeSkywalker.Homeworld)
// 	fmt.Println(LukeSkywalker.Films[0])
// }
