package metasdk

import (
	"encoding/json"
	"log"
)

func (m *Meta) Query(command string, maxRows int, parameters map[string]string) DbResponse {
	dbQuery := dbQuery{
		Database: database{
			Alias: m.DbName,
		},
		DbQuery: dbquery{
			MaxRows:    maxRows,
			Command:    command,
			Parameters: parameters,
			ShardKey:   nil,
		},
	}

	req, err := json.Marshal(dbQuery)
	check(err)

	resp := m.nativeCall("db", "query", "POST", req)
	if resp == nil {
		log.Panic("Произошла ошибка при запросе в БД")
	}

	dbResponse := DbResponse{}
	check(json.Unmarshal(resp, &dbResponse))

	return dbResponse
}

func (m *Meta) One(command string, parameters map[string]string) map[string]interface{} {
	maxRows := 1
	rows := m.Query(command, maxRows, parameters).Rows
	if len(rows) > 0 {
		return rows[0]
	}
	return nil
}

func (m *Meta) All(command string, parameters map[string]string) []map[string]interface{} {
	maxRows := 0
	rows := m.Query(command, maxRows, parameters).Rows
	return rows
}
