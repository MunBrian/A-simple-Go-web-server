package main 


import (
	"fmt"
	"net/http"
	"log"
)


//helloHandler func
func helloHandler(w http.ResponseWriter, r *http.Request){
	//check if path is correct
	if r.URL.Path != "/hello"{
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}

	//check if method is correct
	if r.Method != "GET"{
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}

	//send message to browser
	fmt.Fprintf(w, "Hello")
}


func main(){

	http.HandleFunc("/hello", helloHandler)


	fmt.Printf("starting server at port 8080\n")

	//listenAnsServe - allow us to start webserver and specify the port to listen for incoming request
	if err := http.ListenAndServe(":8080", nil); err != nil{
		log.Fatal(err)
	}
}