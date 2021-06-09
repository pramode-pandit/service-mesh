package main

import (
    "fmt"
    "time"
    "net/http"
//  "io/ioutil"
    "bytes"
//    "math/rand"
    "encoding/json"
//    "strconv"
)


func main() {

    url := "http://localhost:4000/updateMember?id=1234567890&member=3"
    fmt.Println("URL:>", url)

    var reg_json = []byte(`
    {
        "vaccinated":"partial",
        "dose1At":"Lucky Pharm",
        "dose1On":"",
        "dose1By":"Goppa",
        "dose2At":"",
        "dose2On":"",
        "dose2By":""
     }`) 

    //for {

         time.Sleep(time.Second)
         
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

         info["dose1On"] = time.Now().Format("2006-01-02 15:04:05")

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
         // body, _ := ioutil.ReadAll(resp.Body)
         //fmt.Println("response Body:", string(body))
    //}
}