# Project Description
This project has been create as a test for Hahnair, it contains  the following:

2 api endpoints:
- http://localhost:8000/
    Just a welcome page
- http://localhost:8000/cardinfo
    API end point which accepts the information provided by the users

# How to Run The Project
1- install docker compose please refer to the following link:
    https://docs.docker.com/compose/install/

2- run the command "docker-compose p -d" to run the project

3- type in a browser of you choice "http://localhost:8000/" to check that the system is running, the following message should be displayed:
    Welcome to the HomePage!

# Submitting a Credit Card Payment
please use th following API http://localhost:8000/cardinfo.
The body submitted should follow the scheme:

`{
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

# Issues not completed
 - Front End
 - Database Encryption


