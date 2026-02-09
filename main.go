package main

import (
	"blogging/database"
	"blogging/network"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

// As this is a very simple project, I won't extend myself in a complex, difficult setup. I will just arrange the code in a way
// each responsibility doesn't step on the other.

func main() {
	err := godotenv.Load()
	if err != nil {
		panic("unable to load environment variables; impossible to proceed")
	}
	hsa := os.Getenv("HTTP_SERVER_ADDRESS")
	
	db := database.NewInMemoryDatabase()
	
	mux := network.NewHTTPServer(db)

	// TODO: implement Golang concurrent patterns to optimize resources usage
	err = http.ListenAndServe(hsa, mux)
	if err != nil {
		
	}

}