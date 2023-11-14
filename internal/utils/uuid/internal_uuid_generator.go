package uuid

import "github.com/google/uuid"

type GoogleUUIDProvider struct {
}

func (i *GoogleUUIDProvider) GenerateUUID() string {
	return uuid.New().String()
}
