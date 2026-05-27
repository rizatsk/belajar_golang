package helper

import "github.com/gofrs/uuid"

func GenerateUuidV6() string {
	uid, _ := uuid.NewV6()
	return uid.String()
}

func GenerateUuidV4() string {
	uid, _ := uuid.NewV4()
	return uid.String()
}
