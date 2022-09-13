package subject

import "fmt"

func getCommunicationValue(val string) string {
	return fmt.Sprintf("%s.%s", prefixes.Communication, val)
}

var Communication = struct {
	RequestHTTP  string
	ResponseHTTP string
}{
	RequestHTTP:  getCommunicationValue("request_http"),
	ResponseHTTP: getCommunicationValue("response_http"),
}
