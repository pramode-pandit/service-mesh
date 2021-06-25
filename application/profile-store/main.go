package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
	"os"
	"github.com/gomodule/redigo/redis"
)

var redis_host = "localhost"  // = os.Getenv("REDIS_HOST")
var redis_port = "6379" // os.Getenv("REDIS_PORT") 
var redis_password = "foobared" // os.Getenv("REDIS_PASSWORD")


func main() {

    loadEnv()

	// Initialize a connection pool and assign it to the pool global variable.
	pool = &redis.Pool{
		MaxIdle:     10,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", redis_host+":"+redis_port)
		},
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/profile", showProfile)
	log.Println("Listening on :4000...")
	http.ListenAndServe(":4000", mux)
}

func loadEnv()  {

    if rho := os.Getenv("REDIS_HOST"); rho != "" {
		redis_host = rho
	}
	if rpo := os.Getenv("REDIS_PORT"); rpo != "" {
		redis_port = rpo
	}
	if rpa := os.Getenv("REDIS_PASSWORD"); rpa != "" {
		redis_password = rpa
	}
}

func showProfile(w http.ResponseWriter, r *http.Request) {

	// Unless the request is using the GET method, return a 405 'Method
	// Not Allowed' response.
	// log.Println("Request to get profile  " + r.URL.Query().Get("id"))
	if r.Method != http.MethodGet {
		w.Header().Set("Allow", http.MethodGet)
		http.Error(w, http.StatusText(405), 405)
		return
	}

	// Retrieve the id from the request URL query string. If there is
	// no id key in the query string then Get() will return an empty
	// string. We check for this, returning a 400 Bad Request response
	// if it's missing.
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, http.StatusText(400), 400)
		return
	}

	// Validate that the id is a valid integer by trying to convert it,
	// returning a 400 Bad Request response if the conversion fails.
	if _, err := strconv.Atoi(id); err != nil {
		http.Error(w, http.StatusText(400), 400)
		return
	}

	// log.Println("Retreiving profile  " + id )
	// Call the FindProfile() function passing in the user-provided id.
	// If there's no matching profile found, return a 404 Not Found
	// response. In the event of any other errors, return a 500
	// Internal Server Error response.
	bk, err := FindProfile(id)
	if err == ErrNoProfile {
		http.NotFound(w, r)
		return
	} else if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}

	// Write the profile details as plain text to the client.
	 fmt.Fprintf(w, "%s | %s | %s \n", bk.Name, bk.City, bk.Aadhar)
}