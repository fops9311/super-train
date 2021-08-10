package main

import (
	"strings"
	"testing"

	. "github.com/franela/goblin"
)

func TestPredicateSucceed(t *testing.T) {
	fakeTest := testing.T{}
	g := Goblin(&fakeTest)
	g.Describe("Predicate Tests", func() {
		cars := []string{"Reno", "Volvo", "Audi"}
		g.Describe("Starts with Re", func() {
			g.It("Should be true", func() {
				bs := Any(cars, func(st string) bool {
					return strings.HasPrefix(st, "Re")
				})
				g.Assert(bs).Equal(true)
			})
		})
		g.Describe("Starts with lo", func() {
			g.It("Should be false", func() {
				bs := Any(cars, func(st string) bool {
					return strings.HasPrefix(st, "lo")
				})
				g.Assert(bs).Equal(false)
			})
		})
		g.Describe("All dont have letter `b`", func() {
			g.It("Should be true", func() {
				bs := All(cars, func(st string) bool {
					return !strings.Contains(st, "b")
				})
				g.Assert(bs).Equal(true)
			})
		})
		g.Describe("Filter with letter `d`", func() {
			g.It("Should be []string{\"Reno\", \"Volvo\"}", func() {
				bs := Filter(cars, func(st string) bool {
					return !strings.Contains(st, "d")
				})
				g.Assert(bs).Equal([]string{"Reno", "Volvo"})
			})
		})
	})
}
