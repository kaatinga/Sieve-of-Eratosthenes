package sieve

import "log"

// Sieve of Eratosthenes function. Returns a map of prime numbers. Point out maximum number in the returning map.
func Sieve(limit int) map[int]int {
	primes := make(map[int]int, limit)

	// Forming an array of numbers where the maximum number = limit
	primes[2] = 2
	for i := 3; i < limit; i = i + 2 {
		primes[i] = i
	}

	// Sieve of Eratosthenes
	for i2 := 3; i2*i2 <= limit; i2 = i2 + 2 {
		for i3 := i2; i2*i3 <= limit; i3++ {
			if val, ok := primes[i2*i3]; ok {
				log.Println("Deleting", val)
				delete(primes, val)
			}
		}
	}
	return primes
}
