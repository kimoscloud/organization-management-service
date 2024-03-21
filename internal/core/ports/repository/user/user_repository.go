package user

import (
	organization "github.com/kimoscloud/organization-management-service/internal/core/model/entity"
	dto "github.com/kimoscloud/value-types/domain/auth/dto"
)

type Repository interface {
	GetByID(id string) (*organization.User, error)
	GetAllByEmail(emails []string) ([]organization.User, error)
	ValidateToken(signedToken string) (*dto.JWTClaim, error)
}
