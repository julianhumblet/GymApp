package configfile

// Configuration struct
type Config struct {
	Webserver struct {
		Port int `json:"port"`
	} `json:"webserver"`
	Database struct {
		Host         string `json:"host"`
		Port         int    `json:"port"`
		Username     string `json:"username"`
		Password     string `json:"Password"`
		DatabaseName string `json:"databasename"`
	} `json:"database"`
	Session struct {
		SecretSessionKey string `json:"secretsessionkey"`
		SessionName      string `json:"sessionname"`
	} `json:"session"`
}
