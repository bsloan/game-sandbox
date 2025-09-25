package game

import (
	"math"
	"math/rand/v2"
	"slices"

	"github.com/bsloan/game-sandbox/board"
	"github.com/bsloan/game-sandbox/entity"
	"github.com/bsloan/game-sandbox/settings"
	"github.com/jakecoffman/cp"
)

func adjustPlayerWeaponPosition(pWeapon *entity.Entity, px, py float64) {
	if pWeapon == nil {
		return
	}
	switch pWeapon.State {
	case entity.ActiveRight:
		pWeapon.Body.SetPosition(cp.Vector{px + 5, py})
	case entity.ActiveRight2:
		pWeapon.Body.SetPosition(cp.Vector{X: px + 5, Y: py})
	case entity.ActiveRight3:
		pWeapon.Body.SetPosition(cp.Vector{X: px + 5, Y: py - 15})
	case entity.ActiveLeft:
		pWeapon.Body.SetPosition(cp.Vector{px - 15, py})
	case entity.ActiveLeft2:
		pWeapon.Body.SetPosition(cp.Vector{px - 7, py})
	case entity.ActiveLeft3:
		pWeapon.Body.SetPosition(cp.Vector{px - 7, py - 15})
	default:
	}
}

func (g *Game) canClimb(e *entity.Entity) bool {
	tx := int(e.Body.Position().X / settings.TileSize)
	ty := int(e.Body.Position().Y / settings.TileSize)
	return slices.Contains(board.ClimbableTiles, g.board.Map[ty][tx+1]) || slices.Contains(board.ClimbableTiles, g.board.Map[ty][tx])
}

func (g *Game) MovePlayer(p *entity.Entity) {
	// get player's weapon - may be nil if it's not being used
	var pWeapon = g.registry.Query(entity.PlayerWeapon)

	// if we are transitioning from crouching to not-crouching in this update, reset the player's shape to normal
	if p.State != entity.CrouchLeft && p.State != entity.CrouchRight && (p.RememberState == entity.CrouchLeft || p.RememberState == entity.CrouchRight) {
		entity.InitializeNormalPlayerShape(g.space, p)
		p.RememberState = entity.Idle
	}

	// if no input from player, and player is not climbing, be idle
	if !g.inputDown() && !g.inputRight() && !g.inputLeft() && p.State != entity.ActiveRight && p.State != entity.ActiveLeft && p.State != entity.ClimbingIdle && p.State != entity.ClimbingActive && p.Grounded {
		if p.Facing == entity.Right {
			p.State = entity.IdleRight
		} else if p.Facing == entity.Left {
			p.State = entity.IdleLeft
		}
	}

	// no jump input: reset boost if player is on ground, otherwise zero it out if in the air
	if !g.inputJump() {
		if p.Grounded {
			p.Boost = settings.PlayerJumpBoostHeight
		} else {
			p.Boost = 0
		}
	}

	// if player pressed down, there's no other input, and player is not already crouching or attacking, then
	// detach the player's current shape and attach a smaller/shorter one (initiate crouch)
	if g.inputDown() && p.Grounded && !g.inputRight() && !g.inputLeft() && p.State != entity.ActiveRight && p.State != entity.ActiveLeft && p.State != entity.CrouchRight && p.State != entity.CrouchLeft {
		if p.Facing == entity.Right {
			p.State = entity.CrouchRight
		} else {
			p.State = entity.CrouchLeft
		}
		entity.InitializeCrouchPlayerShape(g.space, p)
	}

	// move right
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

	// move left
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

	// climb a ladder
	if g.inputUp() && g.canClimb(p) && p.State != entity.ClimbingActive {
		// initiate climb from some other state
		p.State = entity.ClimbingActive
	} else if !g.inputUp() && g.canClimb(p) && p.State == entity.ClimbingActive {
		// already in a climbing state, on a climbable tile, but up button is released
		p.State = entity.ClimbingIdle
	}

	// jump
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

	// attack and/or run ...
	if g.inputAttack() && pWeapon == nil && p.WeaponAvailable {
		// create a special Entity for the slash and animate it
		pWeapon = entity.InitializePlayerSword(g.space, p.Body.Position().X, p.Body.Position().Y)
		g.registry.AddEntity(pWeapon)

		var weaponShape *cp.Shape = nil

		if p.Facing == entity.Right {
			if g.inputDown() {
				// downslash right
				pWeapon.State = entity.ActiveRight2
				weaponShape = g.space.AddShape(cp.NewBox(pWeapon.Body, 28, 35, 10))
			} else if g.inputUp() {
				// upslash right
				pWeapon.State = entity.ActiveRight3
				weaponShape = g.space.AddShape(cp.NewBox(pWeapon.Body, 28, 35, 10))
			} else {
				// regular slash right
				pWeapon.State = entity.ActiveRight
				weaponShape = g.space.AddShape(cp.NewBox(pWeapon.Body, 35, 28, 10))
			}
			p.State = entity.ActiveRight
		} else {
			if g.inputDown() {
				// downslash left
				pWeapon.State = entity.ActiveLeft2
				weaponShape = g.space.AddShape(cp.NewBox(pWeapon.Body, 28, 35, 10))
			} else if g.inputUp() {
				// upslash left
				pWeapon.State = entity.ActiveLeft3
				weaponShape = g.space.AddShape(cp.NewBox(pWeapon.Body, 28, 35, 10))
			} else {
				// regular slash left
				pWeapon.State = entity.ActiveLeft
				weaponShape = g.space.AddShape(cp.NewBox(pWeapon.Body, 35, 28, 10))
			}
			p.State = entity.ActiveLeft
		}
		weaponShape.SetCollisionType(entity.PlayerSwordCollisionType)
		pWeapon.Shape = weaponShape
		p.WeaponAvailable = false // disable next attack until after the button is released
	} else if g.inputAttack() {
		// attack and horizontal boost are the same button
		if p.Grounded {
			p.Running = true
		}
	} else if !g.inputAttack() {
		p.WeaponAvailable = true
		p.Running = false
	}

	// make sure player's weapon position tracks player's body position each frame, with slight adjustment
	adjustPlayerWeaponPosition(pWeapon, p.Body.Position().X, p.Body.Position().Y)

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
}

func (g *Game) handleDeadEnemy(e *entity.Entity) bool {
	if e.Health <= 0 {
		e.Body.EachShape(func(shape *cp.Shape) {
			e.Body.RemoveShape(shape)
			g.space.RemoveShape(shape)
		})
		e.State = entity.Dying
		e.Body.SetPosition(cp.Vector{X: e.Body.Position().X, Y: e.Body.Position().Y - 1})
		e.Body.SetVelocity(0, 0)
		return true
	}
	return false
}

func (g *Game) MoveSwordDog(swordDog *entity.Entity) {
	// first, check if we're dead
	if g.handleDeadEnemy(swordDog) {
		return
	}

	// find proximity to player
	var swordDogX = swordDog.Body.Position().X
	var swordDogY = swordDog.Body.Position().Y
	var playerX = g.registry.Player().Body.Position().X
	var playerY = g.registry.Player().Body.Position().Y
	xDistance := math.Abs(swordDogX - playerX)
	yDistance := math.Abs(swordDogY - playerY)

	notAttacking := swordDog.State != entity.ActiveRight && swordDog.State != entity.ActiveLeft && swordDog.State != entity.ActiveRight2 && swordDog.State != entity.ActiveLeft2

	// attack damage increases if sword dog is swinging his sword
	if notAttacking {
		swordDog.AttackDamage = 2
	} else {
		swordDog.AttackDamage = 6
	}

	// chase the player if we're close
	if xDistance < 128 && notAttacking {
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

		swordShape := g.space.AddShape(cp.NewBox2(swordDog.Body, cp.BB{
			L: -23,
			B: 10,
			R: 38,
			T: -10,
		}, 3))
		swordShape.SetElasticity(0.4)
		swordShape.SetFriction(0.75)
		swordShape.UserData = 1 // hack to remember later that this shape was added for attack
		swordShape.SetCollisionType(entity.GenericEnemyCollisionType)
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
}

func (g *Game) MoveAlligator(alligator *entity.Entity) {
	// first, check if we're dead
	if g.handleDeadEnemy(alligator) {
		return
	}

	// find proximity to player
	var alligatorX = alligator.Body.Position().X
	var alligatorY = alligator.Body.Position().Y
	var playerX = g.registry.Player().Body.Position().X
	var playerY = g.registry.Player().Body.Position().Y
	xDistance := math.Abs(alligatorX - playerX)
	yDistance := math.Abs(alligatorY - playerY)

	notAttacking := alligator.State != entity.ActiveRight && alligator.State != entity.ActiveLeft

	// attack damage increases when alligator is actively attacking
	if notAttacking {
		alligator.AttackDamage = 2
	} else {
		alligator.AttackDamage = 6
	}

	// alligator paces back and forth, doesn't chase player
	if xDistance < 256 && notAttacking && alligator.State != entity.MovingLeft && alligator.State != entity.MovingRight {
		if playerX > alligatorX {
			alligator.Facing = entity.Right
			alligator.State = entity.MovingRight
		} else {
			alligator.Facing = entity.Left
			alligator.State = entity.MovingLeft
		}
	}

	if ((playerX < alligatorX && xDistance < 22) || (playerX > alligatorX && xDistance < 40)) && yDistance < 20 && notAttacking {
		if alligator.Facing == entity.Right {
			alligator.State = entity.ActiveRight
		} else {
			alligator.State = entity.ActiveLeft
		}
		// TODO: adjust box shape for Alligator?
		swordShape := g.space.AddShape(cp.NewBox2(alligator.Body, cp.BB{
			L: -23,
			B: 10,
			R: 38,
			T: -10,
		}, 3))
		swordShape.SetElasticity(0.4)
		swordShape.SetFriction(0.75)
		swordShape.UserData = 1 // hack to remember later that this shape was added for attack
		swordShape.SetCollisionType(entity.GenericEnemyCollisionType)
	} else if notAttacking {
		// remove any shapes that were added just for attack
		alligator.Body.EachShape(func(shape *cp.Shape) {
			if shape.UserData == 1 {
				alligator.Body.RemoveShape(shape)
				g.space.RemoveShape(shape)
			}
		})
	}

	// move the alligator
	if alligator.State == entity.MovingLeft {
		alligator.Facing = entity.Left
		vx, vy := alligator.Body.Velocity().X, 0.0
		vx -= settings.SwordDogAccelerationStep
		alligator.Body.ApplyForceAtWorldPoint(cp.Vector{X: vx, Y: vy}, alligator.Body.Position())
	} else if alligator.State == entity.MovingRight {
		alligator.Facing = entity.Right
		vx, vy := alligator.Body.Velocity().X, 0.0
		vx += settings.SwordDogAccelerationStep
		alligator.Body.ApplyForceAtWorldPoint(cp.Vector{X: vx, Y: vy}, alligator.Body.Position())
	}
}

func (g *Game) MoveFrog(frog *entity.Entity) {
	// first, check if we're dead
	if g.handleDeadEnemy(frog) {
		return
	}

	var frogX = frog.Body.Position().X
	var playerX = g.registry.Player().Body.Position().X

	// if idle, switch direction to face the player
	if frog.Grounded {
		if playerX > frogX {
			frog.State = entity.IdleRight
			frog.Facing = entity.Right
		} else if playerX <= frogX {
			frog.State = entity.IdleLeft
			frog.Facing = entity.Left
		}
		// jump if occasionally when in view
		if g.vp.InView(frog) {
			frog.TickCounter++
			if frog.TickCounter > 250 {
				frog.Boost = 2
				frog.TickCounter = 0
				if frog.State == entity.IdleLeft {
					frog.State = entity.JumpingLeft
				} else {
					frog.State = entity.JumpingRight
				}
			}
		}
	}

	// if frog is jumping up, adjust velocities accordingly
	if frog.State == entity.JumpingLeft || frog.State == entity.JumpingRight {
		frog.Grounded = false
		frog.Shape.SetFriction(0)

		// randomize velocities a little bit to keep things interesting
		xVelocity := rand.IntN(5000-2500+1) + 2500
		maxYVelocity := 10000
		minYVelocity := 5000
		yVelocity := rand.IntN(maxYVelocity-minYVelocity+1) + minYVelocity

		if frog.State == entity.JumpingLeft {
			xVelocity = -xVelocity
		}
		if frog.Boost > 0 {
			frog.Body.ApplyForceAtLocalPoint(cp.Vector{X: float64(xVelocity), Y: float64(-yVelocity)}, cp.Vector{X: 0, Y: 0})
			frog.Boost--
		}
	}

	// if frog has steady downward velocity, it must be falling
	if frog.Body.Velocity().Y > 70 && !frog.OnSlope {
		if frog.Facing == entity.Right {
			frog.State = entity.FallingRight
		}
		if frog.Facing == entity.Left {
			frog.State = entity.FallingLeft
		}
		frog.Grounded = false
		frog.Shape.SetFriction(0)
	}
}

func (g *Game) MoveEagle(eagle *entity.Entity) {
	// first, check if we're dead
	if g.handleDeadEnemy(eagle) {
		return
	}

	xVelocity := 50.0
	if eagle.State == entity.MovingLeft {
		if eagle.Body.Position().X < eagle.OriginX-80 {
			eagle.State = entity.MovingRight
		} else {
			eagle.Body.ApplyForceAtLocalPoint(cp.Vector{X: -xVelocity, Y: 0}, cp.Vector{X: 0, Y: 0})
		}
	} else if eagle.State == entity.MovingRight {
		if eagle.Body.Position().X > eagle.OriginX+80 {
			eagle.State = entity.MovingLeft
		} else {
			eagle.Body.ApplyForceAtLocalPoint(cp.Vector{X: xVelocity, Y: 0}, cp.Vector{X: 0, Y: 0})
		}
	}

	// TODO: divebomb the player if x and y coordinates are close

}

var EntityBehavior map[entity.EntityType]entity.Behavior

func InitializeEntityBehavior(g *Game) {
	EntityBehavior = make(map[entity.EntityType]entity.Behavior)
	EntityBehavior[entity.Player] = g.MovePlayer
	EntityBehavior[entity.SwordDog] = g.MoveSwordDog
	EntityBehavior[entity.Alligator] = g.MoveAlligator
	EntityBehavior[entity.Frog] = g.MoveFrog
	EntityBehavior[entity.Eagle] = g.MoveEagle
}
