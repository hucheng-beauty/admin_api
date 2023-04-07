package uuid

import "github.com/gofrs/uuid"

func New() string {
	uid, _ := uuid.NewV4()
	return uid.String()
}
