# Ad server

A small ad server written with go-kit

## Getting Started

These instructions will get you the project up and running on your local machine for development and testing purposes.

### Prerequisites

1. This project requires Go v1.22 which can be downloaded from [here](https://golang.org/dl/).
2. Docker is required to create and manage your application environment. If you don't have Docker installed, you can download it from [here](https://docs.docker.com/engine/install/).
3. Ensure that `gojq` is installed. If missing, the `make` targets will automatically install `gojq` at its latest version.

### Running the services

This project consists of two separate services which can be run independently using the following commands:

1. Auction Service:
   ```make auction-service```

2. Bidding Service:
   ```make bidding-service```

To ping your services for a response, use the following command:
```make demo```
