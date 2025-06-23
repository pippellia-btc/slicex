package slicex

import "math/rand/v2"

// RandomElement returns a random element from the slice.
// It panics if the slice is empty.
// Not safe for security purposes.
func RandomElement[E any](s []E) E { return s[rand.IntN(len(s))] }

// Shuffle the provided slice at random.
// It panics if the slice is empty.
// Not safe for security purposes.
func Shuffle[E any](s []E) { rand.Shuffle(len(s), func(i, j int) { s[i], s[j] = s[j], s[i] }) }
