package props

import (
	"slices"
	"testing"
)

func TestProperties_Append(t *testing.T) {
	var props = Properties{}
	var expected = Properties{"flag", "arg=value"}
	props.AppendAll(map[string]string{
		"flag": "",
		"arg":  "value",
	})
	if !slices.Equal(props, expected) && !slices.Equal(props, []string{"arg=value", "flag"}) {
		t.Errorf("Expected %q, but got %q instead", expected, props)
		t.Fail()
	}
}

func TestProperties_Copy(t *testing.T) {
	var props1 = From(map[string]string{"arg1": "value"})
	var props2 = props1.Copy().Append("arg2", "value")
	if len(*props1) == len(*props2) {
		t.Errorf("%q and %q should be different!", props1, props2)
		t.Fail()
	}
}
