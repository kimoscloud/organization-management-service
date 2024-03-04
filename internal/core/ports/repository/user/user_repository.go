package user

import (
	organization "github.com/kimoscloud/organization-management-service/internal/core/model/entity"
)

type Repository interface {
	GetByID(id string) (*organization.User, error)
	GetAllByEmail(emails []string) ([]organization.User, error)
}
