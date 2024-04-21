package pqb

import (
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"github.com/google/uuid"
)

var findPositionalArgsRegex = regexp.MustCompile(`\$(\d+)`)

// IncrementArgs increments the args in a SQL statement by amount.
// eg: $1 becomes $2, $2 becomes $3, etc.
func IncrementArgs(sql string, amount int) string {
	// find all occurrences of $n
	matches := findPositionalArgsRegex.FindAllStringSubmatch(sql, -1)

	if len(matches) == 0 {
		return sql
	}

	// sort matches by $n descending
	sort.Slice(matches, func(i, j int) bool {
		a, _ := strconv.Atoi(matches[i][1])
		b, _ := strconv.Atoi(matches[j][1])
		return a > b
	})

	// temporary replacement markers.
	tempReplacements := make(map[string]string)

	// replace each $n with a unique UUID marker
	// reverse loop to avoid replacing the same $n twice
	for _, match := range matches {
		// olgArg is `$n` eg: `$3`
		oldArg := match[0]

		// oldNum is n in $n eg: 3
		oldNum, _ := strconv.Atoi(match[1])

		// replace the $n with a unique UUID marker
		tempMarker := uuid.New().String()
		sql = strings.ReplaceAll(sql, oldArg, tempMarker)

		// store the UUID marker and the incremented $n
		tempReplacements[tempMarker] = fmt.Sprintf("$%d", oldNum+amount)
	}

	// replace each UUID marker with its incremented value
	for tempMarker, newArg := range tempReplacements {
		sql = strings.ReplaceAll(sql, tempMarker, newArg)
	}

	return sql
}
