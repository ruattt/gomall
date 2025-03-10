package middleware

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/hertz-contrib/sessions"
	frontendUtils "gomall/app/frontend/utils"
)

func GlobalAuth() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		s := sessions.Default(c)
		userId := s.Get("user_id")
		ctx = context.WithValue(ctx, frontendUtils.SessionUserId, userId)

		// s_id := sessions.DefaultMany(c,"user_id")
		// s_em := sessions.DefaultMany(c,"user_email")
		// ctx = context.WithValue(ctx, frontendUtils.SessionUserId, s_id.Get("user_id"))
		// ctx = context.WithValue(ctx,"email", s_em.Get("email"))
		c.Next(ctx)
	}
}

func Auth() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		s := sessions.Default(c)
		userId := s.Get("user_id")
		if userId == nil {
			c.Redirect(consts.StatusFound, []byte("/sign-in?next="+c.FullPath()))
			c.Abort()
			return
		}
		c.Next(ctx)
	}
}
