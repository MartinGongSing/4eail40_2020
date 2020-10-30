package coord

import "testing"

func TestNewCartesian(t *testing.T){
	//t.Error("test failed")
	c := NewCartesian(1, 2)
	if (c.x != 1) || (c.y != 2) {
		// test succes
	} else {
		t.Errorf("expected 1 and 2 as coordinates")
	}
}


func TestCartesian_Coord(t *testing.T) {
	c := NewCartesian(1, 2)

	tests := map[int]int{
		0: 1,
		1: 2,
	}

	for n, want := range tests {
		t.Run(string(rune(n)), func(t *testing.T) {
			got, err := c.Coord(n)
			if err != nil {
				t.Error(err)
			}
			if got != want {
				t.Errorf("expected %d, but got %d", want, got)
			}
		})
	}

	// test for err
	t.Run("err", func(t *testing.T) {
		_, err := c.Coord(2)
		if err == nil {
			t.Errorf("expected and error for n == 2")
		}
	})
}
