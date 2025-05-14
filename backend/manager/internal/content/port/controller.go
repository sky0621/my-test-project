package port

import (
	"context"
	"github.com/sky0621/my-test-project/backend/manager/internal/api"
)

type ContentController interface {
	PostContents(ctx context.Context, request api.PostContentsRequestObject) (api.PostContentsResponseObject, error)
	GetContents(ctx context.Context, request api.GetContentsRequestObject) (api.GetContentsResponseObject, error)
	GetContentsByID(ctx context.Context, request api.GetContentsByIDRequestObject) (api.GetContentsByIDResponseObject, error)
}
