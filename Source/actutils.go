/*++

Copyright (C) 2018 Autodesk Inc. (Original Author)

All rights reserved.

Redistribution and use in source and binary forms, with or without modification,
are permitted provided that the following conditions are met:

1. Redistributions of source code must retain the above copyright notice, this
list of conditions and the following disclaimer.
2. Redistributions in binary form must reproduce the above copyright notice,
this list of conditions and the following disclaimer in the documentation
and/or other materials provided with the distribution.

THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS" AND
ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED
WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE FOR
ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES
(INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES;
LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND
ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
(INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS
SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

--*/

//////////////////////////////////////////////////////////////////////////////////////////////////////
// actutils.go
// Utility functions for ACT
//////////////////////////////////////////////////////////////////////////////////////////////////////

package main

import (
	"errors"
	"fmt"
	"log"
	"os"
)

var (
	ErrPythonBuildFailed   = errors.New("failed to build dynamic Python implementation")
	ErrFileDeletionFailed     = errors.New("failed to write to output file")
	ErrReservedKeyword    = errors.New("failed to generate bindings as you are using a reserved keyword")
)

// Keep a map of reserved keywords in Python
var pythonReservedKeywords = map[string]bool{
	"False": true, "None": true, "True": true, "and": true, "as": true, "assert": true, "async": true,
	"await": true, "break": true, "class": true, "continue": true, "def": true, "del": true, "elif": true,
	"else": true, "except": true, "finally": true, "for": true, "from": true, "global": true, "if": true,
	"import": true, "in": true, "is": true, "lambda": true, "nonlocal": true, "not": true, "or": true,
	"pass": true, "raise": true, "return": true, "try": true, "while": true, "with": true, "yield": true,
}

// FileExists returns true if and only if the file in a given path exists
func FileExists(path string) (bool) {
	_, err := os.Stat(path); 
	return !os.IsNotExist(err);
}

// ReservedKeywordExit logs the formatted error, deletes the partially created file, and returns a named error.
func ReservedKeywordExit(bindingPath string, format string, a ...interface{}) error {
	// Format the message using variadic arguments
	msg := fmt.Sprintf(format, a...)

	// Log the error message
	log.Printf("%s", msg)

	// Attempt to delete the partially created file
	log.Printf("Deleting partially created binding file: %s", bindingPath)
	err := os.Remove(bindingPath)
	if err != nil {
		// Log and return a wrapped error with context
		log.Printf("Failed to delete incomplete file %s: %v", bindingPath, err)
		return fmt.Errorf("%w: %v", ErrFileDeletionFailed, err)
	}

	log.Printf("Deleted binding file")

	// Return the reserved keyword error
	return ErrReservedKeyword
}