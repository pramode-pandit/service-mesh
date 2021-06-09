package main

import (
    "fmt"
    "time"
    "net/http"
    "io/ioutil"
    "bytes"
    "math/rand"
    "encoding/json"
    "strconv"
)



func main() {

    url := "http://myapp.servicemesh.io/register"
    fmt.Println("URL:>", url)

    var reg_json = []byte(`
    {
        "name":"Mark Taylor",
        "mobile":"9876543210",
        "address":"123 kota road, Kolkata 700090",
        "adhar":"",
        "uuid":"",
        "date":""
    }`)

    for {

         time.Sleep(time.Second)
         
         mobile := 1000000000 + rand.Intn(9000000000)
         uuid := 1000000000 + rand.Intn(9000000000)
         
         // fmt.Println(mobile_number)

         // Unmarshal the JSON string into info map variable.
	     var info map[string]interface{}

         // unmarschal JSON
         e := json.Unmarshal(reg_json, &info)

         // panic on error
         if e != nil {
             panic(e)
         }

         // fmt.Println(info["name"])
         // mobile := fmt.Println(info["mobile"]) 

         info["mobile"] = strconv.Itoa(mobile)
         info["uuid"] = strconv.Itoa(uuid)
         info["date"] = time.Now().Format("2006-01-02 15:04:05")

         fmt.Println(info)

         reqBody, error := json.Marshal(info)
         if error != nil {
           panic(error)
           return
         }

         req, err := http.NewRequest("POST", url, bytes.NewBuffer(reqBody))
         
         req.Header.Set("X-Custom-Header", "myvalue")
         req.Header.Set("Content-Type", "application/json")
         
         client := &http.Client{}
         resp, err := client.Do(req)
         if err != nil {
             panic(err)
         }
         defer resp.Body.Close()
         

         fmt.Println("response Status:", resp.Status)
         fmt.Println("response Headers:", resp.Header)
         body, _ := ioutil.ReadAll(resp.Body)
         fmt.Println("response Body:", string(body))
    }
}