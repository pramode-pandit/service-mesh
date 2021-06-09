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
	"io/ioutil"
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
	mux.HandleFunc("/updateMember", update_vaccine)
	mux.HandleFunc("/showMember", show_vaccine)
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

func show_vaccine(w http.ResponseWriter, r *http.Request) {

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

	member := r.URL.Query().Get("member")
	if member == "" {
		http.Error(w, http.StatusText(400), 400)
		return
	}

	if _, err := strconv.Atoi(member); err != nil {
		http.Error(w, http.StatusText(400), 400)
		return
	}

	p, err := FindProfile(id, member)
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


func update_vaccine(w http.ResponseWriter, r *http.Request) {

	// Unless the request is using the GET method, return a 405 'Method
	// Not Allowed' response.
	// log.Println("Request to get profile  " + r.URL.Query().Get("id"))
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		http.Error(w, http.StatusText(405), 405)
		return
	}

	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, http.StatusText(400), 400)
		return
	}
	member := r.URL.Query().Get("member")
	if member == "" {
		http.Error(w, http.StatusText(400), 400)
		return
	}

	// Validate that the id is a valid integer by trying to convert it,
	// returning a 400 Bad Request response if the conversion fails.
	if _, err := strconv.Atoi(id); err != nil {
		http.Error(w, http.StatusText(400), 400)
		return
	}

	if _, err := strconv.Atoi(member); err != nil {
		http.Error(w, http.StatusText(400), 400)
		return
	}


	reqBody, readErr := ioutil.ReadAll(r.Body)
	if readErr != nil {
	 log.Fatal(readErr)
	}

	var info map[string]interface{}

	// unmarschal JSON
    e := json.Unmarshal(reqBody, &info)

    // panic on error
    if e != nil {
        panic(e)
    }

	vaccinated   := fmt.Sprint(info["vaccinated"]) 
	dose1At       := fmt.Sprint(info["dose1At"]) 
	dose1On      := fmt.Sprint(info["dose1On"])
	dose1By        := fmt.Sprint(info["dose1By"])
	dose2At       := fmt.Sprint(info["dose2At"]) 
	dose2On      := fmt.Sprint(info["dose2On"])
	dose2By        := fmt.Sprint(info["dose2By"])

    // connection from the pool.
	conn := pool.Get()

	defer conn.Close()

	_, er := conn.Do("HMSET", "family:"+id+":"+member, "vaccinated", vaccinated, "dose1At", dose1At, "dose1On", dose1On, "dose1By", dose1By, "dose2At", dose2At, "dose2On", dose2On, "dose2By", dose2By )
    if er != nil {
        log.Println(er)
		http.Error(w, http.StatusText(500), 500)
		return
    }

	// Write the profile details as plain text to the client.
	 fmt.Fprintf(w, "%s | %s | %s | %s | %s \n", id, member, dose1At, dose1On, dose1By )
}