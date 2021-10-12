# tweet

The easiest way to send a tweet in Go

![lint](https://github.com/abranhe/tweet/workflows/lint/badge.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/abranhe/tweet)](https://goreportcard.com/report/github.com/abranhe/tweet)
[![CodeQL](https://github.com/abranhe/tweet/actions/workflows/codeql.yml/badge.svg)](https://github.com/abranhe/tweet/actions/workflows/codeql.yml)
[![Go Reference](https://pkg.go.dev/badge/github.com/abranhe/tweet.svg)](https://pkg.go.dev/github.com/abranhe/tweet)

## Install

```console
$ go get github.com/abranhe/tweet
```

## Usage

```go
import "github.com/abranhe/tweet"

func main() {
	credentials := TwitterCredentials{
		Token: "",
		ct0:   "",
		Auth:  "",
	}

	tweet.Tweet("Hello, world!", credentials)
}
```
