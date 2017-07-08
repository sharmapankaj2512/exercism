package bob

import (
	"strings"
	"strconv"
)

const testVersion = 3

type Statement interface {
	reply() string
}

type Question string

type AwkwardSilence string

type Shout string

type None string

func (q Question) reply() string {
	return "Sure."
}

func (n None) reply() string {
	return "Whatever."
}

func (a AwkwardSilence) reply() string {
	return "Fine. Be that way!"
}

func (s Shout) reply() string {
	return "Whoa, chill out!"
}

func Hey(statement string) string {
	return typeOf(statement).reply()
}

func typeOf(statement string) Statement {
	if IsShout(statement) {
		return Shout(statement)
	}
	if IsQuestion(statement) {
		return Question(statement)
	}
	if IsBlank(statement) {
		return AwkwardSilence(statement)
	}
	return None(statement)
}

func IsQuestion(statement string) bool {
	return strings.HasSuffix(statement, "?")
}

func IsBlank(statement string) bool {
	return len(strings.Trim(statement, " ")) == 0
}

func IsShout(statement string) bool {
	text := strings.Replace(strings.Trim(statement, "?"), ", ", "", -1)
	return IsNumber(text) && strings.ToUpper(statement) == statement
}

func IsNumber(n string) bool {
	_, err := strconv.Atoi(n)
	return err == nil
}