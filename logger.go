// Работа с логами
package metasdk

import (
	"encoding/json"
	"github.com/fluent/fluent-logger-golang/fluent"
	"strconv"
	"strings"
)

func (m *Meta) log(level, msg string, context map[string]interface{}) {

	if context == nil {
		context = make(map[string]interface{})
	}

	m.logger.SetPrefix(level + ":")
	m.logger.Printf("%s %s", msg, context)
	if (m.serviceId != "local_debug_serivce") && (m.buildNum != "0") {
		m.fluent(level, msg, context)
	}
	m.logger.SetPrefix("INFO:")
}

func (m *Meta) LogInfo(msg string, context map[string]interface{}) {
	m.log("INFO", msg, context)

}

func (m *Meta) LogError(msg string, context map[string]interface{}) {
	m.log("ERROR", msg, context)
}

func (m *Meta) LogWarning(msg string, context map[string]interface{}) {
	m.log("WARNING", msg, context)
}

func (m *Meta) fluent(level, msg string, context map[string]interface{}) {

	hp := strings.Split(m.gcloudlog, ":")
	port, _ := strconv.Atoi(hp[1])
	logger, _ := fluent.New(
		fluent.Config{
			FluentHost:    hp[0],
			FluentPort:    port,
			FluentNetwork: "tcp",
			MarshalAsJSON: true,
		})

	defer logger.Close()

	j, _ := json.Marshal(context)
	fmsg := FluentMsg{
		Message:  msg,
		Context:  string(j),
		Severity: level,
		ServiceContext: ServiceContext{
			Service: m.serviceId,
			Version: m.buildNum,
		}}
	logger.Post(m.serviceNameSpace+"."+m.serviceId, fmsg)

}
