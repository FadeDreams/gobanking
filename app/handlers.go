package app

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"
)

type Customer struct {
	Name    string `json:"full_name" xml:"name"`
	City    string `json:"city" xml:"city"`
	Zipcode string `json:"zipcode" xml:"zipcode"`
}

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "hello word")
}
func getAllCustomers(w http.ResponseWriter, r *http.Request) {
	customers := []Customer{
		{Name: "name1", City: "city1", Zipcode: "1"},
		{Name: "name2", City: "city2", Zipcode: "2"},
		{Name: "name3", City: "city3", Zipcode: "3"},
		{Name: "name4", City: "city4", Zipcode: "4"},
	}

	if r.Header.Get("Content-Type") == "Application/xml" {
		w.Header().Add("Content-Type", "application/xml")
		xml.NewEncoder(w).Encode(customers) //encoding customers to json format
	} else {
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(customers) //encoding customers to json format
	}
	//json.NewEncoder(w).Encode(customers) //encoding customers to json format

}
