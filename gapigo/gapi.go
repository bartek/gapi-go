package gapi

import (
	"net/http"
	"log"
	"io/ioutil"
	"encoding/json"
	"time"
	"strings"
)

var (
    ApiRoot string = "https://sandbox.gadventures.com"
    ApiProxy string = ""
    ApiKey string = ""
)


// Describe the request made to the API, as well as the result
type ApiRequest struct {
   Url string
   Method string
   Request *http.Request

   // Result containers
   Result interface{} // Unmarshalled response into generic interface
   Error interface{} // Unmarshalled error response
   RawText string // Raw text from the server (after byte conversion)
   Status int // HTTP Status code

   // Meta
   Timestamp time.Time
}

// Initializes the Request, but doesnt process it yet. As it can be added onto
func (r *ApiRequest) buildRequest() {
     reqUri :=  strings.Join([]string{ApiRoot, r.Url}, "")
     req, _ :=  http.NewRequest(r.Method, reqUri, nil);

     // TODO Need to allow custom headers sent through.
     req.Header.Set("X-Application-Key", ApiKey);

     r.Request = req
     r.Url = reqUri
     r.Timestamp = time.Now()

     log.Printf("Planning Query to (%v) %v", r.Method, r.Url)
}

// Given a built ApiRequest object, make the call out to the API
func (r *ApiRequest) Fetch() (status int, err error) {
    if r.Method == "" {
        r.Method = "GET"
    }

    r.buildRequest()

    client := &http.Client{}
    resp, err := client.Do(r.Request);

    if err != nil {
       log.Fatal(err)
    }

    defer resp.Body.Close();

    bytes, _ := ioutil.ReadAll(resp.Body)
    r.RawText = string(bytes)

    if resp.StatusCode == 200 {
       error := json.Unmarshal(bytes, &r.Result)
           if error != nil {
        log.Fatal(error)
       }
    } else {
        log.Fatal("Bad HTTP Status", resp.StatusCode)
    }
    return resp.StatusCode, err
}
