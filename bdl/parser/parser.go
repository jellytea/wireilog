// Copyright 2023-2024 JetERA Creative
// This Source Code Form is subject to the terms of the Mozilla Public License, v. 2.0
// that can be found in the LICENSE file and https://mozilla.org/MPL/2.0/.

package parser

import (
	. "wireilog/bdl/ast"
	"wireilog/bdl/token"
)

type Parser struct {
	Scanner
}

func (p *Parser) ExpectIdent() Ident {
	if p.Token.Kind != token.IDENT {
		panic(UnexpectedNodeError{})
	}

	return Ident{Token: p.Token}
}
