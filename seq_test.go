// Copyright (c) 2020-2022 The seq developers. All rights reserved.
// Project site: https://github.com/goinvest/seq
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE file for the project.

package seq

import (
	"encoding/json"
	"testing"
)

func TestNewInt(t *testing.T) {
	testCases := []struct {
		given    string
		expected []int
	}{
		{"1", []int{1}},
		{"1-2", []int{1, 2}},
		{"0-2", []int{0, 1, 2}},
		{"10-12", []int{10, 11, 12}},
		{"1-2,4", []int{1, 2, 4}},
		{"4,1-2", []int{1, 2, 4}},
		{"1-2,6-7", []int{1, 2, 6, 7}},
		{"13,25,37", []int{13, 25, 37}},
	}
	for _, tc := range testCases {
		calc, err := newInt(tc.given)
		if err != nil {
			t.Errorf("error creating Int: %s", err)
		}
		if len(calc) != len(tc.expected) {
			t.Errorf("len mismatch, got %d, expected %d", len(calc), len(tc.expected))
		} else {
			// If the lengths don't match, skip checking each element to avoid
			// panicing due to an index out of range.
			for i := 0; i < len(calc); i++ {
				if calc[i] != tc.expected[i] {
					t.Errorf("idx %d mismatch: got %d, expected %d", i, calc[i], tc.expected[i])
				}
			}
		}
	}
}
func TestUnmarshalRange(t *testing.T) {
	testCases := []struct {
		given    string
		expected []int
	}{
		{"1", []int{1}},
		{"1-2", []int{1, 2}},
		{"0-2", []int{0, 1, 2}},
		{"10-12", []int{10, 11, 12}},
		{"1-2,4", []int{1, 2, 4}},
		{"4,1-2", []int{1, 2, 4}},
		{"1-2,6-7", []int{1, 2, 6, 7}},
		{"13,25,37", []int{13, 25, 37}},
	}
	for _, tc := range testCases {
		v, err := json.Marshal(tc.given)
		if err != nil {
			t.Errorf("error1: %s", err)
		}
		var calc Int
		if err := json.Unmarshal(v, &calc); err != nil {
			t.Errorf("error: %s", err)
		}
		if len(calc) != len(tc.expected) {
			t.Errorf("len mismatch, got %d, expected %d", len(calc), len(tc.expected))
		} else {
			// If the lengths don't match, skip checking each element to avoid
			// panicing due to an index out of range.
			for i := 0; i < len(calc); i++ {
				if calc[i] != tc.expected[i] {
					t.Errorf("idx %d mismatch: got %d, expected %d", i, calc[i], tc.expected[i])
				}
			}
		}
	}
}
