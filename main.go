package sieve

import "log"

// Sieve of Eratosthenes function. Returns a map of prime numbers. Point out maximum number in the returning map.
func Sieve(limit int) map[int]int {
	primes := make(map[int]int, limit)

	// Forming an array of numbers where the maximum number = limit
	for i := 2; i < limit; i++ {
		primes[i] = i
	}

	// Sieve of Eratosthenes
	for i2 := 2; i2<<1 <= limit; i2++ {
		for i3 := i2; i2*i3 <= limit; i3++ {
			if val, ok := primes[i2*i3]; ok {
				log.Println("Deleting", val)
				delete(primes, val)
			}
		}
	}
	return primes
}
