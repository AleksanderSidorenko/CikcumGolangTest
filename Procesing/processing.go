package main

import (
    "bufio"
    "encoding/csv"
	"strconv"
    "io"
	"log"
	"encoding/json"
	"labix.org/v2/mgo"
    "os"
)

type Person struct {
    Id int64
	Name  string
	Email string
	MobileNumber string
}

func main() {
	csvFile, _ := os.Open("data.csv")
	reader := csv.NewReader(bufio.NewReader(csvFile))
	reader.Read()
    for {
        line, error := reader.Read()
        if error == io.EOF {
            break
        } else if error != nil {
            log.Fatal(error)
		}

		id, err := strconv.ParseInt(string(line[0]), 10, 32)
		if err != nil {
			panic(err)
		}
		result := int64(id)	
			
		person := Person{
			Id: result, 
			Name: string(line[1]), 
			Email: string(line[2]), 
			MobileNumber: "+44" + string(line[3])}

		json, err := json.Marshal(person)
		if err != nil {
			panic (err)
		}

		url := "localhost:9000/persist"
	
		req, err := http.NewRequest("POST", url, bytes.NewBuffer(json))
		req.Header.Set("Content-Type", "application/json")
	
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()

    }
}