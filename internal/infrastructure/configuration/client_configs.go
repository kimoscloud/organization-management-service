package configuration

import "os"

type ClientConfig struct {
	userServerUrl string
}

var clientConfiguration *ClientConfig

func GetClientConfig() *ClientConfig {
	if clientConfiguration == nil {
		initClientConfig()
	}
	return clientConfiguration
}

func initClientConfig() {
	clientConfiguration = &ClientConfig{}
	clientConfiguration.userServerUrl = os.Getenv("USER_SERVER_URL")

}

// Getters for ClientConfig

func (clientConfig *ClientConfig) GetUserServerUrl() string {
	return clientConfig.userServerUrl
}
