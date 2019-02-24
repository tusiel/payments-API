# Payments API

## Summary
This project was created as a technical exercise for Form3. 

The application talks to a MongoDB database, but this could be easily replaced by changing the implementation in the `db` package. Because of how MongoDB handles document ID's this Application does not allow you to specify an ID on creation, unless creating a new primitive ObjectID. 

The request will fail if one is specified, otherwise it will lead to unintended consequences in the future. 

## Project Design
An overall architectural design can be found [here](https://github.com/tusiel/payments-api/blob/master/design/ArchitecturalDiagram.pdf) which shows how the Payments Manager microservice could be implemented from an infrastructure point of view. 

A flow diagram of how the Payments Manager microservice can be found [here](https://github.com/tusiel/payments-api/blob/master/design/flowDiagram.pdf). It also outlines the specified endpoints the application will expose. 

## Running & building the application
MongoDB is a prerequisite for running this application. A `docker-compose.yml` file has been added to the project so, if you don't have it installed locally, you can run it in a docker container by running `docker-compose up`. 

Two scripts have been created to make the process of running this application easier. 

- `run.sh` will run the application with the `--race` flag. This is intended to be for development purposes. 
- `build.sh` will create a `build` folder, which will have `darwin` and `linux` compiled versions of the application. This is intended to be for production purposes. 

The application will run according to the `listenAddress` specified in the `config.json` file, in the project root (defaults to `localhost:3005`). The MongoDB connection string, database name and collection can be configued in the same configuration file. 

## Endpoints

| Method      | Endpoint | Description |
| ----------- | -----------| ----------- |
| GET         | /api/v1/payments | Returns an array of all payments |
| GET         | /api/v1/payment/{id}| Get a payment by ID |
| POST        | /api/v1/payment| Add a payment |
| PUT         | /api/v1/payment/{id} | Update a payment by ID |
| DELETE      | /api/v1/payment/{id}| Delete a payment by ID|

Example of a POST request to `/payments` 

```json
{"type":"Payment","version":0,"organisation_id":"743d5b63-8e6f-432e-a8fa-c5d8d2ee5fcb","attributes":{"amount":"100.21","beneficiary_party":{"account_name":"W Owens","account_number":"31926819","account_number_code":"BBAN","account_type":0,"address":"1 The Beneficiary Localtown SE2","bank_id":"403000","bank_id_code":"GBDSC","name":"Wilfred Jeremiah Owens"},"charges_information":{"bearer_code":"SHAR","sender_charges":[{"amount":"5.00","currency":"GBP"},{"amount":"10.00","currency":"USD"}],"receiver_charges_amount":"1.00","receiver_charges_currency":"USD"},"currency":"GBP","debtor_party":{"account_name":"EJ Brown Black","account_number":"GB29XABC10161234567801","account_number_code":"IBAN","address":"10 Debtor Crescent Sourcetown NE1","bank_id":"203301","bank_id_code":"GBDSC","name":"Emelia Jane Brown"},"end_to_end_reference":"Wil piano Jan","fx":{"contract_reference":"FX123","exchange_rate":"2.00000","original_amount":"200.42","original_currency":"USD"},"numeric_reference":"1002001","payment_id":"123456789012345678","payment_purpose":"Paying for goods/services","payment_scheme":"FPS","payment_type":"Credit","processing_date":"2017-01-18","reference":"Payment for Em's piano lessons","scheme_payment_sub_type":"InternetBanking","scheme_payment_type":"ImmediatePayment","sponsor_party":{"account_number":"56781234","bank_id":"123123","bank_id_code":"GBDSC"}}}
```

## Points of consideration
- The `db_test.go` tests require a connection to MongoDB. These were designed as integration tests, not Unit Tests. 

## Next Steps
- Investigate Go Integration Testing strategies, to ensure this application has implemented them properly.
- Investigate how to perform Unit Test setup/teardown for each test, to avoid duplication (`deleteAll` method called multiple times)