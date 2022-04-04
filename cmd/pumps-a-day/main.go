package main

import (
    "os"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/yonikosiner/track-pumps-a-day/pkg/pumps"
)

// Get the Port from the environment so we can run on Heroku
func GetPort() string {
    var port = os.Getenv("PORT")
    // Set a default port if there is nothing in the environment
    if port == "" {
        port = "8080"
    }
    return ":" + port
}

var p pumps.Pumps

func getPumps(w http.ResponseWriter, r *http.Request) {
    count := p.GetPumps()

    if count == nil {
        w.Write([]byte(p.Count))
    } else {
        log.Fatal(count)
        w.Write([]byte("Sorry there was an error"))
    }
}

func clearPumps(w http.ResponseWriter, r *http.Request) {
    clear := p.ResetFileCount()

    if clear == nil {
        w.Write([]byte("Successfully reset the file to 0"))
    } else {
        log.Fatal(clear)
        w.Write([]byte("Sorry there was an error"))
    }
}

func updatePumpCount(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    count := p.GetPumps()

    if count != nil {
        log.Fatal(count)
        w.Write([]byte("Sorry there was en error"))
    }

    countInt, err := strconv.Atoi(p.Count)
    if err != nil {
        log.Fatal(err)
        w.Write([]byte("Sorry there was en error"))
    }

    newCount, err := strconv.Atoi(params["new-count"])
    if err != nil {
        log.Fatal(err)
        w.Write([]byte("Sorry there was en error"))
    }

    newCountFinal := countInt + newCount

    if newCountFinal >= 10 {
        w.Write([]byte("You have hit your limit for the day, slow down and relex"))
    }

    newCountStr := strconv.Itoa(newCountFinal)
    p.UpdatePumpCount(newCountStr)
}

func main() {
    router := mux.NewRouter()
    router.HandleFunc("/", getPumps).Methods("GET")
    router.HandleFunc("/clear", clearPumps).Methods("PUT")
    router.HandleFunc("/update/{new-count}", updatePumpCount).Methods("POST")
    log.Fatal(http.ListenAndServe(GetPort(), router))
}
