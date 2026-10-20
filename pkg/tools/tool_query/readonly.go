package tool_query

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"unicode"

	fbcontext "github.com/firebolt-db/firebolt-go-sdk/context"
	"github.com/firebolt-db/firebolt-go-sdk/statement"
)

// writeKeywords are rejected anywhere in a read-only query, outside of literals and comments.
// This catches data-modifying CTEs such as `WITH x AS (...) INSERT INTO ...` and `SELECT ... INTO t`,
// which creates a table.
var writeKeywords = map[string]struct{}{
	"INSERT": {}, "UPDATE": {}, "DELETE": {}, "MERGE": {}, "TRUNCATE": {}, "COPY": {},
	"CREATE": {}, "ALTER": {}, "DROP": {}, "UNDROP": {}, "VACUUM": {},
	"GRANT": {}, "REVOKE": {}, "CALL": {}, "INTO": {},
}

var errMultipleStatements = errors.New("bad request: read-only mode: only a single statement is allowed")

// validateReadOnly accepts a single SELECT or WITH statement and rejects anything else.
// It is defence in depth on top of the Firebolt role's grants, not a substitute for them:
// a keyword scan can be bypassed, the role cannot.
func validateReadOnly(ctx context.Context, sql string) error {
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

	// The driver splits statements on the client with its own MySQL-style tokenizer and sends each
	// piece as a separate request, so it must see exactly one statement too.
	stmt, err := statement.MakeStmt(nil, sql, fbcontext.GetPreparedStatementsStyle(ctx))
	if err != nil {
		return fmt.Errorf("bad request: read-only mode: %w", err)
	}
	if len(stmt.Queries) != 1 {
		return errMultipleStatements
	}
	if _, ok := stmt.Queries[0].(*statement.SingleStatement); !ok {
		return errors.New("bad request: read-only mode: only SELECT or WITH statements are allowed")
	}

	return nil
}

// scanWords returns the upper-cased bare words of sql, skipping string literals, quoted
// identifiers and comments. It follows the Firebolt lexer and fails if anything but comments
// follows a `;`.
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
		case c == '$' && i+1 < n && r[i+1] == '$':
			j, err := skipDollarQuoted(r, i)
			if err != nil {
				return nil, err
			}
			i = j
		case c == '\'' || c == '"':
			j, err := skipQuoted(r, i, false)
			if err != nil {
				return nil, err
			}
			i = j
		case isIdentStart(c) || c == '$' && i+1 < n && isIdentStart(r[i+1]):
			j := i + 1
			for j < n && isIdentPart(r[j]) {
				j++
			}
			if j == i+1 && (c == 'E' || c == 'e') && j < n && r[j] == '\'' {
				k, err := skipQuoted(r, j, true)
				if err != nil {
					return nil, err
				}
				i = k
				break
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

// skipDollarQuoted returns the index just past the $$...$$ literal starting at r[start].
// Firebolt has no tagged $tag$ literals: `$tag$` is an identifier.
func skipDollarQuoted(r []rune, start int) (int, error) {
	for i := start + 2; i+1 < len(r); i++ {
		if r[i] == '$' && r[i+1] == '$' {
			return i + 2, nil
		}
	}
	return 0, errors.New("bad request: read-only mode: unterminated dollar-quoted string")
}

// isIdentStart and isIdentPart follow Firebolt's unquoted identifiers, `[A-Za-z_][A-Za-z0-9_$]*`.
func isIdentStart(c rune) bool {
	return c >= 'A' && c <= 'Z' || c >= 'a' && c <= 'z' || c == '_'
}

func isIdentPart(c rune) bool {
	return isIdentStart(c) || c >= '0' && c <= '9' || c == '$'
}
