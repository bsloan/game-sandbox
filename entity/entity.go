package entity

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
	Facing          Direction
	Animations      map[EntityState]*Animation
	StaticImage     *ebiten.Image
	XPos            float64
	YPos            float64
	BoundingOffsetX int
	BoundingOffsetY int
	BoundingWidth   int
	BoundingHeight  int
	Speed           float64
	// other metadata: health, attack damage, points, etc
}

// Image returns the active image in the entity's animation sequence for the
// entity's current state. If the entity is not animated (e.g., Animations is nil)
// then try to return the entity's static image. If no animations or static image
// are defined, return nil.
func (e *Entity) Image() *ebiten.Image {
	if e.Animations != nil {
		return e.Animations[e.State].Frames[e.Animations[e.State].CurrentFrameIndex]
	} else if e.StaticImage != nil {
		return e.StaticImage
	}
	return nil
}

func (e *Entity) TileCollisions(gameboard [][]int) []int {
	// get the 4 corners of the entity's bounding box
	x1 := int(e.XPos) + e.BoundingOffsetX
	y1 := int(e.YPos) + e.BoundingOffsetY
	x2 := x1 + e.BoundingWidth
	y2 := y1 + e.BoundingHeight

	// get the tile coordinates of the top left and lower right tiles that
	// are overlapped by the entity's bounding box
	tileX1 := x1 / 16
	tileY1 := y1 / 16
	tileX2 := x2 / 16
	tileY2 := y2 / 16

	// store the tile types that collide with entity in this slice
	collisions := make([]int, 0)

	if e.State == MovingRight {
		for ty := tileY1; ty <= tileY2; ty++ {
			tile := gameboard[ty][tileX2]
			collisions = append(collisions, tile)
		}
	} else if e.State == MovingLeft {
		for ty := tileY1; ty <= tileY2; ty++ {
			tile := gameboard[ty][tileX1]
			collisions = append(collisions, tile)
		}
	} else if e.State == FallingRight || e.State == FallingLeft {
		for tx := tileX1; tx <= tileX2; tx++ {
			tile := gameboard[tileY2][tx]
			collisions = append(collisions, tile)
		}
	} else if e.State == JumpingLeft || e.State == JumpingRight {
		for tx := tileX1; tx <= tileX2; tx++ {
			tile := gameboard[tileY1][tx]
			collisions = append(collisions, tile)
		}
	}

	// return collisions and let the caller handle them
	return collisions
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

func (r *Registry) Player() *Entity {
	for i, entity := range r.Entities {
		if entity.Type == Player {
			return &r.Entities[i]
		}
	}
	return nil
}

func (r *Registry) DrawableEntities() []*Entity {
	var result []*Entity
	for i, entity := range r.Entities {
		if entity.State != Dead && entity.Image() != nil {
			result = append(result, &r.Entities[i])
		}
	}
	return result
}

type Animation struct {
	Frames            []*ebiten.Image
	CurrentFrameIndex int
	Count             float64
	AnimationSpeed    float64
}

func (a *Animation) Animate() {
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
}
