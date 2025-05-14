package setup

import (
	"database/sql"
	"github.com/sky0621/my-test-project/backend/player/internal/api"
	apiimpl "github.com/sky0621/my-test-project/backend/player/internal/api/impl"
	"github.com/sky0621/my-test-project/backend/player/internal/content"
)

func createHandlers(db *sql.DB) api.ServerInterface {
	return apiimpl.New(
		content.NewController(db),
		nil,
	)
}
