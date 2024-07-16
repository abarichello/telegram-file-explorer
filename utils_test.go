package main

import (
	"math/rand"
	"strconv"
	"testing"

	"github.com/mymmrac/telego"
	"github.com/stretchr/testify/assert"
)

func TestIsUserAdmin(t *testing.T) {
	admin := rand.Int63()
	randomID := rand.Int63()
	adminID = strconv.FormatInt(admin, 10)

	type testCases struct {
		name     string
		id       int64
		expected bool
	}

	testcases := []testCases{
		{
			name:     "is admin",
			id:       admin,
			expected: true,
		},
		{
			name:     "is not an admin",
			id:       randomID,
			expected: false,
		},
	}
	for _, test := range testcases {
		t.Run(test.name, func(t *testing.T) {
			message := &telego.Message{
				From: &telego.User{
					ID: test.id,
				},
			}

			result := isUserAdmin(*message)
			assert.Equal(t, test.expected, result)
		})
	}
}
