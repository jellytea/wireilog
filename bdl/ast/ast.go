// Copyright 2023-2024 JetERA Creative
// This Source Code Form is subject to the terms of the Mozilla Public License, v. 2.0
// that can be found in the LICENSE file and https://mozilla.org/MPL/2.0/.

package ast

import "wireilog/bdl/token"

type PosRange struct {
	From, To token.Position
}

type Token struct {
	PosRange
	Kind    int
	Literal string
}

type Ident struct {
	Token
}

type NetDecl struct {
}

type ResistorExpr struct {
	Value uint64
}

type InductorExpr struct {
	Value uint64
}

type CapacitorExpr struct {
	Value uint64
}
