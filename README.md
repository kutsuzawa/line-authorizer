# line-authorizer
[![Build Status](https://travis-ci.org/kutsuzawa/line-authorizer.svg?branch=master)](https://travis-ci.org/kutsuzawa/line-authorizer) [![Maintainability](https://api.codeclimate.com/v1/badges/d1358f4c069a4275eb34/maintainability)](https://codeclimate.com/github/kutsuzawa/line-authorizer/maintainability) [![Test Coverage](https://api.codeclimate.com/v1/badges/d1358f4c069a4275eb34/test_coverage)](https://codeclimate.com/github/kutsuzawa/line-authorizer/test_coverage)   
A library for getting line access token.

## example
```go
package main

import (
	"fmt"
	"os"

	"github.com/kutsuzawa/line-authorizer"
)

func main() {
	config := authorizer.Config{
		ID:     "channelID",
		Secret: "channelSecret",
	}
	client := authorizer.NewClient(config)
	token, err := client.PublishChannelToken()
	if err != nil {
		os.Exit(1)
	}
	fmt.Printf("token: %s\n", *token)
}

```
