package assemblline

import (
	"chain/definition"
)

type AssemblingLineCR interface {
	Execute(*definition.CarCR)
	SetNext(AssemblingLineCR)
}
