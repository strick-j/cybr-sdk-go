package cybr

import (
	"strconv"
	"testing"
)

type mockOptions struct {
	Bool bool
	Str  string
}

func (m mockOptions) GetDisableHTTPS() bool {
	return m.Bool
}

func TestGetDisableHTTPS(t *testing.T) {
	cases := []struct {
		Options     []interface{}
		ExpectFound bool
		ExpectValue bool
	}{
		{
			Options: []interface{}{struct{}{}},
		},
		{
			Options: []interface{}{mockOptions{
				Bool: false,
			}},
			ExpectFound: true,
			ExpectValue: false,
		},
		{
			Options: []interface{}{mockOptions{
				Bool: true,
			}},
			ExpectFound: true,
			ExpectValue: true,
		},
		{
			Options:     []interface{}{struct{}{}, mockOptions{Bool: true}, mockOptions{Bool: false}},
			ExpectFound: true,
			ExpectValue: true,
		},
	}

	for i, tt := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			value, found := GetDisableHTTPS(tt.Options...)
			if found != tt.ExpectFound {
				t.Fatalf("expect value to not be found")
			}
			if value != tt.ExpectValue {
				t.Errorf("expect %v, got %v", tt.ExpectValue, value)
			}
		})
	}
}
