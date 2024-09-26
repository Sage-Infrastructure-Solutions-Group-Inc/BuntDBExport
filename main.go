package main

import (
	"encoding/json"
	"flag"
	"github.com/tidwall/buntdb"
	"io/ioutil"
	"log"
	"os"
)

func saveJSON(db *buntdb.DB) []byte {
	keys := make(map[string]string)
	db.View(func(tx *buntdb.Tx) error {
		tx.Ascend("", func(key, val string) bool {
			keys[key] = val
			return true
		})
		return nil
	})
	data, _ := json.Marshal(keys)
	return data
}

func main() {
	// Open the data.db file. It will be created if it doesn't exist.
	db_path := flag.String("db-path", "", "The path to the BuntDB database that you would like to export. Exports in JSON format.")
	export_path := flag.String("export-path", "", "The path that the BuntDB data should be stored in.")
	flag.Parse()
	if *db_path == "" || *export_path == "" {
		flag.Usage()
		os.Exit(1)
	}
	log.Printf("Opening database file %s for data export.", db_path)
	db, err := buntdb.Open(*db_path)
	log.Println("Opened the database file.")
	if err != nil {
		log.Fatal(err)
	}
	json_data := saveJSON(db)
	defer db.Close()
	ioutil.WriteFile(*export_path, json_data, 0644)
}
