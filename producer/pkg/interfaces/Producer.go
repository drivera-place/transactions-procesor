package interfaces

import (
	"domain/pkg/domain"
)

type Producer interface {
	Publish(txn *domain.Transaction) error
}
