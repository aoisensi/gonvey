package models

type User struct {
	ID       int
	Username string `sql:"type:varchar(16);unique_index"`
	Nick     string `sql:"size:255"`
	Password string
}

func (u User) Check() bool {
	return CheckUsername(u.Username)
}

func CheckUsername(username string) bool {
	l := len(username)
	if 1 > l || l > 16 {
		return false
	}
	for _, c := range username {
		if !checkUsernameRune(c) {
			return false
		}
	}
	return true
}

func checkUsernameRune(c rune) bool {
	if 'a' <= c && c <= 'z' {
		return true
	}
	if 'A' <= c && c <= 'Z' {
		return true
	}
	if '0' <= c && c <= '9' {
		return true
	}
	return c == '_'
}
