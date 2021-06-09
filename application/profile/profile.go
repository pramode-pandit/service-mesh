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
	Dose1At    string `redis:"dose1_at"`
	Dose1On    string `redis:"dose1_on"`
	Dose1By    string `redis:"dose1_by"`
	Dose2At    string `redis:"dose2_at"`
	Dose2On    string `redis:"dose2_on"`
	Dose2By    string `redis:"dose2_by"`	
}

type AllProfile struct{
	Member1 Profile
	Member2 Profile
	Member3 Profile
	Member4 Profile
}


func FindProfile(id string) (*AllProfile, error) {

	conn := pool.Get()

	defer conn.Close()
	

	var p1, p2, p3, p4 Profile
	var ap AllProfile


	log.Println("loading profile id="+id+"&member=1")

	values, err := redis.Values(conn.Do("HGETALL", "family:"+id+":1"))

	// check error
	if err != nil {
	    log.Println("Error while execution")
		log.Println(err) 
	} else if len(values) == 0 {
		// return nil, ErrNoProfile
		return &ap, nil
	}
	
	err = redis.ScanStruct(values, &p1)
	if err != nil {
	    log.Println("Error while scanning struct member 1")
		return nil, err
	}
	ap.Member1 = p1


	log.Println("loading profile id="+id+"&member=2")
	values, err = redis.Values(conn.Do("HGETALL", "family:"+id+":2"))

	// check error
	if err != nil {
	    log.Println("Error while execution")
		log.Println(err) 
	} else if len(values) == 0 {
		return &ap, nil
	}
	
	err = redis.ScanStruct(values, &p2)
	if err != nil {
	    log.Println("Error while scanning struct member 2")
		return nil, err
	}
	ap.Member2 = p2



	log.Println("loading profile id="+id+"&member=3")
	values, err = redis.Values(conn.Do("HGETALL", "family:"+id+":3"))

	// check error
	if err != nil {
	    log.Println("Error while execution")
		log.Println(err) 
	} else if len(values) == 0 {
		return &ap, nil
	}
	
	err = redis.ScanStruct(values, &p3)
	if err != nil {
	    log.Println("Error while scanning struct member 3")
		return nil, err
	}
	ap.Member3 = p3


	log.Println("loading profile id="+id+"&member=4")
	values, err = redis.Values(conn.Do("HGETALL", "family:"+id+":4"))

	// check error
	if err != nil {
	    log.Println("Error while execution")
		log.Println(err) 
	} else if len(values) == 0 {
		return &ap, nil
	}
	
	err = redis.ScanStruct(values, &p4)
	if err != nil {
	    log.Println("Error while scanning struct member 4")
		return nil, err
	}
	ap.Member4 = p4

	return &ap,nil
}