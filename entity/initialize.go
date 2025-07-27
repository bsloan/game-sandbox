package entity

import (
	"github.com/bsloan/game-sandbox/asset"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/jakecoffman/cp"
)

func InitializePlayer(space *cp.Space, x, y float64) *Entity {
	idleRight := Animation{
		Frames: []*ebiten.Image{
			asset.PlayerIdleRight1,
			asset.PlayerIdleRight2,
			asset.PlayerIdleRight3,
			asset.PlayerIdleRight4,
		},
		AnimationSpeed: 0.1,
	}
	idleLeft := Animation{
		Frames: []*ebiten.Image{
			asset.PlayerIdleLeft1,
			asset.PlayerIdleLeft2,
			asset.PlayerIdleLeft3,
			asset.PlayerIdleLeft4,
		},
		AnimationSpeed: 0.1,
	}
	moveRight := Animation{
		Frames: []*ebiten.Image{
			asset.PlayerMoveRight1,
			asset.PlayerMoveRight2,
			asset.PlayerMoveRight3,
			asset.PlayerMoveRight4,
			asset.PlayerMoveRight5,
			asset.PlayerMoveRight6,
		},
		AnimationSpeed: 0.2,
	}
	moveLeft := Animation{
		Frames: []*ebiten.Image{
			asset.PlayerMoveLeft1,
			asset.PlayerMoveLeft2,
			asset.PlayerMoveLeft3,
			asset.PlayerMoveLeft4,
			asset.PlayerMoveLeft5,
			asset.PlayerMoveLeft6,
		},
		AnimationSpeed: 0.2,
	}
	jumpRight := Animation{
		Frames: []*ebiten.Image{
			asset.PlayerJumpRight1,
		},
	}
	jumpLeft := Animation{
		Frames: []*ebiten.Image{
			asset.PlayerJumpLeft1,
		},
	}
	fallRight := Animation{
		Frames: []*ebiten.Image{
			asset.PlayerFallRight1,
		},
	}
	fallLeft := Animation{
		Frames: []*ebiten.Image{
			asset.PlayerFallLeft1,
		},
	}
	activeRight := Animation{
		Frames: []*ebiten.Image{
			asset.PlayerMoveRight4,
			asset.PlayerMoveRight5,
			asset.PlayerMoveRight6,
			asset.PlayerMoveRight6,
			asset.PlayerMoveRight2,
			asset.PlayerMoveRight2,
		},
		AnimationSpeed:        0.4,
		EntityStateTransition: Idle,
	}
	activeLeft := Animation{
		Frames: []*ebiten.Image{
			asset.PlayerMoveLeft4,
			asset.PlayerMoveLeft5,
			asset.PlayerMoveLeft6,
			asset.PlayerMoveLeft6,
			asset.PlayerMoveLeft2,
			asset.PlayerMoveLeft2,
		},
		AnimationSpeed:        0.4,
		EntityStateTransition: Idle,
	}
	player := Entity{
		Type:   Player,
		State:  Idle,
		Facing: Right,
		Animations: map[EntityState]*Animation{
			Idle:         &idleRight,
			IdleRight:    &idleRight,
			IdleLeft:     &idleLeft,
			MovingRight:  &moveRight,
			MovingLeft:   &moveLeft,
			JumpingRight: &jumpRight,
			JumpingLeft:  &jumpLeft,
			FallingRight: &fallRight,
			FallingLeft:  &fallLeft,
			ActiveRight:  &activeRight,
			ActiveLeft:   &activeLeft,
		},
		Body:  cp.NewBody(1, cp.INFINITY),
		Boost: 0,
	}

	player.Body.UserData = &player
	space.AddBody(player.Body)
	player.Body.SetPosition(cp.Vector{X: x, Y: y})
	//playerShape := space.AddShape(cp.NewBox(player.Body, 8, 7, 8))
	playerShape := space.AddShape(cp.NewCircle(player.Body, 10, cp.Vector{X: 0, Y: 0}))

	playerShape.SetElasticity(0)
	playerShape.SetFriction(0.75)
	playerShape.SetCollisionType(PlayerCollisionType)
	player.Shape = playerShape
	GenericGroundedHandler(space, PlayerCollisionType)
	SlopeHandler(space, PlayerCollisionType)
	return &player
}

func InitializePlayerSword(space *cp.Space, x, y float64) *Entity {
	slashRight := Animation{
		Frames: []*ebiten.Image{
			asset.WhiteSlashRight1,
			asset.WhiteSlashRight2,
			asset.WhiteSlashRight3,
			asset.WhiteSlashRight4,
			asset.WhiteSlashRight5,
			asset.WhiteSlashRight6,
		},
		AnimationSpeed:        0.4,
		EntityStateTransition: IdleRight,
	}
	slashLeft := Animation{
		Frames: []*ebiten.Image{
			asset.WhiteSlashLeft1,
			asset.WhiteSlashLeft2,
			asset.WhiteSlashLeft3,
			asset.WhiteSlashLeft4,
			asset.WhiteSlashLeft5,
			asset.WhiteSlashLeft6,
		},
		AnimationSpeed:        0.4,
		EntityStateTransition: IdleLeft,
	}
	sword := Entity{
		Type:   PlayerWeapon,
		State:  Idle,
		Facing: Right,
		Animations: map[EntityState]*Animation{
			Idle:        nil,
			ActiveRight: &slashRight,
			ActiveLeft:  &slashLeft,
		},
		Body: cp.NewBody(0.001, cp.INFINITY),
	} // mass is set to near-zero to ignore gravity. setting mass to true 0 causes weird physics bugs

	sword.Body.UserData = &sword
	space.AddBody(sword.Body)
	sword.Body.SetPosition(cp.Vector{X: x, Y: y})
	return &sword
}

func InitializeSwordDog(space *cp.Space, x, y float64) *Entity {
	idleRight := Animation{
		Frames: []*ebiten.Image{
			asset.SwordDogIdleRight1,
			asset.SwordDogIdleRight2,
			asset.SwordDogIdleRight3,
			asset.SwordDogIdleRight4,
			asset.SwordDogIdleRight5,
			asset.SwordDogIdleRight6,
		},
		AnimationSpeed: 0.1,
	}
	idleLeft := Animation{
		Frames: []*ebiten.Image{
			asset.SwordDogIdleLeft1,
			asset.SwordDogIdleLeft2,
			asset.SwordDogIdleLeft3,
			asset.SwordDogIdleLeft4,
			asset.SwordDogIdleLeft5,
			asset.SwordDogIdleLeft6,
		},
		AnimationSpeed: 0.1,
	}
	swordDog := Entity{
		Type:   SwordDog,
		State:  IdleLeft,
		Facing: Left,
		Animations: map[EntityState]*Animation{
			IdleRight: &idleRight,
			IdleLeft:  &idleLeft,
		},
		Body: cp.NewBody(1, cp.INFINITY),
	}
	swordDog.Body.UserData = &swordDog
	space.AddBody(swordDog.Body)
	swordDog.Body.SetPosition(cp.Vector{X: x, Y: y})
	swordDogShape := space.AddShape(cp.NewCircle(swordDog.Body, 11, cp.Vector{X: 0, Y: 0}))
	swordDogShape.SetElasticity(0)
	swordDogShape.SetFriction(0.75)
	swordDogShape.SetCollisionType(SwordDogCollisionType)
	swordDog.Shape = swordDogShape
	GenericGroundedHandler(space, SwordDogCollisionType)
	SlopeHandler(space, SwordDogCollisionType)
	return &swordDog
}
