package metasdk

import "encoding/json"

func (m *Meta) dataGet(alias string) map[string]interface{} {
	result := make(map[string]interface{})
	resp := m.nativeCall("settings", "data/get/"+alias, "GET", []byte{})
	check(json.Unmarshal(resp, result))
	m.settingsCache[alias] = result
	return result
}

func (m *Meta) ConfigGet(alias string, dataOnly, useCache bool) map[string]interface{} {

	result := make(map[string]interface{})
	if useCache {
		result = m.settingsCache[alias]
	}

	if result == nil {

		r := m.nativeCall("settings", "data/get/"+alias, "GET", []byte{})
		check(json.Unmarshal(r, result))
		m.settingsCache[alias] = result
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
