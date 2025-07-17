package async

import (
	"testing"
	"time"
)

func TestRunAsync(t *testing.T) {
	c := RunAsync(func() int {
		time.Sleep(time.Second)

		return 1
	})
	t.Log(c.Get())
}
