package providerfunding_test

import (
	"carecontinuity/internal/continuity/providerfunding"
	"reflect"
	"testing"
)

func TestProviderFundingBucketPublicBehavior(t *testing.T) {
	defer func() {
		if v := recover(); v != nil {
			t.Fatalf("empty provider loader panicked: %v", v)
		}
	}()
	c := providerfunding.NewCoordinator()
	if got := c.FirstFunding("provider-new", "grant-1"); !reflect.DeepEqual(got, []string{"grant-1"}) {
		t.Fatalf("unexpected funding bucket %v", got)
	}
}
