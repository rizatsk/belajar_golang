package helper

import "github.com/gofrs/uuid"

func GenerateUuidV6() string {
	uid, _ := uuid.NewV6()
	return uid.String()
}
