package token

import (
	"BE-ecommerce-web-template/utils"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type JWTClaims struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Role     string `json:"role"`
	jwt.RegisteredClaims
}

var (
	API_SECRET string = utils.Getenv("API_SECRET", "secret_key")
)

func GenerateToken(user_id uint, username, role string) (string, error) {
	var (
		issuer = utils.Getenv("JWT_ISSUER", "kelompok2")
		method = jwt.SigningMethodHS256
	)

	token_lifespan, err := strconv.Atoi(utils.Getenv("TOKEN_HOUR_LIFESPAN", "24"))
	if err != nil {
		return "", err
	}

	claims := JWTClaims{
		ID:       user_id,
		Username: username,
		Role:     role,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    issuer,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * time.Duration(token_lifespan))),
		},
	}

	token := jwt.NewWithClaims(method, claims)
	return token.SignedString([]byte(API_SECRET))
}

func TokenValid(c *gin.Context) (*JWTClaims, error) {
	tokenString := ExtractToken(c)
	return getClaimsFromToken(tokenString)
}

func ExtractToken(c *gin.Context) string {
	token := c.Query("token")
	if token != "" {
		return token
	}
	bearerToken := c.Request.Header.Get("Authorization")
	if len(strings.Split(bearerToken, " ")) == 2 {
		return strings.Split(bearerToken, " ")[1]
	}
	return ""
}

func getClaimsFromToken(tokenString string) (*JWTClaims, error) {
	jwtToken, err := VerifyJWT(tokenString)
	if err != nil || !jwtToken.Valid {
		return nil, err
	}

	claims, ok := jwtToken.Claims.(*JWTClaims)
	if !ok {
		return nil, err
	}

	return claims, nil
}

func ExtractTokenID(c *gin.Context) (uint, error) {
	tokenString := ExtractToken(c)
	claims := &JWTClaims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(API_SECRET), nil
	})
	if err != nil {
		return 0, err
	}
	claims, ok := token.Claims.(*JWTClaims)
	if ok && token.Valid {
		return claims.ID, nil
	}
	return 0, nil
}

func VerifyJWT(tokenString string) (*jwt.Token, error) {
	claims := &JWTClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
		return []byte(API_SECRET), nil
	})
	return token, err
}

func ExtractClaims(c *gin.Context) (*JWTClaims, error) {
	authHeader := c.GetHeader("Authorization")
	tokenString := strings.ReplaceAll(authHeader, "Bearer ", "")
	token, err := VerifyJWT(tokenString)
	if err != nil {
		return nil, err
	}

	return token.Claims.(*JWTClaims), nil
}
