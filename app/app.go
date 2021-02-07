package app

import (
	"fmt"
	"github.com/adenapila/market-go/domain"
	"github.com/adenapila/market-go/service"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

func Start() {
	//mux := http.NewServeMux()
	router := mux.NewRouter()

	//wiring
	//ch := CustomerHandlers{service: service.NewCustomerService(domain.NewCustomerRepositoryStub())}
	//accountRepositoryDb := domain.NewAccountRepositoryDb()

	ch := CustomerHandlers{service.NewCustomerService(domain.NewCustomerRepositoryDb())}
	ah := AccountHandler{service.NewAccountService(domain.NewAccountRepositoryDb())}

	//define routes - matcher
	router.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customer_id:[0-9]+}", ch.getCustomer).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customer_id:[0-9]+}/account", ah.NewAccount).Methods(http.MethodPost)

	//router.HandleFunc("/greet", greet).Methods(http.MethodGet)
	//router.HandleFunc("/customers", createCustomers).Methods(http.MethodPost)
	//router.HandleFunc("/customers/{customer_id:[0-9]+}", getCustomer).Methods(http.MethodGet)

	//starting server
	address := os.Getenv("SERVER_ADDRESS")
	port := os.Getenv("SERVER_PORT")
	//log.Fatal(http.ListenAndServe("localhost:8888", router))
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", address, port), router))
}

//func getCustomer(w http.ResponseWriter, r *http.Request) {
//	vars := mux.Vars(r)
//	fmt.Fprint(w, vars["customer_id"])
//}
//
//func createCustomers(w http.ResponseWriter, r *http.Request) {
//	fmt.Fprint(w, "Post request received!")
//}
