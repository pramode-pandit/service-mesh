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
	Name       string  `redis:"name"`
	Aadhar     string  `redis:"adhar"`
	DOB        string  `redis:"dob"`
	Vaccinated string  `redis:"vaccinated"`
	Reg_dt     string  `redis:"reg_dt"`
	CowinId    string `redis:"cowinid"`
	Dose1At    string `redis:"dose1At"`
	Dose1On    string `redis:"dose1On"`
	Dose1By    string `redis:"dose1By"`
	Dose2At    string `redis:"dose2At"`
	Dose2On    string `redis:"dose2On"`
	Dose2By    string `redis:"dose2By"`	
}


func FindProfile(id string, member string) (*Profile, error) {

	conn := pool.Get()

	defer conn.Close()
	
	log.Println("loading profile id="+id+"&member="+member)

	values, err := redis.Values(conn.Do("HGETALL", "family:"+id+":"+member))

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

	return &profile,nil
}