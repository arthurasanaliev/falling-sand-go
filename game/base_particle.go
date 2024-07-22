package game

import (
	"image/color"
)

type BaseParticle struct {
	color          color.Color
	alreadyUpdated bool
}

func (bp *BaseParticle) IsAlreadyUpdated() bool {
	return bp.alreadyUpdated
}

func (bp *BaseParticle) SetAlreadyUpdated(updated bool) {
	bp.alreadyUpdated = updated
}
