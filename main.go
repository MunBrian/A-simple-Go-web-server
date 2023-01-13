package main 


import (
	"fmt"
	"net/http"
	"log"
)


func main(){

	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request){
		//Fprintf -print out formatted string to the w destination
		fmt.Fprintf(w, "hello!!!")
	})


	fmt.Printf("starting server at port 8080\n")

	//listenAnsServe - allow us to start webserver and specify the port to listen for incoming request
	if err := http.ListenAndServe(":8080", nil); err != nil{
		log.Fatal(err)
	}
}