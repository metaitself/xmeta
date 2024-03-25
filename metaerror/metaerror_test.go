package metaerror

import (
	"errors"
	"testing"
)

func TestFromError(t *testing.T) {
	_isDebug = true
	err := Base(1000, "大哭大哭大哭").WithCause(errors.New("dkdkdkdkdkdkdk"))
	t.Logf("%v", err)
}
