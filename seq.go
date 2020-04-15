// Copyright (c) 2020 The seq developers. All rights reserved.
// Project site: https://github.com/goinvest/seq
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE file for the project.

package seq

import (
	"encoding/json"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

// Int converts JSON range strings into a slice of int.s
type Int []int

// Parse parses a string into an integer slice. If the parsing fails a nil
// slice is returned.
func Parse(s string) ([]int, error) {
	return newInt(s)
}

// Parsef parses a string according to a format specifier and returns the
// resulting integer slice.
func Parsef(s string, a ...interface{}) ([]int, error) {
	return newInt(fmt.Sprintf(s, a...))
}

// NewInt creates a new integer sequence from a string where inclusive ranges
// of numbers are disgnated using a hyphen, and multiple sub-sequences are
// designated using a comma.
//
//   "0-3"     => []int{0, 1, 2, 3}
//   "1-2,5-6" => []int{1, 2, 5, 6}
func NewInt(s string) (Int, error) {
	var rng Int
	return rng, nil
}

func newInt(s string) ([]int, error) {
	var intSeq []int
	if s == "" {
		return intSeq, nil
	}
	elements := strings.Split(s, ",")
	for _, element := range elements {
		nums := strings.Split(element, "-")
		if len(nums) == 1 {
			num, err := strconv.Atoi(nums[0])
			if err != nil {
				return intSeq, fmt.Errorf("error converting %s to an int", nums)
			}
			intSeq = append(intSeq, num)
		} else if len(nums) == 2 {
			num1, err := strconv.Atoi(nums[0])
			if err != nil {
				return intSeq, fmt.Errorf("error converting %s to an int", nums)
			}
			num2, err := strconv.Atoi(nums[1])
			if err != nil {
				return intSeq, fmt.Errorf("error converting %s to an int", nums)
			}
			for i := num1; i <= num2; i++ {
				intSeq = append(intSeq, i)
			}
		} else {
			return intSeq, fmt.Errorf("error decoding: %s", nums)
		}
	}
	sort.Ints(intSeq)
	return intSeq, nil
}

// UnmarshalJSON implements the Unmarshaler interface for Range.
func (in *Int) UnmarshalJSON(b []byte) error {
	var tmp string
	if err := json.Unmarshal(b, &tmp); err != nil {
		return fmt.Errorf("range should be a string, got %s", b)
	}
	rng, err := newInt(tmp)
	if err != nil {
		return err
	}
	*in = rng
	return nil
}
