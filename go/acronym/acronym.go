package acronym

import(
    "strings"
    "unicode"
    "unicode/utf8"
)

func Abbreviate(s string) string {
    var acro strings.Builder
    words := strings.FieldsFunc(s, func(r rune) bool {
        return strings.ContainsRune(" \t\n\v-", r)
    })
    if (len(words) == 0) {
        return "";
    }

    for _, w := range words {
        c, _ := utf8.DecodeRuneInString(w)
        if unicode.IsLetter(c) {
            acro.WriteRune(unicode.ToUpper(c))
        }
    }
    return acro.String()
}
