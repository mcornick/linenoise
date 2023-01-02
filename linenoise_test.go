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
	"testing"
)

// Test that a valid length returns a password.
func TestValidLength(t *testing.T) {
	p := Parameters{Length: 16, Upper: true, Lower: true, Digit: true}
	generatedNoise, _ := Noise(p)
	returnedLength := len(generatedNoise)
	if returnedLength != 16 {
		t.Error("expected 16, got ", returnedLength)
	}
}

// Test that an invalid length returns an error.
func TestInvalidLength(t *testing.T) {
	p := Parameters{Length: 0, Upper: true, Lower: true, Digit: true}
	_, err := Noise(p)
	if err == nil {
		t.Error("expected error, got nil")
	}
}

// Test that a length longer than that of the permittedCharacters array returns
// a password.
func TestLongLength(t *testing.T) {
	p := Parameters{Length: 64, Upper: true, Lower: true, Digit: true}
	generatedNoise, _ := Noise(p)
	returnedLength := len(generatedNoise)
	if returnedLength != 64 {
		t.Error("expected 64, got ", returnedLength)
	}
}

// Tests for acceptable passwords.
func TestAcceptableWithAllThree(t *testing.T) {
	p := Parameters{Length: 16, Upper: true, Lower: true, Digit: true}
	if !isAcceptable("Aa1", p) {
		t.Error("expected Aa1 to be acceptable; it is not")
	}
}

func TestUnacceptableWithoutUppercase(t *testing.T) {
	p := Parameters{Length: 16, Upper: true, Lower: true, Digit: true}
	if isAcceptable("aa1", p) {
		t.Error("expected aa1 to be unacceptable; it is not")
	}
}

func TestUnacceptableWithoutLowercase(t *testing.T) {
	p := Parameters{Length: 16, Upper: true, Lower: true, Digit: true}
	if isAcceptable("AA1", p) {
		t.Error("expected AA1 to be unacceptable; it is not")
	}
}

func TestUnacceptableWithoutDigit(t *testing.T) {
	p := Parameters{Length: 16, Upper: true, Lower: true, Digit: true}
	if isAcceptable("Aaa", p) {
		t.Error("expected Aaa to be unacceptable; it is not")
	}
}

func TestAcceptableWithoutUppercase(t *testing.T) {
	p := Parameters{Length: 16, Upper: false, Lower: true, Digit: true}
	if !isAcceptable("aa1", p) {
		t.Error("expected aa1 to be acceptable; it is not")
	}
}

func TestAcceptableWithoutLowercase(t *testing.T) {
	p := Parameters{Length: 16, Upper: true, Lower: false, Digit: true}
	if !isAcceptable("AA1", p) {
		t.Error("expected AA1 to be acceptable; it is not")
	}
}

func TestAcceptableWithoutDigit(t *testing.T) {
	p := Parameters{Length: 16, Upper: true, Lower: true, Digit: false}
	if !isAcceptable("Aaa", p) {
		t.Error("expected Aaa to be acceptable; it is not")
	}
}

// Test that candidateSet repeats permittedCharacters.
func TestCandidateSet(t *testing.T) {
	p := Parameters{Length: 64, Upper: true, Lower: true, Digit: true}
	set, _ := candidateSet(p)
	permitted, _ := permittedCharacters(p)
	if len(set) != 2*len(permitted) {
		t.Error("expected permittedCharacters to be repeated twice; it is not")
	}
}

// Test that short lengths are handled with errors.
func TestNoiseShortLength(t *testing.T) {
	p := Parameters{Length: 0, Upper: true, Lower: true, Digit: true}
	_, err := Noise(p)
	if err == nil {
		t.Error("expected error, got nil")
	}
}

// Convenience function for the next tests.
func expectedText(minimum int) string {
	message := fmt.Sprintf("invalid length - must be an integer greater than %d", minimum)
	return message
}

// Test the tooShortErrorText function.
func TestNoiseShortLengthThreeClasses(t *testing.T) {
	p := Parameters{Length: 16, Upper: true, Lower: true, Digit: true}
	msg := tooShortErrorText(p)
	if msg != expectedText(2) {
		t.Errorf("expected \"%s\", got %s", expectedText(2), msg)
	}
}
func TestNoiseShortLengthTwoClasses(t *testing.T) {
	p := Parameters{Length: 16, Upper: true, Lower: true, Digit: false}
	msg := tooShortErrorText(p)
	if msg != expectedText(1) {
		t.Errorf("expected \"%s\", got %s", expectedText(1), msg)
	}
}
func TestNoiseShortLengthOneClass(t *testing.T) {
	p := Parameters{Length: 16, Upper: true, Lower: false, Digit: false}
	msg := tooShortErrorText(p)
	if msg != expectedText(0) {
		t.Errorf("expected \"%s\", got %s", expectedText(0), msg)
	}
}

// Test that attempting with no classes enabled produces an error.
func TestUnacceptableWithoutAnyClasses(t *testing.T) {
	p := Parameters{Length: 16, Upper: false, Lower: false, Digit: false}
	_, err := Noise(p)
	if err == nil {
		t.Error("expected error, got nil")
	}
}
