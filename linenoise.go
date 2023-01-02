// SPDX-License-Identifier: MIT
//
// Copyright (c) 2017-2023 Mark Cornick
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package linenoise

import (
	"fmt"
	"math/rand"
	"regexp"
	"strings"
	"time"
)

// Parameters is a struct containing the length desired and whether we want the
// different classes of characters in the password.
type Parameters struct {
	Length              int
	Upper, Lower, Digit bool
}

// btoi converts a bool to an int (false = 0, true = 1).
func btoi(b bool) int {
	if b {
		return 1
	}
	return 0
}

// tooShortErrorText returns the text of the error generated when
// the requested length is too short.
func tooShortErrorText(p Parameters) string {
	min := btoi(p.Upper) + btoi(p.Lower) + btoi(p.Digit)
	return fmt.Sprintf("invalid length - must be an integer greater than %d", min-1)
}

// permittedCharacters is the array of characters that can be used in
// passwords. It is a combination of uppercase letters, lowercase letters,
// and digits depending on which classes are requested.
func permittedCharacters(p Parameters) ([]string, error) {
	setUpper := []string{
		"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M",
		"N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z",
	}
	setLower := []string{
		"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m",
		"n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z",
	}
	setDigit := []string{
		"0", "1", "2", "3", "4", "5", "6", "7", "8", "9",
	}
	var set []string

	if !p.Upper && !p.Lower && !p.Digit {
		return nil, fmt.Errorf("must include at least one of -lower, -upper and/or -digit")
	}

	if p.Upper {
		set = append(set, setUpper...)
	}
	if p.Lower {
		set = append(set, setLower...)
	}
	if p.Digit {
		set = append(set, setDigit...)
	}
	return set, nil
}

// candidateSet returns an array of characters, created by concatenating
// multiple instances of permittedCharacters, sufficient to create a password
// of the specified length.
func candidateSet(p Parameters) ([]string, error) {
	var set []string
	permitted, err := permittedCharacters(p)
	if err != nil {
		return nil, err
	}
	multiples := p.Length / len(permitted)
	set = permitted
	for i := 0; i < multiples; i++ {
		set = append(set, permitted...)
	}
	return set, nil
}

// isAcceptable checks whether the specified password meets standards.
// This means it includes each of the requested character classes.
func isAcceptable(pw string, p Parameters) bool {
	if !p.Upper && !p.Lower && !p.Digit {
		return false
	}
	hasUpperCase, _ := regexp.MatchString("[[:upper:]]", pw)
	hasLowerCase, _ := regexp.MatchString("[[:lower:]]", pw)
	hasDigit, _ := regexp.MatchString("[[:digit:]]", pw)
	return ((hasUpperCase || !p.Upper) && (hasLowerCase || !p.Lower) && (hasDigit || !p.Digit))
}

// Noise creates a password of the specified length if it is valid, or an
// error if the length is invalid.
func Noise(p Parameters) (string, error) {
	if p.Length < btoi(p.Upper)+btoi(p.Lower)+btoi(p.Digit) {
		return "", fmt.Errorf(tooShortErrorText(p))
	}
	generatedNoise := "!"
	for !isAcceptable(generatedNoise, p) {
		candidates, err := candidateSet(p)
		if err != nil {
			return "", err
		}
		rand.Seed(time.Now().UnixNano())
		characters := make([]string, p.Length)
		for i := 0; i < p.Length; i++ {
			index := rand.Intn(len(candidates))
			characters[i] = candidates[index]
			candidates[index] = candidates[len(candidates)-1]
			candidates = candidates[:len(candidates)-1]
		}
		generatedNoise = strings.Join(characters[0:p.Length], "")
	}
	return generatedNoise, nil
}
