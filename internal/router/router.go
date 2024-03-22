package router

import "gin-template/internal/router/admin"

type Group struct {
	admin.UserRouter
}

var AllRouter = new(Group)
