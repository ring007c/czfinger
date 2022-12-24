package structs

import "net/http"

// 参考byro07/fwhatweb
type FofaFinger struct {
	RuleID         string `json:"rule_id"`
	Level          string `json:"level"`
	Softhard       string `json:"softhard"`
	Product        string `json:"product"`
	Company        string `json:"company"`
	Category       string `json:"category"`
	ParentCategory string `json:"parent_category"`
	Rules          [][]struct {
		Match   string `json:"match"`
		Content string `json:"content"`
	} `json:"rules"`
}

type Rule struct {
	Match   string
	Content string
}

type RequestOptions struct {
	Timeout int
}

type FetchResult struct {
	Url          string
	Content      []byte
	Headers      http.Header
	HeaderString string
	Certs        []byte
}
