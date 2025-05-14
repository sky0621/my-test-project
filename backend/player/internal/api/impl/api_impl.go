package apiimpl

import (
	"context"
	"github.com/sky0621/my-test-project/backend/player/internal/api"
	contentport "github.com/sky0621/my-test-project/backend/player/internal/content/port"
)

var _ api.StrictServerInterface = (*strictServerImpl)(nil)

func New(
	contentsController contentport.ContentController,
	middlewares []api.StrictMiddlewareFunc,
) api.ServerInterface {
	return api.NewStrictHandler(&strictServerImpl{
		contentsController,
	}, middlewares)
}

type strictServerImpl struct {
	contentsController contentport.ContentController
}

func (s strictServerImpl) GetContents(ctx context.Context, request api.GetContentsRequestObject) (api.GetContentsResponseObject, error) {
	return s.contentsController.GetContents(ctx, request)
}

func (s strictServerImpl) GetContentsByID(ctx context.Context, request api.GetContentsByIDRequestObject) (api.GetContentsByIDResponseObject, error) {
	return s.contentsController.GetContentsByID(ctx, request)
}

func (s strictServerImpl) PostContents(ctx context.Context, request api.PostContentsRequestObject) (api.PostContentsResponseObject, error) {
	return s.contentsController.PostContents(ctx, request)
}
