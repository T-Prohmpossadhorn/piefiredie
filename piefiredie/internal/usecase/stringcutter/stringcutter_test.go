package stringcutter

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetStock(t *testing.T) {
	testcase := []struct {
		name   string
		input  string
		cutter func(rune) bool

		expect struct {
			result []string
		}
	}{
		{
			name:  "Happy case1",
			input: "Fatback t-bone t-bone, pastrami t-bone.  pork, meatloaf jowl enim.  Bresaola t-bone.",
			cutter: func(v rune) bool {
				return !strings.Contains(",. ", string(v))
			},

			expect: struct {
				result []string
			}{
				result: []string{"Fatback", "t-bone", "t-bone", "pastrami", "t-bone", "pork", "meatloaf", "jowl", "enim", "Bresaola", "t-bone"},
			},
		},
	}

	for _, v := range testcase {
		t.Run(v.name, func(t *testing.T) {
			ret := make([]string, 0)
			StringCutter(v.input, v.cutter, func(v string) error {
				ret = append(ret, v)
				return nil
			})

			assert.Equal(t, ret, v.expect.result)
		})
	}
}
