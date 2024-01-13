package main

import (
	"fmt"
	"log"
	"net/http"
)

func main(){
	port := 8084

	fmt.Printf("Service is running on port :%d ...\n", port)

	http.HandleFunc("/api/startup/test", helloHandler)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))

}

func helloHandler(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w, "Service is running")
}
