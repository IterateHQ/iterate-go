# Iterate Go [![GoDoc](http://img.shields.io/badge/godoc-reference-blue.svg)](http://godoc.org/github.com/iteratehq/iterate-go) [![Build Status](https://travis-ci.org/iteratehq/iterate-go.svg?branch=master)](https://travis-ci.org/iteratehq/iterate-go)

## Summary

A Go client for the [Iterate](https://iteratehq.com) API 

## Features

The client currently supports
- Listing your surveys
- Emailing a survey to a user 

Additional API features will be added in the future.

## Installation

```sh
go get github.com/iteratehq/iterate-go
```

## Usage

```go
import "go get github.com/iteratehq/iterate-go"

client := iterate.New("iterate-api-key")

// Send a survey
client.EmailSurvey("survey-id", iterate.SendParams{Email: "art@vandelayindustries.com"})

// Send a survey a week from now
client.EmailSurvey("survey-id", iterate.SendParams{
	Email: "art@vandelayindustries.com",
	Delay: 7 * 24 * time.Hour,
})

// Send a survey at a specific time
client.EmailSurvey("survey-id", iterate.SendParams{
	Email: "art@vandelayindustries.com",
	Date:  time.Date(2017, time.March, 12, 10, 0, 0, 0, time.UTC),
})
```

## Documentation

For details on all the functionality in this package, check out the [GoDoc](http://godoc.org/github.com/iteratehq/iterate-go) documentation.

For a list of all available API endpoints, check out the [API documentation](http://docs.iterate.apiary.io).
