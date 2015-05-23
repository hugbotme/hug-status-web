package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/garyburd/redigo/redis"
)

type Info struct {
	InProgress int `json:"in_progress"`
	Merged     int `json:"merged"`
	Closed     int `json:"closed"`
	Received   int `json:"received"`
}

func main() {

	red, err := redis.Dial("tcp", ":6379")
	if err != nil {
		fmt.Println("error", err)
		return
	}
	defer red.Close()

	http.HandleFunc("/info.json", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Content-Security-Policy", "*")

		in_progress, _ := redis.Int(red.Do("LLEN", "hug:pullrequests"))
		merged, _ := redis.Int(red.Do("LLEN", "hug:pullrequests:merged"))
		closed, _ := redis.Int(red.Do("LLEN", "hug:pullrequests:closed"))
		received, _ := redis.Int(red.Do("LLEN", "hug:queue"))

		pr := Info{
			in_progress,
			merged,
			closed,
			received,
		}
		data, err := json.Marshal(pr)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write(data)
	})

	fmt.Println("Running on http://localhost:8080")
	log.Fatal(http.ListenAndServe("127.0.0.1:8080", nil))
}
