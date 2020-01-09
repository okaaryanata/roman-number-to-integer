package convert

import (
	"strings"
	"testing"
)

func TestConvertData(t *testing.T) {
	var tests = []struct {
		input    string
		expected int
	}{
		{"XXXX", 0},
		{"IX", 9},
		{"VIII", 8},
		{"LXXXII", 82},
		{"CCC", 300},
		{"DCCC", 800},
		{"MMMDCCXXIV", 0},
		{"CMXCIX", 999},
		{"DLXVII", 567},
	}
	for _, test := range tests {
		if out := ConvertData(test.input); out != test.expected {
			t.Error("Test Failed: {} inputted, {} expected, received : {}", test.input, test.expected, out)
		}
	}
}

func TestFindIdx(t *testing.T) {
	var tests = []struct {
		slice    []string
		str      string
		expected int
	}{
		{[]string{"a", "b", "c", "d"}, "a", 0},
		{[]string{"a", "b", "c", "d"}, "c", 2},
		{[]string{"a", "b", "c", "d"}, "d", 3},
		{[]string{"a", "b", "c", "d"}, "b", 1},
	}
	for _, test := range tests {
		if out := FindIdx(test.slice, test.str); out != test.expected {
			t.Error("Test Failed: {} slice, {} string, {} expected, received : {}", test.slice, test.str, test.expected, out)
		}
	}
}

func TestCheckData(t *testing.T) {
	var tests = []struct {
		str      string
		slice    []string
		expected bool
	}{
		{"oka", []string{"a", "b", "c", "d"}, false},
		{"oka", []string{"a", "b", "c", "d", "oka"}, true},
		{"oka", []string{"a", "b", "c", "d", "okaoa"}, false},
		{"oka", []string{"OKA", "b", "c", "d"}, false},
		{"oka", []string{"OKA", "b", "oka", "d"}, true},
	}
	for _, test := range tests {
		if out := CheckData(test.str, test.slice); out != test.expected {
			t.Error("Test Failed: {} slice, {} string, {} expected, received : {}", test.slice, test.str, test.expected, out)
		}
	}
}

func TestValidateInputData(t *testing.T) {
	var tests = []struct {
		str      string
		expected bool
	}{
		{"glob is I", true},
		{"glob is i", false},
		{"glab is II", true},
		{"prok is VVV", false},
		{"prok is V", true},
		{"pish is X", true},
		{"tegj is L", true},
		{"ghas gahsd is I", false},
		{"ghas is P", false},
		{"is ghas X", false},
		{"glob glob Silver is 34 Credits", true},
		{"glob prok gold is 57800 credits", true},
		{"pish pish Iron is 3910 Credits", true},
		{"pish pisha Iron is 3910 Credits", false},
		{"pish oka Iron is 3910 Credits", false},
		{"pish Iron pish is 3910 Credits", false},
		{"pish Iron pish is glob 3910 Credits", false},
		{"pish Iron oka is 3910 Credits", false},
		{"pish is pisha Iron 3910 Credits", false},
		{"pish pisha Iron is 3910", false},
		{"pish pisha Iron is creadits", false},
		{"how much wood could a woodchuck chuck if a woodchuck could chuck wood ? ", false},
	}
	for _, test := range tests {
		arr := strings.Split(test.str, " ")
		lenArr := len(arr)
		for idx, elm := range arr {
			if idx != lenArr-1 || elm == "Credits" {
				arr[idx] = strings.ToLower(elm)
			}
		}
		if out := ValidateInputData(arr); out != test.expected {
			t.Errorf("Test Failed: {%v} string, {%v} expected, received : {%v}", test.str, test.expected, out)
		}
	}
}

func TestValidateReadData(t *testing.T) {
	var tests = []struct {
		str      string
		expected bool
	}{
		{"how much is pish tegj glob glob ?", true},
		{"how much is pish tegj glob glob glob glob ?", false},
		{"how much is pish tegj glob glob ?", true},
		{"how much is prok prok prok ?", false},
		{"how many Credits is glob prok Silver ?", true},
		{"prok is VVV", false},
		{"how many Credits is glob prok Gold ?", true},
		{"how many Credits is glob prok Iron ?", true},
		{"how many Credits is glob glab Iron ?", true},
		{"glob", false},
		{"many how many Credits is glob glab Iron ?", false},
		{"many is how many Credits is glob glab Iron ?", false},
		{"how many Credits is is glob glab Iron ?", false},
		{"many how many is how many Credits is glob glab Iron ?", false},
		{"how many Credits is Iron glob glab Iron ?", false},
		{"how many Credits is Iron glob glab Iron", false},
		{"glob glob Silver is 34 Credits", false},
		{"glob prok gold is 57800 credits", false},
	}
	for _, test := range tests {
		test.str = strings.ToLower(test.str)
		arr := strings.Split(test.str, " ")
		if out := ValidateReadData(arr); out != test.expected {
			t.Errorf("Test Failed: {%v} string, {%v} expected, received : {%v}", test.str, test.expected, out)
		}
	}
}
