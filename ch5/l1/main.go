package main

import (
	"net/http"
)

func getContentType(res *http.Response) string {
	myRes := *res
	return myRes.Header.Get("content-type")
}
