package metasdk

import (
	"encoding/json"
	"log"
)

func (m *Meta) query(options map[string]string, command string, maxRows int, parameters map[string]string) []map[string]interface{} {
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

	return dbResponse.Rows
}

func (m *Meta) one(options map[string]string, command string, maxRows int, parameters map[string]string) map[string]interface{} {
	rows := m.query(options, command, maxRows, parameters)
	if len(rows) > 0 {
		return rows[0]
	}
	return nil
}
