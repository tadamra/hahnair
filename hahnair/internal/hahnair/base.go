package hahnair

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gorilla/mux"
)

const Authtoken = "czZ5SFh5aHR5QzNMRHgxSDpOQW5qUHRuUHRtTGhidWZNUklkMXo5ZWJFN25HZUJGR1R0N1FkcFBDbGhrWXFWb2NRdlo3MHpDOUJiZEx3bWJs"

const authReqDefaultTemplate = `{
	"transactionReference": "unique-transactionReference",
    "merchant": {
        "entity": "default"
    },
    "instruction": {
        "narrative": {
            "line1": "trading name"
        },
        "value": {
            "currency": "GBP",
            "amount": 250
        },
        "paymentInstrument": {
            "type": "card/token",
            "href": "%s"
        }
    }
}`

type Card struct {
	Description       string `json:"description"`
	PaymentInstrument struct {
		Type           string `json:"type"`
		CardHolderName string `json:"cardHolderName"`
		CardNumber     string `json:"cardNumber"`
		CardExpiryDate struct {
			Month int `json:"month"`
			Year  int `json:"year"`
		} `json:"cardExpiryDate"`
		BillingAddress struct {
			Address1    string `json:"address1"`
			Address2    string `json:"address2"`
			Address3    string `json:"address3"`
			PostalCode  string `json:"postalCode"`
			City        string `json:"city"`
			State       string `json:"state"`
			CountryCode string `json:"countryCode"`
		} `json:"billingAddress"`
	} `json:"paymentInstrument"`
}

func parseToken(resMap map[string]interface{}) string {
	tmp, ok := resMap["tokenPaymentInstrument"]
	if !ok {
		panic("could not find tokenPaymentInstrument in JSON")
	}

	tPInstrument, ok := tmp.(map[string]interface{})
	if !ok {
		panic("The Token returned response does not contain a tokenPaymentInstrument of the expected type")
	}

	tmp, ok = tPInstrument["href"]
	token, ok := tmp.(string)
	if !ok {
		panic("Token returned in href is not of expected type")
	}
	return token

}

func getToken(ctx context.Context, c *http.Client, cardInfo Card) (string, error) {
	payloadBytes, err := json.Marshal(cardInfo)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not encode card information: %s", err)
		return "", err
	}

	payload := bytes.NewReader(payloadBytes)
	url := "https://try.access.worldpay.com/tokens"
	req, err := http.NewRequestWithContext(ctx, "POST", url, payload)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Could create request: %s", err)
		return "", err
	}

	req.Header.Add("Authorization", "Basic "+Authtoken)
	req.Header.Add("Content-Type", "application/vnd.worldpay.tokens-v2.hal+json")

	res, err := c.Do(req)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Could retrive Token: %s", err)
		return "", err
	}
	defer res.Body.Close()

	decodedRes := make(map[string]interface{})
	err = json.NewDecoder(res.Body).Decode(&decodedRes)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Could decode token response: %s", err)
		return "", nil
	}

	token := parseToken(decodedRes)
	fmt.Fprintln(os.Stdout, "Token:")
	fmt.Fprintf(os.Stdout, "%s", token)
	return token, nil
}

func AuthorizePayment(ctx context.Context, c *http.Client, token string) (map[string]interface{}, error) {
	tmpPayload := fmt.Sprintf(authReqDefaultTemplate, token)
	payload := strings.NewReader(tmpPayload)
	url := "https://try.access.worldpay.com/payments/authorizations"

	req, err := http.NewRequestWithContext(ctx, "POST", url, payload)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not create request: %s", err)
		return nil, err
	}

	req.Header.Add("Authorization", "Basic "+Authtoken)
	req.Header.Add("Content-Type", "application/vnd.worldpay.tokens-v2.hal+json")

	// Send HTTP request
	res, err := c.Do(req)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error could not sent Request: %s", err)
		return nil, err
	}
	defer res.Body.Close()

	decodedRes := make(map[string]interface{})
	err = json.NewDecoder(res.Body).Decode(&decodedRes)
	if err != nil {
		fmt.Fprintf(os.Stderr, "could not decode response: %s", err)
		return nil, err
	}
	return decodedRes, nil
}

func ProcessPayment(card Card) (map[string]interface{}, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	client := &http.Client{}

	token, err := getToken(ctx, client, card)
	if err != nil {
		return nil, err
	}

	return AuthorizePayment(ctx, client, token)
}

func MyHTTPPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func CardInfo(w http.ResponseWriter, r *http.Request) {
	var card Card

	w.Header().Set("Content-Type", "application/json")
	err := json.NewDecoder(r.Body).Decode(&card)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not decode Card information: %s", err)

		m := make(map[string]string)
		m["error"] = "Invalid JSON"
		json.NewEncoder(w).Encode(&m)
		return
	}

	resp, err := ProcessPayment(card)
	if err != nil {
		m := make(map[string]string)
		m["error"] = "Invalid JSON"
		json.NewEncoder(w).Encode(&m)
	}

	json.NewEncoder(w).Encode(&resp)
}

func RunServer() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", MyHTTPPage)
	myRouter.HandleFunc("/cardinfo", CardInfo).Methods("POST")
	log.Fatal(http.ListenAndServe(":8000", myRouter))
}
