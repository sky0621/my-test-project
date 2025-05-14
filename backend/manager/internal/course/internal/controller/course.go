package controller

import (
	"context"
	"github.com/sky0621/my-test-project/backend/manager/internal/api"
	"github.com/sky0621/my-test-project/backend/manager/internal/course/port"
)

var _ port.CourseController = (*impl)(nil)

func New() port.CourseController {
	return impl{}
}

type impl struct {
}

func (c impl) GetCourses(ctx context.Context, request api.GetCoursesRequestObject) (api.GetCoursesResponseObject, error) {
	//TODO implement me
	panic("implement me")
}

func (c impl) PostCourses(ctx context.Context, request api.PostCoursesRequestObject) (api.PostCoursesResponseObject, error) {
	//TODO implement me
	panic("implement me")
}
