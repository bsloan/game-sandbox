package entity

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/jakecoffman/cp"
	"math"
)

type EntityType int

const (
	Player EntityType = iota
	PlayerWeapon
	SwordDog
	Frog
)

type EntityState int

const (
	Default EntityState = iota
	Idle
	IdleLeft
	IdleRight
	MovingLeft
	MovingRight
	MovingUp
	MovingDown
	JumpingRight
	JumpingLeft
	FallingRight
	FallingLeft
	ActiveRight
	ActiveLeft
	ActiveRight2
	ActiveLeft2
	Dead
)

type Direction int

const (
	Left Direction = iota
	Right
	Up
	Down
)

// https://co0p.github.io/posts/ecs-animation/ provides a good starting point

type Entity struct {
	Type            EntityType
	State           EntityState
	RememberState   EntityState
	Facing          Direction
	Grounded        bool
	OnSlope         bool
	WeaponAvailable bool
	Running         bool
	Boost           float64
	Animations      map[EntityState]*Animation
	StaticImage     *ebiten.Image
	Body            *cp.Body
	Shape           *cp.Shape
	// other metadata: health, attack damage, points, etc
}

func (e *Entity) Position() (float64, float64) {
	return e.Body.Position().X, e.Body.Position().Y
}

// Image returns the active image in the entity's animation sequence for the
// entity's current state. If the entity is not animated (e.g., Animations is nil)
// then try to return the entity's static image. If no animations or static image
// are defined, return nil.
func (e *Entity) Image() *ebiten.Image {
	if e.Animations != nil && e.Animations[e.State] != nil {
		return e.Animations[e.State].Frames[e.Animations[e.State].CurrentFrameIndex]
	} else if e.StaticImage != nil {
		return e.StaticImage
	}
	return nil
}

type Registry struct {
	Entities []*Entity
}

// AddEntity adds a new Entity to the Registry of entities.
// Current implementation is very simple, revisit later to update
// this API appropriately according to performance and/or querying needs.
func (r *Registry) AddEntity(entity *Entity) {
	r.Entities = append(r.Entities, entity)
}

func (r *Registry) Player() *Entity {
	return r.Query(Player)
}

func (r *Registry) Query(entityType EntityType) *Entity {
	for i, entity := range r.Entities {
		if entity.Type == entityType {
			return r.Entities[i]
		}
	}
	return nil
}

func (r *Registry) DrawableEntities() []*Entity {
	var result []*Entity
	for i, entity := range r.Entities {
		if entity.State != Dead && entity.Image() != nil {
			result = append(result, r.Entities[i])
		}
	}
	return result
}

func (r *Registry) RemoveDead(space *cp.Space) {
	for _, entity := range r.Entities {
		if entity.State == Dead {
			entity.Body.EachShape(func(shape *cp.Shape) {
				entity.Body.RemoveShape(shape)
				space.RemoveShape(shape)
			})
		}
	}
}

type Animation struct {
	Frames                []*ebiten.Image
	CurrentFrameIndex     int
	Count                 float64
	AnimationSpeed        float64
	EntityStateTransition EntityState
}

func (a *Animation) Animate() EntityState {
	// advance animation
	a.Count += a.AnimationSpeed
	a.CurrentFrameIndex = int(math.Floor(a.Count))

	// animation cycle is complete. if this Animation has any state transition
	// other than Default, return it to the caller so the owning Entity's state
	// can be updated. Otherwise, repeat the animation indefinitely
	if a.CurrentFrameIndex >= len(a.Frames) {
		a.Count = 0
		a.CurrentFrameIndex = 0
		if a.EntityStateTransition != Default {
			return a.EntityStateTransition
		}
	}
	return Default
}
