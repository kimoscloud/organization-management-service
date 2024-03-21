package rest

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/kimoscloud/organization-management-service/internal/core/ports/logging"
	"github.com/kimoscloud/organization-management-service/internal/infrastructure/configuration"
	"github.com/kimoscloud/value-types/domain/auth/dto"
	"io"
	"net/http"
)

type UserClient struct {
	userServerUrl string
	logger        logging.Logger
}

func NewUserClient(config *configuration.ClientConfig, logger logging.Logger) *UserClient {
	return &UserClient{
		userServerUrl: config.GetUserServerUrl(),
		logger:        logger,
	}
}

func (c *UserClient) CallValidateToken(token string) (*dto.JWTClaim, error) {
	requestBody, err := json.Marshal(map[string]string{"token": token})
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", c.userServerUrl+"/api/v1/auth/validate-token", bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			c.logger.Error("error closing response body at validate token endpoint", err)
		}
	}(resp.Body)

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		c.logger.Error("error reading response body at validate token endpoint", err)
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("error validating token: " + string(responseBody))
	}

	var jwtClaims dto.JWTClaim
	err = json.Unmarshal(responseBody, &jwtClaims)
	if err != nil {
		return nil, err
	}

	return &jwtClaims, nil
}
