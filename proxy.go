package metasdk

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

const apiProxyURL = "http://apiproxy.apis.kb.1ad.ru"

func (m *Meta) CallAPIProxy(engine, method string, payload map[string]interface{}, analyzeErr bool, substrVar []string) (*ApiProxyResponse, error) {
	url := fmt.Sprintf("%s/%s", m.ApiProxyURL, method)

	body := map[string]interface{}{
		"engine":  engine,
		"payload": payload,
	}

	b, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	client := http.Client{Timeout: time.Duration(time.Minute)}
	req, err := http.NewRequest("POST", url, bytes.NewReader([]byte(string(b))))
	check(err)
	req.Header.Add("User-Agent", m.userAgent)
	req.Header.Add("X-App", "META")
	req.Header.Add("X-Worker", m.serviceId)

	for i := 0; i < 20; i++ {
		resp, err := client.Do(req)
		check(err)

		result, err := checkErr(resp, analyzeErr, substrVar)

		if err != nil {
			if err.Error() == "Retry" {
				continue
			} else {
				return nil, err
			}
		} else {
			return result, nil
		}

	}
	return nil, errors.New("End of tries")
}

func checkErr(resp *http.Response, analyzeErr bool, substrVariants []string) (*ApiProxyResponse, error) {
	retryErr := errors.New("Retry")
	result := &ApiProxyResponse{}
	statusCode := resp.StatusCode
	substrVariants = append(substrVariants, "TLSV1_ALERT_ACCESS_DENIED")
	for _, v := range [3]int{502, 503, 504} {
		if statusCode == v {
			return nil, retryErr
		}
	}

	body, _ := ioutil.ReadAll(resp.Body)

	if statusCode >= 400 {
		s := string(body)
		for _, v := range substrVariants {
			if strings.Index(v, s) >= 0 {
				return nil, retryErr
			}
		}
		return nil, errors.New("HTTP request failed")
	}

	if analyzeErr {
		_ = json.Unmarshal(body, result)
		if result.Error.Type != "" {
			for _, v := range substrVariants {
				if strings.Index(v, result.Error.Message) >= 0 {
					return nil, retryErr
				}
			}
			return nil, errors.New("Api Proxy Error")
		}

	}
	return result, nil
}
