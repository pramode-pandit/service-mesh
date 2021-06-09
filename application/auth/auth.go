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


var pool *redis.Pool

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
	mux.HandleFunc("/auth", auth)
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

func auth(w http.ResponseWriter, r *http.Request) {

	// Unless the request is using the GET method, return a 405 'Method
	// Not Allowed' response.
	// log.Println("Request to get profile  " + r.URL.Query().Get("id"))
	if r.Method != http.MethodGet {
		w.Header().Set("Allow", http.MethodGet)
		http.Error(w, http.StatusText(405), 405)
		return
	}

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

    // connection from the pool.
	conn := pool.Get()

	defer conn.Close()

	_, er := conn.Do("SADD", "family", id)
    if er != nil {
        log.Println(er)
		http.Error(w, http.StatusText(500), 500)
		return
    }

	// Write the profile details as plain text to the client.
	 fmt.Fprintf(w, "%s authenticated \n", id)
}