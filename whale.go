package whale

import (
	"container/list"
	"errors"
	"fmt"
	//	"log"
	"reflect"
	"strconv"
	"strings"
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

func ReadFromTokens(tokens []string) (*list.List, error) {
	if len(tokens) == 0 {
		return list.New(),
			fmt.Errorf("%v: unexpected EOF while reading", SyntaxError)
	}

	first := tokens[0]
	rest := append(tokens[:0], tokens[1:]...)

	if first == "(" {
		l := list.New()

		for rest[0] != ")" {
			atom, err := ReadFromTokens(rest)
			if err != nil {
				return list.New(), err
			}

			l.PushBack(atom.Front().Value)
		}

		rest = append(tokens[:0], tokens[1:]...)
		return l, nil

	} else if first == ")" {
		return list.New(),
			fmt.Errorf("%v: unexpected ')' at line %v", SyntaxError)

	} else {
		val, err := GetAtom(first)
		if err != nil {
			return list.New(), err
		}

		wrapped := list.New()
		wrapped.PushBack(val)

		return wrapped, nil
	}
}

func Parse(source string) (*list.List, error) {
	parsed, err := ReadFromTokens(Tokenize(source))
	if err != nil {
		return list.New(), err
	}

	return parsed, nil
}
