package entity

import (
	"github.com/bsloan/game-sandbox/asset"
	"github.com/hajimehoshi/ebiten/v2"
)

func InitializePlayer(x, y float64) *Entity {
	idleRight := Animation{
		Frames: []*ebiten.Image{
			asset.PlayerIdleRight1,
			asset.PlayerIdleRight2,
			asset.PlayerIdleRight3,
			asset.PlayerIdleRight4,
		},
		AnimationSpeed: 0.05,
	}
	idleLeft := Animation{
		Frames: []*ebiten.Image{
			asset.PlayerIdleLeft1,
			asset.PlayerIdleLeft2,
			asset.PlayerIdleLeft3,
			asset.PlayerIdleLeft4,
		},
		AnimationSpeed: 0.05,
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
		AnimationSpeed: 0.1,
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
		AnimationSpeed: 0.1,
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
	player := &Entity{
		Type:   Player,
		State:  Idle,
		XPos:   x,
		YPos:   y,
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
		},
	}
	return player
}
