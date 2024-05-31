package interfaces

type Builder interface {
	CreateFile(lines int) (string, error)
}
