// Copyright 2013 Andreas Louca. All rights reserved.
// Use of this source code is goverend by a BSD-style
// license that can be found in the LICENSE file.

package gonavmap

import (
	"testing"
)

var (
	testData = map[string]interface{}{
		"Int1": 1,
		"Str1": "lala",
		"testing": map[string]interface{}{
			"Int2": 2,
			"Str2": "lolo",
			"further": map[string]interface{}{
				"Int3": 3,
				"Str3": "lele",
			},
		},
	}
)

func TestTraverse(t *testing.T) {
	f := Get(testData, "testing.further")

	if f != nil {
		t.Logf("Got %+v\n", f)
	}
}

func TestValue(t *testing.T) {
	v := Value(testData, "testing.further.Str3")

	if v != nil {
		t.Logf("Got %+v\n", v)
		if str, ok := v.(string); ok {
			if str == "lele" {
				t.Logf("Correctly got string %s\n", str)
				return
			}
		} else {
			t.Fatalf("Got type %t return when expecting string!\n", v)
		}
	} else {
		t.Fatal("Got nil return when expecting value!\n")
	}
}

func TestSet(t *testing.T) {
	e := Set(testData, "testing.further.Lolo", 5)

	if e == nil {
		v := Value(testData, "testing.further.Lolo")
		if num, ok := v.(int); ok {
			t.Logf("Got int %d\n", num)
		} else {
			t.Fatalf("Invalid return type!\n")
		}
	} else {
		t.Fatalf("Error while setting value: %s\n", e.Error())
	}

}
