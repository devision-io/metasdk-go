package metasdk

import (
	"encoding/json"
)


func (m *Meta) DataGet(alias string, dataOnly, useCache bool) map[string]interface{} {

	result := make(map[string]interface{})
	if useCache {
		result = m.settingsCache[alias]
	}

	if result == nil {

		r := m.nativeCall("settings", "data/get/"+alias, "GET", []byte{})
		check(json.Unmarshal(r, &result))

		m.settingsCache = map[string]map[string]interface{}{alias: result}
	}
	if dataOnly {
		return result["form_data"].(map[string]interface{})
	}
	return result
}

func Flatten(m map[string]interface{}) map[string]interface{} {
	o := make(map[string]interface{})
	for k, v := range m {
		switch child := v.(type) {
		case map[string]interface{}:
			nm := Flatten(child)
			for nk, nv := range nm {
				o[k+"."+nk] = nv
			}
		default:
			o[k] = v
		}
	}
	return o
}


func (m *Meta) GetAccess(exAccessId string) *ExternalSystemSettings {
	saveNameDB := m.DbName
	m.DbName = "meta"
	defer func() {	m.DbName = saveNameDB }()
	ess := &ExternalSystemSettings{}
	cryptParams := m.DataGet("crypt_params", true, true)
	secureKey := cryptParams["secureKey"].(string)

	answer := m.One(
		"SELECT ex_system_id, login, token_info, form_data FROM meta.ex_access WHERE id=:id::uuid",
		map[string]string{"id": exAccessId})

	b, _  := json.Marshal(answer)
	check(json.Unmarshal(b, ess))

	if ess.TokenInfo.AccessToken != ""{
		ess.TokenInfo.AccessToken = DecodeJwt(ess.TokenInfo.AccessToken, secureKey)
	}

	if ess.TokenInfo.RefreshToken != "" {
		ess.TokenInfo.RefreshToken = DecodeJwt(ess.TokenInfo.RefreshToken, secureKey)
	}

	return ess
}
