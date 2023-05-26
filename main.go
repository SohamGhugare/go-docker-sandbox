package main

import (
	"fmt"
	"log"
	"net/http"
)

func main(){
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		fmt.Fprintln(w, "root")
	})

	fmt.Println("Server listening at localhost:8081...")
	log.Fatal(http.ListenAndServe(":8081", nil))
}