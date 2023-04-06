package main

import (
	"encoding/csv"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/opt"
)

var (
	outputMode = flag.String("output", "line", "Output mode (line,line_raw,csv,json)")
)

func main() {
	flag.Parse()

	args := flag.Args()
	if len(args) < 1 {
		log.Fatalf("usage: %s <path/to/leveldb>", os.Args[0])
	}

	options := &opt.Options{
		ErrorIfMissing: true,
		ReadOnly:       true,
	}

	db, err := leveldb.OpenFile(args[0], options)
	if err != nil {
		log.Fatalf("open db err: %s", err)
	}

	defer db.Close()
	iter := db.NewIterator(nil, nil)

	jsonout := json.NewEncoder(os.Stdout)
	csvout := csv.NewWriter(os.Stdout)

	type Record struct {
		Key   []byte
		Value []byte
	}

	for iter.Next() {
		key := iter.Key()

		value, err := db.Get([]byte(key), nil)
		if err != nil {
			log.Fatal(err)
		}
		if *outputMode == "json" {
			err = jsonout.Encode([][]byte{key, value})
			if err != nil {
				log.Fatal(err)
			}
		} else if *outputMode == "csv" {
			err = csvout.Write([]string{string(key), string(value)})
			if err != nil {
				log.Fatal(err)
			}
			csvout.Flush()
		} else if *outputMode == "line" {
			fmt.Printf("%q\t%q\n", key, value)
		} else if *outputMode == "line_raw" {
			fmt.Printf("%s\t%s\n", key, value)
		} else {
			log.Fatalf("unknown output mode %s", *outputMode)
		}
	}

	iter.Release()
	err = iter.Error()
	if err != nil {
		log.Fatalf("iter err: %s", err)
	}
}
