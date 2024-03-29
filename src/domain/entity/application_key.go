package entity

import (
	"encoding/json"
	"math/rand"
	"time"

	"github.com/ostafen/clover"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

type ApplicationKey struct {
	Id            string `json:"_id"`
	ApplicationId string `json:"applicationId"`
	Key           string `json:"key"`
}

func (r ApplicationKey) Exist() bool {
	return r.Id != "" && r.ApplicationId != "" && r.Key != ""
}

func (r *ApplicationKey) GenerateKey(length int) {
	var seededRand *rand.Rand = rand.New(
		rand.NewSource(time.Now().UnixNano()))
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	r.Key = string(b)
}

func (r ApplicationKey) MapStringInterface() (map[string]interface{}, error) {
	mapStringIntf := make(map[string]interface{})
	j, err := json.Marshal(r)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(j, &mapStringIntf)
	if err != nil {
		return nil, err
	}
	return mapStringIntf, nil
}

type ApplicationKeys []ApplicationKey

type FilterApplicationKey struct {
	SkipValidation bool
	ApplicationId  string
	Key            string
}

func (f FilterApplicationKey) Validation() bool {
	if f.SkipValidation {
		return true
	}
	return f.ApplicationId != ""
}

func (f FilterApplicationKey) Filter() *clover.Criteria {
	if !f.Validation() {
		return nil
	}
	criterias := make([]*clover.Criteria, 0)

	if f.ApplicationId != "" {
		criterias = append(criterias, clover.Field("applicationId").Eq(f.ApplicationId))
	}

	if f.Key != "" {
		criterias = append(criterias, clover.Field("key").Eq(f.Key))
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
