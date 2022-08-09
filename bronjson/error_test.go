// Copyright (c) 2014 The brsuite developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package bronjson_test

import (
	"testing"

	"github.com/brsuite/brond/bronjson"
)

// TestErrorCodeStringer tests the stringized output for the ErrorCode type.
func TestErrorCodeStringer(t *testing.T) {
	t.Parallel()

	tests := []struct {
		in   bronjson.ErrorCode
		want string
	}{
		{bronjson.ErrDuplicateMethod, "ErrDuplicateMethod"},
		{bronjson.ErrInvalidUsageFlags, "ErrInvalidUsageFlags"},
		{bronjson.ErrInvalidType, "ErrInvalidType"},
		{bronjson.ErrEmbeddedType, "ErrEmbeddedType"},
		{bronjson.ErrUnexportedField, "ErrUnexportedField"},
		{bronjson.ErrUnsupportedFieldType, "ErrUnsupportedFieldType"},
		{bronjson.ErrNonOptionalField, "ErrNonOptionalField"},
		{bronjson.ErrNonOptionalDefault, "ErrNonOptionalDefault"},
		{bronjson.ErrMismatchedDefault, "ErrMismatchedDefault"},
		{bronjson.ErrUnregisteredMethod, "ErrUnregisteredMethod"},
		{bronjson.ErrNumParams, "ErrNumParams"},
		{bronjson.ErrMissingDescription, "ErrMissingDescription"},
		{0xffff, "Unknown ErrorCode (65535)"},
	}

	// Detect additional error codes that don't have the stringer added.
	if len(tests)-1 != int(bronjson.TstNumErrorCodes) {
		t.Errorf("It appears an error code was added without adding an " +
			"associated stringer test")
	}

	t.Logf("Running %d tests", len(tests))
	for i, test := range tests {
		result := test.in.String()
		if result != test.want {
			t.Errorf("String #%d\n got: %s want: %s", i, result,
				test.want)
			continue
		}
	}
}

// TestError tests the error output for the Error type.
func TestError(t *testing.T) {
	t.Parallel()

	tests := []struct {
		in   bronjson.Error
		want string
	}{
		{
			bronjson.Error{Description: "some error"},
			"some error",
		},
		{
			bronjson.Error{Description: "human-readable error"},
			"human-readable error",
		},
	}

	t.Logf("Running %d tests", len(tests))
	for i, test := range tests {
		result := test.in.Error()
		if result != test.want {
			t.Errorf("Error #%d\n got: %s want: %s", i, result,
				test.want)
			continue
		}
	}
}
