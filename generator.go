package memorable

import (
	"math/rand"
	"strconv"
	"strings"
)

type Generator struct {
	r *rand.Rand
	l [][]string
}

// `NewGenerator` takes a seed to use for initialising its random number
// generator and a number of string slices. These slices are used to construct
// strings in the Get() method.
func NewGenerator(seed int64, lists [][]string) *Generator {
	return &Generator{
		r: rand.New(rand.NewSource(seed)),
		l: lists,
	}
}

// Get() returns a string containing your brand new "random" word thing.
func (g Generator) Get() string {
	s := make([]string, len(g.l)+1)

	for i, l := range g.l {
		s[i] = l[g.r.Intn(len(l))]
	}

	s[len(g.l)] = strconv.Itoa(g.r.Intn(9999))

	return strings.Join(s, "-")
}
