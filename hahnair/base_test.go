package hahnair

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
	"testing"
)

const testCard = `{
	"description": "Test Token Description",
	"paymentInstrument": {
		"type": "card/front",
		"cardHolderName": "John Appleseed",
		"cardNumber": "4444333322221111",
		"cardExpiryDate": {
			"month": 5,
			"year": 2035
		},
		"billingAddress": {
			"address1": "Worldpay",
			"address2": "1 Milton Road",
			"address3": "The Science Park",
			"postalCode": "CB4 0WE",
			"city": "Cambridge",
			"state": "Cambridgeshire",
			"countryCode": "GB"
		}
	}
}`

// TestProcessPayment will test the logic behind the operations of getting the token and doing the autherization
func TestProcessPayment(t *testing.T) {
	var card Card

	r := strings.NewReader(testCard)
	err := json.NewDecoder(r).Decode(&card)
	if err != nil {
		t.Fatal(err)
	}

	m, err := ProcessPayment(card)
	if err != nil {
		t.Fatal(err)
	}

	// This will print the result of the unit test
	fmt.Fprintf(os.Stdout, "%v", m)
}

// TestCardInfo will run the server and call the spi end point
func TestCardInfo(t *testing.T) {
	go RunServer()

	r := strings.NewReader(testCard)

	client := &http.Client{}

	req, err := http.NewRequest("POST", "http://127.0.0.1:8000/cardinfo", r)
	if err != nil {
		t.Fatal(err)
	}

	res, err := client.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	defer res.Body.Close()

	decodedRes := make(map[string]interface{})
	err = json.NewDecoder(res.Body).Decode(&decodedRes)
	if err != nil {
		t.Fatal(err)
	}

	// Prints out the returned response
	fmt.Fprintf(os.Stdout, "%v", decodedRes)

}
