package utils

import (
	"context"
)

func GetUserIdFromCtx(ctx context.Context) int32 {
	userId := ctx.Value(SessionUserId)
	if userId == nil {
		return 0
	}
	return userId.(int32)
}

func GetUserEmailFromCtx(ctx context.Context) string {
	Email := ctx.Value("email")
	if Email == nil {
		return ""	
	}
	return Email.(string)
}