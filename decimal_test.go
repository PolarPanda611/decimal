/*
 * @Author: Daniel TAN
 * @Description:
 * @Date: 2021-03-08 18:21:48
 * @LastEditTime: 2021-07-27 18:40:24
 * @LastEditors: Daniel TAN
 * @FilePath: /fr-price-common-pkg/decimal/decimal_test.go
 */
package decimal

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"math"
	"math/big"
	"reflect"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// func Test11(t *testing.T) {
// 	a, _ := new(big.Rat).SetString("0")
// 	b := new(big.Rat).SetInt64(0)
// 	c := new(big.Rat).SetInt(new(big.Int).SetInt64(0))
// 	assert.Equal(t, b, c, "1 wrong ")
// 	assert.Equal(t, a, b, "2 wrong ")
// 	assert.Equal(t, a, c, "3 wrong ")

// }
func TestSum(t *testing.T) {
	assert.Equal(t, NewFromInt64(0), Sum(), "error ")

	assert.Equal(t, NewFromInt64(1800), Sum(
		NewFromInt64(300),
		NewFromInt64(400),
		NewFromInt64(500),
		NewFromInt64(600),
	), "error ")

	assert.Equal(t, NewFromInt64(1000), Sum(
		NewFromInt64(300),
		NewFromInt64(-400),
		NewFromInt64(500),
		NewFromInt64(600),
	), "error ")
}

func TestAvg(t *testing.T) {
	assert.Equal(t, NewFromInt64(0), Avg(), "error ")
	assert.Equal(t, NewFromInt64(5), Avg(
		NewFromInt64(2),
		NewFromInt64(4),
		NewFromInt64(6),
		NewFromInt64(8),
	), "error ")

	assert.Equal(t, NewFromInt64(250), Avg(
		NewFromInt64(300),
		NewFromInt64(-400),
		NewFromInt64(500),
		NewFromInt64(600),
	), "error ")
}

func TestAbs(t *testing.T) {
	a := NewFromInt64(-10)
	//check result
	assert.Equal(t, NewFromInt64(10), a.Abs(), "error ")
	result := a.Abs().value
	//check ptr
	assert.NotEqual(t, fmt.Sprintf("%p", &result), fmt.Sprintf("%p", &a.value), "error ")
	//check origin val
	assert.Equal(t, NewFromInt64(-10), a, "error ")
}

func TestAdd(t *testing.T) {
	a := NewFromInt64(-10)
	b := NewFromInt64(20)
	//check result
	assert.Equal(t, NewFromInt64(10), a.Add(b), "error ")
	result := a.Sub(b).value
	//check ptr
	assert.NotEqual(t, fmt.Sprintf("%p", &result), fmt.Sprintf("%p", &a.value), "error ")
	assert.NotEqual(t, fmt.Sprintf("%p", &result), fmt.Sprintf("%p", &b.value), "error ")
	//check origin val
	assert.Equal(t, NewFromInt64(-10), a, "error ")
	assert.Equal(t, NewFromInt64(20), b, "error ")

}

func TestSub(t *testing.T) {
	a := NewFromInt64(-10)
	b := NewFromInt64(20)
	assert.Equal(t, NewFromInt64(-30), a.Sub(b), "error ")
	result := a.Sub(b).value
	assert.NotEqual(t, fmt.Sprintf("%p", &result), fmt.Sprintf("%p", &a.value), "error ")
	assert.NotEqual(t, fmt.Sprintf("%p", &result), fmt.Sprintf("%p", &b.value), "error ")
	//check origin val
	assert.Equal(t, NewFromInt64(-10), a, "error ")
	assert.Equal(t, NewFromInt64(20), b, "error ")
}

func TestMul(t *testing.T) {
	a := NewFromInt64(-10)
	b := NewFromInt64(20)
	assert.Equal(t, NewFromInt64(-200), a.Mul(b), "error ")
	result := a.Mul(b)
	assert.NotEqual(t, fmt.Sprintf("%p", &result), fmt.Sprintf("%p", &a.value), "error ")
	assert.NotEqual(t, fmt.Sprintf("%p", &result), fmt.Sprintf("%p", &b.value), "error ")
	//check origin val
	assert.Equal(t, NewFromInt64(-10), a, "error ")
	assert.Equal(t, NewFromInt64(20), b, "error ")

}

func TestDiv(t *testing.T) {
	{
		a := NewFromInt64(20)
		b := NewFromInt64(2)
		assert.Equal(t, NewFromInt64(10), a.Div(b), "error ")
		result := a.Mul(b)
		assert.NotEqual(t, fmt.Sprintf("%p", &result), fmt.Sprintf("%p", &a.value), "error ")
		assert.NotEqual(t, fmt.Sprintf("%p", &result), fmt.Sprintf("%p", &b.value), "error ")
		//check origin val
		assert.Equal(t, NewFromInt64(20), a, "error ")
		assert.Equal(t, NewFromInt64(2), b, "error ")
	}
	{
		a, _ := NewFromString("300")
		b, _ := NewFromString("0.3")
		assert.Equal(t, NewFromBigInt(mustBigIntFromString("1000"), -1), a.Div(b), "error ")
		result := a.Mul(b)
		assert.NotEqual(t, fmt.Sprintf("%p", &result), fmt.Sprintf("%p", &a.value), "error ")
		assert.NotEqual(t, fmt.Sprintf("%p", &result), fmt.Sprintf("%p", &b.value), "error ")
		//check origin val
		aC, _ := NewFromString("300")
		bC, _ := NewFromString("0.3")
		assert.Equal(t, aC, a, "error ")
		assert.Equal(t, bC, b, "error ")
	}
}

func TestMod(t *testing.T) {
	{
		time.Now().Format("2006-01-02")
		a := NewFromInt(10)
		b := NewFromInt(3)
		assert.Equal(t, NewFromInt(1), a.Mod(b), "error ")
		result := a.Mod(b)
		assert.NotEqual(t, fmt.Sprintf("%p", &result), fmt.Sprintf("%p", &a.value), "error ")
		assert.NotEqual(t, fmt.Sprintf("%p", &result), fmt.Sprintf("%p", &b.value), "error ")
		//check origin val
		assert.Equal(t, NewFromInt64(10), a, "error ")
		assert.Equal(t, NewFromInt64(3), b, "error ")
	}
	{
		a := NewFromInt(-10)
		b := NewFromInt(3)
		assert.Equal(t, NewFromInt(2), a.Mod(b), "error ")
		result := a.Mod(b)
		assert.NotEqual(t, fmt.Sprintf("%p", &result), fmt.Sprintf("%p", &a.value), "error ")
		assert.NotEqual(t, fmt.Sprintf("%p", &result), fmt.Sprintf("%p", &b.value), "error ")
		//check origin val
		assert.Equal(t, NewFromInt64(-10), a, "error ")
		assert.Equal(t, NewFromInt64(3), b, "error ")
	}
	{
		a := NewFromFloat64(3451204593)
		b := NewFromFloat64(2454495034)
		assert.Equal(t, NewFromFloat64(996709559), a.Mod(b), "error ")
		result := a.Mod(b)
		assert.NotEqual(t, fmt.Sprintf("%p", &result), fmt.Sprintf("%p", &a.value), "error ")
		assert.NotEqual(t, fmt.Sprintf("%p", &result), fmt.Sprintf("%p", &b.value), "error ")
		//check origin val
		assert.Equal(t, NewFromFloat64(3451204593), a, "error ")
		assert.Equal(t, NewFromFloat64(2454495034), b, "error ")
	}

	{
		a, _ := NewFromString("24544.95034")
		b, _ := NewFromString("0.3451204593")
		r, _ := NewFromString("0.3283950433")
		assert.Equal(t, r.String(), a.Mod(b).String(), "error ")
		result := a.Mod(b)
		assert.NotEqual(t, fmt.Sprintf("%p", &result), fmt.Sprintf("%p", &a.value), "error ")
		assert.NotEqual(t, fmt.Sprintf("%p", &result), fmt.Sprintf("%p", &b.value), "error ")
		//check origin val
		aC, _ := NewFromString("24544.95034")
		bC, _ := NewFromString("0.3451204593")
		assert.Equal(t, aC, a, "error ")
		assert.Equal(t, bC, b, "error ")
	}
}

func TestNeg(t *testing.T) {
	{
		a := NewFromInt64(0)
		assert.Equal(t, NewFromInt64(0), a.Neg(), "error ")
		result := a.String()
		assert.NotEqual(t, fmt.Sprintf("%p", &result), fmt.Sprintf("%p", &a.value), "error ")
		//check origin val
		assert.Equal(t, NewFromInt64(0), a, "error ")
	}
	{
		a := NewFromInt64(1)
		assert.Equal(t, NewFromInt64(-1), a.Neg(), "error ")
		result := a.String()
		assert.NotEqual(t, fmt.Sprintf("%p", &result), fmt.Sprintf("%p", &a.value), "error ")
		//check origin val
		assert.Equal(t, NewFromInt64(1), a, "error ")
	}
	{
		a := NewFromInt64(-1)
		assert.Equal(t, NewFromInt64(1), a.Neg(), "error ")
		result := a.String()
		assert.NotEqual(t, fmt.Sprintf("%p", &result), fmt.Sprintf("%p", &a.value), "error ")
		//check origin val
		assert.Equal(t, NewFromInt64(-1), a, "error ")
	}
}

func TestDecimal_Floor(t *testing.T) {
	tests := []struct {
		name  string
		field Decimal
		want  Decimal
	}{

		{"1", NewRequiredFromString("1.999"), NewRequiredFromString("1")},
		{"2", NewRequiredFromString("1"), NewRequiredFromString("1")},
		{"3", NewRequiredFromString("1.01"), NewRequiredFromString("1")},
		{"4", NewRequiredFromString("0"), NewRequiredFromString("0")},
		{"5", NewRequiredFromString("0.9"), NewRequiredFromString("0")},
		{"6", NewRequiredFromString("0.1"), NewRequiredFromString("0")},
		{"7", NewRequiredFromString("-0.9"), NewRequiredFromString("-1")},
		{"8", NewRequiredFromString("-0.1"), NewRequiredFromString("-1")},
		{"9", NewRequiredFromString("-1.00"), NewRequiredFromString("-1")},
		{"10", NewRequiredFromString("-1.01"), NewRequiredFromString("-2")},
		{"11", NewRequiredFromString("-1.999"), NewRequiredFromString("-2")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := tt.field
			if got := d.Floor(); !reflect.DeepEqual(got.String(), tt.want.String()) {
				t.Errorf("Decimal.Floor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDecimal_Max(t *testing.T) {
	tests := []struct {
		name string
		d1   Decimal
		d2   Decimal
		want Decimal
	}{

		{"1", NewRequiredFromString("1.999"), NewRequiredFromString("1"), NewRequiredFromString("1.999")},
		{"2", NewRequiredFromString("1"), NewRequiredFromString("1"), NewRequiredFromString("1")},
		{"3", NewRequiredFromString("1.01"), NewRequiredFromString("1"), NewRequiredFromString("1.01")},
		{"4", NewRequiredFromString("0"), NewRequiredFromString("0"), NewRequiredFromString("0")},
		{"5", NewRequiredFromString("0.9"), NewRequiredFromString("0"), NewRequiredFromString("0.9")},
		{"6", NewRequiredFromString("0.1"), NewRequiredFromString("0"), NewRequiredFromString("0.1")},
		{"7", NewRequiredFromString("-0.9"), NewRequiredFromString("-1"), NewRequiredFromString("-0.9")},
		{"8", NewRequiredFromString("-0.1"), NewRequiredFromString("-1"), NewRequiredFromString("-0.1")},
		{"9", NewRequiredFromString("-1.00"), NewRequiredFromString("-1"), NewRequiredFromString("-1")},
		{"10", NewRequiredFromString("-1.01"), NewRequiredFromString("-2"), NewRequiredFromString("-1.01")},
		{"11", NewRequiredFromString("-1.999"), NewRequiredFromString("-2"), NewRequiredFromString("-1.999")},
		{"11", NewRequiredFromString("-2"), NewRequiredFromString("-1.999"), NewRequiredFromString("-1.999")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.d1.Max(tt.d2); !reflect.DeepEqual(got.String(), tt.want.String()) {
				t.Errorf("Decimal.Max() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDecimal_Min(t *testing.T) {
	tests := []struct {
		name string
		d1   Decimal
		d2   Decimal
		want Decimal
	}{

		{"1", NewRequiredFromString("1.999"), NewRequiredFromString("1"), NewRequiredFromString("1")},
		{"2", NewRequiredFromString("1"), NewRequiredFromString("1"), NewRequiredFromString("1")},
		{"3", NewRequiredFromString("1.01"), NewRequiredFromString("1"), NewRequiredFromString("1")},
		{"4", NewRequiredFromString("0"), NewRequiredFromString("0"), NewRequiredFromString("0")},
		{"5", NewRequiredFromString("0.9"), NewRequiredFromString("0"), NewRequiredFromString("0")},
		{"6", NewRequiredFromString("0.1"), NewRequiredFromString("0"), NewRequiredFromString("0")},
		{"7", NewRequiredFromString("-0.9"), NewRequiredFromString("-1"), NewRequiredFromString("-1")},
		{"8", NewRequiredFromString("-0.1"), NewRequiredFromString("-1"), NewRequiredFromString("-1")},
		{"9", NewRequiredFromString("-1.00"), NewRequiredFromString("-1"), NewRequiredFromString("-1")},
		{"10", NewRequiredFromString("-1.01"), NewRequiredFromString("-2"), NewRequiredFromString("-2")},
		{"11", NewRequiredFromString("-1.999"), NewRequiredFromString("-2"), NewRequiredFromString("-2")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.d1.Min(tt.d2); !reflect.DeepEqual(got.String(), tt.want.String()) {
				t.Errorf("Decimal.Min() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDecimal_StringFixed(t *testing.T) {
	tests := []struct {
		name string
		d    Decimal
		prec int32
		want string
	}{
		{"1", NewRequiredFromString("1"), 0, "1"},
		{"2", NewRequiredFromString("1.824"), 0, "2"},
		{"3", NewRequiredFromString("1.824"), 4, "1.8240"},
		{"4", NewRequiredFromString("1.424"), 0, "1"},
		{"5", NewRequiredFromString("1.524"), 0, "2"},
		{"6", NewRequiredFromString("1.954"), 1, "2.0"},
		{"7", NewRequiredFromString("-1.954"), 1, "-2.0"},
		{"8", NewRequiredFromString("0"), 1, "0.0"},
		{"9", NewRequiredFromString("-0"), 1, "0.0"},
		{"10", NewRequiredFromString("54.954"), -1, "50"},
		{"11", NewRequiredFromString("5466.954"), -3, "5000"},
		{"12", NewRequiredFromString("3.14124125231621521342351262156123231412352315"), 44, "3.14124125231621521342351262156123231412352315"},
		{"13", NewFromFloat64(3.14124125231621), 44, "3.14124125231621000000000000000000000000000000"},
		{"13", NewFromFloat64(3.14124125231621521342351262156123231412352315), 44, "3.14124125231621540000000000000000000000000000"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.d.StringFixed(tt.prec); got != tt.want {
				t.Errorf("Decimal.StringFixed() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDecimal_IsZero(t *testing.T) {
	tests := []struct {
		name string
		d    string
		want bool
	}{
		{"1", "1", false},
		{"2", "1.824", false},
		{"3", "1.424", false},
		{"4", "1.524", false},
		{"5", "1.954", false},
		{"6", "-1.954", false},
		{"7", "0", true},
		{"8", "-0", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := NewRequiredFromString(tt.d)
			if got := d.IsZero(); got != tt.want {
				t.Errorf("Decimal.IsZero() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDecimal_Sign(t *testing.T) {
	tests := []struct {
		name string
		d    string
		want int
	}{
		{"1", "1", 1},
		{"2", "1.824", 1},
		{"3", "1.424", 1},
		{"4", "1.524", 1},
		{"5", "1.954", 1},
		{"6", "-1.954", -1},
		{"7", "0", 0},
		{"8", "-0", 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := NewRequiredFromString(tt.d)
			if got := d.Sign(); got != tt.want {
				t.Errorf("Decimal.Sign() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDecimal_Pow(t *testing.T) {
	tests := []struct {
		name string
		d    string
		d2   int
		want Decimal
	}{
		{"1", "1", 2, NewRequiredFromString("1")},
		{"2", "3", 2, NewRequiredFromString("9")},
		{"3", "0", 0, NewRequiredFromString("1")},
		{"4", "3", 0, NewRequiredFromString("1")},
		{"5", "1.1", 2, NewRequiredFromString("1.21")},
		{"6", "3", 3, NewRequiredFromString("27")},
		{"7", "4", 5, NewRequiredFromString("1024")},
		{"8", "-1.3", 5, NewRequiredFromString("-3.71293")},
		{"9", "5", 4, NewRequiredFromString("625")},
		{"10", "4", -2, NewFromBigRat(new(big.Rat).SetFrac(mustBigIntFromString("1"), mustBigIntFromString("16")), 0)},
		{"11", "2", -2, NewFromBigRat(new(big.Rat).SetFrac(mustBigIntFromString("1"), mustBigIntFromString("4")), 0)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, NewRequiredFromString(tt.d).Pow(tt.d2), "Decimal.Pow() error ")
		})
	}
}

func TestDecimal_Int(t *testing.T) {
	tests := []struct {
		name string
		d    string
		want int
	}{
		{"1", "1", 1},
		{"2", "3", 3},
		{"3", "0", 0},
		{"4", "3", 3},
		{"5", "1.1", 1},
		{"6", "3", 3},
		{"7", "4", 4},
		{"8", "1.3", 1},
		{"9", "-1.3", -2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewRequiredFromString(tt.d).int(); got != tt.want {
				t.Errorf("Decimal.Int() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDecimal_Int32(t *testing.T) {
	tests := []struct {
		name string
		d    string
		want int32
	}{
		{"1", "1", 1},
		{"2", "3", 3},
		{"3", "0", 0},
		{"4", "3", 3},
		{"5", "1.1", 1},
		{"6", "3", 3},
		{"7", "4", 4},
		{"8", "1.3", 1},
		{"9", "-1.3", -2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewRequiredFromString(tt.d).int32(); got != tt.want {
				t.Errorf("Decimal.Int() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDecimal_Int64(t *testing.T) {
	tests := []struct {
		name string
		d    string
		want int64
	}{
		{"1", "1", 1},
		{"2", "3", 3},
		{"3", "0", 0},
		{"4", "3", 3},
		{"5", "1.1", 1},
		{"6", "3", 3},
		{"7", "4", 4},
		{"8", "1.3", 1},
		{"9", "-1.3", -2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewRequiredFromString(tt.d).int64(); got != tt.want {
				t.Errorf("Decimal.Int() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMax(t *testing.T) {
	type args struct {
		ds []Decimal
	}
	tests := []struct {
		name string
		args args
		want Decimal
	}{
		{name: "1", args: args{ds: []Decimal{NewRequiredFromString("1")}}, want: NewRequiredFromString("1")},
		{name: "2", args: args{ds: []Decimal{NewRequiredFromString("1"), NewRequiredFromString("2")}}, want: NewRequiredFromString("2")},
		{name: "3", args: args{ds: []Decimal{NewRequiredFromString("1"), NewRequiredFromString("-2")}}, want: NewRequiredFromString("1")},
		{name: "4", args: args{ds: []Decimal{NewRequiredFromString("-1"), NewRequiredFromString("-2")}}, want: NewRequiredFromString("-1")},
		{name: "5", args: args{ds: []Decimal{NewRequiredFromString("-1.1234"), NewRequiredFromString("-1.1235")}}, want: NewRequiredFromString("-1.1234")},
		{name: "6", args: args{ds: []Decimal{NewRequiredFromString("-1.1234"), NewRequiredFromString("-0"), NewRequiredFromString("-1.1235"), NewRequiredFromString("-1.53123")}}, want: NewRequiredFromString("0")},
		{name: "7", args: args{ds: []Decimal{NewRequiredFromString("-1.1234"), NewRequiredFromString("-0"), NewRequiredFromString("124215.5"), NewRequiredFromString("-121312")}}, want: NewRequiredFromString("124215.5")},
		{name: "8", args: args{ds: []Decimal{NewRequiredFromString("-1.1234"), NewRequiredFromString("-0"), NewRequiredFromString("-0.24"), NewRequiredFromString("1.124215")}}, want: NewRequiredFromString("1.124215")},
		{name: "9", args: args{ds: []Decimal{NewRequiredFromString("-1.1234"), NewRequiredFromString("-0"), NewRequiredFromString("-123.424"), NewRequiredFromString("-1.23211")}}, want: NewRequiredFromString("0")},
		{name: "10", args: args{ds: []Decimal{}}, want: NewRequiredFromString("0")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Max(tt.args.ds...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Max() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMin(t *testing.T) {
	type args struct {
		ds []Decimal
	}
	tests := []struct {
		name string
		args args
		want Decimal
	}{
		{name: "1", args: args{ds: []Decimal{NewRequiredFromString("1")}}, want: NewRequiredFromString("1")},
		{name: "2", args: args{ds: []Decimal{NewRequiredFromString("1"), NewRequiredFromString("2")}}, want: NewRequiredFromString("1")},
		{name: "3", args: args{ds: []Decimal{NewRequiredFromString("1"), NewRequiredFromString("-2")}}, want: NewRequiredFromString("-2")},
		{name: "4", args: args{ds: []Decimal{NewRequiredFromString("-1"), NewRequiredFromString("-2")}}, want: NewRequiredFromString("-2")},
		{name: "5", args: args{ds: []Decimal{NewRequiredFromString("-1.1234"), NewRequiredFromString("-1.1235")}}, want: NewRequiredFromString("-1.1235")},
		{name: "6", args: args{ds: []Decimal{NewRequiredFromString("-1.1234"), NewRequiredFromString("-0"), NewRequiredFromString("-1.1235"), NewRequiredFromString("-1.53123")}}, want: NewRequiredFromString("-1.53123")},
		{name: "7", args: args{ds: []Decimal{NewRequiredFromString("-1.1234"), NewRequiredFromString("-0"), NewRequiredFromString("124215.5"), NewRequiredFromString("-121312")}}, want: NewRequiredFromString("-121312")},
		{name: "8", args: args{ds: []Decimal{NewRequiredFromString("-1.1234"), NewRequiredFromString("-0"), NewRequiredFromString("-0.24"), NewRequiredFromString("1.124215")}}, want: NewRequiredFromString("-1.1234")},
		{name: "9", args: args{ds: []Decimal{NewRequiredFromString("-1.1234"), NewRequiredFromString("-0"), NewRequiredFromString("-123.424"), NewRequiredFromString("-1.23211")}}, want: NewRequiredFromString("-123.424")},
		{name: "10", args: args{ds: []Decimal{}}, want: NewRequiredFromString("0")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Min(tt.args.ds...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Min() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNew(t *testing.T) {
	// 0
	assert.Equal(t, NewFromInt(0), NewFromBigInt(new(big.Int).SetInt64(0), 0), " 0 value not same ")
	assert.Equal(t, Decimal{value: *new(big.Rat).SetInt(mustBigIntFromString("1000")), prec: -1}, NewFromBigInt(mustBigIntFromString("1000"), -1), " 0 value not same ")
	assert.Equal(t, NewFromInt(0), NewFromInt32(0), " 0 value not same ")
	assert.Equal(t, NewFromInt(0), NewFromInt64(0), " 0 value not same ")
	assert.Equal(t, NewFromInt(0), NewFromFloat32(0), " 0 value not same ")
	assert.Equal(t, NewFromInt(0), NewFromFloat64(0), " 0 value not same ")
	assert.Equal(t, NewFromInt(0), NewRequiredFromString("0"), " 0 value not same ")
	assert.Equal(t, NewFromInt(0), NewFromBigRat(big.NewRat(0, 1), 0), " 0 value not same ")
	x, _ := new(big.Rat).SetString("0")
	assert.NotEqual(t, NewFromInt(0), NewFromBigRat(x, 0), " 0 value not same ")
	// 1.5
	assert.Equal(t, NewRequiredFromString("1.5"), NewFromBigRat(big.NewRat(3, 2), -1), " 1.5 value not same ")
	assert.Equal(t, NewRequiredFromString("1.5"), NewFromFloat32(1.5), " 1.5 value not same ")
	assert.Equal(t, NewRequiredFromString("1.5"), NewFromFloat64(1.5), " 1.5 value not same ")
	// -1.5
	assert.Equal(t, NewRequiredFromString("-1.5"), NewFromBigRat(big.NewRat(-3, 2), -1), " 1.5 value not same ")
	assert.Equal(t, NewRequiredFromString("-1.5"), NewFromBigRat(big.NewRat(3, -2), -1), " 1.5 value not same ")
	assert.Equal(t, NewRequiredFromString("-1.5"), NewFromFloat32(-1.5), " 1.5 value not same ")
	assert.Equal(t, NewRequiredFromString("-1.5"), NewFromFloat64(-1.5), " 1.5 value not same ")
	y, _ := NewFromFloat64(3.14124125231621521342351262156123231412352315).Float64()
	assert.Equal(t, 3.14124125231621521342351262156123231412352315, y, " 1.5 value not same ")
	{
		_, err := NewFromString("-1.5")
		assert.Equal(t, nil, err, " 1.5 value not same ")
	}
	{
		_, err := NewFromString("-1.5d")
		assert.Error(t, err, " 1.5 value not same ")
	}
	{
		_, err := NewFromString("-1.5e123s123")
		assert.Error(t, err, " 1.5 value not same ")
	}
	{
		_, err := NewFromString("-1dasd.5")
		assert.Error(t, err, " 1.5 value not same ")
	}
	{
		_, err := NewFromString("a")
		assert.Equal(t, "strconv.ParseInt: parsing \"a\": invalid syntax", err.Error(), " wrong string convert")
	}
	assert.Panics(t, func() { NewRequiredFromString("a") }, " wrong required string convert ")
}
func TestDecimal_Float32(t *testing.T) {
	tests := []struct {
		name  string
		d     string
		wantR float32
		wantB bool
	}{
		{"1", "1", 1, true},
		{"2", "3", 3, true},
		{"3", "0", 0, true},
		{"4", "3", 3, true},
		{"5", "1.1", 1.1, false},
		{"6", "3", 3, true},
		{"7", "4", 4, true},
		{"8", "1.3", 1.3, false},
		{"9", "-1.3", -1.3, false},
		{"10", "-1.324125125212134124124", -1.324125125212134124124, false},
		// {"11", "1/3", 0.3333333333333333, false},
		// {"12", "-1/3", -0.3333333333333333, false},
		{"13", "3.14124125231621521342351262156123231412352315", 3.1412412523162154, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotR, gotB := NewRequiredFromString(tt.d).Float32()
			if gotR != tt.wantR {
				t.Errorf("Decimal.Float32() = %v, want %v", gotR, tt.wantR)
			}
			if gotB != tt.wantB {
				t.Errorf("Decimal.Float32() = %v, want %v", gotB, tt.wantB)
			}
		})
	}
}
func TestDecimal_Float64(t *testing.T) {
	tests := []struct {
		name  string
		d     string
		wantR float64
		wantB bool
	}{
		{"1", "1", 1, true},
		{"2", "3", 3, true},
		{"3", "0", 0, true},
		{"4", "3", 3, true},
		{"5", "1.1", 1.1, false},
		{"6", "3", 3, true},
		{"7", "4", 4, true},
		{"8", "1.3", 1.3, false},
		{"9", "-1.3", -1.3, false},
		{"10", "-1.324125125212134124124", -1.324125125212134124124, false},
		// {"11", "1/3", 0.3333333333333333, false},
		// {"12", "-1/3", -0.3333333333333333, false},
		{"13", "3.14124125231621521342351262156123231412352315", 3.1412412523162154, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotR, gotB := NewRequiredFromString(tt.d).Float64()
			if gotR != tt.wantR {
				t.Errorf("Decimal.Float64() = %v, want %v", gotR, tt.wantR)
			}
			if gotB != tt.wantB {
				t.Errorf("Decimal.Float64() = %v, want %v", gotB, tt.wantB)
			}
		})
	}
}

func TestString(t *testing.T) {
	tests := []struct {
		name string
		d    Decimal
		want string
	}{
		{"1", NewRequiredFromString("1"), "1"},
		{"2", NewRequiredFromString("6.6666666666666667"), "6.6666666666666667"},
		{"3", NewRequiredFromString("6666.6666666666666666667"), "6666.6666666666666666667"},
		{"4", NewRequiredFromString("3.141592653589793"), "3.141592653589793"},
		{"5", NewFromFloat64(3.141592653589793), "3.141592653589793"},
		{"6", NewFromFloat64(-3.141592653589793), "-3.141592653589793"},
		{"7", NewFromFloat64(-0.141592653589793), "-0.141592653589793"},
		{"8", Decimal{value: *new(big.Rat).SetInt64(4444444400), prec: -2}, "4444444400"},
		{"9", NewFromBigRat(new(big.Rat).SetFrac(mustBigIntFromString("1832"), mustBigIntFromString("1000")), -4), "1.832"},
		{"10", NewFromBigRat(new(big.Rat).SetFrac(mustBigIntFromString("1832"), mustBigIntFromString("1000")), -2), "1.83"},
		{"11", NewFromFrac(mustBigIntFromString("1832"), mustBigIntFromString("1000"), 0), "2"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.d.String(); got != tt.want {
				t.Errorf("Decimal.String() = %v, want %v", got, tt.want)
			}
		})
	}

}

func TestNewFromString(t *testing.T) {
	type args struct {
		d string
	}
	tests := []struct {
		name    string
		args    args
		want    Decimal
		wantErr bool
	}{
		{
			name: "1",
			args: args{
				d: "0.3",
			},
			want: Decimal{
				value: *new(big.Rat).SetFrac(mustBigIntFromString("3"), mustBigIntFromString("10")),
				prec:  -1,
			},
		},
		{
			name: "2",
			args: args{
				d: "3e-1",
			},
			want: Decimal{
				value: *new(big.Rat).SetFrac(mustBigIntFromString("3"), mustBigIntFromString("10")),
				prec:  -1,
			},
		},
		{
			name: "3",
			args: args{
				d: "3e5",
			},
			want: Decimal{
				value: *new(big.Rat).SetFrac(mustBigIntFromString("300000"), mustBigIntFromString("1")),
				prec:  0,
			},
		},
		{
			name: "4",
			args: args{
				d: "1.323232424e5",
			},
			want: Decimal{
				value: *new(big.Rat).SetFrac(mustBigIntFromString("1323232424"), mustBigIntFromString("10000")),
				prec:  -4,
			},
		},
		{
			name: "4",
			args: args{
				d: "-1.323232424e5",
			},
			want: Decimal{
				value: *new(big.Rat).SetFrac(mustBigIntFromString("-1323232424"), mustBigIntFromString("10000")),
				prec:  -4,
			},
		},
		{
			name: "5",
			args: args{
				d: "1.29067116156722e-309",
			},
			want: Decimal{
				value: *new(big.Rat).SetFrac(mustBigIntFromString("129067116156722"), mustBigIntFromString("100000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000")),
				prec:  -323,
			},
		},
		{
			name: "6",
			args: args{
				d: ".0e0",
			},
			want: Decimal{
				value: *new(big.Rat).SetFrac(mustBigIntFromString("0"), mustBigIntFromString("1")),
				prec:  -1,
			},
		},
		{
			name: "7",
			args: args{
				d: "5.8339553793802237e+23",
			},
			want: Decimal{
				value: *new(big.Rat).SetFrac(mustBigIntFromString("583395537938022370000000"), mustBigIntFromString("1")),
				prec:  0,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewFromString(tt.args.d)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewFromString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewFromString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_abs(t *testing.T) {
	type args struct {
		n int64
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "1",
			args: args{
				n: 1,
			},
			want: 1,
		},
		{
			name: "2",
			args: args{
				n: -1,
			},
			want: 1,
		},
		{
			name: "3",
			args: args{
				n: 0,
			},
			want: 0,
		},
		{
			name: "4",
			args: args{
				n: -1242141242151251515,
			},
			want: 1242141242151251515,
		},
		{
			name: "4",
			args: args{
				n: math.MaxInt64,
			},
			want: 9223372036854775807,
		},
		{
			name: "5",
			args: args{
				n: -math.MaxInt64,
			},
			want: 9223372036854775807,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := abs(tt.args.n); got != tt.want {
				t.Errorf("abs() = %v, want %v", got, tt.want)
			}
		})
	}
	assert.Panics(t, func() { abs(math.MinInt64) }, " range test ")
}

func Test_stringToBigInt(t *testing.T) {
	type args struct {
		value string
	}
	tests := []struct {
		name    string
		args    args
		want    *big.Int
		wantErr bool
	}{
		{
			name: "1",
			args: args{
				value: "1",
			},
			want: new(big.Int).SetInt64(1),
		},
		{
			name: "2",
			args: args{
				value: fmt.Sprintf("%v", math.MaxInt64),
			},
			want: new(big.Int).SetInt64(math.MaxInt64),
		},
		{
			name: "3",
			args: args{
				value: fmt.Sprintf("%v", math.MinInt64),
			},
			want: new(big.Int).SetInt64(math.MinInt64),
		},
		{
			name: "4",
			args: args{
				value: "-19223372036854775808",
			},
			want: mustBigIntFromString("-19223372036854775808"),
		},
		{
			name: "5",
			args: args{
				value: "-19223372036854775808",
			},
			want: mustBigIntFromString("-19223372036854775808"),
		},
		{
			name: "5",
			args: args{
				value: "-19223372s036854775808",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := stringToBigInt(tt.args.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("stringToBigInt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("stringToBigInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func mustBigIntFromString(str string) *big.Int {
	i, ok := new(big.Int).SetString(str, 10)
	if !ok {
		panic(fmt.Errorf("convert to big.Int error from string , %v", str))
	}
	return i

}

func Test_min(t *testing.T) {
	type args struct {
		x int32
		y int32
	}
	tests := []struct {
		name string
		args args
		want int32
	}{
		{
			name: "1",
			args: args{
				x: 1,
				y: 1,
			},
			want: 1,
		},
		{
			name: "2",
			args: args{
				x: 2,
				y: 1,
			},
			want: 1,
		},
		{
			name: "3",
			args: args{
				x: 1,
				y: 2,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := min(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("min() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDecimal_RoundDown(t *testing.T) {
	type args struct {
		d2 int32
	}
	tests := []struct {
		name   string
		fields Decimal
		args   args
		want   Decimal
	}{
		{
			name:   "1",
			fields: NewFromBigRat(new(big.Rat).SetFrac(mustBigIntFromString("2"), mustBigIntFromString("3")), 0),
			args: args{
				d2: 2,
			},
			want: NewFromBigRat(new(big.Rat).SetFrac(mustBigIntFromString("66"), mustBigIntFromString("100")), -2),
		},
		{
			name:   "2",
			fields: NewFromBigRat(new(big.Rat).SetFrac(mustBigIntFromString("200"), mustBigIntFromString("3")), 0),
			args: args{
				d2: -1,
			},
			want: NewFromBigRat(new(big.Rat).SetFrac(mustBigIntFromString("60"), mustBigIntFromString("1")), 1),
		},
		{
			name:   "3",
			fields: NewFromBigRat(new(big.Rat).SetFrac(mustBigIntFromString("-200"), mustBigIntFromString("3")), 0),
			args: args{
				d2: -1,
			},
			want: NewFromBigRat(new(big.Rat).SetFrac(mustBigIntFromString("-70"), mustBigIntFromString("1")), 1),
		},
		{
			name:   "4",
			fields: NewRequiredFromString("1.34"),
			args: args{
				d2: 1,
			},
			want: NewRequiredFromString("1.3"),
		},
		{
			name:   "5",
			fields: NewRequiredFromString("1.454"),
			args: args{
				d2: 0,
			},
			want: NewRequiredFromString("1"),
		},
		{
			name:   "6",
			fields: NewRequiredFromString("1.454"),
			args: args{
				d2: 1,
			},
			want: NewRequiredFromString("1.4"),
		},
		{
			name:   "7",
			fields: NewRequiredFromString("1.454"),
			args: args{
				d2: 2,
			},
			want: NewRequiredFromString("1.45"),
		},
		{
			name:   "8",
			fields: NewRequiredFromString("1.454"),
			args: args{
				d2: 3,
			},
			want: NewRequiredFromString("1.454"),
		},
		{
			name:   "9",
			fields: NewRequiredFromString("1.454"),
			args: args{
				d2: 4,
			},
			want: NewRequiredFromString("1.4540"),
		},
		{
			name:   "10",
			fields: NewRequiredFromString("1.454"),
			args: args{
				d2: -1,
			},
			want: NewFromBigInt(mustBigIntFromString("0"), 1),
		},
		{
			name:   "11",
			fields: NewRequiredFromString("1.454"),
			args: args{
				d2: -2,
			},
			want: NewFromBigInt(mustBigIntFromString("0"), 2),
		},
		{
			name:   "12",
			fields: NewRequiredFromString("-1.454"),
			args: args{
				d2: 0,
			},
			want: NewRequiredFromString("-2"),
		},
		{
			name:   "13",
			fields: NewRequiredFromString("-1.454"),
			args: args{
				d2: 1,
			},
			want: NewRequiredFromString("-1.5"),
		},
		{
			name:   "14",
			fields: NewRequiredFromString("-1.454"),
			args: args{
				d2: 2,
			},
			want: NewRequiredFromString("-1.46"),
		},
		{
			name:   "15",
			fields: NewRequiredFromString("-1.454"),
			args: args{
				d2: 3,
			},
			want: NewRequiredFromString("-1.454"),
		},
		{
			name:   "16",
			fields: NewRequiredFromString("-1.454"),
			args: args{
				d2: 4,
			},
			want: NewRequiredFromString("-1.4540"),
		},
		{
			name:   "17",
			fields: NewRequiredFromString("-1.454"),
			args: args{
				d2: -1,
			},
			want: NewFromBigInt(mustBigIntFromString("-10"), 1),
		},
		{
			name:   "18",
			fields: NewRequiredFromString("-1.454"),
			args: args{
				d2: -2,
			},
			want: NewFromBigInt(mustBigIntFromString("-100"), 2),
		},
		{
			name:   "19",
			fields: NewRequiredFromString("545"),
			args: args{
				d2: -1,
			},
			want: NewFromBigInt(mustBigIntFromString("540"), 1),
		},
		{
			name:   "20",
			fields: NewRequiredFromString("545"),
			args: args{
				d2: -2,
			},
			want: NewFromBigInt(mustBigIntFromString("500"), 2),
		},
		{
			name:   "21",
			fields: NewRequiredFromString("-545"),
			args: args{
				d2: -1,
			},
			want: NewFromBigInt(mustBigIntFromString("-550"), 1),
		},
		{
			name:   "22",
			fields: NewRequiredFromString("-545"),
			args: args{
				d2: -2,
			},
			want: NewFromBigInt(mustBigIntFromString("-600"), 2),
		},
		{
			name:   "23",
			fields: NewRequiredFromString("1.45"),
			args: args{
				d2: -2,
			},
			want: NewFromBigInt(mustBigIntFromString("0"), 2),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.fields.RoundDown(tt.args.d2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Decimal.RoundDown() = %v, want %v", got, tt.want)
			}
			if got := tt.fields.Round(RoundDown, tt.args.d2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Decimal.RoundDown() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDecimal_RoundUp(t *testing.T) {
	type args struct {
		d2 int32
	}
	tests := []struct {
		name   string
		fields Decimal
		args   args
		want   Decimal
	}{
		{
			name:   "1",
			fields: NewFromBigRat(new(big.Rat).SetFrac(mustBigIntFromString("2"), mustBigIntFromString("3")), 0),
			args: args{
				d2: 2,
			},
			want: NewFromBigRat(new(big.Rat).SetFrac(mustBigIntFromString("67"), mustBigIntFromString("100")), -2),
		},
		{
			name:   "2",
			fields: NewFromBigRat(new(big.Rat).SetFrac(mustBigIntFromString("200"), mustBigIntFromString("3")), 0),
			args: args{
				d2: -1,
			},
			want: NewFromBigRat(new(big.Rat).SetFrac(mustBigIntFromString("70"), mustBigIntFromString("1")), 1),
		},
		{
			name:   "3",
			fields: NewFromBigRat(new(big.Rat).SetFrac(mustBigIntFromString("-200"), mustBigIntFromString("3")), 0),
			args: args{
				d2: -1,
			},
			want: NewFromBigRat(new(big.Rat).SetFrac(mustBigIntFromString("-60"), mustBigIntFromString("1")), 1),
		},
		{
			name:   "4",
			fields: NewRequiredFromString("1.34"),
			args: args{
				d2: 1,
			},
			want: NewRequiredFromString("1.4"),
		},
		{
			name:   "5",
			fields: NewRequiredFromString("1.454"),
			args: args{
				d2: 0,
			},
			want: NewRequiredFromString("2"),
		},
		{
			name:   "6",
			fields: NewRequiredFromString("1.454"),
			args: args{
				d2: 1,
			},
			want: NewRequiredFromString("1.5"),
		},
		{
			name:   "7",
			fields: NewRequiredFromString("1.454"),
			args: args{
				d2: 2,
			},
			want: NewRequiredFromString("1.46"),
		},
		{
			name:   "8",
			fields: NewRequiredFromString("1.454"),
			args: args{
				d2: 3,
			},
			want: NewRequiredFromString("1.454"),
		},
		{
			name:   "9",
			fields: NewRequiredFromString("1.454"),
			args: args{
				d2: 4,
			},
			want: NewRequiredFromString("1.4540"),
		},
		{
			name:   "10",
			fields: NewRequiredFromString("1.454"),
			args: args{
				d2: -1,
			},
			want: NewFromBigInt(mustBigIntFromString("10"), 1),
		},
		{
			name:   "11",
			fields: NewRequiredFromString("1.454"),
			args: args{
				d2: -2,
			},
			want: NewFromBigInt(mustBigIntFromString("100"), 2),
		},
		{
			name:   "12",
			fields: NewRequiredFromString("-1.454"),
			args: args{
				d2: 0,
			},
			want: NewRequiredFromString("-1"),
		},
		{
			name:   "13",
			fields: NewRequiredFromString("-1.454"),
			args: args{
				d2: 1,
			},
			want: NewRequiredFromString("-1.4"),
		},
		{
			name:   "14",
			fields: NewRequiredFromString("-1.454"),
			args: args{
				d2: 2,
			},
			want: NewRequiredFromString("-1.45"),
		},
		{
			name:   "15",
			fields: NewRequiredFromString("-1.454"),
			args: args{
				d2: 3,
			},
			want: NewRequiredFromString("-1.454"),
		},
		{
			name:   "16",
			fields: NewRequiredFromString("-1.454"),
			args: args{
				d2: 4,
			},
			want: NewRequiredFromString("-1.4540"),
		},
		{
			name:   "17",
			fields: NewRequiredFromString("-1.454"),
			args: args{
				d2: -1,
			},
			want: NewFromBigInt(mustBigIntFromString("0"), 1),
		},
		{
			name:   "18",
			fields: NewRequiredFromString("-1.454"),
			args: args{
				d2: -2,
			},
			want: NewFromBigInt(mustBigIntFromString("0"), 2),
		},
		{
			name:   "19",
			fields: NewRequiredFromString("545"),
			args: args{
				d2: -1,
			},
			want: NewFromBigInt(mustBigIntFromString("550"), 1),
		},
		{
			name:   "20",
			fields: NewRequiredFromString("545"),
			args: args{
				d2: -2,
			},
			want: NewFromBigInt(mustBigIntFromString("600"), 2),
		},
		{
			name:   "21",
			fields: NewRequiredFromString("-545"),
			args: args{
				d2: -1,
			},
			want: NewFromBigInt(mustBigIntFromString("-540"), 1),
		},
		{
			name:   "22",
			fields: NewRequiredFromString("-545"),
			args: args{
				d2: -2,
			},
			want: NewFromBigInt(mustBigIntFromString("-500"), 2),
		},
		{
			name:   "22",
			fields: NewRequiredFromString("9.09"),
			args: args{
				d2: 0,
			},
			want: NewFromBigInt(mustBigIntFromString("10"), 0),
		},
		{
			name:   "23",
			fields: NewRequiredFromString("9.00000000000009"),
			args: args{
				d2: 0,
			},
			want: NewFromBigInt(mustBigIntFromString("10"), 0),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.fields.RoundUp(tt.args.d2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Decimal.RoundUp() = %v, want %v", got, tt.want)
			}
			if got := tt.fields.Round(RoundUp, tt.args.d2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Decimal.Round(RoundUp ) = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDecimal_RoundHalfUp(t *testing.T) {
	type args struct {
		d2 int32
	}
	tests := []struct {
		name   string
		fields Decimal
		args   args
		want   Decimal
	}{
		{
			name:   "1",
			fields: NewFromBigRat(new(big.Rat).SetFrac(mustBigIntFromString("2"), mustBigIntFromString("3")), 0),
			args: args{
				d2: 2,
			},
			want: NewFromBigRat(new(big.Rat).SetFrac(mustBigIntFromString("67"), mustBigIntFromString("100")), -2),
		},
		{
			name:   "2",
			fields: NewFromBigRat(new(big.Rat).SetFrac(mustBigIntFromString("200"), mustBigIntFromString("3")), 0),
			args: args{
				d2: -1,
			},
			want: NewFromBigRat(new(big.Rat).SetFrac(mustBigIntFromString("70"), mustBigIntFromString("1")), 1),
		},
		{
			name:   "3",
			fields: NewFromBigRat(new(big.Rat).SetFrac(mustBigIntFromString("-200"), mustBigIntFromString("3")), 0),
			args: args{
				d2: -1,
			},
			want: NewFromBigRat(new(big.Rat).SetFrac(mustBigIntFromString("-70"), mustBigIntFromString("1")), 1),
		},
		{
			name:   "4",
			fields: NewRequiredFromString("1.34"),
			args: args{
				d2: 1,
			},
			want: NewRequiredFromString("1.3"),
		},
		{
			name:   "5",
			fields: NewRequiredFromString("1.454"),
			args: args{
				d2: 0,
			},
			want: NewRequiredFromString("1"),
		},
		{
			name:   "6",
			fields: NewRequiredFromString("1.454"),
			args: args{
				d2: 1,
			},
			want: NewRequiredFromString("1.5"),
		},
		{
			name:   "7",
			fields: NewRequiredFromString("1.454"),
			args: args{
				d2: 2,
			},
			want: NewRequiredFromString("1.45"),
		},
		{
			name:   "8",
			fields: NewRequiredFromString("1.454"),
			args: args{
				d2: 3,
			},
			want: NewRequiredFromString("1.454"),
		},
		{
			name:   "9",
			fields: NewRequiredFromString("1.454"),
			args: args{
				d2: 4,
			},
			want: NewRequiredFromString("1.4540"),
		},
		{
			name:   "10",
			fields: NewRequiredFromString("1.454"),
			args: args{
				d2: -1,
			},
			want: NewFromBigInt(mustBigIntFromString("0"), 1),
		},
		{
			name:   "11",
			fields: NewRequiredFromString("1.454"),
			args: args{
				d2: -2,
			},
			want: NewFromBigInt(mustBigIntFromString("0"), 2),
		},
		{
			name:   "12",
			fields: NewRequiredFromString("-1.454"),
			args: args{
				d2: 0,
			},
			want: NewRequiredFromString("-1"),
		},
		{
			name:   "13",
			fields: NewRequiredFromString("-1.454"),
			args: args{
				d2: 1,
			},
			want: NewRequiredFromString("-1.5"),
		},
		{
			name:   "14",
			fields: NewRequiredFromString("-1.454"),
			args: args{
				d2: 2,
			},
			want: NewRequiredFromString("-1.45"),
		},
		{
			name:   "15",
			fields: NewRequiredFromString("-1.454"),
			args: args{
				d2: 3,
			},
			want: NewRequiredFromString("-1.454"),
		},
		{
			name:   "16",
			fields: NewRequiredFromString("-1.454"),
			args: args{
				d2: 4,
			},
			want: NewRequiredFromString("-1.4540"),
		},
		{
			name:   "17",
			fields: NewRequiredFromString("-1.454"),
			args: args{
				d2: -1,
			},
			want: NewFromBigInt(mustBigIntFromString("0"), 1),
		},
		{
			name:   "18",
			fields: NewRequiredFromString("-1.454"),
			args: args{
				d2: -2,
			},
			want: NewFromBigInt(mustBigIntFromString("0"), 2),
		},
		{
			name:   "19",
			fields: NewRequiredFromString("545"),
			args: args{
				d2: -1,
			},
			want: NewFromBigInt(mustBigIntFromString("550"), 1),
		},
		{
			name:   "20",
			fields: NewRequiredFromString("545"),
			args: args{
				d2: -2,
			},
			want: NewFromBigInt(mustBigIntFromString("500"), 2),
		},
		{
			name:   "21",
			fields: NewRequiredFromString("-545"),
			args: args{
				d2: -1,
			},
			want: NewFromBigInt(mustBigIntFromString("-550"), 1),
		},
		{
			name:   "22",
			fields: NewRequiredFromString("-545"),
			args: args{
				d2: -2,
			},
			want: NewFromBigInt(mustBigIntFromString("-500"), 2),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.fields.RoundHalfUp(tt.args.d2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Decimal.RoundHalfUp() = %v, want %v", got, tt.want)
			}
			if got := tt.fields.Round(RoundHalfUp, tt.args.d2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Decimal.RoundHalfUp() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDecimal_RoundTowardZero(t *testing.T) {
	type args struct {
		d2 int32
	}
	tests := []struct {
		name   string
		fields Decimal
		args   args
		want   Decimal
	}{
		{
			name:   "1",
			fields: NewFromBigRat(new(big.Rat).SetFrac(mustBigIntFromString("2"), mustBigIntFromString("3")), 0),
			args: args{
				d2: 2,
			},
			want: NewFromBigRat(new(big.Rat).SetFrac(mustBigIntFromString("66"), mustBigIntFromString("100")), -2),
		},
		{
			name:   "2",
			fields: NewFromBigRat(new(big.Rat).SetFrac(mustBigIntFromString("200"), mustBigIntFromString("3")), 0),
			args: args{
				d2: -1,
			},
			want: NewFromBigRat(new(big.Rat).SetFrac(mustBigIntFromString("60"), mustBigIntFromString("1")), 1),
		},
		{
			name:   "3",
			fields: NewFromBigRat(new(big.Rat).SetFrac(mustBigIntFromString("-200"), mustBigIntFromString("3")), 0),
			args: args{
				d2: -1,
			},
			want: NewFromBigRat(new(big.Rat).SetFrac(mustBigIntFromString("-60"), mustBigIntFromString("1")), 1),
		},
		{
			name:   "4",
			fields: NewRequiredFromString("1.34"),
			args: args{
				d2: 1,
			},
			want: NewRequiredFromString("1.3"),
		},
		{
			name:   "5",
			fields: NewRequiredFromString("1.454"),
			args: args{
				d2: 0,
			},
			want: NewRequiredFromString("1"),
		},
		{
			name:   "6",
			fields: NewRequiredFromString("1.454"),
			args: args{
				d2: 1,
			},
			want: NewRequiredFromString("1.4"),
		},
		{
			name:   "7",
			fields: NewRequiredFromString("1.454"),
			args: args{
				d2: 2,
			},
			want: NewRequiredFromString("1.45"),
		},
		{
			name:   "8",
			fields: NewRequiredFromString("1.454"),
			args: args{
				d2: 3,
			},
			want: NewRequiredFromString("1.454"),
		},
		{
			name:   "9",
			fields: NewRequiredFromString("1.454"),
			args: args{
				d2: 4,
			},
			want: NewRequiredFromString("1.4540"),
		},
		{
			name:   "10",
			fields: NewRequiredFromString("1.454"),
			args: args{
				d2: -1,
			},
			want: NewFromBigInt(mustBigIntFromString("0"), 1),
		},
		{
			name:   "11",
			fields: NewRequiredFromString("1.454"),
			args: args{
				d2: -2,
			},
			want: NewFromBigInt(mustBigIntFromString("0"), 2),
		},
		{
			name:   "12",
			fields: NewRequiredFromString("-1.454"),
			args: args{
				d2: 0,
			},
			want: NewRequiredFromString("-1"),
		},
		{
			name:   "13",
			fields: NewRequiredFromString("-1.454"),
			args: args{
				d2: 1,
			},
			want: NewRequiredFromString("-1.4"),
		},
		{
			name:   "14",
			fields: NewRequiredFromString("-1.454"),
			args: args{
				d2: 2,
			},
			want: NewRequiredFromString("-1.45"),
		},
		{
			name:   "15",
			fields: NewRequiredFromString("-1.454"),
			args: args{
				d2: 3,
			},
			want: NewRequiredFromString("-1.454"),
		},
		{
			name:   "16",
			fields: NewRequiredFromString("-1.454"),
			args: args{
				d2: 4,
			},
			want: NewRequiredFromString("-1.4540"),
		},
		{
			name:   "17",
			fields: NewRequiredFromString("-1.454"),
			args: args{
				d2: -1,
			},
			want: NewFromBigInt(mustBigIntFromString("0"), 1),
		},
		{
			name:   "18",
			fields: NewRequiredFromString("-1.454"),
			args: args{
				d2: -2,
			},
			want: NewFromBigInt(mustBigIntFromString("0"), 2),
		},
		{
			name:   "19",
			fields: NewRequiredFromString("545"),
			args: args{
				d2: -1,
			},
			want: NewFromBigInt(mustBigIntFromString("540"), 1),
		},
		{
			name:   "20",
			fields: NewRequiredFromString("545"),
			args: args{
				d2: -2,
			},
			want: NewFromBigInt(mustBigIntFromString("500"), 2),
		},
		{
			name:   "21",
			fields: NewRequiredFromString("-545"),
			args: args{
				d2: -1,
			},
			want: NewFromBigInt(mustBigIntFromString("-540"), 1),
		},
		{
			name:   "22",
			fields: NewRequiredFromString("-545"),
			args: args{
				d2: -2,
			},
			want: NewFromBigInt(mustBigIntFromString("-500"), 2),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.fields.RoundTowardZero(tt.args.d2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Decimal.RoundTowardZero() = %v, want %v", got, tt.want)
			}
			if got := tt.fields.Round(RoundTowardZero, tt.args.d2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Decimal.Round(RoundTowardZero) = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDecimal_Mod2(t *testing.T) {
	tests := []struct {
		name string
		d1   Decimal
		d2   Decimal
		want Decimal
	}{
		{
			name: "1",
			d1:   NewRequiredFromString("7.5"),
			d2:   NewRequiredFromString("2"),
			want: NewRequiredFromString("1.5"),
		},
		{
			name: "2",
			d1:   NewRequiredFromString("-7.5"),
			d2:   NewRequiredFromString("2"),
			want: NewRequiredFromString("0.5"),
		},
		{
			name: "3",
			d1:   NewRequiredFromString("-7.5"),
			d2:   NewRequiredFromString("-2"),
			want: NewRequiredFromString("-1.5"),
		},
		{
			name: "4",
			d1:   NewRequiredFromString("7.5"),
			d2:   NewRequiredFromString("-2"),
			want: NewRequiredFromString("-0.5"),
		},
		{
			name: "5",
			d1:   NewRequiredFromString("0.499999999999999999"),
			d2:   NewRequiredFromString("0.25"),
			want: NewRequiredFromString("0.249999999999999999"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.d1.Mod(tt.d2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Decimal.Mod() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDecimal_Rat(t *testing.T) {
	tests := []struct {
		name string
		d    *big.Rat
		want *big.Rat
	}{

		{
			name: "1",
			d:    new(big.Rat).SetInt64(1),
			want: new(big.Rat).SetInt64(1),
		},
		{
			name: "2",
			d:    new(big.Rat).SetInt64(-1),
			want: new(big.Rat).SetInt64(-1),
		},
		{
			name: "3",
			d:    new(big.Rat).SetInt64(0),
			want: new(big.Rat).SetInt64(0),
		},
		{
			name: "4",
			d:    new(big.Rat).SetInt64(1241241212),
			want: new(big.Rat).SetInt64(1241241212),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := NewFromBigRat(tt.d, 1)
			if got := d.Rat(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Decimal.Rat() = %v, want %v", got, tt.want)
			}
			assert.NotEqual(t, fmt.Sprintf("%p", tt.d), fmt.Sprintf("%p", d.Rat()))
		})
	}
}

func TestUnmarshalJSONNull(t *testing.T) {
	var doc struct {
		Amount Decimal `json:"amount"`
	}
	docStr := `{"amount": null}`
	err := json.Unmarshal([]byte(docStr), &doc)
	if err != nil {
		t.Errorf("error unmarshaling %s: %v", docStr, err)
	} else if doc.Amount.Cmp(NewFromInt(0)) != 0 {
		t.Errorf("expected Zero, got %s (%s, %d)",
			doc.Amount.String(),
			doc.Amount.value.String(), doc.Amount.prec)
	}
}

func TestDecimal_UnmarshalJSON(t *testing.T) {
	type test struct {
		Amount Decimal `json:"amount"`
	}
	type args struct {
		Json string
	}
	tests := []struct {
		name    string
		args    args
		want    Decimal
		wantErr bool
	}{
		{
			name: "1",
			args: args{
				Json: `{"amount": null}`,
			},
			want: NewRequiredFromString("0"),
		},
		{
			name: "2",
			args: args{
				Json: `{"amount": "null"}`,
			},
			want: NewRequiredFromString("0"),
		},
		{
			name: "3",
			args: args{
				Json: `{"amount": "1"}`,
			},
			want: NewRequiredFromString("1"),
		},
		{
			name: "4",
			args: args{
				Json: `{"amount": "\"1\""}`,
			},
			wantErr: true,
		},
		{
			name: "5",
			args: args{
				Json: `{"amount": "-1"}`,
			},
			want: NewRequiredFromString("-1"),
		},
		{
			name: "6",
			args: args{
				Json: `{"amount": "-1"}`,
			},
			want: NewRequiredFromString("-1"),
		},
		{
			name: "7",
			args: args{
				Json: `{"amount": "1.99999999999999"}`,
			},
			want: NewRequiredFromString("1.99999999999999"),
		},
		{
			name: "8",
			args: args{
				Json: `{"amount": "1125125e10"}`,
			},
			want: NewRequiredFromString("1125125e10"),
		},
		{
			name: "9",
			args: args{
				Json: `{"amount": 12412515}`,
			},
			want: NewRequiredFromString("12412515"),
		},
		{
			name: "10",
			args: args{
				Json: `{"amount": ""}`,
			},
			wantErr: true,
		},
		{
			name: "11",
			args: args{
				Json: `{"amount": "123"}`,
			},
			want: NewRequiredFromString("123"),
		},
		{
			name: "12",
			args: args{
				Json: `{"amount": 123}`,
			},
			want: NewRequiredFromString("123"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := new(test)
			if err := json.Unmarshal([]byte(tt.args.Json), res); (err != nil) != tt.wantErr {
				t.Errorf("Decimal.UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
			} else {
				assert.Equal(t, tt.want, res.Amount, "unmarshall wrong  , expected = %v , actual = %v ", tt.want, res.Amount)
			}
		})
	}
}

func TestDecimal_MarshalJSON(t *testing.T) {
	type test struct {
		Amount Decimal `json:"amount"`
	}
	type fields struct {
		Amount                   test
		MarshalJSONWithoutQuotes bool
	}
	tests := []struct {
		name    string
		fields  fields
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "1",
			fields: fields{
				Amount: test{
					Amount: NewRequiredFromString("123"),
				},
			},
			want:    []byte(`{"amount":"123"}`),
			wantErr: false,
		},
		{
			name: "2",
			fields: fields{
				Amount: test{
					Amount: NewRequiredFromString("123"),
				},
				MarshalJSONWithoutQuotes: true,
			},
			want:    []byte(`{"amount":123}`),
			wantErr: false,
		},
		{
			name: "3",
			fields: fields{
				Amount: test{
					Amount: NewRequiredFromString("1231241.123124125"),
				},
				MarshalJSONWithoutQuotes: true,
			},
			want:    []byte(`{"amount":1231241.123124125}`),
			wantErr: false,
		},
		{
			name: "4",
			fields: fields{
				Amount: test{
					Amount: NewRequiredFromString("1231241.123124125"),
				},
				MarshalJSONWithoutQuotes: false,
			},
			want:    []byte(`{"amount":"1231241.123124125"}`),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			MarshalJSONWithoutQuotes = tt.fields.MarshalJSONWithoutQuotes
			got, err := json.Marshal(tt.fields.Amount)
			if (err != nil) != tt.wantErr {
				t.Errorf("Decimal.MarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Decimal.MarshalJSON() = %v, want %v", string(got), string(tt.want))
			}
		})
	}
}

func TestDecimal_Scan(t *testing.T) {
	a := "124"
	type args struct {
		value interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    Decimal
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "1",
			args: args{
				value: int(123),
			},
			want: NewRequiredFromString("123"),
		},
		{
			name: "2",
			args: args{
				value: int32(123),
			},
			want: NewRequiredFromString("123"),
		},
		{
			name: "3",
			args: args{
				value: int64(123),
			},
			want: NewRequiredFromString("123"),
		},
		{
			name: "4",
			args: args{
				value: float32(123.12312),
			},
			want: NewRequiredFromString("123.12312"),
		},
		{
			name: "5",
			args: args{
				value: float64(123.123124),
			},
			want: NewRequiredFromString("123.123124"),
		},
		{
			name: "6",
			args: args{
				value: "123.123124",
			},
			want: NewRequiredFromString("123.123124"),
		},
		{
			name: "7",
			args: args{
				value: struct{}{},
			},
			wantErr: true,
		},
		{
			name: "8",
			args: args{
				value: &a,
			},
			wantErr: true,
		},
		{
			name: "9",
			args: args{
				value: []byte("12344"),
			},
			want: NewRequiredFromString("12344"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := Decimal{}
			if err := d.Scan(tt.args.value); (err != nil) != tt.wantErr {
				t.Errorf("Decimal.Scan() error = %v, wantErr %v", err, tt.wantErr)
			} else {
				assert.EqualValuesf(t, tt.want, d, "Decimal.Scan() value not equal , actual = %v ( %v , %v ), want %v ( %v , %v ) ", d, d.value, d.prec, tt.want, tt.want.value, tt.want.prec)
			}
		})
	}
}

func TestDecimal_Value(t *testing.T) {
	tests := []struct {
		name    string
		fields  Decimal
		want    driver.Value
		wantErr bool
	}{
		{
			name:   "1",
			fields: NewRequiredFromString("1"),
			want:   "1",
		},
		{
			name:   "2",
			fields: NewRequiredFromString("1.1241241"),
			want:   "1.1241241",
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.fields.Value()
			if (err != nil) != tt.wantErr {
				t.Errorf("Decimal.Value() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Decimal.Value() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDecimal_Round(t *testing.T) {
	type args struct {
		roundStrategy RoundStrategy
		d2            int32
	}
	tests := []struct {
		name         string
		fields       Decimal
		args         args
		want         Decimal
		wantPanic    bool
		wantPanicMsg string
	}{
		// TODO: Add test cases
		{
			name: "1",
			args: args{
				roundStrategy: RoundStrategy("1234"),
				d2:            2,
			},
			wantPanic:    true,
			wantPanicMsg: "unsupported rounding strategy, actual: 1234",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.wantPanic {
				assert.PanicsWithValue(t, tt.wantPanicMsg, func() { tt.fields.Round(tt.args.roundStrategy, tt.args.d2) }, "want panic error  ")
			} else {
				if got := tt.fields.Round(tt.args.roundStrategy, tt.args.d2); !reflect.DeepEqual(got, tt.want) {
					t.Errorf("Decimal.Round() = %v, want %v", got, tt.want)
				}
			}

		})
	}
}

func TestRoundStrategy_IsValid(t *testing.T) {
	tests := []struct {
		name    string
		tr      RoundStrategy
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:    "1",
			tr:      RoundDown,
			wantErr: false,
		},
		{
			name:    "1",
			tr:      RoundUp,
			wantErr: false,
		},
		{
			name:    "1",
			tr:      RoundHalfUp,
			wantErr: false,
		},
		{
			name:    "1",
			tr:      RoundTowardZero,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.tr.IsValid(); (err != nil) != tt.wantErr {
				t.Errorf("RoundStrategy.IsValid() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestClone(t *testing.T) {
	a := NewFromInt(0)
	b := a

	a = a.Add(NewFromInt(10))
	assert.Equal(t, NewFromInt(0), b, "wrong copy")
	assert.Equal(t, NewFromInt(10), a, "wrong source")
}
