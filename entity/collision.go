package entity

import (
	"github.com/jakecoffman/cp"
)

const (
	PlayerCollisionType cp.CollisionType = iota
	BlockCollisionType
	RSlopeCollisionType
)

func GenericGroundedHandler(space *cp.Space, collisionType cp.CollisionType) {
	handler := space.NewCollisionHandler(collisionType, BlockCollisionType)

	handler.BeginFunc = func(arb *cp.Arbiter, space *cp.Space, data interface{}) bool {
		n := arb.Normal()
		grounded := n.Y > 0
		if grounded {
			body1, body2 := arb.Bodies()
			if body1.UserData != nil {
				body1.UserData.(*Entity).Grounded = true
				body1.UserData.(*Entity).Shape.SetFriction(0.75)
			} else {
				body2.UserData.(*Entity).Grounded = true
				body1.UserData.(*Entity).Shape.SetFriction(0.75)
			}
		}
		return true
	}
}

func SlopeHandler(space *cp.Space, collisionType cp.CollisionType) {
	handler := space.NewCollisionHandler(collisionType, RSlopeCollisionType)

	handler.BeginFunc = func(arb *cp.Arbiter, space *cp.Space, data interface{}) bool {
		n := arb.Normal()
		body1, body2 := arb.Bodies()
		grounded := n.Y > 0
		ascending := n.X > 0

		if ascending {
			if body1.UserData != nil {
				body := body1.UserData.(*Entity).Body

				body.SetVelocity(body.Velocity().X, body.Velocity().Y-50)

				body1.UserData.(*Entity).OnSlope = true
			} else {
				body := body2.UserData.(*Entity).Body
				body.SetVelocity(body.Velocity().X, body.Velocity().Y-50)
				body2.UserData.(*Entity).OnSlope = true
			}
		}
		if grounded {
			if body1.UserData != nil {
				body1.UserData.(*Entity).Grounded = true
				body1.UserData.(*Entity).Shape.SetFriction(0.75)
			} else {
				body2.UserData.(*Entity).Grounded = true
				body1.UserData.(*Entity).Shape.SetFriction(0.75)
			}
		}
		return true
	}

	handler.SeparateFunc = func(arb *cp.Arbiter, space *cp.Space, data interface{}) {
		body1, body2 := arb.Bodies()
		if body1.UserData != nil {
			body1.UserData.(*Entity).OnSlope = false
		} else {
			body2.UserData.(*Entity).OnSlope = false
		}
	}
}
