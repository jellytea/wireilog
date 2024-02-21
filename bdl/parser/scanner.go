// Copyright 2023-2024 JetERA Creative
// This Source Code Form is subject to the terms of the Mozilla Public License, v. 2.0
// that can be found in the LICENSE file and https://mozilla.org/MPL/2.0/.

package parser

import (
	"wireilog/bdl/ast"
	"wireilog/bdl/scanner"
)

type Scanner struct {
	scanner.Scanner
	Token ast.Token
}

func (s *Scanner) Scan() {
	begin := s.Scanner.Position

	kind, lit, err := s.ScanToken()
	if err != nil {
		panic(err)
	}

	s.Token = ast.Token{
		PosRange: ast.PosRange{From: begin, To: s.Position},
		Kind:     kind,
		Literal:  string(lit),
	}
}
