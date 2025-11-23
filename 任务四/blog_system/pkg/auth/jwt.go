package auth

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWTService struct {
    secret     string
    expiration time.Duration
}

type Claims struct {
    UserID   uint   `json:"user_id"`
    Name string `json:"Name"`
    jwt.RegisteredClaims
}

func NewJWTService(secret string, expiration string) *JWTService {
    exp, err := time.ParseDuration(expiration)
    if err != nil {
        exp = 24 * time.Hour // 默认24小时
    }
    
    return &JWTService{
        secret:     secret,
        expiration: exp,
    }
}

func (s *JWTService) GenerateToken(userID uint, Name string) (string, error) {
    claims := &Claims{
        UserID:   userID,
        Name: Name,
        RegisteredClaims: jwt.RegisteredClaims{
            ExpiresAt: jwt.NewNumericDate(time.Now().Add(s.expiration)),
            IssuedAt:  jwt.NewNumericDate(time.Now()),
            Issuer:    "blog_system",
        },
    }
    
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString([]byte(s.secret))
}

func (s *JWTService) ValidateToken(tokenString string) (*Claims, error) {
    token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
        if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, errors.New("unexpected signing method")
        }
        return []byte(s.secret), nil
    })
    
    if err != nil {
        return nil, err
    }
    
    if claims, ok := token.Claims.(*Claims); ok && token.Valid {
        return claims, nil
    }
    
    return nil, errors.New("invalid token")
}