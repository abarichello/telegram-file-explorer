package main

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/enescakir/emoji"
	"github.com/mymmrac/telego"
	"github.com/stretchr/testify/assert"

	tutil "github.com/mymmrac/telego/telegoutil"
)

func makeRandomKeyboardButtons(n int) []telego.InlineKeyboardButton {
	var buttons []telego.InlineKeyboardButton
	for i := 0; i < n; i++ {
		button := tutil.InlineKeyboardButton(emoji.LeftArrow.String() + " " + strconv.Itoa(i))
		buttons = append(buttons, button)
	}
	return buttons
}

func TestMakeInlineKeyboard(t *testing.T) {
	type testCases struct {
		name         string
		n            int
		expectedGrid []int
	}
	testcases := []testCases{
		{
			name:         "return a line of maxColumns",
			n:            maxColumns,
			expectedGrid: []int{maxColumns},
		},
		{
			name:         "return a line of maxColumns and a second line with a lone button",
			n:            maxColumns + 1,
			expectedGrid: []int{maxColumns, 1},
		},
		{
			name:         "return a line of maxColumns and a second line with two buttons",
			n:            maxColumns + 2,
			expectedGrid: []int{maxColumns, 2},
		},
		{
			name:         "return two lines of maxColumns",
			n:            2 * maxColumns,
			expectedGrid: []int{maxColumns, maxColumns},
		},
		{
			name:         "return five lines of maxColumns and a sixth line with two buttons",
			n:            5*maxColumns + 2,
			expectedGrid: []int{maxColumns, maxColumns, maxColumns, maxColumns, maxColumns, 2},
		},
	}
	for _, test := range testcases {
		t.Run(test.name, func(t *testing.T) {
			buttons := makeRandomKeyboardButtons(test.n)
			result := makeInlineKeyboard(buttons)

			for i := 0; i < len(test.expectedGrid); i++ {
				t.Run(fmt.Sprintf("i: %d", i), func(t *testing.T) {
					assert.Equal(t, test.expectedGrid[i], len(result.InlineKeyboard[i]))
				})
			}
		})
	}
}
