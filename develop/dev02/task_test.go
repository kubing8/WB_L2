package main

import (
	"testing"
)

func TestDecopositionString(t *testing.T) {
	m := make(map[string]string)
	m["a4bc2d5e"] = "aaaabccddddde"
	m["abcd"] = "abcd"
	m["45"] = ""
	m[""] = ""

	// На вход нужно указывать двойные облатные слэши -> на выходе тоже получаем двойные
	m["ed\\\\2se"] = "ed\\\\se" // "ed\\2se" => "ed\\se"
	m["ed\\4se"] = "ed4se"
	m["ed\\\\4se"] = "ed\\\\\\\\se"

	m["qwe\\4\\5"] = "qwe45"
	m["qwe\\45"] = "qwe44444"
	m["qwe\\\\2"] = "qwe\\\\"

	m["eh2k2l\\4n\\\\2dg2"] = "ehhkkl4n\\\\dgg"

	for k, v := range m {
		if s, _ := decompos(k); s != v {
			t.Error("Error in string: ", k)
		}
	}
}
