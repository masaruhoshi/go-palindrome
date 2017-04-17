package main

import(
	"testing"
)

func TestNewToReturnGoPal(t *testing.T) {
	r := New()

	ExpectNotNil(t, r)
}