package main

import (
 "encoding/json"
 "fmt"
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
}

// main function
func main() {
 http.HandleFunc("/", handler) // root endpoint
 http.ListenAndServe("localhost:9000", nil)
}

// handler function
func handler(w http.ResponseWriter, r *http.Request) {
 // getting the queryString variable timezone
 timeZone := r.URL.Query().Get("timezone")
 // calculating difference
 timeDifference, _ := conversionMap[timeZone]
 // converting difference
 currentTimeConverted, _ := getCurrentTimeByTimeDifference(timeDifference)
 // convertion
 tzc := new(timeZoneConvertion)
 tzc.CurrentTime = currentTimeConverted
 tzc.TimeZone = timeZone
 // converting response to JSON
 jsonResponse, _ := json.Marshal(tzc)
 // writing header
 w.WriteHeader(http.StatusOK)
 // printing
 fmt.Fprintf(w, string(jsonResponse))
}

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
