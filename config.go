// Работа с конфигом
package metasdk

import (
	"encoding/json"
	"os"
)

const (
	devSettingsPath = "/.rwmeta/developer_settings.json"
	authCacheFile   = "/.rwmeta_auth_cache.json"
)

// чтение конфига в струтуру
func (ds *developerSettings) readConfig(path string) {
	file, err := os.Open(buildPath(path))
	check(err)

	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(ds)
	check(err)
}

// запись конфига
func (ds *developerSettings) writeConfig(path string) {
	config, err := json.Marshal(ds)
	check(err)

	file, err := os.OpenFile(buildPath(path), os.O_RDWR, 0644)
	check(err)

	defer file.Close()
	_, err = file.WriteAt(config, 0)
	check(err)
}

func (ds *developerSettings) readFromEnv() {
	value := os.Getenv("META_SERVICE_ACCOUNT_SECRET")
	if value == "" {
		value = os.Getenv("X-META-Developer-Settings")
	}
	if value != "" {
		err := json.Unmarshal([]byte(value), ds)
		check(err)
	}
}

func buildPath(path string) string {
	dir, _ := os.UserHomeDir()
	return dir + path
}

func readDeveloperSettings() *developerSettings {
	ds := &developerSettings{}
	ds.readConfig(devSettingsPath)
	ds.readFromEnv()
	return ds
}
