package data

import (
	"bytes"
	"testing"
)

func TestReadPasswordData(t *testing.T) {
	data := []byte(`1-3 a: abcde`)
	r := bytes.NewReader(data)

	entries, _ := ReadPasswordData(r)
	if len(*entries) != 1 {
		t.Errorf("unexpected number of entries %d; expected %d", len(*entries), 1)
	}
}
