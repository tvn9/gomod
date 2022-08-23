package toolkit

import (
	"crypto/rand"
	"fmt"
)

const randomStringSource = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789_+"

// Tools is a type used to instantiate this module. Any veriable of this type will have access
// to all the methods with the receiver *Tools
type Tools struct {
}

// RandomString returns a string of random characters of length n, using RandomStringSource
// as the source for the string
func (t *Tools) RandomString(n int) string {
	s, r := make([]rune, n), []rune(randomStringSource)
	fmt.Printf("s = %v, r = %v\n", s, r)
	fmt.Printf("len of r %d\n", len(r))
	for i := range s {
		p, _ := rand.Prime(rand.Reader, len(r))
		fmt.Println(p)
		x, y := p.Uint64(), uint64(len(r))
		s[i] = r[x%y]
	}
	return string(s)
}
