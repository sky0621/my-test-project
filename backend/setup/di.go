package setup

import (
	"database/sql"
	"github.com/sky0621/my-test-project/backend/internal/api"
	apiimpl "github.com/sky0621/my-test-project/backend/internal/api/impl"
	"github.com/sky0621/my-test-project/backend/internal/content/application/command"
	contentsController "github.com/sky0621/my-test-project/backend/internal/content/controller"
	"github.com/sky0621/my-test-project/backend/internal/content/infrastructure"
	coursesController "github.com/sky0621/my-test-project/backend/internal/course/controller"
)

func createHandlers(db *sql.DB) api.ServerInterface {
	return apiimpl.New(
		contentsController.NewContent(
			infrastructure.NewSearchContents(db),
			infrastructure.NewGetContent(db),
			command.NewSaveContent(
				infrastructure.NewContentRepository(db),
			),
		),
		coursesController.Course{},
		nil,
	)
}
