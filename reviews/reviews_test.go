package reviews_test

import (
	"os"
	"testing"

	"github.com/jacoelho/openblind/reviews"
)

func TestParse(t *testing.T) {
	fixture, err := os.Open("../fixtures/reviews.html")
	if err != nil {
		t.Fatalf("failed to read fixture: %v", err)
	}

	defer fixture.Close()

	res, err := reviews.Parse(fixture)
	if err != nil {
		t.Fatalf("failed to parse: %v", err)
	}

	if len(res) != 10 {
		t.Errorf("expect %d, got %d", 10, len(res))
	}
}
