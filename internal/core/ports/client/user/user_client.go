package user

import "github.com/kimoscloud/value-types/domain/auth/dto"

type Client interface {
	CallValidateToken(token string) (*dto.JWTClaim, error)
}
