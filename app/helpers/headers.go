package helpers

import (
	"context"
	"fmt"

	"google.golang.org/grpc/metadata"
)

type ContextKey string

//	type Header struct {
//			ContentTypeCtxKey *contextKey
//			LanguageCtxKey *contextKey
//			AcceptctxKey *contextKey
//			AuthorizationCtxKey *contextKey
//	}
const (
	ContentTypeCtxKey   ContextKey = "Content-Type"
	LanguageCtxKey      ContextKey = "Accept-Language"
	AcceptctxKey        ContextKey = "Accept"
	AuthorizationCtxKey ContextKey = "authorization"
	UserId              ContextKey = "X-User-ID"
	UserRole            ContextKey = "x-userRole"
	UserFirstName       ContextKey = "x-userFirstName"
	UserLastName        ContextKey = "x-userLastName"
	UserEmail           ContextKey = "x-userEmail"
)

func AddMetadataToContext(ctx context.Context) (context.Context, error) {
	userId, ok := ctx.Value(UserId).(string)
	if !ok {
		return nil, fmt.Errorf("user not found")
	}
	userEmail, ok := ctx.Value(UserEmail).(string)
	if !ok {
		return nil, fmt.Errorf("user not found")
	}
	userRole, ok := ctx.Value(UserRole).(string)
	if !ok {
		return nil, fmt.Errorf("user not found")
	}

	userFirstname, ok := ctx.Value(UserFirstName).(string)
	if !ok {
		return nil, fmt.Errorf("user not found")
	}

	md := metadata.Pairs(
		string(UserId), userId,
		string(UserEmail), userEmail,
		string(UserRole), userRole,
		string(UserFirstName), userFirstname,
	)

	userLastName, ok := ctx.Value(UserLastName).(string)
	if ok {
		md.Append(string(UserLastName), userLastName)
	}

	ctx = metadata.NewOutgoingContext(ctx, md)

	return ctx, nil
}
