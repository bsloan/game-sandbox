package entities

import (
	"github.com/hajimehoshi/ebiten/v2"
	"math"
)

type EntityType int

const (
	Player EntityType = iota
	Bird
	Frog
)

type EntityState int

const (
	Idle EntityState = iota
	MovingLeft
	MovingRight
	MovingUp
	MovingDown
	JumpingUp
	Falling
)

// https://co0p.github.io/posts/ecs-animation/ provides a good starting point

type Entity struct {
	Type        EntityType
	State       EntityState
	ActiveImage *ebiten.Image
	Animations  map[EntityState]*Animation
	XPos        int
	YPos        int
	// other metadata: health, attack damage, points, etc
}

type Registry struct {
	Entities []Entity
}

// AddEntity adds a new Entity to the Registry of entities.
// Current implementation is very simple, revisit later to update
// this API appropriately according to performance and/or querying needs.
func (r *Registry) AddEntity(entity Entity) {
	r.Entities = append(r.Entities, entity)
}

type Animation struct {
	Frames            []*ebiten.Image
	CurrentFrameIndex int
	Count             float64
	AnimationSpeed    float64
}

func (a *Animation) Animate() (*ebiten.Image, error) {
	// advance animation
	a.Count += a.AnimationSpeed
	a.CurrentFrameIndex = int(math.Floor(a.Count))

	// restart the animation
	// TODO: not all animations cycle indefinitely, some may end and then
	//  trigger a side effect (e.g., change state of its owning Entity).
	if a.CurrentFrameIndex >= len(a.Frames) {
		a.Count = 0
		a.CurrentFrameIndex = 0
	}

	// return reference to the next image in animation sequence
	return a.Frames[a.CurrentFrameIndex], nil
}
