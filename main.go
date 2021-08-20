// Copyright 2021 The Origami Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
)

func main() {
	number, primes := make([][]uint, 256), make([]uint, 0, 8)
	primes = append(primes, 2, 3)
	for i := uint(4); i < uint(len(number)); i++ {
		isPrime := true
		for j := range primes {
			if i%primes[j] == 0 {
				isPrime = false
				number[j] = append(number[j], primes[j])
			}
		}
		if isPrime {
			primes = append(primes, i)
		}
	}
	fmt.Println(primes)
}
