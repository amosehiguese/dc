package coreauth

import "fmt"

func GetRoleCredentials(role Role) ([]string, error) {
	var credentials []string

	switch role {
	case AdminRole:
		credentials = []string{
			// User
			UserCreateCredential,
			UserReadCredential,
			UserUpdateCredential,
			UserDeleteCredential,

			// Appointment
			AppointmentCreateCredential,
			AppointmentReadCredential,
			AppointmentUpdateCredential,
			AppointmentDeleteCredential,

			// Doctor
			DoctorCreateCredential,
			DoctorReadCredential,
			DoctorUpdateCredential,
			DoctorDeleteCredential,

			// Slot
			SlotCreateCredential,
			SlotReadCredential,
			SlotUpdateCredential,
			SlotDeleteCredential,
		}
	case DoctorRole:
		credentials = []string{
			// Appointment
			AppointmentCreateCredential,
			AppointmentReadCredential,
			AppointmentUpdateCredential,
			AppointmentDeleteCredential,

			// Doctor
			DoctorReadCredential,
			DoctorUpdateCredential,

			// Slot
			SlotCreateCredential,
			SlotReadCredential,
			SlotUpdateCredential,
			SlotDeleteCredential,
		}
	case UserRole:
		credentials = []string{
			// Appointment
			AppointmentCreateCredential,
			AppointmentReadCredential,
			AppointmentUpdateCredential,
			AppointmentDeleteCredential,

			// Doctor
			DoctorReadCredential,

			// Slot
			SlotReadCredential,
			SlotDeleteCredential,
		}
	default:
		return nil, fmt.Errorf("role '%v' does not exist", role)
	}

	return credentials, nil
}
