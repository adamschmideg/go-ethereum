// +build flaky

package les

import (
	"testing"
)

func TestFlakyGuy(t *testing.T) {
	if testing.Short() {
		t.Errorf("short and flaky and always fails")
	}
}
