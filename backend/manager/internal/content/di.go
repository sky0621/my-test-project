package content

import (
	"database/sql"
	"github.com/sky0621/my-test-project/backend/manager/internal/content/internal/application/command"
	"github.com/sky0621/my-test-project/backend/manager/internal/content/internal/controller"
	"github.com/sky0621/my-test-project/backend/manager/internal/content/internal/infrastructure"
	"github.com/sky0621/my-test-project/backend/manager/internal/content/port"
)

func NewController(db *sql.DB) port.ContentController {
	return controller.NewContentController(
		infrastructure.NewSearchContents(db),
		infrastructure.NewGetContent(db),
		command.NewSaveContent(
			infrastructure.NewContentRepository(db),
		),
	)
}
