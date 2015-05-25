package main

import (
	"encoding/json"
	"github.com/raitucarp/bing-search"
	"github.com/raitucarp/websearch-server/log"
	"net/http"
	"os"
	"strconv"
	"time"
)

// Items for output
type Item struct {
	search.Item
	fileSize     int `json:"fileSize"`
	lastModified int `json:"lastModified"`
}

type Output struct {
	Success bool          `json:"success"`
	Items   []search.Item `json:"items"`
	Count   int           `json:"count"`
}

type Err struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
}

// Main Handler
func MainHandler(res http.ResponseWriter, req *http.Request) {
	params := req.URL.Query()
	q := params.Get("q")
	keyword := params.Get("keyword")
	advancedQuery := params.Get("advancedQuery")
	countQuery := params.Get("count")
	torQuery := params.Get("tor")
	startTime := time.Now()

	// either keyword or q
	if keyword != "" || q != "" {
		var (
			count int
			tor   bool
		)
		if countQuery == "" {
			countQuery = "10"
		}

		count, _ = strconv.Atoi(countQuery)
		tor, _ = strconv.ParseBool(torQuery)

		if keyword != "" {
			q = keyword + " " + advancedQuery
		}

		options := search.Options{
			Query: q,
			Count: count,
			Tor:   tor,
		}

		result, _ := search.WebSearch(options)

		if params.Get("header") == "true" {
			result.GetHeaders()
		}

		output := Output{
			Success: true,
			Items:   result,
			Count:   len(result),
		}

		if len(result) < 1 {
			output.Success = false
			output.Items = []search.Item{}
		}

		data, err := json.Marshal(output)
		if err != nil {
			log.Danger(req.Method, err.Error(), startTime)
			http.Error(res, err.Error(), http.StatusInternalServerError)
			return
		}

		var message string
		if keyword != "" {
			message = keyword
		} else {
			message = q
		}

		if len(result) < 1 {
			log.Warning(req.Method, message, startTime)
		} else {
			log.Success(req.Method, message, startTime)
		}
		res.Header().Set("Content-Type", "application/json; charset=utf-8")
		res.Write(data)
	} else {
		output := Output{}
		res.Header().Set("Content-Type", "application/json")
		data, err := json.Marshal(output)
		if err != nil {
			log.Danger(req.Method, err.Error(), startTime)
			http.Error(res, err.Error(), http.StatusInternalServerError)
			return
		}
		log.Warning(req.Method, "!", startTime)
		res.Write(data)
	}
}

func main() {
	// listening to PORT in environment
	port := os.Getenv("PORT")

	// add / pattern
	http.HandleFunc("/", MainHandler)
	http.NotFoundHandler()

	// listen to port
	log.Standard("WebSearch Listening on port: " + port)
	err := http.ListenAndServe(":"+port, nil)
	// if error
	if err != nil {
		log.Danger("Fatal", "ListenAndServe: "+err.Error(), time.Now())
	}
}
