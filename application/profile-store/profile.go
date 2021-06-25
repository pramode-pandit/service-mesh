package main

import (
	"errors"
	"log"
	"github.com/gomodule/redigo/redis"
)


// Declare a pool variable to hold the pool of Redis connections.
var pool *redis.Pool

var ErrNoProfile = errors.New("no album found")

// Define a custom struct to hold Album data.
type Profile struct {
	Name    string  `redis:"Name"`
	City    string  `redis:"City"`
	Aadhar  string  `redis:"Aadhar"`
}


func FindProfile(id string) (*Profile, error) {
	// Use the connection pool's Get() method to fetch a single Redis
	// connection from the pool.
	conn := pool.Get()

	// Importantly, use defer and the connection's Close() method to
	// ensure that the connection is always returned to the pool before
	// FindProfile() exits.
	defer conn.Close()

    
	log.Println("Authenticating...")
	_, er := conn.Do("AUTH", "default", "foobared")
    if er != nil {
        log.Println(er)
    }
	
	// Fetch the details of a specific album. If no album is found
	// the given id, the []interface{} slice returned by redis.Values
	// will have a length of zero. So check for this and return an
	// ErrNoProfile error as necessary.

	log.Println("loading profile from redis : " + id )
	values, err := redis.Values(conn.Do("HGETALL", "profile:"+id))

	// check error
	if err != nil {
	    log.Println("Error while execution")
		log.Println(err) 
	} else if len(values) == 0 {
		return nil, ErrNoProfile
	}
	
	var profile Profile
	err = redis.ScanStruct(values, &profile)
	if err != nil {
		return nil, err
	}

	return &profile, nil
}