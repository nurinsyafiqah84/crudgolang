// main.go
package main

import (
	"fmt"
	productApi "goWeb/apis/product"
	"goWeb/config"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	config.GetDb()

	router := mux.NewRouter()

	router.HandleFunc("/", productApi.FindAll)
	router.HandleFunc("/add", productApi.Add)
	router.HandleFunc("/processAdd", productApi.ProcessAdd)
	router.HandleFunc("/processSearch", productApi.ProcessSearch)
	router.HandleFunc("/delete", productApi.Delete)
	router.HandleFunc("/edit", productApi.Edit)
	router.HandleFunc("/update", productApi.Update)

	err := http.ListenAndServe(":5000", router)
	if err != nil {
		fmt.Println(err)
	}
}
