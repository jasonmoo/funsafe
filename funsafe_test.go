package funsafe

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/jasonmoo/funsafe/private"
)

func TestMakeSettable(t *testing.T) {

	v := &private.Example{}

	av := reflect.ValueOf(v).Elem().FieldByName("a")

	if av.CanSet() {
		t.Error("Unexported value found settable by default")
	}

	const (
		before = `&private.Example{a:0}`
		after  = `&private.Example{a:1}`
	)

	if val := fmt.Sprintf("%#v", v); val != before {
		t.Errorf("Expected %q, got %q", before, val)
	}

	MakeSettable(&av)

	if !av.CanSet() {
		t.Error("MakeSettable failed to make value settable")
	}

	av.SetInt(1)

	if val := fmt.Sprintf("%#v", v); val != after {
		t.Errorf("Expected %q, got %q", after, val)
	}

}
