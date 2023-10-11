package main


import (
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
)


func main(){
	r := mux.NewRouter()
	r.HandleFunc("/",handle).Methods("GET")
	http.Handle("/",r)
	http.ListenAndServe(":8080",nil)
}

func handle(w http.ResponseWriter , r *http.Request){
	fmt.Fprint(w,"Hello World!")
}