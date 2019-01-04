package hlcuptester

import "testing"
import "regexp"

func TestOptions(t *testing.T) {
	if !EndOfPath("ru/la/la")("/tru/la/la") {
		t.FailNow()
	}

	if !EndOfPath("la/la")("/tru/la/la?test") {
		println(1)

		t.FailNow()
	}

	if EndOfPath("ala/la")("/tru/la/la?test") {
		t.FailNow()
	}

	if !PathRegexp(regexp.MustCompile(`/accounts/\d+/`))("/accounts/8213/?query_id=0") {
		t.FailNow()
	}
}