package hook

import (
	"context"
	"fmt"
	"testing"

	"github.com/vcaesar/tt"
)

func TestAdd(t *testing.T) {
	fmt.Println("hook test...")
	ctx := context.Background()
	e := Start(ctx, BothHookEnabled)
	tt.NotNil(t, e)
}

func TestKey(t *testing.T) {
	k := RawcodetoKeychar(0)
	tt.Equal(t, "error", k)

	r := KeychartoRawcode("error")
	tt.Equal(t, 0, r)
}
