package interfaces

type Builder interface {
	Create(lines int) (string, error)
}
