package algo

import (
	"golang.org/x/exp/constraints"
)

type numbers interface {
	constraints.Float | constraints.Integer
}
