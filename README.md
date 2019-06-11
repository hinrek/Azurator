# Azure-migrator

Azur e DevOps project migration tool

This tool is in very early stage!

Not to be used if no releases available!

# Instructions

## Prerequisites

- Install Go from https://golang.org
- Clone the project `$ go get github.com/hinrek/Azure-migrator`
    - This will clone the project to default `$GOPATH` = `$HOME/go` (Unix)

## Configuration

[Configuration](configs/README.md)

## Building and running the application (Unix)

Change directory into the cloned Azure-migrator project:

`$ cd $HOME/go/src/github.com/hinrek/Azure-migrator/`

> NB! If you did not use the default $GOPATH then change the command accordingly!

### 1 Method

Build: 

`$ go build`

Run after build:

`$ ./Azure-migrator`

### 2 Method

Run application:

`$ go run main.go`


## Running tests
`$ go test ./...`

# CI build statuses

[![Build Status](https://dev.azure.com/hinrek/Azure-migrator%20pipelines/_apis/build/status/hinrek.Azure-migrator?branchName=develop)](https://dev.azure.com/hinrek/Azure-migrator%20pipelines/_build/latest?definitionId=1&branchName=develop)

[![Build Status](https://travis-ci.com/hinrek/Azure-migrator.svg?branch=develop)](https://travis-ci.com/hinrek/Azure-migrator)
