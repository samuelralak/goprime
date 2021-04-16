## **Go Prime**

Given a number input, Go Prime will return the highest prime number lower than the input

## Installation

Clone this repository

```shell
$ git clone git@github.com:samuelralak/goprime.git
```

Switch to GoPrime directory then build and start server

```shell
$ cd goprime 
$ ./build && ./start
```

The above command will run the GoPrime image interactively by default. The server should be available at `localhost:8080`

## Usage

GoPrime provides the following endpoints

```shell
# Root endpoint returns app version
http://localhost:8080/

# Number input returns highest primer number lower than input(:number)
http://localhost:8080/:number 
```

### Examples

Get root endpoint
```shell
$ curl --location --request GET 'localhost:8080/'

# output
Go Prime API v0.0
```

Get highest prime number lower than then given input
```shell
$ curl --location --request GET 'localhost:8080/55'

# output
53
```

## Go Live

A Live version of the service can be found at:
```shell
https://sam-goprime.herokuapp.com
```
