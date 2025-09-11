package entity

import (
	"math"
	"slices"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/jakecoffman/cp"
)

type EntityType int

const (
	Player EntityType = iota
	PlayerWeapon
	SwordDog
	Frog
	Alligator
	Eagle
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
	ActiveRight3
	ActiveLeft3
	Dying
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
	Type             EntityType
	State            EntityState
	RememberState    EntityState
	Facing           Direction
	Grounded         bool
	OnSlope          bool
	WeaponAvailable  bool
	Running          bool
	Damaged          int // number of ticks to show damaged color scale when entity takes damage
	Boost            float64
	Animations       map[EntityState]*Animation
	StaticImage      *ebiten.Image
	Body             *cp.Body
	Shape            *cp.Shape
	MaxHealth        int
	Health           int
	AttackDamage     int
	TickCounter      int // general-use tick counter for timing entity actions or events
	OriginX, OriginY float64
	// other metadata: points, etc
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

type Behavior func(e *Entity)

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
		if entity != nil && entity.Type == entityType {
			return r.Entities[i]
		}
	}
	return nil
}

func (r *Registry) DrawableEntities() []*Entity {
	var result []*Entity
	for i, entity := range r.Entities {
		if entity != nil && entity.State != Dead && entity.Image() != nil {
			result = append(result, r.Entities[i])
		}
	}
	return result
}

func (r *Registry) RemoveDead(space *cp.Space) {
	var entitiesToRemove []int
	for i, entity := range r.Entities {
		if entity != nil && entity.State == Dead {
			entity.Body.EachShape(func(shape *cp.Shape) {
				entity.Body.RemoveShape(shape)
				space.RemoveShape(shape)
			})
			entitiesToRemove = append(entitiesToRemove, i)
		}
	}
	for _, i := range entitiesToRemove {
		slices.Delete(r.Entities, i, i+1)
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
