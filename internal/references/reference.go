package references

import (
	"github.com/dgraph-io/badger/v4"
	"github.com/robogg133/rinha-backend-2026/pkg/vector"
)

const ReferencesFileName string = "references.bin"

type Reference struct {
	Vector vector.Vector
	Fraud  bool
}

var references []Reference

func References() []Reference {
	return references
}

var db *badger.DB

func Database() *badger.DB { return db }

func Init() error {

	var err error
	db, err = badger.Open(badger.DefaultOptions("references_database"))
	if err != nil {
		return err
	}

	return nil
}

/*
func Init() error {

	file, err := os.Open(ReferencesFileName)
	if err != nil {
		return nil
	}
	defer file.Close()

	var ref Reference

	for {
		var err error
		ref.Vector, err = vector.ReadBinary(file)
		if err != nil {
			break
		}

		b := make([]byte, 1)
		if _, err := io.ReadFull(file, b); err != nil {
			return err
		}

		switch b[0] {
		case 1:
			ref.Fraud = true
		}

		references = append(references, ref)
	}

	return nil
}
*/
