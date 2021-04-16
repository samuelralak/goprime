## **Go Prime**

Given a number input, Go Prime will return the highest prime number lower than the input

## Prerequisites

- Install Golang (Go Prime uses Go version 1.15)
- Create your Golang workspace 
  ```shell
  $ mkdir -p $HOME/go/{bin,src}
  ```

## Installation

Fetch GoPrime by running the following command:

```shell
$ go get github.com/samuelralak/goprime
```

Switch to GoPrime directory and start application

```shell
$ cd $HOME/go/src/github.com/samuelralak/goprime && go run .
```

The server should be available at `localhost:8080`

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