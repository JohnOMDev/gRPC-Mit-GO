package main

import (
	"flag"
	"log"
	"net"

	datasvr "csvdata/pb"

	"google.golang.org/grpc"
)

// Svr for the instance of the csv file data read and serve
var svr *Csv
var filename string

func init() {
	flag.StringVar(&filename, "f", "", "csv file name. eg. /home/data.csv")
	flag.Parse()
}

func main() {
	// set log format
	log.SetFlags(log.Lshortfile)
	// init the svr value
	svr, e := NewCsv(filename)
	if e != nil {
		log.Println(e)
		return
	}
	// start listen
	lis, e := net.Listen("tcp", ":9098")
	if e != nil {
		log.Println(e)
		return
	}

	// log.Println(svr.File, svr.Lines)
	rpcsvr := grpc.NewServer()

	datasvr.RegisterCsverServer(rpcsvr, svr)

	e = rpcsvr.Serve(lis)
	if e != nil {
		log.Println(e)
	}

}
