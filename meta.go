// Для основного импортируемого функционала
package metasdk

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

const (
	metaURL = "http://apimeta.devision.io"
	apiPath = "/api/meta/v1/"
)

func (m *Meta) Meta() {
	if m.ApiProxyURL == "" {
		m.ApiProxyURL = apiProxyURL
	}

	if m.MetaURL == "" {
		m.MetaURL = metaURL
	}

	m.developerSettings = readDeveloperSettings()

	sns := os.Getenv("SERVICE_NAMESPACE")
	if sns == "" {
		sns = "appscript"
	}
	m.serviceNameSpace = sns

	sid := os.Getenv("SERVICE_ID")
	if sid == "" {
		sid = "local_debug_serivce"
	}
	m.serviceId = sid

	bn := os.Getenv("BUILD_NUM")
	if bn == "" {
		bn = "0"
	}
	m.buildNum = bn

	createDefaultHeader(m)

}

func createDefaultHeader(m *Meta) {
	headers := make(map[string]string, 4)
	headers["content-type"] = "application/json;charset=UTF-8"
	headers["User-Agent"] = fmt.Sprintf("%s | b%s | %s", m.serviceId, m.buildNum, m.postfix)
	headers["X-META-Developer-Login"] = m.developerSettings.ApiHeaders.Login
	headers["X-META-Developer-Token"] = m.developerSettings.ApiHeaders.Token
	m.defaultHeaders = headers
}

func (m *Meta) nativeCall(service, method, httpMethod string, data []byte) []byte {
	url := fmt.Sprintf("%s%s%s/%s", m.MetaURL, apiPath, service, method)
	client := http.Client{Timeout: time.Duration(time.Minute)}
	req, err := http.NewRequest(httpMethod, url, bytes.NewReader(data))
	check(err)
	for k, v := range m.defaultHeaders {
		req.Header.Set(k, v)
	}
	for i := 0; i < 20; i++ {
		resp, err := client.Do(req)
		check(err)

		body, _ := ioutil.ReadAll(resp.Body)
		if resp.StatusCode != 200 {
			if strings.Contains(string(body), "Служба частично или полностью недоступна") {
				continue
			} else {
				log.Panic("Не удалось получить данные")
			}
		} else {
			return body
		}

	}

	log.Panic("Не удалось связатьсяс ервисом")
	return nil
}
