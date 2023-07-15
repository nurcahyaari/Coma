package dto_test

import (
	"testing"

	"github.com/coma/coma/src/domains/application/dto"
	"github.com/coma/coma/src/domains/application/model"
	"github.com/stretchr/testify/assert"
	"gopkg.in/guregu/null.v4"
)

func TestNewResponseGetClientConfiguration(t *testing.T) {
	testCases := []struct {
		name      string
		haveError bool
		expected  dto.ResponseGetConfigurationViewTypeJSON
		actual    func() (dto.ResponseGetConfigurationViewTypeJSON, error)
	}{
		{
			name:      "test1",
			haveError: false,
			expected: dto.ResponseGetConfigurationViewTypeJSON{
				ClientKey: "1",
				Data:      []byte(`{"APPLICATION":{"NAME":"coma","PORT":"3000"},"NAME":"coma"}`),
			},
			actual: func() (dto.ResponseGetConfigurationViewTypeJSON, error) {
				configuration := model.Configurations{
					{
						Id:        "1",
						ClientKey: "1",
						Field:     "APPLICATION",
						Value:     null.String{},
					},
					{
						Id:        "2",
						ClientKey: "1",
						Field:     "PORT",
						Value:     null.StringFrom("3000"),
					},
					{
						Id:        "3",
						ClientKey: "1",
						Field:     "NAME",
						Value:     null.StringFrom("coma"),
					},
					{
						Id:        "4",
						ClientKey: "1",
						Field:     "NAME",
						Value:     null.StringFrom("coma"),
					},
				}

				response := dto.NewResponseGetConfigurationViewTypeJSON("1")
				err := response.SetData(configuration)
				return response, err
			},
		},
		{
			name:      "test2",
			haveError: false,
			expected: dto.ResponseGetConfigurationViewTypeJSON{
				ClientKey: "1",
				Data:      []byte(`{"APPLICATION":{"NAME":"coma","PORT":"3000"}}`),
			},
			actual: func() (dto.ResponseGetConfigurationViewTypeJSON, error) {
				configuration := model.Configurations{
					{
						Id:        "1",
						ClientKey: "1",
						Field:     "APPLICATION",
						Value:     null.String{},
					},
					{
						Id:        "2",
						ClientKey: "1",
						Field:     "PORT",
						Value:     null.StringFrom("3000"),
					},
					{
						Id:        "3",
						ClientKey: "1",
						Field:     "NAME",
						Value:     null.StringFrom("coma"),
					},
					{
						Id:        "4",
						ClientKey: "1",
						Field:     "NAME",
						Value:     null.StringFrom("coma"),
					},
				}

				response := dto.NewResponseGetConfigurationViewTypeJSON("1")
				err := response.SetData(configuration)
				return response, err
			},
		},
		{
			name:      "test3",
			haveError: false,
			expected: dto.ResponseGetConfigurationViewTypeJSON{
				ClientKey: "1",
				Data:      []byte(`{"APPLICATION":{"INTERNAL":{"NAME":"coma","PORT":"3000"}}}`),
			},
			actual: func() (dto.ResponseGetConfigurationViewTypeJSON, error) {
				configuration := model.Configurations{
					{
						Id:        "1",
						ClientKey: "1",
						Field:     "APPLICATION",
						Value:     null.String{},
					},
					{
						Id:        "2",
						ClientKey: "1",
						Field:     "PORT",
						Value:     null.StringFrom("3000"),
					},
					{
						Id:        "3",
						ClientKey: "1",
						Field:     "NAME",
						Value:     null.StringFrom("coma"),
					},
					{
						Id:        "4",
						ClientKey: "1",
						Field:     "INTERNAL",
					},
				}

				response := dto.NewResponseGetConfigurationViewTypeJSON("1")
				err := response.SetData(configuration)
				return response, err
			},
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			actual, err := test.actual()

			if test.haveError {
				assert.Error(t, err)
			}

			assert.Equal(t, test.expected, actual)
		})
	}
}

func TestNewResponseGetConfiguration(t *testing.T) {
	testCases := []struct {
		name      string
		expected  dto.ResponseGetConfigurationViewTypeJSON
		haveError bool
		actual    func() (dto.ResponseGetConfigurationViewTypeJSON, error)
	}{
		{
			name:      "response from model.Configurations when empty",
			haveError: false,
			expected: dto.ResponseGetConfigurationViewTypeJSON{
				Data: []byte(nil),
			},
			actual: func() (dto.ResponseGetConfigurationViewTypeJSON, error) {
				response := dto.NewResponseGetConfigurationViewTypeJSON("")
				return response, nil
			},
		},
		{
			name:      "response from model.Configurations",
			haveError: false,
			expected: dto.ResponseGetConfigurationViewTypeJSON{
				ClientKey: "1",
				Data:      []byte(`{"age":"1","name":"test"}`),
			},
			actual: func() (dto.ResponseGetConfigurationViewTypeJSON, error) {
				response := dto.NewResponseGetConfigurationViewTypeJSON("")
				response.SetData(model.Configurations{
					{
						Id:        "1",
						ClientKey: "1",
						Field:     "name",
						Value:     null.StringFrom("test"),
					},
					{
						Id:        "2",
						ClientKey: "1",
						Field:     "age",
						Value:     null.StringFrom("1"),
					},
				})
				return response, nil
			},
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			actual, err := test.actual()

			if test.haveError {
				assert.Error(t, err)
			}

			assert.Equal(t, test.expected, actual)
		})
	}
}
