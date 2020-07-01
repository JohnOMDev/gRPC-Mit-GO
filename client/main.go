package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	datacsv "csvdata/pb"

	"google.golang.org/grpc"
)

// define datacsv client
var client datacsv.CsverClient

func main() {
	// connect to grpc serve
	conn, e := grpc.Dial("127.0.0.1:9098", grpc.WithInsecure())
	if e != nil {
		log.Println(e)
		return
	}
	defer conn.Close()

	client = datacsv.NewCsverClient(conn)

	// start a webservice to watch data changes
	http.HandleFunc("/", Get)
	http.Handle("/view/", http.StripPrefix("/view/", http.FileServer(http.Dir("webpage"))))
	http.ListenAndServe(":8080", nil)
}

// Get function write the csv data to webside in json format
func Get(w http.ResponseWriter, r *http.Request) {
	rs, e := client.Get(context.Background(), &datacsv.Req{})
	if e != nil {
		log.Println(e)
		return
	}
	dat, e := json.Marshal(rs)
	if e != nil {
		log.Println(e)
	}
	w.Header().Set("content-type", "text/json")
	w.Write(dat)
}
