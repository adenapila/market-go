package app

import (
	"encoding/json"
	"encoding/xml"
	"github.com/adenapila/market-go/service"
	"net/http"
)

type Customer struct {
	Name string `json:"full_name" xml:"name"`
	City string `json:"city" xml:"city"`
	Zip  string `json:"zip_code" xml:"zip"`
}

//func greet(w http.ResponseWriter, r *http.Request) {
//	fmt.Fprint(w, "Hello world!")
//}

type CustomerHandlers struct {
	service service.CustomerService
}

func (ch *CustomerHandlers) getAllCustomers(w http.ResponseWriter, r *http.Request) {
	//customers := []Customer {
	//	{Name: "Ade", City: "Tangerang", Zip: "15418"},
	//	{Name: "Napila", City: "Jakarta", Zip: "15215"},
	//}

	customers, _ := ch.service.GetAllCustomer()

	if r.Header.Get("Content-Type") == "application/xml" {
		//merubah plain/text menjadi application/xml
		w.Header().Add("Content-type", "application/xml")
		xml.NewEncoder(w).Encode(customers)
	} else {
		//merubah plain/text menjadi application/json
		w.Header().Add("Content-type", "application/json")
		json.NewEncoder(w).Encode(customers)
	}
}
