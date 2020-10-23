package coord

import "testing"

func TestNewCartesian(t *testing.T){
	//t.Error("test failed")
	c := NewCartesian(1, 2)
	if (c.x == 1) && (c.y == 2) {
		// test succes
	} else {
		t.Errorf("expected 1 and 2 as coordinates")
	}
}