// Copyright 2023 LangVM Project
// Copyright 2024 JetERA Creative
// This Source Code Form is subject to the terms of the Mozilla Public License, v. 2.0
// that can be found in the LICENSE file and https://mozilla.org/MPL/2.0/.

package scanner

import (
	"unicode"
	"wireilog/bdl/token"
)

type BufferScanner struct {
	token.Position // Cursor
	Buffer         []rune
}

func IsMark(ch rune) bool {
	return unicode.IsPunct(ch) || unicode.IsSymbol(ch)
}

// PrintCursor print current line and the cursor position for debug use.
func (s *BufferScanner) PrintCursor() {
	println(string(s.Buffer[s.Offset-s.Column : s.Offset]))
	for i := 0; i < s.Column; i++ {
		print(" ")
	}
	println("^")
}

// Move returns current char and move cursor to the next.
// Move does not error if GetChar does not error.
func (s *BufferScanner) Move() (rune, error) {
	ch, err := s.GetChar()
	if err != nil {
		return 0, err
	}

	if ch == '\n' {
		s.Column = 0
		s.Line++
	} else {
		s.Column++
	}

	s.Offset++

	return ch, nil
}

func (s *BufferScanner) GetChar() (rune, error) {
	if s.Offset == len(s.Buffer) {
		return 0, EOFError{Pos: s.Position}
	}
	return s.Buffer[s.Offset], nil
}

type Scanner struct {
	BufferScanner
}

func (s *Scanner) GotoNextLine() error {
	for {
		ch, err := s.GetChar()
		if err != nil {
			return err
		}
		if ch == '\n' {
			return nil
		}
		_, _ = s.Move()
	}
}

func (s *Scanner) SkipWhitespace() error {
	for {
		ch, err := s.GetChar()
		if err != nil {
			return err
		}
		switch ch {
		case ' ':
		case '\t':
		case '\r':
		default:
			return nil
		}
		_, _ = s.Move()
	}
}

// ScanLineComment scans line comment.
func (s *Scanner) ScanLineComment() ([]rune, error) {
	begin := s.Offset
	err := s.GotoNextLine()
	if err != nil {
		return nil, err
	}
	return s.Buffer[begin:s.Offset], nil
}

// ScanQuotedComment scans until "*/".
// Escape char does NOT affect.
func (s *Scanner) ScanQuotedComment() ([]rune, error) {
	begin := s.Offset
	for {
		end := s.Offset
		ch, err := s.Move()
		if err != nil {
			return nil, err
		}
		if ch == '*' {
			ch, err := s.Move()
			if err != nil {
				return nil, err
			}
			if ch == '/' {
				return s.Buffer[begin:end], nil
			}
		}
	}
}

// ScanComment scans and distinguish line comment or quoted comment.
func (s *Scanner) ScanComment() ([]rune, error) {
	ch, err := s.Move()
	if err != nil {
		return nil, err
	}
	switch ch {
	case '/':
		return s.ScanLineComment()
	case '*':
		return s.ScanQuotedComment()
	default:
		return nil, FormatError{Pos: s.Position}
	}
}

func (s *Scanner) ScanDigit() ([]rune, error) {
	ch, _ := s.Move()

	digits := []rune{ch}

	// TODO

	return digits, nil
}

// ScanWord scans and accepts only letters, digits and underlines.
// No valid string found when returns empty []rune.
func (s *Scanner) ScanWord() (str []rune, err error) {
	for {
		ch, err := s.GetChar()
		if err != nil {
			return nil, err
		}
		switch {
		case unicode.IsDigit(ch):
		case unicode.IsLetter(ch):
		case ch == '_':
		default: // Terminate
			if len(str) == 0 {
				return nil, FormatError{Pos: s.Position}
			}
			return str, nil
		}

		_, _ = s.Move()
		str = append(str, ch)
	}
}

// ScanMarkSeq scans CONSEQUENT marks except/until delimiters.
func (s *Scanner) ScanMarkSeq() (str []rune, err error) {
	for {
		ch, err := s.GetChar()
		if err != nil {
			return nil, err
		}
		if IsMark(ch) && !token.IsDelimiter(ch) {
			str = append(str, ch)
			_, _ = s.Move()
		} else {
			if len(str) == 0 {
				return nil, FormatError{Pos: s.Position}
			}
			return str, nil
		}
	}
}

// ScanToken decides the next way to scan by the cursor.
func (s *Scanner) ScanToken() (int, []rune, error) {
	err := s.SkipWhitespace()
	if err != nil {
		return 0, nil, err
	}

	ch, err := s.GetChar()
	if err != nil {
		return 0, nil, err
	}

	switch {
	case unicode.IsDigit(ch): // Digital literal value
		digit, err := s.ScanDigit()
		if err != nil {
			return 0, nil, err
		}
		return token.INT, digit, nil

	case unicode.IsLetter(ch) || ch == '_': // Keyword OR Ident
		tok, err := s.ScanWord()

		switch err := err.(type) {
		case nil:
		case FormatError:
		default:
			return 0, nil, err
		}

		if keyword := token.KeywordEnums[string(tok)]; keyword != 0 {
			return keyword, tok, nil
		}

		return token.IDENT, tok, nil

	case ch == '/': // Comment
		_, err := s.ScanComment() // TODO
		if err != nil {
			return 0, nil, err
		}
		return s.ScanToken()

	case token.IsDelimiter(ch):
		_, _ = s.Move()
		return token.Delimiters[ch], []rune{ch}, nil

	case IsMark(ch): // Operator
		tok, err := s.ScanMarkSeq()

		switch err := err.(type) {
		case nil:
		case FormatError:
		default:
			return 0, nil, err
		}

		if keyword := token.KeywordEnums[string(tok)]; keyword != 0 {
			return keyword, tok, nil
		}

		return 0, nil, UnknownOperatorError{Pos: s.Position}
	default:
		panic("impossible")
	}
}
