package leap

func isDivisibleBy(a, b int) bool {
    return a % b == 0
}

func IsLeapYear(y int) bool {
    if isDivisibleBy(y, 4) {
        if isDivisibleBy(y, 100) && !isDivisibleBy(y, 400) {
            return false
        }
        return true
    }
    return false
}
