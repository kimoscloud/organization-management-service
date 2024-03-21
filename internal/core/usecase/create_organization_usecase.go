package organization

import (
	"github.com/kimoscloud/organization-management-service/internal/core/model/entity"
	request "github.com/kimoscloud/organization-management-service/internal/core/model/request"
	"github.com/kimoscloud/organization-management-service/internal/core/ports/logging"
	repository "github.com/kimoscloud/organization-management-service/internal/core/ports/repository"
	roleRepository "github.com/kimoscloud/organization-management-service/internal/core/ports/repository/role"
	userOrganizationRepository "github.com/kimoscloud/organization-management-service/internal/core/ports/repository/user-organization"
	"github.com/kimoscloud/organization-management-service/internal/core/utils"
	"github.com/kimoscloud/value-types/domain"
	"github.com/kimoscloud/value-types/errors"
	"time"
)

type CreateOrganizationUseCase struct {
	organizationRepository repository.Repository
	userOrganizationRepo   userOrganizationRepository.Repository
	roleRepo               roleRepository.Repository
	logger                 logging.Logger
}

func NewCreateOrganizationUseCase(
	organizationRepository repository.Repository,
	userOrganizationRepo userOrganizationRepository.Repository,
	roleRepo roleRepository.Repository,
	logger logging.Logger,
) *CreateOrganizationUseCase {
	return &CreateOrganizationUseCase{
		organizationRepository: organizationRepository,
		userOrganizationRepo:   userOrganizationRepo,
		roleRepo:               roleRepo,
		logger:                 logger,
	}
}

func (cu CreateOrganizationUseCase) Handler(
	userId string,
	request *request.CreateOrganizationRequest,
) (*organization.Organization, *errors.AppError) {
	tx := cu.organizationRepository.BeginTransaction()
	defer tx.Rollback()

	organizationResult, err := cu.organizationRepository.Create(
		&organization.Organization{
			Name:         request.Name,
			BillingEmail: request.BillingEmail,
			CreatedBy:    userId,
			Slug:         utils.CreateSlug(request.Name),
		},
		tx,
	)

	if err != nil {
		tx.Rollback()
		return nil, errors.NewInternalServerError(
			"Error getting user by email",
			"",
			errors.ErrorCreatingUser,
		).AppError
	}

	_, err = cu.userOrganizationRepo.Create(&organization.UserOrganization{
		OrganizationID: organizationResult.ID,
		UserID:         userId,
		RoleID:         domain.ORGANIZATION_ADMIN,
		//TODO send email to user with different template to the invite
		InvitedAt: time.Now(),
	}, tx)
	if err != nil {
		tx.Rollback()
		return nil, errors.NewInternalServerError(
			"Error creating user organization",
			"",
			errors.ErrorCreatingUser,
		).AppError
	}

	tx.Commit()
	return organizationResult, nil
}
