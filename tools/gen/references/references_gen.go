//go:generate go run .
package main

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/dgraph-io/badger/v4"
	"github.com/robogg133/rinha-backend-2026/pkg/vector"
)

type rawEntry struct {
	Vector []float32 `json:"vector"`
	Label  string    `json:"label"`
}

func main() {

	db, err := badger.Open(badger.DefaultOptions(filepath.Join("..", "..", "..", "references_database")))
	if err != nil {
		panic(err)
	}
	defer db.Close()

	var raw []rawEntry

	f, err := os.Open("blobs/references.json.gz")
	if err != nil {
		panic(err)
	}

	reader, err := gzip.NewReader(f)
	if err != nil {
		f.Close()
		panic(err)
	}

	if err := json.NewDecoder(reader).Decode(&raw); err != nil {
		panic(err)
	}
	reader.Close()
	f.Close()

	wb := db.NewWriteBatch()
	defer wb.Cancel()
	for i, e := range raw {
		var a byte
		switch e.Label {
		case "fraud":
			a = 1
		}

		vec := vector.Vector(e.Vector)

		var buffer bytes.Buffer

		if err := vec.WriteBinary(&buffer); err != nil {
			panic(err)
		}

		if err := wb.Set(buffer.Bytes(), []byte{a}); err != nil {
			panic(err)
		}
		fmt.Println("INSERT:", i)
	}
	if err := wb.Flush(); err != nil {
		panic(err)
	}

	fmt.Println("SUCESS! DATABASE FLUSHED!!")

}
