package coreauth

import "fmt"

type Role string

const (
	AdminRole  Role = "admin"
	DoctorRole Role = "doctor"
	UserRole   Role = "user"
)

func VerifyRole(rl string) (Role, error) {
	var role Role
	switch rl {
	case string(AdminRole):
		role = AdminRole
	case string(DoctorRole):
		role = DoctorRole
	case string(UserRole):
		role = UserRole
	default:
		return "", fmt.Errorf("role '%v' does not exist", rl)
	}

	return role, nil
}
