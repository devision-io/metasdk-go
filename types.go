// для создания пользовательских типов
package metasdk

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
