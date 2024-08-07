package storage

import (
	"context"
	"database/sql"
	"log"
	"testing"

	"github.com/Masterminds/squirrel"
	"github.com/stretchr/testify/suite"

	genprotos "olympy/auth-service/genproto/auth_service"
	"olympy/auth-service/internal/config"
)

type AuthServiceTestSuite struct {
	suite.Suite
	service     *AuthService
	db          *sql.DB
	cleanupFunc func()
}

func (s *AuthServiceTestSuite) SetupSuite() {
	var err error

	configs, err := config.New()
	if err != nil {
		log.Fatal(err)
	}

	s.db, err = ConnectDB(*configs)
	s.Require().NoError(err)

	s.service = &AuthService{
		db:              s.db,
		queryBuilder:    squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
		jwtSecret:       configs.JWT.Secret,
		accessTokenExp:  configs.JWT.AccessTokenExp,
		refreshTokenExp: configs.JWT.RefreshTokenExp,
	}

	s.cleanupFunc = func() {
		s.db.Close()
	}
}

func (s *AuthServiceTestSuite) TearDownSuite() {
	s.cleanupFunc()
}

func (s *AuthServiceTestSuite) TestRegisterUser() {
	ctx := context.Background()

	req := &genprotos.RegisterUserRequest{
		Username: "fghjk",
		Password: "testpass",
		Role:     "user",
	}

	resp, err := s.service.RegisterUser(ctx, req)
	s.Require().NoError(err)
	s.Require().NotNil(resp)
	s.Equal(req.Username, resp.User.Username)
	s.Equal(req.Role, resp.User.Role)
}

func (s *AuthServiceTestSuite) TestLoginUser() {
	ctx := context.Background()

	// Register a user first
	req := &genprotos.RegisterUserRequest{
		Username: "88888",
		Password: "testpass",
		Role:     "user",
	}
	_, err := s.service.RegisterUser(ctx, req)
	s.Require().NoError(err)

	loginReq := &genprotos.LoginUserRequest{
		Username: "88888",
		Password: "testpass",
	}

	resp, err := s.service.LoginUser(ctx, loginReq)
	s.Require().NoError(err)
	s.Require().NotNil(resp)
	s.Equal(req.Username, resp.User.Username)
	s.Equal(req.Role, resp.User.Role)
	s.NotEmpty(resp.AccessToken)
	s.NotEmpty(resp.RefreshToken)
}

// func (s *AuthServiceTestSuite) TestRefreshToken() {
// 	ctx := context.Background()

// 	// Register and login a user first
// 	req := &genprotos.RegisterUserRequest{
// 		Username: "asdfghj",
// 		Password: "testpass",
// 		Role:     "user",
// 	}
// 	_, err := s.service.RegisterUser(ctx, req)
// 	s.Require().NoError(err)

// 	loginReq := &genprotos.LoginUserRequest{
// 		Username: "asdfghj",
// 		Password: "testpass",
// 	}
// 	loginResp, err := s.service.LoginUser(ctx, loginReq)
// 	s.Require().NoError(err)
// 	s.Require().NotNil(loginResp)

// 	refreshReq := &genprotos.RefreshTokenRequest{
// 		RefreshToken: loginResp.RefreshToken,
// 	}
// 	refreshResp, err := s.service.RefreshToken(ctx, refreshReq)
// 	s.Require().NoError(err)
// 	s.Require().NotNil(refreshResp)
// 	s.NotEmpty(refreshResp.AccessToken)
// }

func TestAuthServiceTestSuite(t *testing.T) {
	suite.Run(t, new(AuthServiceTestSuite))
}
