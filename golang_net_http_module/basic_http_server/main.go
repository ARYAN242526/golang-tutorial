package main

import (
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter , r *http.Request){
	if r.Method != http.MethodGet {
		http.Error(w , "only Get is allowed" , http.StatusMethodNotAllowed);
		return
	}

	_ , _ = w.Write([]byte("Hello from GO net/http server"))
}



func main(){
	http.HandleFunc("/hello" , helloHandler)

	fmt.Println("try going to 4000 port")

	err := http.ListenAndServe(":4000" , nil)

	fmt.Println(err)
}