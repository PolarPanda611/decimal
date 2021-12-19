/*
 * @Author: Daniel TAN
 * @Description:
 * @Date: 2021-04-13 10:44:13
 * @LastEditTime: 2021-07-14 16:55:43
 * @LastEditors: Daniel TAN
 * @FilePath: /fr-price-common-pkg/decimal/decimal_benchmark_test.go
 */
package decimal

import (
	"math/big"
	"testing"
)

func Benchmark_decimal_Decimal_NewFromFloat64(b *testing.B) {

	b.ReportAllocs()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		NewFromFloat64(1000.123)
	}
}

func Benchmark_decimal_Decimal_NewFromFloat32(b *testing.B) {

	b.ReportAllocs()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		NewFromFloat32(1000.123)
	}
}

func Benchmark_decimal_Decimal_NewFromInt(b *testing.B) {

	b.ReportAllocs()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		NewFromInt(12451412421)
	}
}

func Benchmark_decimal_Decimal_NewFromInt32(b *testing.B) {

	b.ReportAllocs()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		NewFromInt32(2342)
	}
}

func Benchmark_decimal_Decimal_NewFromInt64(b *testing.B) {

	b.ReportAllocs()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		NewFromInt64(12451412421)
	}
}

func Benchmark_decimal_Decimal_NewFromBigInt(b *testing.B) {

	b.ReportAllocs()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		NewFromBigInt(mustBigIntFromString("1232115"), 2)
	}
}

func Benchmark_decimal_Decimal_NewFromFrac(b *testing.B) {

	b.ReportAllocs()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		NewFromFrac(mustBigIntFromString("1232115"), mustBigIntFromString("12124121"), 2)
	}
}

func Benchmark_decimal_Decimal_NewFromBigRat(b *testing.B) {

	b.ReportAllocs()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		NewFromBigRat(big.NewRat(2, 51231), 2)
	}
}

func Benchmark_decimal_Decimal_NewFromStringWithoutDecimal(b *testing.B) {

	b.ReportAllocs()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		NewFromString("1234211255")
	}
}
func Benchmark_decimal_Decimal_NewFromStringWithDecimal(b *testing.B) {

	b.ReportAllocs()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		NewFromString("12.12412512123")
	}
}
func Benchmark_decimal_Decimal_NewFromStringWithScience(b *testing.B) {

	b.ReportAllocs()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		NewFromString("12e10")
	}
}

func Benchmark_decimal_Decimal_NewRequiredFromStringWithoutDecimal(b *testing.B) {

	b.ReportAllocs()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		NewRequiredFromString("1234211255")
	}
}
func Benchmark_decimal_Decimal_NewRequiredFromStringWithDecimal(b *testing.B) {

	b.ReportAllocs()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		NewRequiredFromString("12.12412512123")
	}
}
func Benchmark_decimal_Decimal_NewRequiredFromStringWithScience(b *testing.B) {

	b.ReportAllocs()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		NewRequiredFromString("12e10")
	}
}
func Benchmark_decimal_Decimal_Add_same_precision(b *testing.B) {
	d1 := NewFromFloat64(1000.123)
	d2 := NewFromFloat64(500.123)

	b.ReportAllocs()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		d1.Add(d2)
	}
}

func Benchmark_decimal_Decimal_Add_different_precision(b *testing.B) {
	d1 := NewFromFloat64(1000.123)
	d2 := NewFromFloat64(500).Mul(NewFromFloat64(0.12))

	b.ReportAllocs()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		d1.Add(d2)
	}
}
func Benchmark_decimal_Decimal_Sub_different_precision(b *testing.B) {
	d1 := NewFromFloat64(1000.123)
	d2 := NewFromFloat64(500).Mul(NewFromFloat64(0.12))

	b.ReportAllocs()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		d1.Sub(d2)
	}
}
func Benchmark_decimal_Decimal_Sub_same_precision(b *testing.B) {
	d1 := NewFromFloat64(1000.12)
	d2 := NewFromFloat64(500).Mul(NewFromFloat64(0.12))

	b.ReportAllocs()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		d1.Sub(d2)
	}
}

func Benchmark_decimal_Decimal_Mul_different_precision(b *testing.B) {
	d1 := NewFromFloat64(1000.123)
	d2 := NewFromFloat64(500).Mul(NewFromFloat64(0.12))

	b.ReportAllocs()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		d1.Mul(d2)
	}
}
func Benchmark_decimal_Decimal_Mul_same_precision(b *testing.B) {
	d1 := NewFromFloat64(1000.12)
	d2 := NewFromFloat64(500).Mul(NewFromFloat64(0.12))

	b.ReportAllocs()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		d1.Mul(d2)
	}
}

func Benchmark_decimal_Decimal_Div_different_precision(b *testing.B) {
	d1 := NewFromFloat64(1000.123)
	d2 := NewFromFloat64(500).Mul(NewFromFloat64(0.12))

	b.ReportAllocs()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		d1.Div(d2)
	}
}

func Benchmark_decimal_Decimal_Div_same_precision(b *testing.B) {
	d1 := NewFromFloat64(1000.12)
	d2 := NewFromFloat64(500).Mul(NewFromFloat64(0.12))

	b.ReportAllocs()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		d1.Div(d2)
	}
}

func Benchmark_decimal_Decimal_Mod_different_precision(b *testing.B) {
	d1 := NewFromFloat64(1000.123)
	d2 := NewFromFloat64(500).Mul(NewFromFloat64(0.12))

	b.ReportAllocs()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		d1.Mod(d2)
	}
}
func Benchmark_decimal_Decimal_Mod_same_precision(b *testing.B) {
	d1 := NewFromFloat64(1000.12)
	d2 := NewFromFloat64(500).Mul(NewFromFloat64(0.12))

	b.ReportAllocs()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		d1.Mod(d2)
	}
}
func Benchmark_decimal_Decimal_RoundDown(b *testing.B) {
	d1 := NewFromFloat64(1000.123)

	b.ReportAllocs()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		d1.RoundDown(1)
	}
}

func Benchmark_decimal_Decimal_RoundUp(b *testing.B) {
	d1 := NewFromFloat64(1000.123)

	b.ReportAllocs()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		d1.RoundUp(1)
	}
}

func Benchmark_decimal_Decimal_RoundHalfUp(b *testing.B) {
	d1 := NewFromFloat64(1000.123)

	b.ReportAllocs()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		d1.RoundHalfUp(1)
	}
}

func Benchmark_decimal_Decimal_RoundTowardZero(b *testing.B) {
	d1 := NewFromFloat64(1000.123)

	b.ReportAllocs()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		d1.RoundTowardZero(1)
	}
}

func Benchmark_decimal_Decimal_Cmp_different_precision(b *testing.B) {
	d1 := NewFromFloat64(1000.123)
	d2 := NewFromFloat64(500).Mul(NewFromFloat64(0.12))

	b.ReportAllocs()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		d1.Cmp(d2)
	}
}

func Benchmark_decimal_Decimal_Sign(b *testing.B) {
	d1 := NewFromFloat64(1000.123)

	b.ReportAllocs()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		d1.Sign()
	}
}

func Benchmark_decimal_Decimal_Neg(b *testing.B) {
	d1 := NewFromFloat64(1000.123)

	b.ReportAllocs()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		d1.Neg()
	}
}

func Benchmark_decimal_Decimal_Abs(b *testing.B) {
	d1 := NewFromFloat64(1000.123)

	b.ReportAllocs()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		d1.Abs()
	}
}

func Benchmark_decimal_Decimal_Pow(b *testing.B) {
	d1 := NewFromFloat64(1000.123)

	b.ReportAllocs()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		d1.Pow(23)
	}
}

func Benchmark_decimal_RoundStrategy_IsValid_Success(b *testing.B) {
	t := RoundDown

	b.ReportAllocs()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		t.IsValid()
	}
}

func Benchmark_decimal_RoundStrategy_IsValid_Error(b *testing.B) {
	t := RoundStrategy("123")

	b.ReportAllocs()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		t.IsValid()
	}
}
