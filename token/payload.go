package token

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

var (
	ErrInvalidToken = errors.New("token is invalid")
	ErrExpiredToken = errors.New("token has expired")
)

type Payload struct {
	ID       uuid.UUID `json:"id"`
	Username string    `json:"username"`
	// Role      string    `json:"role"`
	IssuedAt  time.Time `json:"issued_at"`
	ExpiredAt time.Time `json:"expired_at"`
}

func NewPayload(username string, duration time.Duration) (*Payload, error) {
	tokenID, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}
	Payload := &Payload{
		ID:       tokenID,
		Username: username,
		// Role:      role,
		IssuedAt:  time.Now(),
		ExpiredAt: time.Now().Add(duration),
	}
	return Payload, nil
}

// 验证令牌是否失效
func (payload *Payload) Valid() error {
	if time.Now().After(payload.ExpiredAt) {
		return ErrExpiredToken
	}
	return nil
}

// 获取过期时间
func (payload *Payload) GetExpirationTime() (*jwt.NumericDate, error) {
	return &jwt.NumericDate{
		Time: payload.ExpiredAt,
	}, nil
}

// 获取注册时间
func (payload *Payload) GetIssuedAt() (*jwt.NumericDate, error) {
	return &jwt.NumericDate{
		Time: payload.IssuedAt,
	}, nil
}

// 获取发布时间
func (payload *Payload) GetNotBefore() (*jwt.NumericDate, error) {
	return &jwt.NumericDate{
		Time: payload.IssuedAt,
	}, nil
}

// 发布人
func (payload *Payload) GetIssuer() (string, error) {
	return "", nil
}

// 主题
func (payload *Payload) GetSubject() (string, error) {
	return "", nil
}

// 受众
func (payload *Payload) GetAudience() (jwt.ClaimStrings, error) {
	return jwt.ClaimStrings{}, nil
}
