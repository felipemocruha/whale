package whale

import (
	"errors"
	"fmt"
	//	"log"
	"reflect"
	"strconv"
	"strings"

	"github.com/Workiva/go-datastructures/list"
)

const (
	tInt    = iota
	tFloat  = iota
	tSymbol = iota
	tList   = iota
)

var (
	SyntaxError = errors.New("syntax error")
)

type Symbol string

type Number interface {
	Value() interface{}
	Type() interface{}
}

type Int int
type Float float64

func (i Int) Value() interface{} {
	return i
}

func (i Int) Type() interface{} {
	return reflect.TypeOf(i)
}

func (f Float) Value() interface{} {
	return f
}

func (f Float) Type() interface{} {
	return reflect.TypeOf(f)
}

func popn(tokens []string, n int) (string, []string) {
	var p string
	var out []string
	for i, v := range tokens {
		if i == n {
			p = v
		} else {
			out = append(out, v)
		}
	}

	return p, out
}

type Atom struct {
	Value interface{}
}

func getLines(s string) []string {
	return strings.Split(s, "\n")
}

type TokenType int

type Token struct {
	Type    TokenType
	Literal string
	Value   interface{}
}

func tryInt(token string) (Token, error) {
	val, err := strconv.Atoi(token)
	if err != nil {
		return Token{}, nil
	}

	return Token{tInt, token, Int(val)}, nil
}

func tryFloat(token string) (Token, error) {
	val, err := strconv.ParseFloat(token, 64)
	if err != nil {
		return Token{}, nil
	}

	return Token{tFloat, token, Float(val)}, nil
}

func symbolToToken(token string) Token {
	return Token{tSymbol, token, Symbol(token)}
}

func GetAtom(token string) (interface{}, error) {
	tk, err := tryInt(token)
	if (err != nil) || (tk == Token{}) {
		tk, err = tryFloat(token)
		if (err != nil) || (tk == Token{}) {
			return symbolToToken(token).Value, nil
		}
	}

	return tk.Value, nil
}

func splitSpace(tokens string) []string {
	var split []string
	for _, c := range tokens {
		if string(c) != " " {
			split = append(split, string(c))
		}
	}

	return split
}

func Tokenize(line string) []string {
	open := strings.Replace(line, "(", " ( ", -1)
	close := strings.Replace(open, ")", " ) ", -1)

	return splitSpace(close)
}

func ReadFromTokens(tokens []string) (interface{}, error) {
	if len(tokens) == 0 {
		return nil,
			fmt.Errorf("%v: unexpected EOF while reading", SyntaxError)
	}

	token, tokens := popn(tokens, 0)
	if token == "(" {
		l := list.Empty
		for ok := true; ok; ok = (tokens[0] == ")") {
			atom, err := ReadFromTokens(tokens)
			if err != nil {
				return nil, err
			}
			l = l.Add(atom)
		}

		popn(tokens, 0)
		return l, nil

	} else if token == ")" {
		return nil,
			fmt.Errorf("%v: unexpected ')' at line %v", SyntaxError)

	} else {
		val, err := GetAtom(token)
		if err != nil {
			return nil, err
		}

		return val, nil
	}
}
