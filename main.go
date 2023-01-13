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


//formHandler func
func formHandler(w http.ResponseWriter, r *http.Request){
	//check for error
	//ParseForm pass the raw query and update value to r.PostForm and r.Form
	if err := r.ParseForm(); err != nil{
		fmt.Fprintf(w, "ParseForm() err: %v", err)
	}

	//if no error
	fmt.Fprintf(w, "Post request successful")


	//get name and address value using Form value
	name := r.FormValue("name")
	address := r.FormValue("address")

	fmt.Fprintf(w, "Name =%s\n", name)
	fmt.Fprintf(w, "Address = %s\n", address)
}


func main(){

	fileServer := http.FileServer(http.Dir("./static"))
	//Handle - accept path and fileserver
	//display index.html file
	http.Handle("/", fileServer)

	http.HandleFunc("/hello", helloHandler)

	http.HandleFunc("/form", f)


	fmt.Printf("starting server at port 8080\n")

	//listenAnsServe - allow us to start webserver and specify the port to listen for incoming request
	if err := http.ListenAndServe(":8080", nil); err != nil{
		log.Fatal(err)
	}
}