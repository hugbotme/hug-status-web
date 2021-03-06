package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"github.com/hugbotme/hug-status-web/config"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
)

var (
	flagConfigFile *string
	flagPidFile    *string
	flagVersion    *bool
)

const (
	majorVersion = 1
	minorVersion = 0
	patchVersion = 0
)

type Info struct {
	InProgress int `json:"in_progress"`
	Merged     int `json:"merged"`
	Closed     int `json:"closed"`
	Queue      int `json:"queue"`
}

// Init function to define arguments
func init() {
	flagConfigFile = flag.String("config", "", "Configuration file")
	flagPidFile = flag.String("pidfile", "", "Write the process id into a given file")
	flagVersion = flag.Bool("version", false, "Outputs the version number and exits")
}

func main() {
	flag.Parse()

	// Output the version and exit
	if *flagVersion {
		fmt.Printf("hug-status-web v%d.%d.%d\n", majorVersion, minorVersion, patchVersion)
		return
	}

	// Check for configuration file
	if len(*flagConfigFile) <= 0 {
		log.Fatal("No configuration file found. Please add the --config parameter")
	}

	// PID-File
	if len(*flagPidFile) > 0 {
		ioutil.WriteFile(*flagPidFile, []byte(strconv.Itoa(os.Getpid())), 0644)
	}

	config, err := config.NewConfiguration(flagConfigFile)
	if err != nil {
		log.Fatal("Configuration initialisation failed:", err)
	}

	http.HandleFunc("/info.json", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Content-Security-Policy", "*")
		w.Header().Set("Content-Type", "application/json; charset=utf-8")

		red := config.ConnectRedis()
		defer red.Close()

		inProgress, _ := redis.Int(red.Do("LLEN", "hug:pullrequests"))
		merged, _ := redis.Int(red.Do("LLEN", "hug:pullrequests:merged"))
		closed, _ := redis.Int(red.Do("LLEN", "hug:pullrequests:closed"))
		queue, _ := redis.Int(red.Do("LLEN", "hug:queue"))

		pr := Info{
			inProgress,
			merged,
			closed,
			queue,
		}
		data, err := json.Marshal(pr)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write(data)
	})

	fmt.Printf("Running on http://%s\n", config.StatusWeb.URL)
	log.Fatal(http.ListenAndServe(config.StatusWeb.URL, nil))
}
