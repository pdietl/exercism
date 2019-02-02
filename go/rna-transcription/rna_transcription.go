package strand

import (
	"strings"
)

// ToRNA returns a DNA strand's RNA complement
func ToRNA(dna string) string {
	var str strings.Builder
	for _, x := range dna {
		switch x {
		case 'G':
			str.WriteRune('C')
		case 'C':
			str.WriteRune('G')
		case 'T':
			str.WriteRune('A')
		case 'A':
			str.WriteRune('U')
		}
	}
	return str.String()
}
