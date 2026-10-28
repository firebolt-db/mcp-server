package tool_query

import (
	"errors"
	"fmt"
	"strings"
	"unicode"
)

// writeKeywords are rejected anywhere in a read-only query, outside of literals and comments.
// This catches data-modifying CTEs such as `WITH x AS (...) INSERT INTO ...`.
var writeKeywords = map[string]struct{}{
	"INSERT": {}, "UPDATE": {}, "DELETE": {}, "MERGE": {}, "TRUNCATE": {}, "COPY": {},
	"CREATE": {}, "ALTER": {}, "DROP": {}, "UNDROP": {}, "VACUUM": {},
	"GRANT": {}, "REVOKE": {}, "CALL": {},
}

var errMultipleStatements = errors.New("bad request: read-only mode: only a single statement is allowed")

// validateReadOnly accepts a single SELECT or WITH statement and rejects anything else.
// It is defence in depth on top of the Firebolt role's grants, not a substitute for them:
// a keyword scan can be bypassed, the role cannot.
func validateReadOnly(sql string) error {
	words, err := scanWords(sql)
	if err != nil {
		return err
	}
	if len(words) == 0 {
		return errors.New("bad request: read-only mode: query contains no statement")
	}

	if first := words[0]; first != "SELECT" && first != "WITH" {
		return fmt.Errorf("bad request: read-only mode: only SELECT or WITH statements are allowed, got %s", first)
	}
	for _, w := range words {
		if _, ok := writeKeywords[w]; ok {
			return fmt.Errorf("bad request: read-only mode: %s is not allowed", w)
		}
	}

	return nil
}

// scanWords returns the upper-cased bare words of sql, skipping string literals, quoted
// identifiers and comments. It fails if anything but comments follows a `;`.
func scanWords(sql string) ([]string, error) {
	var (
		words []string
		ended bool
		r     = []rune(sql)
		n     = len(r)
	)

	for i := 0; i < n; {
		c := r[i]
		switch {
		case unicode.IsSpace(c):
			i++
		case c == '-' && i+1 < n && r[i+1] == '-':
			for i < n && r[i] != '\n' {
				i++
			}
		case c == '/' && i+1 < n && r[i+1] == '*':
			j := i + 2
			for j+1 < n && (r[j] != '*' || r[j+1] != '/') {
				j++
			}
			if j+1 >= n {
				return nil, errors.New("bad request: read-only mode: unterminated comment")
			}
			i = j + 2
		case ended:
			return nil, errMultipleStatements
		case c == ';':
			ended = true
			i++
		case c == '$':
			j, err := skipDollarQuoted(r, i)
			if err != nil {
				return nil, err
			}
			i = j
		case c == '\'' || c == '"':
			backslashEscapes := c == '\'' && isEscapeStringPrefix(r, i)
			j, err := skipQuoted(r, i, backslashEscapes)
			if err != nil {
				return nil, err
			}
			i = j
		case isIdentStart(c):
			j := i
			for j < n && isIdentPart(r[j]) {
				j++
			}
			words = append(words, strings.ToUpper(string(r[i:j])))
			i = j
		default:
			i++
		}
	}

	return words, nil
}

// skipQuoted returns the index just past the quoted run starting at r[start].
// A doubled quote is an escaped quote; in E'...' strings a backslash also escapes.
func skipQuoted(r []rune, start int, backslashEscapes bool) (int, error) {
	quote := r[start]
	for i := start + 1; i < len(r); i++ {
		switch {
		case backslashEscapes && r[i] == '\\':
			i++
		case r[i] != quote:
		case i+1 < len(r) && r[i+1] == quote:
			i++
		default:
			return i + 1, nil
		}
	}
	return 0, errors.New("bad request: read-only mode: unterminated quoted string")
}

// skipDollarQuoted returns the index just past a $$...$$ or $tag$...$tag$ literal starting
// at r[start], or start+1 when the `$` does not open one (for example a $1 parameter).
func skipDollarQuoted(r []rune, start int) (int, error) {
	j := start + 1
	if j < len(r) && isIdentStart(r[j]) {
		for j < len(r) && (isIdentStart(r[j]) || unicode.IsDigit(r[j])) {
			j++
		}
	}
	if j >= len(r) || r[j] != '$' {
		return start + 1, nil
	}
	tag := r[start : j+1]
	for k := j + 1; k+len(tag) <= len(r); k++ {
		if string(r[k:k+len(tag)]) == string(tag) {
			return k + len(tag), nil
		}
	}
	return 0, errors.New("bad request: read-only mode: unterminated dollar-quoted string")
}

// isEscapeStringPrefix reports whether the quote at r[i] opens an E'...' escape string.
func isEscapeStringPrefix(r []rune, i int) bool {
	if i == 0 || (r[i-1] != 'E' && r[i-1] != 'e') {
		return false
	}
	return i == 1 || !isIdentPart(r[i-2])
}

func isIdentStart(c rune) bool {
	return unicode.IsLetter(c) || c == '_'
}

// isIdentPart excludes `$`: Firebolt identifiers cannot contain it, so `LIKE$$...$$` is a
// keyword followed by a literal.
func isIdentPart(c rune) bool {
	return unicode.IsLetter(c) || unicode.IsDigit(c) || c == '_'
}
