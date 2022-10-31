package main

import (
	"reflect"
	"testing"
)

func TestSetAnagrams(t *testing.T) {
	anagrams := []string{"чистка", "считка", "частик", "кОлба", "бокал", "балок", "Канистра", "старинка",
		"одно", "пятка", "тяпка"}

	testArr := make(map[string][]string)
	/*
		testArr["колба"] = append(testArr["колба"], "балок", "бокал", "колба")
		testArr["канистра"] = append(testArr["канистра"], "канистра", "старинка")
		testArr["пятка"] = append(testArr["пятка"], "пятка", "тяпка")
		testArr["чистка"] = append(testArr["чистка"], "частик", "чистка")

	*/

	testArr["колба"] = []string{"балок", "бокал", "колба"}
	testArr["канистра"] = []string{"канистра", "старинка"}
	testArr["пятка"] = []string{"пятка", "тяпка"}
	testArr["чистка"] = []string{"считка", "частик", "чистка"}

	m := searchAnagram(&anagrams)
	for k := range *m {
		if ok := reflect.DeepEqual(testArr[k], (*m)[k]); !ok {
			t.Error("Not equal: ", testArr[k], "AND: ", (*m)[k])
		}
	}
}
