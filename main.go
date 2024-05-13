package main

import "log"


func main(){
	// http.HandleFunc("GET /", withSomething(homePage) )

	// http.HandleFunc("POST /", func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Fprintf(w, "POST /")
	// })
	store, err := NewPostgresStore()
	if err != nil {
		log.Fatal(err)
	}
	if err := store.Init(); err != nil {
		log.Fatal(err)
		return 
	}

	log.Printf("%+v", store)

	server := NewApiServer(":3000", store)
	server.Run()
}