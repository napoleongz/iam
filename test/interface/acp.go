package main

import "net/http"

type Acp struct {
}

func (a *Acp) GetReq(r *Request) http.Request {
	return http.Request{}
}
