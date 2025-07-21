package entity

import (
	"github.com/jakecoffman/cp"
)

const (
	PlayerCollisionType cp.CollisionType = iota
	BlockCollisionType
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
			} else {
				body2.UserData.(*Entity).Grounded = true
			}
		}
		return true
	}
}
