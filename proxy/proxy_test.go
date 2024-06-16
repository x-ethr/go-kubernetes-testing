package proxy

import (
	"testing"
)

func Test(t *testing.T) {

	t.Run("New", func(t *testing.T) {
		t.Run("Default-Port", func(t *testing.T) {
			instance := New()

			if instance == nil {
				t.Fatalf("Instance is nil")
			} else if instance.options.Port != 8080 {
				t.Fatalf("Instance options port should be 8080, is %d", instance.options.Port)
			}
		})
	})
}
