package repository

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testStock = map[string]int32{"a": int32(1), "b": int32(2), "c": int32(3)}

func TestAddtoStock(t *testing.T) {
	testcase := []struct {
		name   string
		stock  map[string]int32
		input1 string
		input2 []int32

		expect struct {
			result  int32
			isError bool
		}
	}{
		{
			name:   "Happy case1",
			input1: "a",
			input2: []int32{},

			expect: struct {
				result  int32
				isError bool
			}{
				result:  1,
				isError: false,
			},
		}, {
			name:   "Happy case2",
			input1: "a",
			input2: []int32{2},

			expect: struct {
				result  int32
				isError bool
			}{
				result:  2,
				isError: false,
			},
		}, {
			name:   "Happy case3",
			stock:  map[string]int32{"a": 1},
			input1: "a",
			input2: []int32{2},

			expect: struct {
				result  int32
				isError bool
			}{
				result:  3,
				isError: false,
			},
		}, {
			name:   "Multiple value input error",
			input1: "a",
			input2: []int32{2, 3},

			expect: struct {
				result  int32
				isError bool
			}{
				result:  0,
				isError: true,
			},
		}, {
			name:   "Invalid value input error",
			input1: "a",
			input2: []int32{-5},

			expect: struct {
				result  int32
				isError bool
			}{
				result:  0,
				isError: true,
			},
		},
	}

	for _, v := range testcase {
		t.Run(v.name, func(t *testing.T) {
			repo := CreateBeefStock()
			if v.stock != nil {
				repo.Stock = v.stock
			}
			ret, err := repo.AddToStock(v.input1, v.input2...)

			if !v.expect.isError {
				assert.Nil(t, err)
				assert.Equal(t, ret, v.expect.result)
			} else {
				assert.NotNil(t, err)
			}
		})
	}
}

func TestRemoveFromStock(t *testing.T) {
	testcase := []struct {
		name   string
		stock  map[string]int32
		input1 string
		input2 []int32

		expect struct {
			result  int32
			isError bool
		}
	}{
		{
			name:   "Happy case1",
			stock:  map[string]int32{"a": int32(1), "b": int32(2), "c": int32(3)},
			input1: "a",
			input2: []int32{},

			expect: struct {
				result  int32
				isError bool
			}{
				result:  0,
				isError: false,
			},
		}, {
			name:   "Happy case2",
			stock:  map[string]int32{"a": int32(1), "b": int32(2), "c": int32(3)},
			input1: "c",
			input2: []int32{2},

			expect: struct {
				result  int32
				isError bool
			}{
				result:  1,
				isError: false,
			},
		}, {
			name:   "Multiple value input error",
			stock:  map[string]int32{"a": int32(1), "b": int32(2), "c": int32(3)},
			input1: "a",
			input2: []int32{2, 3},

			expect: struct {
				result  int32
				isError bool
			}{
				result:  0,
				isError: true,
			},
		}, {
			name:   "Invalid value input error",
			stock:  map[string]int32{"a": int32(1), "b": int32(2), "c": int32(3)},
			input1: "a",
			input2: []int32{-2},

			expect: struct {
				result  int32
				isError bool
			}{
				result:  0,
				isError: true,
			},
		}, {
			name:   "Stock not found error",
			stock:  map[string]int32{"a": int32(1), "b": int32(2), "c": int32(3)},
			input1: "g",
			input2: []int32{5},

			expect: struct {
				result  int32
				isError bool
			}{
				result:  0,
				isError: true,
			},
		}, {
			name:   "Stock not enough error",
			stock:  map[string]int32{"a": int32(1), "b": int32(2), "c": int32(3)},
			input1: "a",
			input2: []int32{2},

			expect: struct {
				result  int32
				isError bool
			}{
				result:  0,
				isError: true,
			},
		},
	}

	for _, v := range testcase {
		t.Run(v.name, func(t *testing.T) {
			repo := CreateBeefStock()
			repo.Stock = v.stock

			ret, err := repo.RemoveFromStock(v.input1, v.input2...)

			if !v.expect.isError {
				assert.Nil(t, err)
				assert.Equal(t, ret, v.expect.result)
			} else {
				assert.NotNil(t, err)
			}
		})
	}
}

func TestGetStock(t *testing.T) {
	testcase := []struct {
		name  string
		stock map[string]int32
		input []string

		expect struct {
			result  map[string]int32
			isError bool
		}
	}{
		{
			name:  "Happy case1",
			stock: map[string]int32{"a": int32(1), "b": int32(2), "c": int32(3)},
			input: []string{},

			expect: struct {
				result  map[string]int32
				isError bool
			}{
				result:  map[string]int32{"a": int32(1), "b": int32(2), "c": int32(3)},
				isError: false,
			},
		}, {
			name:  "Happy case2",
			stock: map[string]int32{"a": int32(1), "b": int32(2), "c": int32(3)},
			input: []string{"a", "b"},

			expect: struct {
				result  map[string]int32
				isError bool
			}{
				result:  map[string]int32{"a": int32(1), "b": int32(2)},
				isError: false,
			},
		}, {
			name:  "Unknown Beef error",
			stock: map[string]int32{"a": int32(1), "b": int32(2), "c": int32(3)},
			input: []string{"a", "b", "x"},

			expect: struct {
				result  map[string]int32
				isError bool
			}{
				result:  map[string]int32{},
				isError: true,
			},
		},
	}

	for _, v := range testcase {
		t.Run(v.name, func(t *testing.T) {
			repo := CreateBeefStock()
			repo.Stock = v.stock

			ret, err := repo.GetStock(v.input...)

			if !v.expect.isError {
				assert.Nil(t, err)
				assert.Equal(t, ret, v.expect.result)
			} else {
				assert.NotNil(t, err)
			}
		})
	}
}
