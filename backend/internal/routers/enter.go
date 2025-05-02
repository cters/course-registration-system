package routers

import "github.com/QuanCters/backend/internal/routers/user"

type RouterGroup struct {
	User user.UserRouterGroup
}

var RouterGroupApp = new(RouterGroup)