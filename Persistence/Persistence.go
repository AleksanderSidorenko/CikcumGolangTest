package main


import "C"
import (
	"io"
	"fmt"
	"net/http"
	"encoding/json"
	"labix.org/v2/mgo"
)

type Person struct {
    Id int64
	Name  string
	Email string
	MobileNumber string
}


func persist(res http.ResponseWriter, req *http.Request) {
	var data Person
	json.Unmarshal(req, &data)

	mongoSess, err := mgo.Dial("localhost:27017")
	if err != nil {
		panic(err)
	}
	defer mongoSess.Close()

	c := mongoSess.DB("test").C("roster")
	c.Insert(data)
}

func main() {
	http.HandleFunc("/persist", persist)
	http.ListenAndServe(":9000", nil)
}
