package valid

import "regexp"

var (
	UsernameRegex = regexp.MustCompile(`^[\u4e00-\u9fa5a-zA-Z0-9_@]{4,16}$`)
	PasswordRegex = regexp.MustCompile(`^[a-zA-Z0-9_@]{6,18}$`)
)
