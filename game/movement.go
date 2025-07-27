package game

import (
	"github.com/bsloan/game-sandbox/entity"
	"github.com/bsloan/game-sandbox/settings"
	"github.com/ebitenui/ebitenui/input"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/jakecoffman/cp"
)

func inputAny() bool {
	return input.AnyKeyPressed()
}

func inputJump() bool {
	return ebiten.IsKeyPressed(ebiten.KeySpace)
}

func inputLeft() bool {
	return ebiten.IsKeyPressed(ebiten.KeyLeft) || ebiten.IsKeyPressed(ebiten.KeyA)
}

func inputRight() bool {
	return ebiten.IsKeyPressed(ebiten.KeyRight) || ebiten.IsKeyPressed(ebiten.KeyD)
}

func inputAttack() bool {
	return ebiten.IsKeyPressed(ebiten.KeyK) || ebiten.IsKeyPressed(ebiten.KeyAlt)
}

func (g *Game) MovePlayer() {
	var p = g.registry.Player()
	var pWeapon = g.registry.Query(entity.PlayerWeapon)

	if !inputAny() && p.Grounded {
		if p.Facing == entity.Right {
			p.State = entity.IdleRight
		} else if p.Facing == entity.Left {
			p.State = entity.IdleLeft
		}
	}

	if !inputJump() {
		if p.Grounded {
			p.Boost = settings.PlayerJumpBoostHeight
		} else {
			p.Boost = 0
		}
	}

	if inputRight() {
		p.Facing = entity.Right
		if p.Grounded {
			p.State = entity.MovingRight
		} else if p.State == entity.JumpingLeft {
			p.State = entity.JumpingRight
		}
		vx, vy := p.Body.Velocity().X, 0.0
		vx += settings.PlayerAccelerationStepX
		p.Body.ApplyForceAtWorldPoint(cp.Vector{vx, vy}, p.Body.Position())
	}

	if inputLeft() {
		p.Facing = entity.Left
		if p.Grounded {
			p.State = entity.MovingLeft
		} else if p.State == entity.JumpingRight {
			p.State = entity.JumpingLeft
		}
		vx, vy := p.Body.Velocity().X, 0.0
		vx -= settings.PlayerAccelerationStepX
		p.Body.ApplyForceAtWorldPoint(cp.Vector{vx, vy}, p.Body.Position())
	}

	if ebiten.IsKeyPressed(ebiten.KeyDown) {
		// TODO: crouch
	}

	if inputJump() && p.Boost > 0 {
		if p.State == entity.JumpingRight || p.State == entity.JumpingLeft {
			// player is already in a jump, diminish boost
			p.Boost--
			p.Body.ApplyForceAtWorldPoint(cp.Vector{0, -settings.PlayerJumpContinueVelocity}, p.Body.Position())
		} else {
			// player is in some other state, so must be initiating the jump
			if p.Facing == entity.Left {
				p.State = entity.JumpingLeft
			} else {
				p.State = entity.JumpingRight
			}
			p.Body.ApplyForceAtWorldPoint(cp.Vector{0, -settings.PlayerJumpInitialVelocity}, p.Body.Position())
			p.Grounded = false
			p.Shape.SetFriction(0)
		}
		if p.Boost <= 0 {
			p.Boost = 0
			if p.Facing == entity.Left {
				p.State = entity.FallingLeft
			} else {
				p.State = entity.FallingRight
			}
		}
	}

	if inputAttack() && pWeapon.State != entity.ActiveRight && pWeapon.State != entity.ActiveLeft {
		// create a special Shape for the slash, and show/animate it
		if p.Facing == entity.Right {
			pWeapon.State = entity.ActiveRight
			pWeapon.Body.SetPosition(cp.Vector{p.Body.Position().X + 20, p.Body.Position().Y + 10})
		} else {
			pWeapon.State = entity.ActiveLeft
			pWeapon.Body.SetPosition(cp.Vector{p.Body.Position().X - 5, p.Body.Position().Y + 10})
		}
		weaponShape := g.space.AddShape(cp.NewBox(pWeapon.Body, 64, 47, 10))
		weaponShape.SetSensor(true)
		pWeapon.Shape = weaponShape

		// TODO: collision detection for the slash Shape
	}

	// make sure player's weapon position tracks player's body position each frame
	if p.Facing == entity.Right {
		pWeapon.Body.SetPosition(cp.Vector{p.Body.Position().X + 20, p.Body.Position().Y + 10})
	} else {
		pWeapon.Body.SetPosition(cp.Vector{p.Body.Position().X - 5, p.Body.Position().Y + 10})
	}

	// determine if player is falling, change friction and sprite animation accordingly
	if p.Body.Velocity().Y > 70 && !p.OnSlope {
		// player has steady downward velocity and is falling
		if p.Facing == entity.Right {
			p.State = entity.FallingRight
		}
		if p.Facing == entity.Left {
			p.State = entity.FallingLeft
		}
		p.Grounded = false
		p.Shape.SetFriction(0)
	} else if p.Body.Velocity().Y >= 0.01 && !p.OnSlope {
		// player has a little bit of downward velocity but may not be falling
		// this helps prevent player from becoming "grounded" a.k.a. stuck on
		// vertically stacked tiles
		p.Grounded = false
		p.Shape.SetFriction(0)
	}
	if p.Grounded && !p.OnSlope {
		// reset the player to normal friction and y velocity for being on solid ground
		p.Shape.SetFriction(0.75)
		p.Body.SetVelocity(p.Body.Velocity().X, 0)
	}
	if p.Grounded && p.OnSlope && p.Body.Velocity().Y > 2 {
		// if the player is sliding down a slope, add extra friction to control it
		p.Shape.SetFriction(4.0)
	}

	// enforce maximum velocity in each direction
	if p.Body.Velocity().X > settings.PlayerMaxVelocityX {
		p.Body.SetVelocity(settings.PlayerMaxVelocityX, p.Body.Velocity().Y)
	}
	if p.Body.Velocity().X < -settings.PlayerMaxVelocityX {
		p.Body.SetVelocity(-settings.PlayerMaxVelocityX, p.Body.Velocity().Y)
	}
	if p.Body.Velocity().Y > settings.PlayerMaxVelocityY {
		p.Body.SetVelocity(p.Body.Velocity().X, settings.PlayerMaxVelocityY)
	}
	if p.Body.Velocity().Y < -settings.PlayerJumpVelocityLimit {
		p.Body.SetVelocity(p.Body.Velocity().X, -settings.PlayerJumpVelocityLimit)
	}
}
