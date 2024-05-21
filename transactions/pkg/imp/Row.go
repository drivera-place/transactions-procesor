package imp

import (
	"time"
)

type Row struct {
	Id          int
	Date        time.Time
	Transaction float64
}
