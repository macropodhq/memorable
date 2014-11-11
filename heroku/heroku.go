// Package heroku provides a strategy for memorable to generate heroku-style
// adjective-noun-number strings.
package heroku

import (
	"github.com/BugHerd/memorable"
)

func New(seed int64) *memorable.Generator {
	return memorable.NewGenerator(seed, [][]string{adjectives, nouns})
}

var globalGenerator = New(0)

func Get() string {
	return globalGenerator.Get()
}
