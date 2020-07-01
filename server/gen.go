package main

import (
	"context"
	"encoding/csv"
	"log"
	"os"

	datasvr "csvdata/pb"

	"github.com/fsnotify/fsnotify"
)

// Csv define the csv file datasource as a service
// it will implement the Get() interface which defined
// in the serv.proto
type Csv struct {
	File  string
	Lines []*datasvr.Line
}

// NewCsv create a csv instance with csv file path
func NewCsv(f string) (r *Csv, e error) {
	// validate file path
	_, e = os.Stat(f)
	if e != nil {
		return
	}
	r = &Csv{}
	r.File = f
	// init load data from vsc file
	r.LoadData()
	// watch file change
	r.watchChange()
	return
}

// Get function is implement the interface which defined in
// the serv.proto file this function must impement, otherwise
// the rpc serive can't lanch (cause not match with the prc
// defineation)
func (c *Csv) Get(ct context.Context, req *datasvr.Req) (ls *datasvr.Res, e error) {
	select {
	case <-ct.Done():
		return
	default:
		ls = &datasvr.Res{}
		ls.Data = c.Lines
	}
	return
}

// LoadData read the data from csv file, and store the data in
// the 'lines'
func (c *Csv) LoadData() {
	f, e := os.OpenFile(c.File, os.O_RDONLY, 066)
	if e != nil {
		log.Println(e)
		return
	}
	defer f.Close()

	// clear old data
	c.Lines = make([]*datasvr.Line, 0)

	r := csv.NewReader(f)
	lines, e := r.ReadAll()
	if e != nil {
		log.Println(e)
		return
	}
	// parse line into line
	for _, v := range lines {
		if len(v) == 2 {
			t := &datasvr.Line{}
			t.Time = v[0]
			t.EleUsage = v[1]
			c.Lines = append(c.Lines, t)
		}
	}

}

// watchChange to watch the csv file changes, once the file
// change, service will reload the data into memory
func (c *Csv) watchChange() {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Println(err)
		return
	}

	err = watcher.Add(c.File)
	if err != nil {
		log.Println(err)
		return
	}

	go func() {
		for {
			log.Println("wait for csv file change")
			select {
			case even, ok := <-watcher.Events:
				if !ok {
					log.Println("failed", even)
					watcher.Close()
					break
				}

				// reload csv data
				c.LoadData()
			}
		}
	}()

}
