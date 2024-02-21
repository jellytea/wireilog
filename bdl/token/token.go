// Copyright 2023 LangVM Project
// Copyright 2024 JetERA Creative
// This Source Code Form is subject to the terms of the Mozilla Public License, v. 2.0
// that can be found in the LICENSE file and https://mozilla.org/MPL/2.0/.

package token

const (
	ILLEGAL int = iota

	IDENT // main
	INT   // 12345

	KEYWORD_BEGIN

	OPERATOR_BEGIN

	AND     // &
	OR      // |
	XOR     // ^
	AND_NOT // &^
	NOT     // !

	OPERATOR_END

	STMT_END

	PERIOD   // .
	ELLIPSIS // ...

	KEYWORD_END

	DELIMITER_BEGIN

	LPAREN // (
	LBRACK // [
	LBRACE // {
	COMMA  // ,

	RPAREN    // )
	RBRACK    // ]
	RBRACE    // }
	SEMICOLON // ;
	COLON     // :

	DELIMITER_END
)

var KeywordLiterals = [...]string{

	AND:     "&",
	OR:      "|",
	XOR:     "^",
	AND_NOT: "&^",
	NOT:     "!",

	LPAREN: "(",
	LBRACK: "[",
	LBRACE: "{",
	COMMA:  ",",
	PERIOD: ".",

	RPAREN:    ")",
	RBRACK:    "]",
	RBRACE:    "}",
	SEMICOLON: ";",
	COLON:     ":",
}

var KeywordEnums = map[string]int{}

func IsKeyword(term int) bool { return KEYWORD_BEGIN <= term && term <= KEYWORD_END }

var Delimiters = map[rune]int{
	'{': LBRACE,
	'}': RBRACE,
	'[': LBRACK,
	']': RBRACK,
	'(': LPAREN,
	')': RPAREN,

	',': COMMA,
	';': SEMICOLON,

	'\n': 1, // Newline, might be a statement terminator.
}

func IsDelimiter(ch rune) bool {
	return Delimiters[ch] != 0
}

func init() {
	for i := KEYWORD_BEGIN; i < KEYWORD_END; i++ {
		KeywordEnums[KeywordLiterals[i]] = i
	}
}
