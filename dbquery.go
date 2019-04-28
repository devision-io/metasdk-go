package metasdk

import (
	"encoding/json"
	"log"
)

func (m *Meta) Query(options map[string]string, command string, maxRows int, parameters map[string]string) dbResponse {
	dbQuery := dbQuery{}
	dbQuery.Database = map[string]string{
		"alias": options["dbAlias"],
	}
	dbQuery.DbQuery = map[string]interface{}{
		"maxRows":    maxRows,
		"command":    command,
		"parameters": parameters,
		"shardKey":   options["shardKey"],
	}

	req, err := json.Marshal(dbQuery)
	check(err)

	resp := m.nativeCall("db", "query", "POST", req)
	if resp == nil {
		log.Panic("Произошла ошибка при запросе в БД")
	}

	dbResponse := dbResponse{}
	check(json.Unmarshal(resp, &dbResponse))

	return dbResponse
}

func (m *Meta) One(options map[string]string, command string, maxRows int, parameters map[string]string) map[string]interface{} {
	rows := m.Query(options, command, maxRows, parameters).Rows
	if len(rows) > 0 {
		return rows[0]
	}
	return nil
}

func (m *Meta) All(options map[string]string, command string, parameters map[string]string) []map[string]interface{} {
	maxRows := 0
	rows := m.Query(options, command, maxRows, parameters).Rows
	return rows
}
