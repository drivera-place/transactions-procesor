package interfaces

import (
	"domain/pkg/domain"
)

type TxnsDB interface {
	Save(txn *domain.Transaction) (string, error)
}