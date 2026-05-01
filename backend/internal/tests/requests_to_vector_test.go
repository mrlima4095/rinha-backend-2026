package tests

import (
	"encoding/json"
	"os"
	"testing"
)

func TestRequestsToVector(t *testing.T) {

	fBlob, err := os.ReadFile("data/requests_to_vector.test.json")
	if err != nil {
		t.Fatal(err)
	}

	var allJSons []json.RawMessage

	if err := json.Unmarshal(fBlob, &allJSons); err != nil {
		t.Fatal(err)
	}

}
