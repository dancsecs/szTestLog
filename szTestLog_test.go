/*
   Szerszam Log Utility: szLog.
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

//nolint:goCritic // ok
package szTestLog

import (
	"fmt"
	"os"
	"testing"

	"github.com/dancsecs/szLog"
)

func TestSzTest_CaptureNothing(t *testing.T) {
	chk := CaptureNothing(t)
	defer chk.Release()
}

func TestSzTest_CaptureStdout(t *testing.T) {
	chk := CaptureStdout(t)
	defer chk.Release()

	fmt.Printf("Here is the line")

	chk.Stdout("Here is the line")
}

func TestSzTest_CaptureLog(t *testing.T) {
	chk := CaptureLog(t)
	defer chk.Release()

	szLog.Debug("This is the debug Mesage")
	szLog.Info("This is the info Mesage")
	szLog.Warn("This is the warn Mesage")
	szLog.Error("This is the error Mesage")

	chk.Log(`
		D: This is the debug Mesage
		I: This is the info Mesage
		W: This is the warn Mesage
		E: This is the error Mesage
		`)
}

func TestSzTest_CaptureLogAndStdout(t *testing.T) {
	chk := CaptureLogAndStdout(t)
	defer chk.Release()

	fmt.Printf("Here is the next line")

	szLog.Debug("This is the debug Mesage")
	szLog.Info("This is the info Mesage")
	szLog.Warn("This is the warn Mesage")
	szLog.Error("This is the error Mesage")

	chk.Stdout("Here is the next line")

	chk.Log(`
		D: This is the debug Mesage
		I: This is the info Mesage
		W: This is the warn Mesage
		E: This is the error Mesage
		`)
}

func TestSzTest_CaptureLogAndStderr(t *testing.T) {
	chk := CaptureLogAndStderr(t)
	defer chk.Release()

	fmt.Fprintf(os.Stderr, "Here is the stderr line")

	szLog.Debug("This is the debug Mesage")
	szLog.Info("This is the info Mesage")
	szLog.Warn("This is the warn Mesage")
	szLog.Error("This is the error Mesage")

	chk.Log(`
		D: This is the debug Mesage
		I: This is the info Mesage
		W: This is the warn Mesage
		E: This is the error Mesage
		`)

	chk.Stderr("Here is the stderr line")
}

func TestSzTest_CaptureLogAndStderrAndStdout(t *testing.T) {
	chk := CaptureLogAndStderrAndStdout(t)
	defer chk.Release()

	fmt.Printf("here is the stdout line")

	fmt.Fprintf(os.Stderr, "Here is the stderr line")

	szLog.Debug("This is the debug Mesage")
	szLog.Info("This is the info Mesage")
	szLog.Warn("This is the warn Mesage")
	szLog.Error("This is the error Mesage")

	chk.Stdout("here is the stdout line")

	chk.Log(`
		D: This is the debug Mesage
		I: This is the info Mesage
		W: This is the warn Mesage
		E: This is the error Mesage
		`)

	chk.Stderr("Here is the stderr line")
}

func TestSzTest_CaptureLogWithStderr(t *testing.T) {
	chk := CaptureLogWithStderr(t)
	defer chk.Release()

	fmt.Fprintf(os.Stderr, "Here is the before stderr line\n")

	szLog.Debug("This is the debug Mesage")
	szLog.Info("This is the info Mesage")
	szLog.Warn("This is the warn Mesage")
	szLog.Error("This is the error Mesage")

	fmt.Fprintf(os.Stderr, "Here is the after stderr line\n")

	chk.Log(`
		Here is the before stderr line
		D: This is the debug Mesage
		I: This is the info Mesage
		W: This is the warn Mesage
		E: This is the error Mesage
		Here is the after stderr line
		`)
}

func TestSzTest_CaptureLogWithStderrAndStdout(t *testing.T) {
	chk := CaptureLogWithStderrAndStdout(t)
	defer chk.Release()

	fmt.Printf("here is the stdout line")

	fmt.Fprintf(os.Stderr, "Here is the before stderr line\n")

	szLog.Debug("This is the debug Mesage")
	szLog.Info("This is the info Mesage")
	szLog.Warn("This is the warn Mesage")
	szLog.Error("This is the error Mesage")

	fmt.Fprintf(os.Stderr, "Here is the after stderr line\n")

	chk.Stdout("here is the stdout line")

	chk.Log(`
		Here is the before stderr line
		D: This is the debug Mesage
		I: This is the info Mesage
		W: This is the warn Mesage
		E: This is the error Mesage
		Here is the after stderr line
		`)
}

func TestSzTest_CaptureStderr(t *testing.T) {
	chk := CaptureStderr(t)
	defer chk.Release()

	fmt.Fprintf(os.Stderr, "Here is the only stderr line")

	chk.Stderr("Here is the only stderr line")
}

func TestSzTest_CaptureStderrAndStdout(t *testing.T) {
	chk := CaptureStderrAndStdout(t)
	defer chk.Release()

	fmt.Printf("here is the only stdout line")

	fmt.Fprintf(os.Stderr, "Here is the only stderr line")

	chk.Stdout("here is the only stdout line")

	chk.Stderr("Here is the only stderr line")
}
