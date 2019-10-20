# Kramp
Kramp

# Requirements
- Install Go binary with it's standard tooling support.

# Dependencies
- The server utilizes gin-gonic - a small memory footprint - HTTP server written in Go, for accomplishing request routing and request parsing.

# Instructions to run
- The project follows standard Go directory structure, hence unzip the project and point the $GOPATH to this location.
    - After unzipping the directory structure should look like this: `$GOPATH/src/github.com/nawazish-github/kramp`, `$GOPATH/bin`, `$GOPATH/pkg`.
- The `config` directory has `config.json` file which has attributes to control the behavior of the server; these can be tweaked as per requirements.
- cd to `src/github.com/nawazish-github/kramp/server`.
- run `go build` (or`go install`).
- run `go run main.go`; this would start up the Kramp server on port 8080 (by default).


# Instructions to run requests
- From any `HTTP` client, execute a `HTTP GET` Request: `http://localhost:8080/query/searchString`.
- Based on the settings pre-configured in the `config.json` file, the response would be returned.

# Resiliency, stability and performance
- In case if Kramp server could not execute requests on remote servers, it does not crashes, rather it returns an empty response to the client.
- In case of any remote service failure, Kramp server responds to the clients with `cached` data respecting the entries from the `config.json` file.
- In an attempt to keep up with the performance requirements, Kramp server `cancels` any remote requests which takes more than the `maxTimeout` time from `config.json` file.


# What is not covered but known
- There are no tests. 

# Conventions used
- The function, method and variable names follow camel case formatting.

