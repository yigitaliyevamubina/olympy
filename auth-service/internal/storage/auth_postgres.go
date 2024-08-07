package storage

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/Masterminds/squirrel"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"

	genprotos "olympy/auth-service/genproto/auth_service"
	"olympy/auth-service/internal/config"
)

type AuthService struct {
	db              *sql.DB
	queryBuilder    squirrel.StatementBuilderType
	jwtSecret       string
	accessTokenExp  time.Duration
	refreshTokenExp time.Duration
}

func NewAuthService(cfg *config.Config) (*AuthService, error) {
	db, err := ConnectDB(*cfg)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to DB: %v", err)
	}

	return &AuthService{
		db:              db,
		queryBuilder:    squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
		jwtSecret:       cfg.JWT.Secret,
		accessTokenExp:  cfg.JWT.AccessTokenExp,
		refreshTokenExp: cfg.JWT.RefreshTokenExp,
	}, nil
}

func (a *AuthService) RegisterUser(ctx context.Context, req *genprotos.RegisterUserRequest) (*genprotos.RegisterUserResponse, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("failed to hash password: %v", err)
	}

	data := map[string]interface{}{
		"username": req.Username,
		"password": string(hashedPassword),
		"role":     req.Role,
	}

	query, args, err := a.queryBuilder.Insert("users").
		SetMap(data).
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("failed to build SQL query: %v", err)
	}

	if _, err := a.db.ExecContext(ctx, query, args...); err != nil {
		return nil, fmt.Errorf("failed to execute SQL query: %v", err)
	}

	return &genprotos.RegisterUserResponse{
		User: &genprotos.User{
			Username: req.Username,
			Role:     req.Role,
		},
		Message: "User registered successfully",
	}, nil
}

func (a *AuthService) LoginUser(ctx context.Context, req *genprotos.LoginUserRequest) (*genprotos.LoginUserResponse, error) {
	var (
		hashedPassword string
		role           string
		userId         string
	)

	err := a.db.QueryRowContext(ctx, "SELECT id, password, role FROM users WHERE username = $1", req.Username).
		Scan(&userId, &hashedPassword, &role)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("username not found")
		}
		return nil, fmt.Errorf("failed to fetch user: %v", err)
	}

	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(req.Password))
	if err != nil {
		return nil, fmt.Errorf("invalid password")
	}

	accessToken, err := a.generateToken(userId, role, a.accessTokenExp)
	if err != nil {
		return nil, fmt.Errorf("failed to generate access token: %v", err)
	}

	refreshToken, err := a.generateToken(userId, role, a.refreshTokenExp)
	if err != nil {
		return nil, fmt.Errorf("failed to generate refresh token: %v", err)
	}

	return &genprotos.LoginUserResponse{
		User: &genprotos.User{
			Id:       userId,
			Username: req.Username,
			Role:     role,
		},
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		Message:      "Login successful",
	}, nil
}

func (a *AuthService) generateToken(userId, role string, expiration time.Duration) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userId,
		"role":    role,
		"exp":     time.Now().Add(expiration).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(a.jwtSecret))
}

func (a *AuthService) RefreshToken(ctx context.Context, req *genprotos.RefreshTokenRequest) (*genprotos.RefreshTokenResponse, error) {
	token, err := jwt.Parse(req.RefreshToken, func(token *jwt.Token) (interface{}, error) {
		return []byte(a.jwtSecret), nil
	})
	if err != nil || !token.Valid {
		return nil, fmt.Errorf("invalid refresh token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("invalid token claims")
	}

	userId, ok := claims["user_id"].(string)
	if !ok {
		return nil, fmt.Errorf("invalid user ID")
	}

	role, ok := claims["role"].(string)
	if !ok {
		return nil, fmt.Errorf("invalid role")
	}

	// Generate a new access token
	accessToken, err := a.generateToken(userId, role, a.accessTokenExp)
	if err != nil {
		return nil, fmt.Errorf("failed to generate new access token: %v", err)
	}

	return &genprotos.RefreshTokenResponse{
		AccessToken: accessToken,
	}, nil
}
