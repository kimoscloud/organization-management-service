package organization

import (
	organizationResponse "github.com/kimoscloud/organization-management-service/internal/core/model/response"
	"github.com/kimoscloud/organization-management-service/internal/core/ports/logging"
	repository "github.com/kimoscloud/organization-management-service/internal/core/ports/repository"
	"github.com/kimoscloud/value-types/errors"
)

type GetOrganizationsByUserUseCase struct {
	organizationRepository repository.Repository
	logger                 logging.Logger
}

func NewGetOrganizationsByUserUseCase(
	organizationRepository repository.Repository,
	logger logging.Logger,
) *GetOrganizationsByUserUseCase {
	return &GetOrganizationsByUserUseCase{
		organizationRepository: organizationRepository,
		logger:                 logger,
	}
}

func (cu GetOrganizationsByUserUseCase) Handler(
	userId string,
) ([]organizationResponse.OrganizationListLightElement, *errors.AppError) {
	organizationResult, err := cu.organizationRepository.GetAllByUserId(userId)
	if err != nil {
		return nil, errors.NewInternalServerError(
			"Error getting user organizations",
			"",
			errors.ErrorCreatingUser,
		).AppError
	}
	//Map result to response
	var response []organizationResponse.OrganizationListLightElement = make([]organizationResponse.OrganizationListLightElement, 0)
	for _, organization := range organizationResult {
		response = append(
			response, organizationResponse.OrganizationListLightElement{
				ID:       organization.ID,
				Name:     organization.Name,
				Slug:     organization.Slug,
				ImageUrl: organization.LogoURL,
			},
		)
	}
	return response, nil
}
