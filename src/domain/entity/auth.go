package entity

import (
	"encoding/json"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/ostafen/clover"
)

type Apikey struct {
	Id   uint64 `json:"id"`
	Name string `json:"name"`
	Key  string `json:"key"`
}

type UserAuth struct {
	Id                    string    `json:"_id"`
	UserId                string    `json:"userId"`
	AccessToken           string    `json:"accessToken"`
	RefreshToken          string    `json:"refreshToken"`
	AccessTokenExpiredAt  time.Time `json:"accessTokenExpiredAt"`
	RefreshTokenExpiredAt time.Time `json:"refreshTokenExpiredAt"`
}

func CreateUserAuth(userId string) UserAuth {
	id := uuid.New()
	return UserAuth{
		Id:     id.String(),
		UserId: userId,
	}
}

func (a UserAuth) RefreshTokenExpired(now time.Time) bool {
	return a.RefreshTokenExpiredAt.Before(now)
}

func (a UserAuth) AccessTokenExpired(now time.Time) bool {
	return a.AccessTokenExpiredAt.Before(now)
}

func (a UserAuth) MapStringInterface() (map[string]interface{}, error) {
	mapStringIntf := make(map[string]interface{})
	j, err := json.Marshal(a)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(j, &mapStringIntf)
	if err != nil {
		return nil, err
	}
	return mapStringIntf, nil
}

type FilterUserAuth struct {
	AccessToken string
	UserId      string
}

func (f *FilterUserAuth) Filter() *clover.Criteria {
	criterias := make([]*clover.Criteria, 0)

	if f.AccessToken != "" {
		criterias = append(criterias, clover.Field("accessToken").Eq(f.AccessToken))
	}

	if f.UserId != "" {
		criterias = append(criterias, clover.Field("userId").Eq(f.UserId))
	}

	filter := &clover.Criteria{}

	if len(criterias) == 0 {
		return nil
	}

	for idx, criteria := range criterias {
		if idx == 0 {
			filter = criteria
			continue
		}

		filter = filter.And(criteria)
	}

	return filter
}

type LocalUserAuthToken struct {
	Id   string    `json:"id"`
	Iat  time.Time `json:"iat"`
	Exp  time.Time `json:"exp"`
	Key  string    `json:"key"`
	Type TokenType `json:"tokenType"`
}

func (a LocalUserAuthToken) GenerateToken(key string) (string, error) {
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp":    a.Exp,
		"iat":    a.Iat,
		"userId": a.Id,
	})
	token, err := jwtToken.SignedString([]byte(key))
	if err != nil {
		return "", err
	}

	return token, nil
}

type TokenType string

var (
	AccessToken  TokenType = "accessToken"
	RefreshToken TokenType = "refreshToken"
)

type AuthenticationTokenType string

var (
	BearerAuthenticationToken AuthenticationTokenType = "Bearer"
)