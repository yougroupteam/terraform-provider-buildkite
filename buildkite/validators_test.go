package buildkite

import (
	"testing"
)

func TestValidateStepType(t *testing.T) {
	testCases := []struct {
		val         interface{}
		expectedErr string
	}{
		{
			val: "manual",
		},
		{
			val: "script",
		},
		{
			val: "trigger",
		},
		{
			val: "waiter",
		},
		{
			val:         "my_val",
			expectedErr: "expected resource.property to be one of [manual script trigger waiter], got my_val",
		},
	}

	matchErr := func(errs []error, s string) bool {
		// err must match one provided
		for _, err := range errs {
			if s == err.Error() {
				return true
			}
		}

		return false
	}

	for i, tc := range testCases {
		_, errs := validateStepType()(tc.val, "resource.property")

		if len(errs) == 0 && tc.expectedErr == "" {
			continue
		}

		if len(errs) != 0 && tc.expectedErr == "" {
			t.Fatalf("expected test case %d to produce no errors, got %v", i, errs)
		}

		if !matchErr(errs, tc.expectedErr) {
			t.Fatalf("expected test case %d to produce error matching \"%s\", got %v", i, tc.expectedErr, errs)
		}
	}
}
