## magic-go

This is a sample go server that serves up a /health endpoint and any json files it finds in a data/ directory as separate endpoints. It makes use of the [Go Dep tool](https://github.com/golang/dep) to install Go dependencies. A simple Makefile is provided to run and build the Go server.

To install all dependencies run the following cmd:
```
make setup
```

The following cmd runs the server:
```
make run
```

THe following cmd builds the Go binary:
```
make build
```
