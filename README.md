# decimal
this is the package which doing the price calculation in go 


## Benchmark
```
goos: darwin
goarch: amd64
pkg: github.com/fastretailing/fr-price-common-pkg/decimal
cpu: Intel(R) Core(TM) i5-8279U CPU @ 2.40GHz
Benchmark_decimal_Decimal_NewFromFloat64-8                        	 1409138	       828.8 ns/op	     176 B/op	      14 allocs/op
Benchmark_decimal_Decimal_NewFromFloat32-8                        	 1000000	      1225 ns/op	     176 B/op	      14 allocs/op
Benchmark_decimal_Decimal_NewFromInt-8                            	26887730	        42.70 ns/op	      16 B/op	       2 allocs/op
Benchmark_decimal_Decimal_NewFromInt32-8                          	28672197	        42.04 ns/op	      16 B/op	       2 allocs/op
Benchmark_decimal_Decimal_NewFromInt64-8                          	34635019	        34.95 ns/op	      16 B/op	       2 allocs/op
Benchmark_decimal_Decimal_NewFromBigInt-8                         	 5122410	       287.6 ns/op	      88 B/op	       5 allocs/op
Benchmark_decimal_Decimal_NewFromFrac-8                           	 1797223	       665.5 ns/op	     176 B/op	      10 allocs/op
Benchmark_decimal_Decimal_NewFromBigRat-8                         	10449814	       114.0 ns/op	      32 B/op	       4 allocs/op
Benchmark_decimal_Decimal_NewFromStringWithoutDecimal-8           	 3201276	       368.6 ns/op	     128 B/op	       9 allocs/op
Benchmark_decimal_Decimal_NewFromStringWithDecimal-8              	 1418436	       816.3 ns/op	     192 B/op	      12 allocs/op
Benchmark_decimal_Decimal_NewFromStringWithScience-8              	 1599679	       793.2 ns/op	     232 B/op	      19 allocs/op
Benchmark_decimal_Decimal_NewRequiredFromStringWithoutDecimal-8   	 3168177	       396.7 ns/op	     128 B/op	       9 allocs/op
Benchmark_decimal_Decimal_NewRequiredFromStringWithDecimal-8      	 1450674	       864.3 ns/op	     192 B/op	      12 allocs/op
Benchmark_decimal_Decimal_NewRequiredFromStringWithScience-8      	 1551255	       779.6 ns/op	     232 B/op	      19 allocs/op
Benchmark_decimal_Decimal_Add_same_precision-8                    	 3652807	       333.3 ns/op	     208 B/op	       6 allocs/op
Benchmark_decimal_Decimal_Add_different_precision-8               	 4178389	       289.4 ns/op	     208 B/op	       6 allocs/op
Benchmark_decimal_Decimal_Sub_different_precision-8               	 4047906	       287.3 ns/op	     168 B/op	       6 allocs/op
Benchmark_decimal_Decimal_Sub_same_precision-8                    	 4599156	       261.8 ns/op	     168 B/op	       6 allocs/op
Benchmark_decimal_Decimal_Mul_different_precision-8               	 4609089	       262.4 ns/op	     112 B/op	       4 allocs/op
Benchmark_decimal_Decimal_Mul_same_precision-8                    	 5497916	       219.7 ns/op	     112 B/op	       4 allocs/op
Benchmark_decimal_Decimal_Div_different_precision-8               	 4640180	       261.2 ns/op	     112 B/op	       4 allocs/op
Benchmark_decimal_Decimal_Div_same_precision-8                    	 5744563	       212.7 ns/op	     112 B/op	       4 allocs/op
Benchmark_decimal_Decimal_Mod_different_precision-8               	 1391182	       868.6 ns/op	     424 B/op	      18 allocs/op
Benchmark_decimal_Decimal_Mod_same_precision-8                    	 1530318	       788.5 ns/op	     424 B/op	      18 allocs/op
Benchmark_decimal_Decimal_RoundDown-8                             	 4159063	       289.7 ns/op	     104 B/op	       8 allocs/op
Benchmark_decimal_Decimal_RoundUp-8                               	  962616	      1313 ns/op	     704 B/op	      28 allocs/op
Benchmark_decimal_Decimal_RoundHalfUp-8                           	  655545	      1896 ns/op	     928 B/op	      41 allocs/op
Benchmark_decimal_Decimal_RoundTowardZero-8                       	 3988682	       298.2 ns/op	     104 B/op	       8 allocs/op
Benchmark_decimal_Decimal_Cmp_different_precision-8               	14034086	        87.49 ns/op	      96 B/op	       2 allocs/op
Benchmark_decimal_Decimal_Sign-8                                  	630356682	         1.888 ns/op	       0 B/op	       0 allocs/op
Benchmark_decimal_Decimal_Neg-8                                   	24549660	        48.57 ns/op	      16 B/op	       2 allocs/op
Benchmark_decimal_Decimal_Abs-8                                   	25765116	        48.55 ns/op	      16 B/op	       2 allocs/op
Benchmark_decimal_Decimal_Pow-8                                   	  116383	     10427 ns/op	    3185 B/op	      64 allocs/op
Benchmark_decimal_RoundStrategy_IsValid_Success-8                 	488689292	         2.443 ns/op	       0 B/op	       0 allocs/op
Benchmark_decimal_RoundStrategy_IsValid_Error-8                   	13089566	        91.84 ns/op	      64 B/op	       2 allocs/op
```