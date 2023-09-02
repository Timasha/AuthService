package dependencies

type UseCasesConfig interface {
	GetMinLoginLen() int
	GetMinPasswordLen() int
}
