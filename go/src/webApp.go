package main

import ("fmt"
		  "net/http"
		)

		func handleFunc(w http.ResponseWriter, r *http.Request)  {
			fmt.Fprint(w,"<h1> welcome to my sit</h1>")
		}

		func main(){
			http.HandleFunc("/ram", handleFunc)

			http.ListenAndServe(":3000" , nil)
		}