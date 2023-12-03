package utilities

import (
	"reflect"
	"testing"
)

func TestGetLines(t *testing.T) {
	input := `This is
	a multiline string
	with a few lines`

	got := GetLines(input)
	want := []string{"This is", "a multiline string", "with a few lines"}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, wanted %v", got, want)
	}
}
