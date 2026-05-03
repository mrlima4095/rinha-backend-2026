package tests

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"

	"github.com/robogg133/rinha-backend-2026/pkg/payments"
)

func TestRequestsToVector(t *testing.T) {

	// Init
	fBlob, err := os.ReadFile("data/requests_to_vector.test.json")
	if err != nil {
		t.Fatal(err)
	}

	var allJsons []json.RawMessage

	if err := json.Unmarshal(fBlob, &allJsons); err != nil {
		t.Fatal(err)
	}
	t.Log("Starting to test...")
	// Real testing

	for _, req := range allJsons {

		p := new(payments.Payment)

		if err := json.Unmarshal(req, p); err != nil {
			t.Fatal(err)
		}

		data, err := json.Marshal(*p.ToVector())
		if err != nil {
			t.Fatal(err)
		}
		fmt.Println(string(data))
	}

}
