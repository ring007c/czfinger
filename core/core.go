package core

import (
	"czfinger/structs"
	"encoding/json"
	"fmt"
	"github.com/asmcos/requests"
	"log"
	"net/http"
	"regexp"
	"strings"
	"time"
)

func Detect(result *structs.FetchResult) {
	products := make([]string, 0)
	//获取网页返回数据并赋值
	web_Content := strings.ToLower(string(result.Content))
	//web_Headers := resp.Headers
	certString := string(result.Certs)
	web_Certs := result.Certs
	web_HeaderString := result.HeaderString
	headerServerString := fmt.Sprintf("Server : %v\n", result.Headers["Server"])
	fofajson, _ := Parse("fofa.json")
	for _, fp := range fofajson {
		//fofa指纹中的最后一项
		rules := fp.Rules
		matchFlag := false
		//其中的match
		//matchFlag := false
		//对每个json的最后一项进行迭代
		for _, onerule := range rules {

			//控制继续器
			ruleMatchContinueFlag := true

			for _, rule := range onerule {
				if !ruleMatchContinueFlag {
					break
				}
				lowerRuleContent := strings.ToLower(rule.Content)

				switch strings.Split(rule.Match, "_")[0] {

				case "banner":
					reBanner := regexp.MustCompile(`(?im)<\s*banner.*>(.*?)<\s*/\s*banner>`)
					matchResults := reBanner.FindAllString(web_Content, -1)
					if len(matchResults) == 0 {
						ruleMatchContinueFlag = false
						break
					}

					for _, matchResult := range matchResults {
						if !strings.Contains(strings.ToLower(matchResult), lowerRuleContent) {
							ruleMatchContinueFlag = false
							break
						}

					}

				case "title":
					reTitle := regexp.MustCompile(`(?im)<\s*title.*>(.*?)<\s*/\s*title>`)
					matchResults := reTitle.FindAllString(web_Content, -1)
					if len(matchResults) == 0 {
						ruleMatchContinueFlag = false
						break
					}

					for _, matchResult := range reTitle.FindAllString(web_Content, -1) {
						if !strings.Contains(strings.ToLower(matchResult), lowerRuleContent) {
							ruleMatchContinueFlag = false
						}
					}
				case "body":
					if !strings.Contains(web_Content, lowerRuleContent) {
						ruleMatchContinueFlag = false
					}
				case "header":
					if !strings.Contains(web_HeaderString, rule.Content) {
						ruleMatchContinueFlag = false
					}
				case "server":
					if !strings.Contains(headerServerString, rule.Content) {
						ruleMatchContinueFlag = false
					}
				case "cert":
					if (web_Certs == nil) || (web_Certs != nil && !strings.Contains(certString, rule.Content)) {
						ruleMatchContinueFlag = false
					}
				default:
					ruleMatchContinueFlag = false

				}
				// 单个rule之间是AND关系，匹配成功
				if ruleMatchContinueFlag {
					matchFlag = true
					break
				}

			}

		}
		// 多个rule之间是OR关系，匹配成功
		if matchFlag {
			products = append(products, fp.Product)
		}

	}
	PrintResult(result.Url, products)

}

func PrintResult(target string, products []string) {
	fmt.Printf("[+] %s %s\n", target, strings.Join(products, ", "))
}

func Reqdata(url string) *structs.FetchResult {
	req := requests.Requests()
	req.SetTimeout(time.Duration(10))
	resp, err := req.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	var headerString string
	respData := structs.FetchResult{
		Url:          url,
		Content:      resp.Content(),
		Headers:      resp.R.Header,
		HeaderString: headerString,
		Certs:        getCerts(resp.R),
	}
	return &respData
}

// 获取证书内容，参考byro07/fwhatweb
func getCerts(resp *http.Response) []byte {
	var certs []byte
	if resp.TLS != nil {
		cert := resp.TLS.PeerCertificates[0]
		var str string
		if js, err := json.Marshal(cert); err == nil {
			certs = js
		}
		str = string(certs) + cert.Issuer.String() + cert.Subject.String()
		certs = []byte(str)
	}
	return certs
}
