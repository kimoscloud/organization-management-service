package user

import (
	organization "github.com/kimoscloud/organization-management-service/internal/core/model/entity"
	userClient "github.com/kimoscloud/organization-management-service/internal/core/ports/client/user"
	"github.com/kimoscloud/organization-management-service/internal/core/ports/repository/user"
	dto "github.com/kimoscloud/value-types/domain/auth/dto"
)

type RepositoryRest struct {
	client userClient.Client
}

func NewUserRepositoryRest(client userClient.Client) user.Repository {
	return &RepositoryRest{
		client: client,
	}
}

func (repo *RepositoryRest) ValidateToken(signedToken string) (*dto.JWTClaim, error) {
	return repo.client.CallValidateToken(signedToken)
}

func (repo *RepositoryRest) GetByID(id string) (*organization.User, error) {
	//TODO implement me
	panic("implement me")
}

func (repo *RepositoryRest) GetAllByEmail(emails []string) ([]organization.User, error) {

	//var users []organization.User
	//_, err := repo.client
	//	//SetResult(&users).
	//	//SetQueryParams(map[string]string{"emails": strings.Join(emails, ",")}).
	//	//Get("/users")
	//if err != nil {
	//	return nil, err
	//}
	//return users, nil
	//TODO implement me
	return nil, nil
}
