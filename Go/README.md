# MobileESP Golang

## How To Use

example check if request from an android phone
```go
package main

import (
	mdetect "github.com/fari-99/mobileesp/Go"
	"log"
	"net/http"
)

func yourHandler(w http.ResponseWriter, r *http.Request) {
	detect := mdetect.NewMDetect(r)
	if detect.IsAndroidPhone == 0 {
		log.Printf("i'm an android phone")
		return
	}
	
	log.Printf("not android phone")
	return
}

```