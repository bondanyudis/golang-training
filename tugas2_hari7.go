package main

// import (
// 	"encoding/json"
// 	"fmt"
// 	"net/http"
// )

// type manusia struct {
// 	NamaDepan    string
// 	NamaBelakang string
// 	Umur         int
// 	Tinggi       int
// }

// var data = []manusia{
// 	manusia{"yudistira", "sugandi", 21, 178},
// 	manusia{"bondan", "yudis", 25, 182},
// 	manusia{"tira", "bondan", 25, 182},
// }

// func hooman(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	if r.Method == "GET" {
// 		var result, err = json.Marshal(data)
// 		if err != nil {
// 			http.Error(w, err.Error(), http.StatusInternalServerError)
// 			return
// 		}
// 		w.Write(result)
// 		return
// 	}
// 	http.Error(w, "", http.StatusBadRequest)
// }

// func main() {
// 	http.HandleFunc("/hooman", hooman)
// 	fmt.Println("starting web server at http://localhost:8000")
// 	http.ListenAndServe(":3002", nil)

// }
