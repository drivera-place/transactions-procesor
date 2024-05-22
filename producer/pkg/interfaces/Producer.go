package interfaces

import (
	"transactions/pkg/imp"
)
type Producer interface {
	Push(r *imp.Row) error
}
