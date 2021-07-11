package main

import (
  "encoding/json"
  "fmt"
  "log"
  "net/http"
  "time"
)

// convertion struct (format)
type timeZoneConvertion struct {
  TimeZone string
  CurrentTime string
}

// convertion map values for calculation
var conversionMap = map[string] string {
  "ASR": "-3h",
  "EST": "-5h",
  "BST": "+1h",
  "IST": "+5h30m",
  "HKT": "+8h",
  "ART": "-3h",
  "DEL": "null", // forcing an error
}

// middleware definition
type handlerFunc func(w http.ResponseWriter, r *http.Request)

// main function
func main() {
  http.HandleFunc("/convert", loggingMiddleware(handler)) // middleware with custom handler
  http.HandleFunc("/", loggingMiddleware(notFoundHandler)) // root endpoint
  http.ListenAndServe("localhost:9000", nil)
}

// logging Middleware
func loggingMiddleware(handler handlerFunc) handlerFunc {
  fn := func(w http.ResponseWriter, r *http.Request) {
    // logging to console
    log.Printf("%s - %s - %s", time.Now().Format("2019-01-25 14:32:58"), r.Method, r.URL.String())
    // resolving the request
    handler(w, r)
  }
  return fn
}

// handler function
func handler(w http.ResponseWriter, r *http.Request) {
  // getting the queryString variable timezone
  timeZone := r.URL.Query().Get("timezone")
  // required parameter exception
  if timeZone == "" {
    w.WriteHeader(http.StatusBadRequest)
    fmt.Fprintf(w, "Error 400: timezone query parameter is required")
    return
  }
  // calculating difference
  timeDifference, ok := conversionMap[timeZone]
  // time convertion exception
  if !ok {
   w.WriteHeader(http.StatusNotFound)
   fmt.Fprintf(w, `Error 404: The timezone value "%s" does not correspond to an existing timezone value.`, timeZone)
   return
  }
  // converting difference
  currentTimeConverted, err := getCurrentTimeByTimeDifference(timeDifference)
  // time convertion exception
  if err != nil {
    w.WriteHeader(http.StatusInternalServerError)
    fmt.Fprintf(w, "Error: Server Error")
    return
  }
  // writing header
   w.WriteHeader(http.StatusOK)
  // convertion
  tzc := new(timeZoneConvertion)
  tzc.CurrentTime = currentTimeConverted
  tzc.TimeZone = timeZone
  // converting response to JSON
  jsonResponse, err := json.Marshal(tzc)
  if err != nil {
    w.WriteHeader(http.StatusInternalServerError)
    fmt.Fprintf(w, "Error: Server error.")
    return
  }
  // printing
  fmt.Fprintf(w, string(jsonResponse))
}

// 404
func notFoundHandler(w http.ResponseWriter, r *http.Request) {
  // forced 404
  w.WriteHeader(http.StatusNotFound)
  // output to console
  fmt.Fprintf(w, "Error 404: The request URL does not exist")
}

// comparting times
func getCurrentTimeByTimeDifference(timeDifference string) (string, error) {
  // converting actual time into UTC format
  now := time.Now().UTC()
  // calculating the difference
  difference, err := time.ParseDuration(timeDifference)
  // exception
  if err != nil {
    return "", err
  }
  // modifying in certain format
  now = now.Add(difference)
  return now.Format("15:04:05"), nil
}
