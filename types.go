// для создания пользовательских типов
package metasdk

import (
	"log"
)

// Основной тип для работы с Мета хранит все данные для нормальной работы
type Meta struct {
	MetaURL           string
	ApiProxyURL       string
	DbName            string
	AuthUserId        int8
	logger            *log.Logger
	developerSettings *developerSettings
	defaultHeaders    map[string]string
	gcloudlog         string
	serviceNameSpace  string
	userAgent         string
	serviceId         string
	buildNum          string
	postfix           string
	settingsCache     map[string]map[string]interface{}
}

//тип для вычитывания настроек разработчика
type developerSettings struct {
	ApiHeaders   apiHeaders `json:"api_headers,omitempty"`
	GcloudDev    gcloudDev  `json:"gcloudDev,omitempty"`
	RefreshToken string     `json:"refreshToken,omitempty"`
	AccessToken  string     `json:"accessToken,omitempty"`
	ClientId     string     `json:"clientId,omitempty"`
	ClientSecret string     `json:"clientSecret,omitempty"`
}

// Данные для доступа в мету
type apiHeaders struct {
	Login string `json:"X-META-Developer-Login,omitempty"`
	Token string `json:"X-META-Developer-Token,omitempty"`
}

// Данные для доступа в gcloud
type gcloudDev struct {
	Project string `json:"project,omitempty"`
	Prefix  string `json:"prefix,omitempty"`
}

// Данные для запроса в базу данных через dbquery
type dbQuery struct {
	Database database `json:"database"`
	DbQuery  dbquery  `json:"dbQuery"`
}

type dbquery struct {
	MaxRows    int               `json:"maxRows"`
	Command    string            `json:"command"`
	Parameters map[string]string `json:"parameters"`
	ShardKey   interface{}       `json:"shard_key,omitempty"`
}
type database struct {
	Alias string `json:"alias"`
}

// Формат ответа от базы данных при запросе через dbquery
type DbResponse struct {
	MetaData interface{}              `json:"metaData"`
	Rows     []map[string]interface{} `json:"rows"`
}

type ApiProxyResponse struct {
	Error  ApiProxyError `json:"error"`
	Result interface{}   `json:"result"`
}

type ApiProxyError struct {
	Type        string `json:"type"`
	Message     string `json:"message"`
	WaitingTime int    `json:"waiting_time"`
}

type ExternalSystemSettings struct {
	ExSystemId string                 `json:"ex_system_id,omitempty"`
	Login      string                 `json:"login,omitempty"`
	TokenInfo  TokenInfo              `json:"token_info,omitempty"`
	FormData   map[string]interface{} `json:"form_data,omitempty"`
}

type TokenInfo struct {
	AccessToken  string `json:"accessToken,omitempty"`
	RefreshToken string `json:"refreshToken,omitempty"`
}

type FluentMsg struct {
	Message        string         `json:"message"`
	Context        string         `json:"context"`
	Severity       string         `json:"severity"`
	ServiceContext ServiceContext `json:"serviceContext"`
}

type ServiceContext struct {
	Service string `json:"service"`
	Version string `json:"version"`
}
