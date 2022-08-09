// Copyright (c) 2017 The brsuite developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package netsync

import "github.com/brsuite/bronlog"

// log is a logger that is initialized with no output filters.  This
// means the package will not perform any logging by default until the caller
// requests it.
var log bronlog.Logger

// DisableLog disables all library log output.  Logging output is disabled
// by default until either UseLogger or SetLogWriter are called.
func DisableLog() {
	log = bronlog.Disabled
}

// UseLogger uses a specified Logger to output package logging info.
// This should be used in preference to SetLogWriter if the caller is also
// using bronlog.
func UseLogger(logger bronlog.Logger) {
	log = logger
}
