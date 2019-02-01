package bob

import(
    "unicode"
    "regexp"
)

var question_regex *regexp.Regexp
var whitespace_regex *regexp.Regexp

func isShouting(s string) bool {
    letter_seen := false
    for _, c := range s {
        if unicode.IsLetter(c) && unicode.IsLower(c) {
            return false
        }
        if (unicode.IsLetter(c)) {
            letter_seen = true
        }
    }
    return letter_seen
}

func isQuestion(s string) bool {
    return question_regex.MatchString(s)
}

func isAllSpace(s string) bool {
    return whitespace_regex.MatchString(s)
}

func compileRegexes() {
    question_regex = regexp.MustCompile(`\?\s*$`)
    whitespace_regex = regexp.MustCompile(`^\s*$`)
}

func Hey(remark string) string {
    compileRegexes()

    isShouting := isShouting(remark)
    isQuestion := isQuestion(remark)
    isAllSpace := isAllSpace(remark)

    if isShouting && isQuestion {
        return "Calm down, I know what I'm doing!"
    } else if isShouting {
        return "Whoa, chill out!"
    } else if isQuestion {
        return "Sure."
    } else if isAllSpace {
        return "Fine. Be that way!"
    } else {
        return "Whatever."
    }
}
