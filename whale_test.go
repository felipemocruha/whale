package whale

import (
	"container/list"
	"fmt"
	"reflect"
	"testing"
)

func MakeErrorMsg(t *testing.T, want, got interface{}) {
	t.Errorf("Unexpected output.\n\t Want (%v),\n\t Got (%v)", want, got)
}

func TestPopn(t *testing.T) {
	tokens := []string{"0", "1", "2"}
	want1 := "0"
	want2 := []string{"1", "2"}

	p, res := popn(tokens, 0)

	if p != want1 {
		MakeErrorMsg(t, want1, res)
	}

	for i := 0; i < len(want2); i++ {
		if want2[i] != res[i] {
			MakeErrorMsg(t, want2, res)
			break
		}
	}
}

func TestGetAtomInt(t *testing.T) {
	res, _ := GetAtom("1")
	res = reflect.TypeOf(res)
	want := reflect.TypeOf(Int(1))

	if res != want {
		MakeErrorMsg(t, want, res)
	}
}

func TestGetAtomFloat(t *testing.T) {
	res, _ := GetAtom("3.141592")
	res = reflect.TypeOf(res)
	want := reflect.TypeOf(Float(1.2))

	if res != want {
		MakeErrorMsg(t, want, res)
	}
}

func TestGetAtomSymbol(t *testing.T) {
	res, _ := GetAtom("lisp")
	res = reflect.TypeOf(res)
	want := reflect.TypeOf(Symbol(1))

	if res != want {
		MakeErrorMsg(t, want, res)
	}
}

func TestGetAtomSymbolPlus(t *testing.T) {
	res, _ := GetAtom("+")
	res = reflect.TypeOf(res)
	want := reflect.TypeOf(Symbol("+"))

	if res != want {
		MakeErrorMsg(t, want, res)
	}
}

func TestTokenize(t *testing.T) {
	want := []string{"(", "+", "1", "2", ")"}
	res := Tokenize("(+ 1 2)")

	for i := 0; i < len(want); i++ {
		if want[i] != res[i] {
			MakeErrorMsg(t, want, res)
			break
		}
	}
}

func TestSplitSpace(t *testing.T) {
	want := []string{"(", "+", "1", "2", ")"}
	res := splitSpace(" ( + 1 2 ) ")

	for i := 0; i < len(want); i++ {
		if want[i] != res[i] {
			MakeErrorMsg(t, want, res)
			break
		}
	}
}

func TestReadFromTokens(t *testing.T) {
	res, _ := ReadFromTokens(Tokenize("(+ 1 2)"))

	want := list.New()
	want.PushBack(Symbol("+"))
	want.PushBack(Int(1))
	want.PushBack(Int(2))

	for w, r := want.Front(), res.Front(); r != nil; w, r = w.Next(), r.Next() {
		if w.Value != r.Value {
			MakeErrorMsg(t, w.Value, r.Value)
			break
		}
	}
}

func TestReadFromTokensNested(t *testing.T) {
	res, _ := ReadFromTokens(Tokenize("(+ (+ 1 2) 1)"))

	win := list.New()
	win.PushBack(Symbol("+"))
	win.PushBack(Int(1))
	win.PushBack(Int(2))

	wout := list.New()
	wout.PushBack(Symbol("+"))
	wout.PushBack(win)
	wout.PushBack(Int(1))

	for r := res.Front(); r != nil; r = r.Next() {
		fmt.Println(r.Value)
	}

}
