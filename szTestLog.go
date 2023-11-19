/*
   Szerszam Test Log Utility: szTestLog.
   Copyright (C) 2023  Leslie Dancsecs

   This program is free software: you can redistribute it and/or modify
   it under the terms of the GNU General Public License as published by
   the Free Software Foundation, either version 3 of the License, or
   (at your option) any later version.

   This program is distributed in the hope that it will be useful,
   but WITHOUT ANY WARRANTY; without even the implied warranty of
   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
   GNU General Public License for more details.

   You should have received a copy of the GNU General Public License
   along with this program.  If not, see <https://www.gnu.org/licenses/>.
*/

//nolint:goDot // Ok.
/*
Package szTestLog provides a convenience package for linking Szerszam packages
together:

- github.com/dancsecs/szTest
- github.com/dancsecs/szLog

*/
package szTestLog

import (
	"testing"

	"github.com/dancsecs/szLog"
	"github.com/dancsecs/szTest"
)

func new(chk *szTest.Chk) *szTest.Chk {
	chk.T().Helper()
	origLevel := szLog.SetLevel(szLog.DebugLevel)
	chk.PushPreReleaseFunc(func() error {
		szLog.SetLevel(origLevel)
		return nil
	})
	return chk
}

// CaptureNothing returns a new szTest object without any logs or standard io
// being captured.  Further the szLog package's state is stored and set to
// capture debug messages without line locations.
func CaptureNothing(t *testing.T) *szTest.Chk {
	t.Helper()
	return new(szTest.CaptureNothing(t))
}

// CaptureStdout returns a new szTest object with os.Stdout being captured for
// testing by the Stdout() method.  Further the szLog package's state is stored
// and set to capture debug messages without line locations.
func CaptureStdout(t *testing.T) *szTest.Chk {
	t.Helper()
	return new(szTest.CaptureStdout(t))
}

// CaptureLog returns a new szTest object with log.Setout() being captured for
// testing by the Log() method.  Further the szLog package's state is stored
// and set to capture debug messages without line locations.
func CaptureLog(t *testing.T) *szTest.Chk {
	t.Helper()
	return new(szTest.CaptureLog(t))
}

// CaptureLogAndStdout returns a new szTest object with os.Stdout and
// log.Setout() being captured for testing by the Stdout() and Log() methods
// respectfully.  Further the szLog package's state is stored and set to
// capture debug messages without line locations.
func CaptureLogAndStdout(t *testing.T) *szTest.Chk {
	t.Helper()
	return new(szTest.CaptureLogAndStdout(t))
}

// CaptureLogAndStderr returns a new szTest object with os.Stderr and
// log.Setout() being captured for testing by the Stderr() and Log() methods
// respectfully.  Further the szLog package's state is stored and set to
// capture debug messages without line locations.
func CaptureLogAndStderr(t *testing.T) *szTest.Chk {
	t.Helper()
	return new(szTest.CaptureLogAndStderr(t))
}

// CaptureLogAndStderrAndStdout returns a new szTest object with os.Stdout,
// os.Stderr and log.Setout() being captured for testing by the Stdout(),
// Stderr() and Log() methods respectfully.  Further the szLog package's state
// is stored and set to capture debug messages without line locations.
func CaptureLogAndStderrAndStdout(t *testing.T) *szTest.Chk {
	t.Helper()
	return new(szTest.CaptureLogAndStderrAndStdout(t))
}

// CaptureLogWithStderr returns a new szTest object with os.Stderr and
// log.Setout() being captured in the same log for testing by either the
// Stderr() or Log() methods.  Further the szLog package's state is stored and
// set to capture debug messages without line locations.
func CaptureLogWithStderr(t *testing.T) *szTest.Chk {
	t.Helper()
	return new(szTest.CaptureLogWithStderr(t))
}

// CaptureLogWithStderrAndStdout returns a new szTest object with os.Stdout and
// os.Stderr with log.Setout() being captured in the same log for testing by
// the Stdout() and either the Stderr() or Log() methods respectuffly.  Further
// the szLog package's state is stored and set to capture debug messages
// without line locations.
func CaptureLogWithStderrAndStdout(t *testing.T) *szTest.Chk {
	t.Helper()
	return new(szTest.CaptureLogWithStderrAndStdout(t))
}

// CaptureStderr returns a new szTest object with os.Stderr being captured for
// testing by the Stderr() method.  Further the szLog package's state is stored
// and set to capture debug messages without line locations.
func CaptureStderr(t *testing.T) *szTest.Chk {
	t.Helper()
	return new(szTest.CaptureStderr(t))
}

// CaptureStderrAndStdout returns a new szTest object with os.Stdout and
// log.Seterr() being captured for testing by the Stdout() and Stderr() methods
// respectfully.  Further the szLog package's state is stored and set to
// capture debug messages without line locations.
func CaptureStderrAndStdout(t *testing.T) *szTest.Chk {
	t.Helper()
	return new(szTest.CaptureStderrAndStdout(t))
}
