package main

import (
	"bytes"
	"fmt"

	"github.com/dgraph-io/badger/v4"
	"github.com/robogg133/rinha-backend-2026/internal/references"
	"github.com/robogg133/rinha-backend-2026/pkg/vector"
)

func FindFraudScore(query vector.Vector) float32 {
	var top5 [5]Match
	db := references.Database()

	db.View(func(txn *badger.Txn) error {

		it := txn.NewIterator(badger.DefaultIteratorOptions)
		defer it.Close()

		for it.Rewind(); it.Valid(); it.Next() {

			vec, err := vector.ReadBinary(bytes.NewReader(it.Item().Key()))
			if err != nil {
				return err
			}

			isFraud := false

			it.Item().Value(func(val []byte) error {

				switch val[0] {
				case 1:
					isFraud = true
				}

				return nil
			})

			dist := compareVector(query, vec)
			fmt.Printf("==%.4f==\n", dist)
			for i, v := range top5 {
				if dist < v.dist {
					top5[i].dist = dist
					top5[i].fraud = isFraud
				}
			}

		}

		return nil
	})

	frauds := 0
	for i, m := range top5 {
		fmt.Printf("=> top%d: %.4f\n", i, m.dist)
		if m.fraud {
			frauds++
		}
	}

	return float32(frauds) / 5.0
}

func compareVector(a, b vector.Vector) float32 {
	var sum float32
	for i := range 14 {
		sum += euclidean(a[i], b[i])
	}
	return sum
}

func euclidean(a, b float32) float32 {
	return square(square(a - b))
}

func square(a float32) float32 {
	return a * a
}

type Match struct {
	dist  float32
	fraud bool
}
