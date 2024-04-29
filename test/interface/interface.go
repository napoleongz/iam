package main

import "net/http"

type Request struct {
	Cluster   string
	Namespace string
	Token     string
}

type Cluster_in interface {
	GetReq(r *Request) http.Request
}
