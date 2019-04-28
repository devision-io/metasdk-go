// для создания пользовательских типов
package metasdk

// Основной тип для работы с Мета хранит все данные для нормальной работы
type Meta struct {
	MetaURL           string
	ApiProxyURL       string
	AuthUserId        int8
	developerSettings *developerSettings
	defaultHeaders    map[string]string
	serviceNameSpace  string
	serviceId         string
	buildNum          string
	postfix           string
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

type AutoGenerated struct {
	APIHeaders struct {
		XMETADeveloperLogin string `json:"X-META-Developer-Login"`
		XMETADeveloperToken string `json:"X-META-Developer-Token"`
	} `json:"api_headers"`
	GcloudDev struct {
		Project string `json:"project"`
		Prefix  string `json:"prefix"`
	} `json:"gcloudDev"`
	RefreshToken string `json:"refreshToken"`
	ClientID     string `json:"clientId"`
	ClientSecret string `json:"clientSecret"`
}

// Данные для запроса в базу данных через dbquery
type dbQuery struct {
	Database map[string]string      `json:"database"`
	DbQuery  map[string]interface{} `json:"dbQuery"`
}

// Формат ответа от базы данных при запросе через dbquery
type dbResponse struct {
	MetaData interface{}              `json:"metaData"`
	Rows     []map[string]interface{} `json:"rows"`
}
