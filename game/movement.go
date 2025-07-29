package game

import (
	"github.com/bsloan/game-sandbox/entity"
	"github.com/bsloan/game-sandbox/settings"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/jakecoffman/cp"
	"math"
)

const (
	NonstandardGamepadButtonLeft  = 18
	NonstandardGamepadButtonRight = 16
	NonstandardGamepadButtonX     = 3
	NonstandardGamepadButtonA     = 0
)

func (g *Game) gamepadAvailable() bool {
	return len(g.gamepadIds) > 0
}

func (g *Game) inputAny() bool {
	return g.inputJump() || g.inputLeft() || g.inputRight() || g.inputAttack()
}

func (g *Game) inputJump() bool {
	return ebiten.IsKeyPressed(ebiten.KeySpace) || (g.gamepadAvailable() && ebiten.IsGamepadButtonPressed(g.gamepadIds[0], NonstandardGamepadButtonA))
}

func (g *Game) inputLeft() bool {
	return ebiten.IsKeyPressed(ebiten.KeyLeft) || ebiten.IsKeyPressed(ebiten.KeyA) || (g.gamepadAvailable() && ebiten.IsGamepadButtonPressed(g.gamepadIds[0], NonstandardGamepadButtonLeft))
}

func (g *Game) inputRight() bool {
	return ebiten.IsKeyPressed(ebiten.KeyRight) || ebiten.IsKeyPressed(ebiten.KeyD) || (g.gamepadAvailable() && ebiten.IsGamepadButtonPressed(g.gamepadIds[0], NonstandardGamepadButtonRight))
}

func (g *Game) inputAttack() bool {
	return ebiten.IsKeyPressed(ebiten.KeyK) || ebiten.IsKeyPressed(ebiten.KeyAlt) || (g.gamepadAvailable() && ebiten.IsGamepadButtonPressed(g.gamepadIds[0], NonstandardGamepadButtonX))
}

func (g *Game) MovePlayer() {
	var p = g.registry.Player()
	var pWeapon = g.registry.Query(entity.PlayerWeapon)

	g.gamepadIds = ebiten.AppendGamepadIDs(g.gamepadIds[:0])

	if !g.inputRight() && !g.inputLeft() && p.State != entity.ActiveRight && p.State != entity.ActiveLeft && p.Grounded {
		if p.Facing == entity.Right {
			p.State = entity.IdleRight
		} else if p.Facing == entity.Left {
			p.State = entity.IdleLeft
		}
	}

	if !g.inputJump() {
		if p.Grounded {
			p.Boost = settings.PlayerJumpBoostHeight
		} else {
			p.Boost = 0
		}
	}

	if g.inputRight() {
		p.Facing = entity.Right
		if p.Grounded {
			p.State = entity.MovingRight
		} else if p.State == entity.JumpingLeft {
			p.State = entity.JumpingRight
		}
		vx, vy := p.Body.Velocity().X, 0.0
		if p.Running {
			vx += settings.PlayerRunningAccelerationStep
		} else {
			vx += settings.PlayerAccelerationStep
		}
		p.Body.ApplyForceAtWorldPoint(cp.Vector{X: vx, Y: vy}, p.Body.Position())
	}

	if g.inputLeft() {
		p.Facing = entity.Left
		if p.Grounded {
			p.State = entity.MovingLeft
		} else if p.State == entity.JumpingRight {
			p.State = entity.JumpingLeft
		}
		vx, vy := p.Body.Velocity().X, 0.0
		if p.Running {
			vx -= settings.PlayerRunningAccelerationStep
		} else {
			vx -= settings.PlayerAccelerationStep
		}
		p.Body.ApplyForceAtWorldPoint(cp.Vector{X: vx, Y: vy}, p.Body.Position())
	}

	if ebiten.IsKeyPressed(ebiten.KeyDown) {
		// TODO: crouch
	}

	if g.inputJump() && p.Boost > 0 {
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
	}

	if g.inputAttack() && pWeapon.State != entity.ActiveRight && pWeapon.State != entity.ActiveLeft && p.WeaponAvailable {
		// create a special Shape for the slash, and show/animate it
		if p.Facing == entity.Right {
			pWeapon.State = entity.ActiveRight
			pWeapon.Body.SetPosition(cp.Vector{p.Body.Position().X + 20, p.Body.Position().Y + 10})
			p.State = entity.ActiveRight
		} else {
			pWeapon.State = entity.ActiveLeft
			pWeapon.Body.SetPosition(cp.Vector{p.Body.Position().X - 5, p.Body.Position().Y + 10})
			p.State = entity.ActiveLeft
		}
		weaponShape := g.space.AddShape(cp.NewBox(pWeapon.Body, 64, 47, 10))
		weaponShape.SetSensor(true)
		pWeapon.Shape = weaponShape
		p.WeaponAvailable = false // disable next attack until after the button is released
		// TODO: collision detection for the slash Shape
	} else if g.inputAttack() {
		// attack and horizontal boost are the same button
		if p.Grounded {
			p.Running = true
		}
	} else if !g.inputAttack() {
		p.WeaponAvailable = true
		p.Running = false
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
	if p.Running && p.Body.Velocity().X > settings.PlayerMaxRunningVelocityX {
		p.Body.SetVelocity(settings.PlayerMaxRunningVelocityX, p.Body.Velocity().Y)
	} else if !p.Running && p.Body.Velocity().X > settings.PlayerMaxVelocityX {
		p.Body.SetVelocity(settings.PlayerMaxVelocityX, p.Body.Velocity().Y)
	}
	if p.Running && p.Body.Velocity().X < -settings.PlayerMaxRunningVelocityX {
		p.Body.SetVelocity(-settings.PlayerMaxRunningVelocityX, p.Body.Velocity().Y)
	} else if !p.Running && p.Body.Velocity().X < -settings.PlayerMaxVelocityX {
		p.Body.SetVelocity(-settings.PlayerMaxVelocityX, p.Body.Velocity().Y)
	}
	if p.Body.Velocity().Y > settings.PlayerMaxVelocityY {
		p.Body.SetVelocity(p.Body.Velocity().X, settings.PlayerMaxVelocityY)
	}
	if p.Body.Velocity().Y < -settings.PlayerJumpVelocityLimit {
		p.Body.SetVelocity(p.Body.Velocity().X, -settings.PlayerJumpVelocityLimit)
	}
}

func (g *Game) MoveSwordDog(swordDog *entity.Entity) {
	// find proximity to player
	var swordDogX = swordDog.Body.Position().X
	var swordDogY = swordDog.Body.Position().Y
	var playerX = g.registry.Player().Body.Position().X
	var playerY = g.registry.Player().Body.Position().Y
	xDistance := math.Abs(swordDogX - playerX)
	yDistance := math.Abs(swordDogY - playerY)

	notAttacking := swordDog.State != entity.ActiveRight && swordDog.State != entity.ActiveLeft && swordDog.State != entity.ActiveRight2 && swordDog.State != entity.ActiveLeft2

	// chase the player if we're close
	if xDistance < 32 && notAttacking {
		if playerX > swordDogX {
			swordDog.Facing = entity.Right
			swordDog.State = entity.MovingRight
		} else {
			swordDog.Facing = entity.Left
			swordDog.State = entity.MovingLeft
		}
	}

	// attack the player if we're really close and not already in an attacking state
	if ((playerX < swordDogX && xDistance < 22) || (playerX > swordDogX && xDistance < 31)) && yDistance < 20 && notAttacking {
		if swordDog.Facing == entity.Right {
			if swordDog.RememberState == entity.ActiveRight || swordDog.RememberState == entity.ActiveLeft {
				swordDog.State = entity.ActiveRight2
				swordDog.RememberState = entity.ActiveRight2
			} else {
				swordDog.State = entity.ActiveRight
				swordDog.RememberState = entity.ActiveRight
			}
		} else {
			if swordDog.RememberState == entity.ActiveRight || swordDog.RememberState == entity.ActiveLeft {
				swordDog.State = entity.ActiveLeft2
				swordDog.RememberState = entity.ActiveLeft2
			} else {
				swordDog.State = entity.ActiveLeft
				swordDog.RememberState = entity.ActiveLeft
			}
		}

		// TODO: attach a larger shape to the dog while it's attacking
		//swordShape := g.space.AddShape(cp.NewBox(swordDog.Body, 45, 19, 3))
		swordShape := g.space.AddShape(cp.NewBox2(swordDog.Body, cp.BB{
			L: -23,
			B: 10,
			R: 38,
			T: -10,
		}, 3))
		swordShape.SetElasticity(0.4)
		swordShape.SetFriction(0.75)
		swordShape.UserData = 1 // hack to remember later that this shape was added for attack
		// TODO: collision type and handler
		//swordShape.SetCollisionType(SwordDogCollisionType)
		//GenericGroundedHandler(space, SwordDogCollisionType)
	} else if notAttacking {
		// remove any shapes that were added just for attack
		swordDog.Body.EachShape(func(shape *cp.Shape) {
			if shape.UserData == 1 {
				swordDog.Body.RemoveShape(shape)
				g.space.RemoveShape(shape)
			}
		})
	}

	// move the dog
	if swordDog.State == entity.MovingLeft {
		swordDog.Facing = entity.Left
		vx, vy := swordDog.Body.Velocity().X, 0.0
		vx -= settings.SwordDogAccelerationStep
		swordDog.Body.ApplyForceAtWorldPoint(cp.Vector{X: vx, Y: vy}, swordDog.Body.Position())
	} else if swordDog.State == entity.MovingRight {
		swordDog.Facing = entity.Right
		vx, vy := swordDog.Body.Velocity().X, 0.0
		vx += settings.SwordDogAccelerationStep
		swordDog.Body.ApplyForceAtWorldPoint(cp.Vector{X: vx, Y: vy}, swordDog.Body.Position())
	}

	// enforce velocity constraints
	if swordDog.Body.Velocity().X < -settings.SwordDogMaxVelocityX {
		swordDog.Body.SetVelocity(-settings.SwordDogMaxVelocityX, swordDog.Body.Velocity().Y)
	}
	if swordDog.Body.Velocity().X > settings.SwordDogMaxVelocityX {
		swordDog.Body.SetVelocity(settings.SwordDogMaxVelocityX, swordDog.Body.Velocity().Y)
	}
}
