package apiimpl

import (
	"context"
	"github.com/sky0621/my-test-project/app/internal/api"
	contentsController "github.com/sky0621/my-test-project/app/internal/content/controller"
	coursesController "github.com/sky0621/my-test-project/app/internal/course/controller"
)

var _ api.StrictServerInterface = (*strictServerImpl)(nil)

func New(
	contentsController contentsController.Content,
	coursesController coursesController.Course,
	middlewares []api.StrictMiddlewareFunc,
) api.ServerInterface {
	return api.NewStrictHandler(&strictServerImpl{
		contentsController,
		coursesController,
	}, middlewares)
}

type strictServerImpl struct {
	contentsController contentsController.Content
	coursesController  coursesController.Course
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
