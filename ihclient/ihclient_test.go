package ihclient

import (
	"strings"
	"testing"
)

func TestGetLiteratureInfoById(t *testing.T) {
	want := "In the context of planar holography"
	if got := GetLiteratureInfoById("1794340"); !strings.Contains(got, want) {
		t.Errorf("GetLiteratureInfoById(451647) = %q, \n\nResponse didn't contain value %q", got, want)
	}
}

func TestGetLiteratureInfoByArxiv(t *testing.T) {
	if got := GetLiteratureInfoByArxiv("2005.01735"); !strings.Contains(got, "Massive Conformal Symmetry and Integrability for Feynman Integrals") {
		t.Errorf("GetLiteratureInfoByArxiv(2005.01735) = %q, \n\n\nResponse didn't contain the title", got)
	}
}
