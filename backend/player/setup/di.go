package setup

import (
	"database/sql"
	"github.com/sky0621/my-test-project/backend/player/internal/api"
	apiimpl "github.com/sky0621/my-test-project/backend/player/internal/api/impl"
	"github.com/sky0621/my-test-project/backend/player/internal/content/application/command"
	contentsController "github.com/sky0621/my-test-project/backend/player/internal/content/controller"
	"github.com/sky0621/my-test-project/backend/player/internal/content/infrastructure"
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
		nil,
	)
}
