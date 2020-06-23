package matcher

import (
	"bytes"
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/praveen4g0/comparator/pkg/format"
)

type JSONMatcher struct {
	JSONToMatch interface{}
}

func (matcher *JSONMatcher) Match(actual interface{}) (success bool, err error) {
	actualString, expectedString, err := matcher.prettyPrint(actual)
	if err != nil {
		return false, err
	}

	var aval interface{}
	var eval interface{}

	json.Unmarshal([]byte(actualString), &aval)
	json.Unmarshal([]byte(expectedString), &eval)

	return reflect.DeepEqual(aval, eval), nil
}

func (matcher *JSONMatcher) prettyPrint(actual interface{}) (actualFormatted, expectedFormatted string, err error) {
	actualString, aok := toString(actual)
	expectedString, eok := toString(matcher.JSONToMatch)

	if !(aok && eok) {
		return "", "", fmt.Errorf("JSONMatcher requires a string or stringer.  Got:\n%s", format.Object(actual, 1))
	}

	abuf := new(bytes.Buffer)
	ebuf := new(bytes.Buffer)

	if err := json.Indent(abuf, []byte(actualString), "", "  "); err != nil {
		return "", "", err
	}

	if err := json.Indent(ebuf, []byte(expectedString), "", "  "); err != nil {
		return "", "", err
	}

	return actualString, expectedString, nil
}
