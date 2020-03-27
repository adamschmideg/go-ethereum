// +build flaky

package les

import (
	"testing"
)

func TestHandshake_flaky(t *testing.T) {
	t.Errorf("short and flaky and always fails")
}
