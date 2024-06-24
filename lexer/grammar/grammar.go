package grammar

func Subsequent(r rune) bool {
    return Alpha(r) || Digit(r)
}

func Initial(r rune) bool {
    return Alpha(r)
}

func Digit(r rune) bool {
    return r >= '0' && r <= '9'
}

func Lowercase(r rune) bool {
    return r >= 'a' && r <= 'z'
}

func Uppercase(r rune) bool {
    return r >= 'A' && r <= 'Z'
}

func Alpha(r rune) bool {
    return Uppercase(r) || Lowercase(r)
}
