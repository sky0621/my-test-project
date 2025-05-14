package port

import (
	"context"
	"github.com/sky0621/my-test-project/backend/manager/internal/api"
)

type CourseController interface {
	GetCourses(ctx context.Context, request api.GetCoursesRequestObject) (api.GetCoursesResponseObject, error)
	PostCourses(ctx context.Context, request api.PostCoursesRequestObject) (api.PostCoursesResponseObject, error)
}
