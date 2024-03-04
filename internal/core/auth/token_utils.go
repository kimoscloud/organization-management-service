package auth

import (
	"errors"
	"github.com/kimoscloud/organization-management-service/internal/core/model/dto"
)

func ValidateToken(signedToken string) (*dto.JWTClaim, error) {
	//TODO vaidate token againt the authentication server
	return nil, errors.New("invalid token")
}
