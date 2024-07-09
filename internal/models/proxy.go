package models

import "strings"

type ProxyRequest struct {
	Method  string            `json:"method" default:"GET"`
	URL     string            `json:"url"`
	Headers map[string]string `json:"headers,omitempty"`
}

type ProxyResponse struct {
	ID      string            `json:"id"`
	Status  int               `json:"status"`
	Headers map[string]string `json:"headers"`
	Length  int               `json:"length"`
}

func (prq *ProxyRequest) MethodToUpper() {
	prq.Method = strings.ToUpper(prq.Method)
}
