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
}

func main() {

	red, err := redis.Dial("tcp", ":6379")
	if err != nil {
		fmt.Println("error", err)
		return
	}
	defer red.Close()

	http.HandleFunc("/info.json", func(w http.ResponseWriter, r *http.Request) {
		in_progress, _ := redis.Int(red.Do("LLEN", "hugbot:pullrequests"))
		merged, _ := redis.Int(red.Do("LLEN", "hugbot:pullrequests:merged"))
		closed, _ := redis.Int(red.Do("LLEN", "hugbot:pullrequests:closed"))

		pr := Info{
			in_progress, merged, closed,
		}
		data, err := json.Marshal(pr)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write(data)
	})

	fmt.Println("Running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
