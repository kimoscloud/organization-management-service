package user

import (
	"github.com/go-resty/resty/v2"
	organization "github.com/kimoscloud/organization-management-service/internal/core/model/entity"
	"github.com/kimoscloud/organization-management-service/internal/core/ports/repository/user"
	dto "github.com/kimoscloud/value-types/domain/auth/dto"
	"strings"
)

type RepositoryRest struct {
	client *resty.Client
}

func NewUserRepositoryRest() user.Repository {
	return &RepositoryRest{
		client: resty.New(),
	}
}

func (repo *RepositoryRest) ValidateToken(signedToken string) (*dto.JWTClaim, error) {
	//TODO implement me
	panic("implement me")
}

func (repo *RepositoryRest) GetByID(id string) (*organization.User, error) {
	//TODO implement me
	panic("implement me")
}

func (repo *RepositoryRest) GetAllByEmail(emails []string) ([]organization.User, error) {
	var users []organization.User
	_, err := repo.client.R().
		SetResult(&users).
		SetQueryParams(map[string]string{"emails": strings.Join(emails, ",")}).
		Get("/users")
	if err != nil {
		return nil, err
	}
	return users, nil
}
