package toolkit

import (
	"crypto/rand"
	"fmt"
)

const randomStringSource = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789_+"
// Tools is the type used to instantiate this module. Any variables of this type will have access to
// all the methods with te receiver *Tools
type Tools struct{}
// RandomString returns a string of random characters of length n, using randomStringSource
// as the source for the string

func (t *Tools) RandomString(n int) string {
	s, r := make([]rune, n), []rune(randomStringSource)
	for i := range s {
		p, _ := rand.Prime(rand.Reader, len(r))
		fmt.Println("Value of P is:", p)
		x, y := p.Uint64(), uint64(len(r))
		fmt.Printf("Value of x,y are%v,%v\n",x,y)
		s[i] = r[x%y]
		fmt.Println("s[i] is",string(s[i]))
	}
	return string(s)
}
