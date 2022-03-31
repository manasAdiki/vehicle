package Errors

import "log"

func ErrorHandler(err error) {

	if err != nil {
		log.Fatalf("error reading HTTP response body: %v", err)
	}

}
