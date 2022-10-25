package app

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Start() {

	// mux := http.NewServeMux()
	router := mux.NewRouter()
	//route
	router.HandleFunc("/greet", greet).Methods(http.MethodGet)
	router.HandleFunc("/customers", getAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers", createCustomer).Methods(http.MethodPost)
	// router.HandleFunc("/customers/{customer_id}", getCustomer)
	router.HandleFunc("/customers/{customer_id:[0-9]+}", getCustomer).Methods(http.MethodGet)


	//starting server
	log.Fatal(http.ListenAndServe("localhost:8000", router))

}

func getCustomer(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	fmt.Fprint(w, vars["customer_id"])

}
func createCustomer(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w,"post req rec")

}


// func Start_2() {
// 	//http.HandleFunc("/greet", func(w http.ResponseWriter, r *http.Request) {
// 	//	fmt.Fprint(w, "hello word")
// 	//})

// 	mux := http.NewServeMux()
// 	//route
// 	mux.HandleFunc("/greet", greet)
// 	mux.HandleFunc("/customers", getAllCustomers)

// 	//starting server
// 	log.Fatal(http.ListenAndServe("localhost:8000", mux))

// }

//func Start_old1() {
//	//http.HandleFunc("/greet", func(w http.ResponseWriter, r *http.Request) {
//	//	fmt.Fprint(w, "hello word")
//	//})
//
//	//route
//	http.HandleFunc("/greet", greet)
//	http.HandleFunc("/customers", getAllCustomers)
//
//	//starting server
//	log.Fatal(http.ListenAndServe("localhost:8000", nil))
//
//}
