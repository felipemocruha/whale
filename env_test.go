package whale

import (
	"testing"
)

func TestAddInt(t *testing.T) {
	want := Int(2)
	res := Add(Int(1), Int(1))

	if res != want {
		MakeErrorMsg(t, want, res)
	}
}
