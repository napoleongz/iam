package main

import "net/http"

type Ace struct {
}

func (a *Ace) GetReq(r *Request) http.Request {
	return http.Request{}
}
