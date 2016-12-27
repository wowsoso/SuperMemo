package SuperMemo

import "testing"

func testSuperMemo(t *testing.T) {
	v := CalcInterval(5, CalcNewFactor(4, 6))
	if v != 95 {
		t.Errorf("Wanted: 95, Got: %f", v)
	} else {
        t.Log("Test passed")
    }
}