## **Go Prime**

Given a number input, Go Prime will return the highest prime number lower than the input

## Design Decisions & Trade-offs

GoPrime follows a REST style architecture powered by the [Gin Web Framework](https://gin-gonic.com/). 

The choice in web framework was inspired by a need for a simple yet fully featured API without compromising on speed and size.

Building GoPrime as a REST API service reduces its UI dependence, allowing it to provide data over an API instead of being coupled to a particular web framework.

GoPrime follows a very minimalistic approach as can be seen in the choice of framework, this makes it un-opinionated thus having the freedom to follow whatever design pattern. This works pretty well for GoPrime's current functions but will not work if for some it evolves into larger enterprise service.   

## Prerequisites

- Ensure that you have docker installed

## Installation

Clone this repository

```shell
$ git clone git@github.com:samuelralak/goprime.git
```

Switch to GoPrime directory then build and start server

```shell
$ cd goprime && make
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
