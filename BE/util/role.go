package util

const (
	Admin = "admin"
	User  = "user"
)

func IsValidRoleOrNot(role string) bool {
	switch role {
	case Admin, User:
		return true
	}
	return false
}
