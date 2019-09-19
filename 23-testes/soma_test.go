package soma

import (
	"testing"
)

func TestSoma(t *testing.T) {

	if i := Soma(4, 4); i != 8 {
		t.Errorf("Deveria ser %d mas recebi %d", 8, i)
		return
	}

	if i := Soma(1, 2); i != 3 {
		t.Errorf("Deveria ser %d mas recebi %d", 3, i)
		return
	}

}