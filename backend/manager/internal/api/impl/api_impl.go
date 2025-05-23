package apiimpl

import (
	"context"
	"github.com/sky0621/my-test-project/backend/manager/internal/api"
	contentport "github.com/sky0621/my-test-project/backend/manager/internal/content/port"
	courseport "github.com/sky0621/my-test-project/backend/manager/internal/course/port"
)

var _ api.StrictServerInterface = (*strictServerImpl)(nil)

func New(
	contentsController contentport.ContentController,
	coursesController courseport.CourseController,
	middlewares []api.StrictMiddlewareFunc,
) api.ServerInterface {
	return api.NewStrictHandler(&strictServerImpl{
		contentsController,
		coursesController,
	}, middlewares)
}

type strictServerImpl struct {
	contentsController contentport.ContentController
	coursesController  courseport.CourseController
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

func (s strictServerImpl) GetCourses(ctx context.Context, request api.GetCoursesRequestObject) (api.GetCoursesResponseObject, error) {
	return s.coursesController.GetCourses(ctx, request)
}

func (s strictServerImpl) PostCourses(ctx context.Context, request api.PostCoursesRequestObject) (api.PostCoursesResponseObject, error) {
	return s.coursesController.PostCourses(ctx, request)
}
