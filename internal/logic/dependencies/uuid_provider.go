package dependencies

type UUIDProvider interface {
	GenerateUUID() string
}
