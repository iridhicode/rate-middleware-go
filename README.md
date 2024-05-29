# Rate Limiting Middleware

- This project implements a rate limiting middleware for HTTP handlers in Go.
- It provides a way to limit the number of requests processed per client based on their IP address.
- The rate limiting is implemented using the leaky bucket algorithm.
  
## Features

- Whitelisting and blacklisting of IP addresses
- Leaky bucket algorithm for rate limiting
- Middleware for easy integration with HTTP handlers

## Prerequisites

- Go 1.16 or later

## Installation

- Clone the repository:

```bash
git clone https://github.com/iridhicode/rate-limiting-middleware.git
```

- Change into the project directory:

```bash
cd rate-limiting-middleware
```

- Install the dependencies:

```bash
go mod download
```

## Configuration

The rate limiting middleware can be configured using a .env file. Create a file named .env in the root directory of the project and add the following variables:

- Copy codePORT=4000
- CAPACITY=10
- LEAK_RATE=100ms
- WHITELIST_IPS=127.0.0.1,::1
- BLACKLIST_IPS=192.168.0.1,10.0.0.1

- PORT: The port number on which the server will run.
- CAPACITY: The capacity of the leaky bucket rate limiter. It represents the maximum number of requests allowed in a burst.
- LEAK_RATE: The leak rate of the leaky bucket rate limiter. It determines how quickly the bucket empties. It is specified as a duration string (e.g., "100ms").
- WHITELIST_IPS: A comma-separated list of IP addresses that are whitelisted and exempt from rate limiting.
- BLACKLIST_IPS: A comma-separated list of IP addresses that are blacklisted and always blocked.
