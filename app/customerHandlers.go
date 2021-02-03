package app

import (
	"encoding/json"
	"github.com/adenapila/market-go/service"
	"github.com/gorilla/mux"
	"net/http"
)

//type Customer struct {
//	Name string `json:"full_name" xml:"name"`
//	City string `json:"city" xml:"city"`
//	Zip  string `json:"zip_code" xml:"zip"`
//}

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
	status := r.URL.Query().Get("status")

	customers, err := ch.service.GetAllCustomer(status)

	if err != nil {
		writeResponse(w, err.Code, err.AsMessage())
	} else {
		writeResponse(w, http.StatusOK, customers)

	}
}

func (ch *CustomerHandlers) getCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["customer_id"]

	customer, err := ch.service.GetCustomer(id)
	if err != nil {
		writeResponse(w, err.Code, err.AsMessage())
	} else {
		writeResponse(w, http.StatusOK, customer)

	}
}

func writeResponse(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		panic(err)
	}
}
