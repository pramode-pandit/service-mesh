package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
	"os"
	"github.com/gomodule/redigo/redis"
	"encoding/json"
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

	if _, err := strconv.Atoi(id); err != nil {
		http.Error(w, http.StatusText(400), 400)
		return
	}


	p, err := FindProfile(id)
	if err == ErrNoProfile {
		http.NotFound(w, r)
		return
	} else if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}


	var jsonData []byte
    jsonData, error := json.MarshalIndent(p,"", "    ")
    if error != nil {
        log.Println(error)
    }
    //fmt.Println(string(jsonData))


	// Write the profile details as plain text to the client.
	// fmt.Fprintf(w, "%s | %s | %s \n%s | %s | %s \n", p.Member1.Name, p.Member1.Aadhar, p.Member1.DOB, p.Member2.Name, p.Member2.Aadhar, p.Member2.DOB)
	fmt.Fprintf(w, string(jsonData))
}