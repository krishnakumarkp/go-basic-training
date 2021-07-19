package mascot_test

import (
	"testing"

	"cybage.com/go-demo/mascot"
)

func TestBestMascot(t *testing.T) {
	if mascot.BestMascot() != "Gopher" {
		t.Fatal("wrong mascot!")
	}
}
