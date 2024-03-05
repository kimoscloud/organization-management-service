package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/kimoscloud/organization-management-service/internal/controller"
	logging2 "github.com/kimoscloud/organization-management-service/internal/core/ports/logging"
	organizationRepository "github.com/kimoscloud/organization-management-service/internal/core/ports/repository"
	roleRepository "github.com/kimoscloud/organization-management-service/internal/core/ports/repository/role"
	teamRepository "github.com/kimoscloud/organization-management-service/internal/core/ports/repository/team"
	teamMemberRepository "github.com/kimoscloud/organization-management-service/internal/core/ports/repository/team-member"
	userRepository "github.com/kimoscloud/organization-management-service/internal/core/ports/repository/user"
	userOrganizationRepository "github.com/kimoscloud/organization-management-service/internal/core/ports/repository/user-organization"
	organization "github.com/kimoscloud/organization-management-service/internal/core/usecase"
	"github.com/kimoscloud/organization-management-service/internal/infrastructure/configuration"
	"github.com/kimoscloud/organization-management-service/internal/infrastructure/db"
	"github.com/kimoscloud/organization-management-service/internal/infrastructure/logging"
	organizationRepositoryPostgres "github.com/kimoscloud/organization-management-service/internal/infrastructure/repository/postgres"
	roleRepositoryPostgres "github.com/kimoscloud/organization-management-service/internal/infrastructure/repository/postgres/role"
	teamRepositoryPostgres "github.com/kimoscloud/organization-management-service/internal/infrastructure/repository/postgres/team"
	teamMemberRepositoryPostgres "github.com/kimoscloud/organization-management-service/internal/infrastructure/repository/postgres/team-member"
	userOrganizationRepositoryPostgres "github.com/kimoscloud/organization-management-service/internal/infrastructure/repository/postgres/user-organization"
	userRepositoryRest "github.com/kimoscloud/organization-management-service/internal/infrastructure/repository/rest/user"
	"github.com/kimoscloud/organization-management-service/internal/infrastructure/server"
	"github.com/kimoscloud/organization-management-service/internal/middleware"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	if os.Getenv("ENV") == "dev" {
		err := godotenv.Load(".env")
		if err != nil {
			panic("Error loading .env file")
		}
	}
	// Create a new instance of the Gin router
	instance := gin.New()
	instance.Use(gin.Recovery())
	conn, err := db.NewConnection()
	if err != nil {
		log.Fatalf("failed to new database err=%s\n", err.Error())
	}
	logger, err := logging.NewLogger()
	if err != nil {
		log.Fatalf("failed to new logger err=%s\n", err.Error())
	}
	// Create the UserRepository
	userRepo := userRepositoryRest.NewUserRepositoryRest()
	orgRepo := organizationRepositoryPostgres.NewOrganizationRepository(conn)
	userOrgRepo := userOrganizationRepositoryPostgres.NewUserOrganizationRepository(conn)
	roleRepo := roleRepositoryPostgres.NewRoleRepository(conn)
	teamRepo := teamRepositoryPostgres.NewTeamRepository(conn)
	teamMemberRepo := teamMemberRepositoryPostgres.NewTeamMemberRepository(conn)

	initOrganizationController(
		instance,
		orgRepo,
		userOrgRepo,
		roleRepo,
		teamRepo,
		teamMemberRepo,
		userRepo,
		middleware.NewAuthMiddleware(userRepo),
		logger,
	)
	httpServer := server.NewHttpServer(
		instance,
		configuration.GetHttpServerConfig(),
	)
	httpServer.Start()
	defer httpServer.Stop()

	// Listen for OS signals to perform a graceful shutdown
	log.Println("listening signals...")
	c := make(chan os.Signal, 1)
	signal.Notify(
		c,
		os.Interrupt,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGQUIT,
		syscall.SIGTERM,
	)
	<-c
	log.Println("graceful shutdown...")
}

func initOrganizationController(
	instance *gin.Engine,
	orgRepo organizationRepository.Repository,
	userOrgRepo userOrganizationRepository.Repository,
	roleRepo roleRepository.Repository,
	teamRepo teamRepository.Repository,
	teamMemberRepo teamMemberRepository.Repository,
	userRepo userRepository.Repository,
	middleware *middleware.AuthMiddleware,
	logger logging2.Logger,
) {
	createOrganizationUseCase := organization.NewCreateOrganizationUseCase(
		orgRepo,
		userOrgRepo,
		roleRepo,
		logger,
	)
	getOrgByUserIdAndOrgIdUseCase := organization.NewGetOrganizationByOrgIdAndUserIdUseCase(
		orgRepo,
		logger,
	)
	getOrganizationsByUserIdUseCase := organization.NewGetOrganizationsByUserUseCase(
		orgRepo,
		logger,
	)

	checkIfUserHasEnoughPermissionsUseCase := organization.NewCheckUserHasPermissionsToMakeAction(
		userOrgRepo,
		logger)

	createOrganizationUserUseCase := organization.NewCreateOrganizationMemberUseCase(
		orgRepo,
		userOrgRepo,
		roleRepo,
		userRepo,
		checkIfUserHasEnoughPermissionsUseCase,
		logger,
	)
	removeOrganizationUserUseCase := organization.NewRemoveOrganizationMemberUseCase(
		orgRepo,
		userOrgRepo,
		logger,
	)

	createTeamUsecase := organization.NewCreateTeamUseCase(
		userOrgRepo,
		teamRepo,
		teamMemberRepo,
		checkIfUserHasEnoughPermissionsUseCase,
		logger)

	addMemberToTeamUseCase := organization.NewAddTeamMembersUseCase(
		userOrgRepo,
		teamRepo,
		teamMemberRepo,
		checkIfUserHasEnoughPermissionsUseCase,
		logger)

	organizationController := controller.NewOrganizationController(
		instance,
		logger,
		createOrganizationUseCase,
		getOrgByUserIdAndOrgIdUseCase,
		getOrganizationsByUserIdUseCase,
		createOrganizationUserUseCase,
		removeOrganizationUserUseCase,
		createTeamUsecase,
		addMemberToTeamUseCase,
		middleware,
	)
	organizationController.InitRouter()
}
