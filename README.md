# Payments API

## Summary
This project was created as a technical exercise for Form3. 

## Project Design
An overall architectural design can be found [here](https://github.com/tusiel/payments-api/blob/master/design/ArchitecturalDiagram.pdf) which shows how the Payments Manager microservice could be implemented from an infrastructure point of view. 

A flow diagram of how the Payments Manager microservice can be found [here](https://github.com/tusiel/payments-api/blob/master/design/flowDiagram.pdf). It also outlines the specified endpoints the application will expose. 

## Running & building the application
Two scripts have been created to make the process of running this application easier. 

- `run.sh` will run the application with the `--race` flag. This is intended to be for development purposes. 
- `build.sh` will create a `build` folder, which will have `darwin` and `linux` compiled versions of the application. This is intended to be for production purposes. 

