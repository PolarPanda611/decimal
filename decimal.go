/*
 * @Author: Daniel TAN
 * @Description:
 * @Date: 2021-03-08 18:21:48
 * @LastEditTime: 2021-07-29 16:59:48
 * @LastEditors: Daniel TAN
 * @FilePath: /fr-price-common-pkg/decimal/decimal.go
 */
package decimal

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"math"
	"math/big"
	"reflect"
	"strconv"
	"strings"
	"sync"
)

// DivisionPrecision is the number of decimal places in the result when it
// doesn't divide exactly.
//
// Example:
//
//     d1 := decimal.NewFromFloat(2).Div(decimal.NewFromFloat(3))
//     d1.String() // output: "0.6666666666666667"
//     d2 := decimal.NewFromFloat(2).Div(decimal.NewFromFloat(30000))
//     d2.String() // output: "0.0000666666666667"
//     d3 := decimal.NewFromFloat(20000).Div(decimal.NewFromFloat(3))
//     d3.String() // output: "6666.6666666666666667"
//     decimal.DivisionPrecision = 3
//     d4 := decimal.NewFromFloat(2).Div(decimal.NewFromFloat(3))
//     d4.String() // output: "0.667"
//
var (
	DivisionPrecision        int32     = 16
	MarshalJSONWithoutQuotes bool      = false
	oneInt                   *big.Int  = new(big.Int).SetInt64(1)
	fiveInt                  *big.Int  = new(big.Int).SetInt64(5)
	tenInt                   *big.Int  = new(big.Int).SetInt64(10)
	oneDecimal               Decimal   = NewFromInt(1)
	fiveDecimal              Decimal   = NewFromInt(5)
	tenDecimal               Decimal   = NewFromInt(10)
	builderPool              sync.Pool = sync.Pool{
		New: func() interface{} {
			return new(strings.Builder)
		},
	}
)

const (
	COMMA            = "."
	ZERO             = "0"
	ONE              = "1"
	EXPONENT_PATTERN = "eE"
)

type RoundStrategy string

const (
	RoundUp         RoundStrategy = "ROUND_UP"
	RoundDown       RoundStrategy = "ROUND_DOWN"
	RoundHalfUp     RoundStrategy = "ROUND_HALFUP"
	RoundTowardZero RoundStrategy = "ROUND_TOWARD_ZERO"
)

func (t *RoundStrategy) IsValid() error {
	switch *t {
	case RoundUp, RoundDown, RoundHalfUp, RoundTowardZero:
		return nil
	default:
		return errors.New(stringBuilder("unsupported rounding strategy, actual: ", string(*t)))
	}
}

type Decimal struct {
	// the presentation of the fraction
	// it always be an INT / INT
	value big.Rat
	// the precision of the fraction
	// like 0.1/0.3,if we present with fraction , it will be 1/3 , and we
	// don't know the precision of the actual number.
	// so we provide the precision to describe the accuracy
	// the prec is more like the exponent
	// -3 means 10e-3
	// 2 means 10e2
	prec int32
}

// NewFromBigRat new decimal from big.rat
// d : the fraction of the value
// prec : the precision of the fraction
// special case:
// because the Rat.SetString("0") will not set the Denom ,
// which will cause the difference as the Rat.SetInt64 and Rat.SetInt
// So be carefull when you want to initial a Decimal with 0 Rat
// x, _ := new(big.Rat).SetString("0")
// assert.NotEqual(t, NewFromInt(0), NewFromBigRat(x), " 0 value not same ")
func NewFromBigRat(value *big.Rat, prec int32) Decimal {
	return Decimal{
		value: *value,
		prec:  prec,
	}
}

// NewFromFrac new decimal from frac
// represent num/denom
func NewFromFrac(num, denom *big.Int, prec int32) Decimal {
	return Decimal{
		value: *new(big.Rat).SetFrac(num, denom),
		prec:  prec,
	}
}

// NewFromBigInt new decimal from big.int
// represent value/1 ,prec
func NewFromBigInt(value *big.Int, prec int32) Decimal {
	return Decimal{
		value: *new(big.Rat).SetInt(value),
		prec:  prec,
	}
}

// NewFromInt new decimal from int
// represent value/1 ,prec
func NewFromInt(value int) Decimal {
	return NewFromInt64(int64(value))
}

// NewFromInt32 new decimal from int
// represent value/1 ,prec
func NewFromInt32(value int32) Decimal {
	return NewFromInt64(int64(value))
}

// NewFromInt64 new decimal from int64
// represent value/1 ,prec
func NewFromInt64(value int64) Decimal {
	return Decimal{
		value: *new(big.Rat).SetInt64(value),
		prec:  0,
	}
}

// NewFromFloat64 new decimal from float64
// if float64 is more than 16 prec , please using Scientific notation
// the value will be represented by the fracion
func NewFromFloat64(value float64) Decimal {
	return NewRequiredFromString(fmt.Sprintf("%v", value))
}

// NewFromFloat32 new decimal from float32
// if float32 is more than 16 prec , please using Scientific notation
// the value will be represented by the fracion
func NewFromFloat32(value float32) Decimal {
	return NewRequiredFromString(fmt.Sprintf("%v", value))
}

// NewFromString new decimal from string
// the string support the science notation like "1234e10"
// TODO support the fraction like "123124/123"
func NewFromString(value string) (Decimal, error) {
	var intValue string = value
	var prec int64
	eIndex := strings.IndexAny(value, EXPONENT_PATTERN)
	if eIndex != -1 {
		intValue = value[:eIndex]
		e, err := strconv.Atoi(value[eIndex+1:])
		if err != nil {
			return Decimal{}, err
		}
		prec = int64(e)
	}

	strSlice := strings.Split(intValue, COMMA)
	if len(strSlice) == 1 {
		if prec > 0 {
			for i := 0; i < int(prec); i++ {
				strSlice[0] += ZERO
			}
			prec = 0
		}
		// if int
		i, err := stringToBigInt(strSlice[0])
		if err != nil {
			return Decimal{}, err
		}
		precStr := ONE + strings.Repeat(ZERO, int(abs(int64(prec))))
		precI, _ := stringToBigInt(precStr)
		return Decimal{
			value: *new(big.Rat).SetFrac(i, precI),
			prec:  int32(prec),
		}, nil
	} else {
		// if has precision
		intValue = strSlice[0] + strSlice[1]
		prec = int64(-len(strSlice[1]) + int(prec))
		if int64(prec) > 0 {
			intValue += strings.Repeat(ZERO, int(abs(int64(prec))))
			prec = 0
		}
		i, err := stringToBigInt(intValue)
		if err != nil {
			return Decimal{}, err
		}
		precStr := ONE
		if prec < 0 {
			precStr += strings.Repeat(ZERO, int(abs(int64(prec))))
		}
		precI, _ := stringToBigInt(precStr)
		return Decimal{
			value: *new(big.Rat).SetFrac(i, precI),
			prec:  int32(prec),
		}, nil
	}
}

// NewRequiredFromString new decimal from string
// this is a util function which help you to
// convert from string without error
// if got error ,will raise the panic
func NewRequiredFromString(d string) Decimal {
	r, err := NewFromString(d)
	if err != nil {
		panic(err)
	}
	return r
}

// Sum return the sum of ds
// implementation with Decimal.Add
func Sum(ds ...Decimal) Decimal {
	total := NewFromInt64(0)
	for i := 0; i < len(ds); i++ {
		total = total.Add(ds[i])
	}
	return total
}

// Avg return average of ds
// Implementation with Decimal.Add and Decimal.Div
// if len(ds) ==0, return 0
// if len(ds) >0 , return sum(ds)/len(ds)
func Avg(ds ...Decimal) Decimal {
	if len(ds) == 0 {
		return NewFromInt64(0)
	}
	return Sum(ds...).Div(NewFromInt(len(ds)))
}

// Max return max of ds
// Implementation with Decimal.Add and Decimal.Div
// if len(ds) ==0, return 0
// if len(ds) >0 , return max(ds)
func Max(ds ...Decimal) Decimal {
	if len(ds) == 0 {
		return NewFromInt64(0)
	}
	if len(ds) >= 2 {
		return ds[0].Max(Max(ds[1:]...))
	}
	return ds[0]
}

// Min return min of ds
// Implementation with Decimal.Add and Decimal.Div
// if len(ds) ==0, return 0
// if len(ds) >0 , return min(ds)
func Min(ds ...Decimal) Decimal {
	if len(ds) == 0 {
		return NewFromInt64(0)
	}
	if len(ds) >= 2 {
		return ds[0].Min(Min(ds[1:]...))
	}
	return ds[0]
}

// Abs returns the absolute value of the decimal  |d|
// implementation with Rat.Abs()
// e.g :
// "-1" => "1"
// "1" => "1"
// "0" => "0"
func (d Decimal) Abs() Decimal {
	return Decimal{value: *new(big.Rat).Abs(&d.value), prec: d.prec}
}

// Add returns d + d2.
// implementation with Rat.Add()
// the precision will get the minimal of the precision
// e.g :
// "-1" + "1" => "0"
// "1" + "1" => "2"
// "0" + "1" => "1"
// "1/3"+"1/2" => "5/6"
// ("1/4",-2) +( "1/2" , -1 ) = ("3/4",-2)
func (d Decimal) Add(d2 Decimal) Decimal {
	return Decimal{value: *new(big.Rat).Add(&d.value, &d2.value), prec: min(d.prec, d2.prec)}
}

// Sub returns d - d2.
// implementation with Rat.Sub()
// the precision will get the minimal of the precision
// e.g :
// "-1" - "1" => "-2"
// "1" - "1" => "0"
// "0" - "1" => "-1"
// "1/3"-"1/2" => "-1/6"
// ("1/4",-2) -( "1/2" , -1 ) = ("-1/4",-2)
func (d Decimal) Sub(d2 Decimal) Decimal {
	return Decimal{value: *new(big.Rat).Sub(&d.value, &d2.value), prec: min(d.prec, d2.prec)}
}

// Mul returns d * d2.
// implementation with Rat.Mul()
// the precision will get the sum of the precision
// e.g :
// "-1" * "1" => "-1"
// "1" * "1" => "1"
// "0" * "1" => "0"
// "1/3"*"1/2" => "1/6"
// ("1/4",-2) *( "1/2" , -1 ) = ("1/8",-3)
func (d Decimal) Mul(d2 Decimal) Decimal {
	return Decimal{value: *new(big.Rat).Mul(&d.value, &d2.value), prec: d.prec + d2.prec}
}

// Div returns d / d2.
// implementation with Rat.Quo()
// the precision will get the sum of the precision
// attention :
// because the division will have unlimited precision some times
// so the precision present with the summary of the precisions is not enough
// but the value will be accurately presented, so when we meet this case
// you can continue do the calculation, but when you try to get the data with string , or int ,
// please take care to handle the precision
// e.g :
// "-1" / "1" => "-1"
// "1" / "1" => "1"
// "0" / "1" => "0"
// "1/3"/"1/2" => "2/3"
// ("1/4",-2) / ( "1/2" , -1 ) = ("1/2",-3)
// ("1/3" , -3) / ("1/2" , -1 )  => ("2/3", -4 ), actual precision is unlimited , but we put -4 here .
func (d Decimal) Div(d2 Decimal) Decimal {
	return Decimal{value: *new(big.Rat).Quo(&d.value, &d2.value), prec: d.prec + d2.prec}
}

// Neg returns -d
// implementation with Rat.Neg()
// "-1".Neg()=> "1"
// "1".Neg() => "-1"
// "0".Neg() => "0"
func (d Decimal) Neg() Decimal {
	return Decimal{value: *new(big.Rat).Neg(&d.value), prec: d.prec}
}

// Floor return the Euclidean division of decimal
// use the num / denom
// implement with the big.Int.Div(Euclidean division)
// e.g :
// "0.9" => "0",
// "0.1" => "0",
// "-0.9" => "-1",
// "-0.1" => "-1",
// "-1.00" => "-1",
// "-1.01" => "-2",
//"-1.999" => "-2",
func (d Decimal) Floor() Decimal {
	return NewFromBigInt(new(big.Int).Div(d.value.Num(), d.value.Denom()), 0)
}

// Mod return d % d2
// implementation with the Euclidean division , then count the mod
// calculation steps :
// 1. x =  Euclidean division(d / d2)
// 2. y = x * d2
// 3. return d - y
// e.g :
// 7.5 % 2 = 1.5
// -7.5 % 2 = 0.5
// -7.5 % -2 = -1.5
// 7.5 % -2 = -0.5
// 0.499999999999999999 % 0.25 = 0.249999999999999999
func (d Decimal) Mod(d2 Decimal) Decimal {
	tm := d.Div(d2).Floor()
	tm2 := tm.Mul(d2)
	r := d.Sub(tm2)
	r.prec = min(d.prec, d2.prec)
	return r
}

// Pow return d^d2
// the Pow only support the int value for the moment
// 0 ^ 0 == 1
// 2 ^ 2 == 4
// 2 ^ -2 == 0.25
func (d Decimal) Pow(d2 int) Decimal {
	if d2 == 0 {
		return oneDecimal
	}
	tmp := d.Pow(d2 / 2)
	if d2%2 == 0 {
		return tmp.Mul(tmp)
	}
	if d2 > 0 {
		return tmp.Mul(tmp).Mul(d)
	}
	return tmp.Mul(tmp).Div(d)
}

// RoundDown return the rounddown of the decimal with precision
// reference docs
// https://en.wikipedia.org/wiki/Rounding#Rounding_down
// https://www.mathsisfun.com/numbers/rounding-methods.html
// attention: the negative number behavior between our implementation and shopspring decimal is difference.
// according to the docs , we thing the rounddown should be more like the floor instead of ceiling for negative number.
// 1.454.RoundDown(3)  1.454
// 1.454.RoundDown(2)  1.45
// 1.454.RoundDown(1)  1.4
// 1.454.RoundDown(0)  1
// 1.454.RoundDown(-1)  0
// 1.454.RoundDown(-2)  0
// 1.454.RoundDown(-3)  0
// -1.454.RoundDown(3)  -1.454
// -1.454.RoundDown(2)  -1.46
// -1.454.RoundDown(1)  -1.5
// -1.454.RoundDown(0)  -2
// -1.454.RoundDown(-1)  -10
// -1.454.RoundDown(-2)  -100
// -1.454.RoundDown(-3)  -1000
func (d Decimal) RoundDown(d2 int32) Decimal {
	var num *big.Int = new(big.Int)
	scale := new(big.Int).Exp(tenInt, big.NewInt(abs(int64(d2))), nil)
	if d2 < 0 {
		num.Mul(num.Div(num.Div(d.value.Num(), scale), d.value.Denom()), scale)
		scale = oneInt
	} else {
		num.Div(num.Mul(d.value.Num(), scale), d.value.Denom())
	}
	return Decimal{
		value: *new(big.Rat).SetFrac(num, scale),
		prec:  -d2,
	}
}

// RoundUp return the round up of the decimal with precision
// reference docs
// https://en.wikipedia.org/wiki/Rounding#Rounding_up
// https://www.mathsisfun.com/numbers/rounding-methods.html
// attention: the negative number behavior between our implementation and shopspring decimal is difference.
// 1.454.RoundUp(3)  1.454
// 1.454.RoundUp(2)  1.46
// 1.454.RoundUp(1)  1.5
// 1.454.RoundUp(0)  2
// 1.454.RoundUp(-1)  10
// 1.454.RoundUp(-2)  100
// 1.454.RoundUp(-3)  1000
// -1.454.RoundUp(3)  -1.454
// -1.454.RoundUp(2)  -1.45
// -1.454.RoundUp(1)  -1.4
// -1.454.RoundUp(0)  -1
// -1.454.RoundUp(-1)  0
// -1.454.RoundUp(-2)  0
// -1.454.RoundUp(-3)  0
func (d Decimal) RoundUp(d2 int32) Decimal {
	if d.Sign() > 0 {
		scale := d.RoundDown(d2)
		if d.Sub(scale).Sign() > 0 {
			newAdd := tenDecimal.Pow(int(-d2))
			newAdd.prec = -d2
			if d.Sign() > 0 {
				return scale.Add(newAdd)
			}
		}
		return scale
	} else {
		return d.Abs().RoundDown(d2).Neg()
	}

	// return d.Add(NewFromBigInt(nineInt, 0).Mul(scale)).RoundDown(d2)
}

// RoundHalfUp return the round half up of the decimal with precision
// reference docs
// https://en.wikipedia.org/wiki/Rounding#Rounding_half_up
// https://www.mathsisfun.com/numbers/rounding-methods.html
// attention: the negative number behavior between our implementation and shopspring decimal is difference.
// 1.454.RoundHalfUp(3)  1.454
// 1.454.RoundHalfUp(2)  1.45
// 1.454.RoundHalfUp(1)  1.5
// 1.454.RoundHalfUp(0)  1
// 1.454.RoundHalfUp(-1)  0
// 1.454.RoundHalfUp(-2)  0
// 1.454.RoundHalfUp(-3)  0
// -1.454.RoundHalfUp(3)  -1.454
// -1.454.RoundHalfUp(2)  -1.45
// -1.454.RoundHalfUp(1)  -1.5
// -1.454.RoundHalfUp(0)  -1
// -1.454.RoundHalfUp(-1)  0
// -1.454.RoundHalfUp(-2)  0
// -1.454.RoundHalfUp(-3)  0
func (d Decimal) RoundHalfUp(d2 int32) Decimal {
	sign := d.Sign()
	var scaleExp int32 = d2 + 1
	d = d.Abs().RoundDown(scaleExp)
	scale := tenDecimal.Pow(int(-scaleExp))
	if sign > 0 {
		return d.Add(fiveDecimal.Mul(scale)).RoundDown(d2)
	}
	return d.Add(fiveDecimal.Mul(scale)).RoundDown(d2).Neg()
}

// RoundTowardZero return the round toward zero of the decimal with precision
// reference docs
// https://en.wikipedia.org/wiki/Rounding#Rounding_toward_zero
// https://www.mathsisfun.com/numbers/rounding-methods.html
// attention:this function is more like the shopspring/decimal RoundDown
// 1.454.RoundTowardZero(3)  1.454
// 1.454.RoundTowardZero(2)  1.45
// 1.454.RoundTowardZero(1)  1.4
// 1.454.RoundTowardZero(0)  1
// 1.454.RoundTowardZero(-1)  0
// 1.454.RoundTowardZero(-2)  0
// 1.454.RoundTowardZero(-3)  0
// -1.454.RoundTowardZero(3)  -1.454
// -1.454.RoundTowardZero(2)  -1.45
// -1.454.RoundTowardZero(1)  -1.4
// -1.454.RoundTowardZero(0)  -1
// -1.454.RoundTowardZero(-1)  0
// -1.454.RoundTowardZero(-2)  0
// -1.454.RoundTowardZero(-3)  0
func (d Decimal) RoundTowardZero(d2 int32) Decimal {
	sign := d.Sign()
	if sign < 0 {
		return d.Abs().RoundDown(d2).Neg()
	}
	var num *big.Int = new(big.Int)
	scale := new(big.Int).Exp(tenInt, big.NewInt(abs(int64(d2))), nil)
	if d2 < 0 {
		num.Mul(num.Div(num.Div(d.value.Num(), scale), d.value.Denom()), scale)
		scale = oneInt
	} else {
		num.Div(num.Mul(d.value.Num(), scale), d.value.Denom())
	}
	return Decimal{
		value: *new(big.Rat).SetFrac(num, scale),
		prec:  -d2,
	}
}

// Round
// Round with the specific rounding strategy
// if unsupported rounding strategy , it will rase panic
func (d Decimal) Round(roundStrategy RoundStrategy, d2 int32) Decimal {
	switch roundStrategy {
	case RoundDown:
		return d.RoundDown(d2)
	case RoundUp:
		return d.RoundUp(d2)
	case RoundHalfUp:
		return d.RoundHalfUp(d2)
	case RoundTowardZero:
		return d.RoundTowardZero(d2)
	default:
		panic(roundStrategy.IsValid().Error())
	}
}

// String return string of the value
// the value will return the string represent the -precision value
// by default , it using the rounding half up
// so it is better for you to do the rounding as your expectation before String
func (d Decimal) String() string {
	return d.RoundHalfUp(-d.prec).string(true)
}

// string return the string represent the decimal , using the prec
// the trimZero is the option to enable if we need to trim the zero and the comma if necessary
func (d Decimal) string(trimZero bool) string {
	absExp := abs(int64(d.prec))
	scale := new(big.Int).Exp(tenInt, new(big.Int).SetInt64(absExp), nil)
	afterScale := new(big.Int).Mul(d.value.Num(), scale)
	str := new(big.Int).Abs(new(big.Int).Div(afterScale, d.value.Denom())).String()
	dExpInt := int(d.prec)
	var intPart, fractionalPart string
	if int64(len(str)) > absExp {
		intPart = str[:len(str)-int(absExp)]
		fractionalPart = str[len(str)-int(absExp):]
	} else {
		intPart = ZERO
		num0s := -dExpInt - len(str)
		fractionalPart = strings.Repeat(ZERO, num0s) + str
	}
	number := intPart
	if len(fractionalPart) > 0 && d.prec < 0 {
		number += COMMA + fractionalPart
	}
	if d.Sign() < 0 {
		number = "-" + number
	}
	if trimZero && strings.Contains(number, COMMA) {
		return strings.TrimRight(strings.TrimRight(number, ZERO), COMMA)
	}
	return number
}

// StringFixed  string fix with prec
// the value will return the string represent the -precision value
// by default , it using the rounding half up
// so it is better for you to do the rounding as your expectation before StringFixed
// NewFromFloat(0).StringFixed(2) // output: "0.00"
// NewFromFloat(0).StringFixed(0) // output: "0"
// NewFromFloat(5.45).StringFixed(0) // output: "5"
// NewFromFloat(5.45).StringFixed(1) // output: "5.5"
// NewFromFloat(5.45).StringFixed(2) // output: "5.45"
// NewFromFloat(5.45).StringFixed(3) // output: "5.450"
// NewFromFloat(545).StringFixed(-1) // output: "550"
func (d Decimal) StringFixed(prec int32) string {
	if d.prec > -prec {
		return NewFromBigRat(&d.value, -prec).string(false)
	}
	v := d.RoundHalfUp(prec).value
	return NewFromBigRat(&v, -prec).string(false)
}

// Cmp compares the numbers represented by d and d2 and returns:
// 1 if d > d2
// 0 if d == d2
// -1 if d < d2
func (d Decimal) Cmp(d2 Decimal) int {
	return d.value.Cmp(&d2.value)
}

// Max return the max of d and d2
// implementation with Decimal.Cmp
func (d Decimal) Max(d2 Decimal) Decimal {
	if d.Cmp(d2) >= 0 {
		return d
	}
	return d2
}

// Min return the min of d and d2
// implementation with Decimal.Cmp
func (d Decimal) Min(d2 Decimal) Decimal {
	if d.Cmp(d2) <= 0 {
		return d
	}
	return d2
}

// IsZero return whether the numbers represented by d and 0 are equal
// true if d == 0
// false if d != 0
func (d Decimal) IsZero() bool {
	return d.Cmp(NewRequiredFromString(ZERO)) == 0
}

// Sign return the sign of the value
// -1 if d <  0
//  0 if d == 0
// +1 if d >  0
// implementation with Rat.Sign()
func (d Decimal) Sign() int {
	return d.value.Sign()
}

// Int return the Euclidean division of the rat
// Div sets z to the quotient x/y for y != 0 and returns z. If y == 0, a division-by-zero run-time panic occurs.
// Div implements Euclidean division
func (d Decimal) int() int {
	return int(new(big.Int).Div(d.value.Num(), d.value.Denom()).Int64())
}

// Int32 return the Euclidean division of the rat
// Div sets z to the quotient x/y for y != 0 and returns z. If y == 0, a division-by-zero run-time panic occurs. Div implements Euclidean division
func (d Decimal) int32() int32 {
	return int32(new(big.Int).Div(d.value.Num(), d.value.Denom()).Int64())
}

// Int64 return the Euclidean division of the rat
// Div sets z to the quotient x/y for y != 0 and returns z. If y == 0, a division-by-zero run-time panic occurs. Div implements Euclidean division
func (d Decimal) int64() int64 {
	return int64(new(big.Int).Div(d.value.Num(), d.value.Denom()).Int64())
}

// Float32 return Float32 of the rat
// Float32 returns the nearest float32 value
// for x and a bool indicating whether f represents x exactly.
// If the magnitude of x is too large to be represented by a float32,
// f is an infinity and exact is false. The sign of f always matches the sign of x, even if f == 0
func (d Decimal) Float32() (float32, bool) {
	return d.value.Float32()
}

// Float64 return Float64 of the rat
// Float64 returns the nearest float64 value for x
// and a bool indicating whether f represents x exactly.
// If the magnitude of x is too large to be represented by a float64,
// f is an infinity and exact is false. The sign of f always matches the sign of x, even if f == 0.
func (d Decimal) Float64() (float64, bool) {
	return d.value.Float64()
}

// Rat return rat
func (d Decimal) Rat() *big.Rat {
	return new(big.Rat).Set(&d.value)
}

func min(x, y int32) int32 {
	if x < y {
		return x
	}
	return y
}

func abs(n int64) int64 {
	if n == math.MinInt64 {
		panic(fmt.Errorf("unsupported scope , should be -%v ~ %v ", math.MaxInt64, math.MaxInt64))
	}
	y := n >> 63
	return (n ^ y) - y
}

// stringToBigInt string to big int
// when len < 18, parse Int has better performance
// if len > 18 , using setString to generate big.Int
func stringToBigInt(value string) (*big.Int, error) {
	if len(value) < 18 {
		precInt, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return nil, err
		}
		return new(big.Int).SetInt64(precInt), nil
	} else {
		precInt, ok := new(big.Int).SetString(value, 10)
		if !ok {
			return nil, errors.New(stringBuilder("string convert to big.Int error ,actual: ", value))
		}
		return precInt, nil
	}
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (d *Decimal) UnmarshalJSON(decimalBytes []byte) error {
	str := string(decimalBytes)
	str = strings.Trim(str, "\"")
	if str == "null" {
		*d = NewFromInt(0)
		return nil
	}
	decimal, err := NewFromString(str)
	if err != nil {
		return fmt.Errorf("error decoding string '%s': %s", string(decimalBytes), err)
	}
	*d = decimal
	return nil
}

// MarshalJSON implements the json.Marshaler interface.
func (d Decimal) MarshalJSON() ([]byte, error) {
	var str string
	if MarshalJSONWithoutQuotes {
		str = d.String()
	} else {
		str = stringBuilder("\"", d.String(), "\"")
	}
	return []byte(str), nil
}

// Scan implements the sql.Scanner interface for database deserialization.
func (d *Decimal) Scan(value interface{}) error {
	// first try to see if the data is stored in database as a Numeric datatype
	switch v := value.(type) {

	case float32:
		*d = NewFromFloat32(v)
		return nil

	case float64:
		// numeric in sqlite3 sends us float64
		*d = NewFromFloat64(v)
		return nil

	case int64:
		// at least in sqlite3 when the value is 0 in db, the data is sent
		// to us as an int64 instead of a float64 ...
		*d = NewFromInt64(v)
		return nil
	case int32:
		*d = NewFromInt32(v)
		return nil
	case int:
		*d = NewFromInt(v)
		return nil
	case string:
		var err error
		str := strings.Trim(v, "\"")
		*d, err = NewFromString(str)
		return err
	case []byte:
		var err error
		str := strings.Trim(string(v), "\"")
		*d, err = NewFromString(str)
		return err
	default:
		return fmt.Errorf("Decimal.Scan error , value = %v , type = %v , ", value, reflect.TypeOf(v).Kind())
	}
}

// Value implements the driver.Valuer interface for database serialization.
func (d Decimal) Value() (driver.Value, error) {
	return d.String(), nil
}

func stringBuilder(text ...string) string {
	builder := builderPool.Get().(*strings.Builder)
	defer func() {
		builder.Reset()
		builderPool.Put(builder)
	}()
	for _, v := range text {
		builder.WriteString(v)
	}
	return builder.String()
}
