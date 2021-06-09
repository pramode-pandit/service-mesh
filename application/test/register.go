// https://www.golangprograms.com/example-to-handle-get-and-post-request-in-golang.html
// https://stackoverflow.com/questions/46515876/how-to-access-key-and-value-from-json-in-golang

   package main
   
   import (
   "fmt"
   "log"
   "net/http"
   "io/ioutil"
   "os"
   "encoding/json"
   )

   func register(w http.ResponseWriter, r *http.Request) {
    
       log.Printf(r.Method+":  "+r.URL.Path)
       //log.Printf(string(r.URL.RawQuery))
       
       switch r.Method {

	        case "GET":		
	        	 fmt.Fprintf(w, "POST methods are supported.")

	        case "POST":

	        	reqBody, readErr := ioutil.ReadAll(r.Body)
	            if readErr != nil {
		            log.Fatal(readErr)
	            }

                log.Printf(string(reqBody))

                // Unmarshal the JSON string into info map variable.
	            var info map[string]interface{}

                // unmarschal JSON
                e := json.Unmarshal(reqBody, &info)

                // panic on error
                if e != nil {
                    panic(e)
                }

                // fmt.Println(info["name"])
                mobile := fmt.Sprint(info["mobile"]) 

                ioutil.WriteFile("mobile+".json", reqBody, os.ModePerm)
                fmt.Fprintf(w, "registered !!!")

	        default:
	        	fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
	   }

       
   }

   
   func main() {
       
       http.HandleFunc("/register", register)
   
       http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    
        if r.URL.Path != "/" {
        http.NotFound(w, r)
        return
        }
    
        w.Write([]byte("<h1>Welcome to my web server!</h1>"))
        })


       
       if err := http.ListenAndServe(":8080", nil); err != nil {
           log.Fatal(err)
       }
   }
