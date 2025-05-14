package course

import (
	"github.com/sky0621/my-test-project/backend/manager/internal/course/internal/controller"
	"github.com/sky0621/my-test-project/backend/manager/internal/course/port"
)

func NewController() port.CourseController {
	return controller.New()
}
