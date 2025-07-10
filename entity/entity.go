package entity

import (
	"github.com/bsloan/game-sandbox/asset"
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
	Animations  map[EntityState]*Animation
	StaticImage *ebiten.Image
	XPos        float64
	YPos        float64
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

func InitializePlayer(x, y float64) *Entity {
	idle := Animation{
		Frames: []*ebiten.Image{
			asset.PlayerIdle1,
			asset.PlayerIdle2,
		},
		AnimationSpeed: 0.02,
	}
	moveRight := Animation{
		Frames: []*ebiten.Image{
			asset.PlayerRun1,
			asset.PlayerRun2,
			asset.PlayerRun3,
			asset.PlayerRun4,
			asset.PlayerRun5,
			asset.PlayerRun6,
		},
	}
	player := &Entity{
		Type:  Player,
		State: Idle,
		XPos:  x,
		YPos:  y,
		Animations: map[EntityState]*Animation{
			Idle:        &idle,
			MovingRight: &moveRight,
		},
	}
	return player
}
