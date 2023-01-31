package roman

import "testing"

func TestRoman(t *testing.T) {
	s := "MCMXCIV"
	i := 1994

	rti := RomanToInt(s)
	itr := IntToRoman(i)
	if rti != i || itr != s {
		t.Error("Unexpected result:", t)
	}
}
