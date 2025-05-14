package setup

import (
	"database/sql"
	"github.com/sky0621/my-test-project/backend/manager/internal/api"
	apiimpl "github.com/sky0621/my-test-project/backend/manager/internal/api/impl"
	"github.com/sky0621/my-test-project/backend/manager/internal/content"
	coursesController "github.com/sky0621/my-test-project/backend/manager/internal/course/controller"
)

func createHandlers(db *sql.DB) api.ServerInterface {
	return apiimpl.New(
		content.NewController(db),
		coursesController.Course{},
		nil,
	)
}
