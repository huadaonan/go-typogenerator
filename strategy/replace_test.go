package strategy_test

import (
	"testing"

	"zenithar.org/go/typogenerator/mapping"
	"zenithar.org/go/typogenerator/strategy"
)

func TestReplace(t *testing.T) {
	out, err := strategy.Replace(mapping.French).Generate("zenithar.org")
	if err != nil {
		t.Fail()
		t.Fatal("Error should not occurs !", err)
	}

	if len(out) == 0 {
		t.FailNow()
	}

	if len(out) != 77 {
		t.FailNow()
	}
}