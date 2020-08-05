package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func pingHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("ping endpoint has been hit")
}

func pongHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("pong endpoint has been hit")
}

func itemsHandler(w http.ResponseWriter, r *http.Request) {
	switch httpMethod := r.Method; httpMethod {
	case "POST":
		item := r.FormValue("item")
		f, err := os.OpenFile("items.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		_, err = f.Write([]byte(fmt.Sprintf("%s\n", item)))
		if err != nil {
			log.Println(err)
		}
		err = f.Close()
		if err != nil {
			log.Println(err)
		}
		log.Printf("the following item has been logged in items.txt: %s\n", item)
	case "GET":
		f, err := ioutil.ReadFile("items.txt")
		if err != nil {
			log.Println(err)
		}
		_, err = w.Write(f)
		if err != nil {
			log.Println(err)
		}
	default:
		fmt.Fprintf(w, "%s is not a valid method for this endpoint\n", r.Method)
	}
}

func main() {
	http.HandleFunc("/ping", pingHandler)
	http.HandleFunc("/pong", pongHandler)
	http.HandleFunc("/items", itemsHandler)
	http.ListenAndServe(":8080", nil)
}
