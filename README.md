# gosuapi

A Work in Progress framework designed around being able to get started using the osu!apiv2 quickly and easily. Consider this module VERY unstable and likely to change as the api gets developed. 

# Quick Start

This program will retrieve my user information (ponpar) from the api and print it.

```
package main

import (
	"fmt"

	"github.com/CCPupp/gosuapi"
)

func main() {
	//Set Client Info
	gosuapi.Client.ID = 1234
	gosuapi.Client.Secret = "secret key"

    //Fetch an OAuth Token from the api
	token := gosuapi.GetClientToken()
    //Pass this token to make authenticated calls
	user := gosuapi.GetUserById("3055278", "osu", token)
	fmt.Println(user)
}
```

# Supported Functions

* GetUserByID(id, gamemode, token) | Returns: User