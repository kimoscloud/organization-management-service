package configuration

import "os"

type ClientConfig struct {
	userServerUrl    string
	userClientId     string
	userClientSecret string
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
	clientConfiguration.userClientId = os.Getenv("USER_CLIENT_ID")
	clientConfiguration.userClientSecret = os.Getenv("USER_CLIENT_SECRET")

}

// Getters for ClientConfig

func (clientConfig *ClientConfig) GetUserServerUrl() string {
	return clientConfig.userServerUrl
}

func (clientConfig *ClientConfig) GetUserClientId() string {
	return clientConfig.userClientId
}

func (clientConfig *ClientConfig) GetUserClientSecret() string {
	return clientConfig.userClientSecret
}
