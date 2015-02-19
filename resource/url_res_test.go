package resource

import (
	"testing"
)

func TestURLRes(t *testing.T) {
	u1, _ := ParseURLRes("/", "")
	u2, _ := ParseURLRes("/usr", "")
	u3, _ := ParseURLRes("/usr/lib", "")
	u4, _ := ParseURLRes("/var", "")
	u2c, _ := ParseURLRes("/usr", "")
	if u1.Equals(u2) {
		t.Fatal("u1 should not equal u2")
	}
	if !u1.Contains(u2, u3, u4) {
		t.Fatal("u1 should contain u2 u3 u4")
	}
	if !u2c.Equals(u2) {
		t.Fatal("u2 and u2c should be equal")
	}
}
