package entity

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/ostafen/clover"
	"golang.org/x/crypto/bcrypt"
)

type UserType string

var (
	UserTypeRoot UserType = "root"
	UserTypeUser UserType = "user"
)

func NewUserType(userType string) (UserType, error) {
	var (
		resp UserType
		err  error
	)

	switch userType {
	case string(UserTypeRoot):
		resp = UserTypeRoot
	case string(UserTypeUser):
		resp = UserTypeUser
	default:
		err = errors.New("err: user type isn't exists")
	}

	return resp, err
}

type UserRbac struct {
	Create bool `json:"create"`
	Delete bool `json:"delete"`
	Update bool `json:"update"`
}

type User struct {
	Id       string    `json:"_id"`
	Username string    `json:"username"`
	Password string    `json:"password"`
	UserType UserType  `json:"userType"`
	Rbac     *UserRbac `json:"rbac"`
}

func (a User) Empty() bool {
	return a.Id == "" && a.Username == ""
}

func (a User) UserAdmin() bool {
	return a.UserType == UserTypeRoot
}

func (a *User) HasRbacAccess(method string) bool {
	hasAccess := false
	switch method {
	case "GET":
		hasAccess = true
	case "POST":
		hasAccess = a.Rbac.Create
	case "PUT", "PATCH":
		hasAccess = a.Rbac.Update
	case "DELETE":
		hasAccess = a.Rbac.Delete
	}

	return hasAccess
}

func (a *User) Update(u User) {
	a.Username = u.Username
}

func (a *User) PatchUserPassword(password string) error {
	a.Password = password
	if err := a.HashPassword(); err != nil {
		return err
	}
	return nil
}

func (a *User) HashPassword() error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(a.Password), 12)
	if err != nil {
		return err
	}
	a.Password = string(bytes)
	return nil
}

func (a *User) ComparePassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(a.Password), []byte(password))
}

func (a User) LocalUserAuthToken(tokenType TokenType, duration time.Duration) LocalUserAuthToken {
	now := time.Now()
	exp := now.Add(duration)

	return LocalUserAuthToken{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(exp),
			IssuedAt:  jwt.NewNumericDate(now),
		},
		Id:       a.Id,
		Type:     tokenType,
		UserType: string(a.UserType),
	}
}

func (a User) MapStringInterface() (map[string]interface{}, error) {
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

type Users []User

type FilterUser struct {
	Id       string
	Username string
	UserType UserType
}

func (f *FilterUser) Filter() *clover.Criteria {
	criterias := make([]*clover.Criteria, 0)

	if f.Id != "" {
		criterias = append(criterias, clover.Field("_id").Eq(f.Id))
	}

	if f.Username != "" {
		criterias = append(criterias, clover.Field("username").Eq(f.Username))
	}

	if f.UserType != "" {
		criterias = append(criterias, clover.Field("userType").Eq(f.UserType))
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
