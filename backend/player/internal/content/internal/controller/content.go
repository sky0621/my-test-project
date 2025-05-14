package controller

import (
	"context"
	"github.com/sky0621/my-test-project/backend/player/internal/api"
	"github.com/sky0621/my-test-project/backend/player/internal/content/internal/application/command"
	"github.com/sky0621/my-test-project/backend/player/internal/content/internal/application/query"
	"github.com/sky0621/my-test-project/backend/player/internal/content/internal/domain/model"
	"github.com/sky0621/my-test-project/backend/player/internal/content/port"
	"github.com/sky0621/my-test-project/backend/shared/converter"
	"github.com/sky0621/my-test-project/backend/shared/service"
)

var _ port.ContentController = (*impl)(nil)

func New(searchContents query.SearchContents, getContent query.GetContent, saveContent command.SaveContent) port.ContentController {
	return impl{searchContents: searchContents, getContent: getContent, saveContent: saveContent}
}

type impl struct {
	searchContents query.SearchContents
	getContent     query.GetContent
	saveContent    command.SaveContent
}

func (c impl) PostContents(ctx context.Context, request api.PostContentsRequestObject) (api.PostContentsResponseObject, error) {
	newContentID := service.MustCreateNewID()
	name := request.Body.Name
	programs := make([]model.ProgramWriteModel, len(request.Body.Programs))
	for i, program := range request.Body.Programs {
		programID := service.MustCreateNewID()
		programs[i] = model.ProgramWriteModel{
			ID:       programID,
			Question: program.Question,
			Answer:   program.Answer,
		}
	}

	if err := c.saveContent.Save(ctx, model.SaveContentWriteModel{
		ID:       newContentID,
		Name:     name,
		Programs: programs,
	}); err != nil {
		return nil, err
	}
	return api.PostContents201JSONResponse(api.ContentResponse{
		ContentID: newContentID.String(),
		Name:      name,
		Programs: func() []api.ProgramResponse {
			ret := make([]api.ProgramResponse, len(programs))
			for i, program := range programs {
				ret[i] = api.ProgramResponse{
					ProgramID: program.ID.String(),
					Question:  &program.Question,
					Answer:    &program.Answer,
				}
			}
			return ret
		}(),
	}), nil
}

func (c impl) GetContents(ctx context.Context, _ api.GetContentsRequestObject) (api.GetContentsResponseObject, error) {
	contents, err := c.searchContents.Exec(ctx)
	if err != nil {
		return nil, err
	}
	responses := make([]api.ContentResponse, len(contents))
	for i, content := range contents {
		responses[i] = api.ContentResponse{
			ContentID: content.ID.String(),
			Name:      content.Name,
			Programs: func() []api.ProgramResponse {
				programs := make([]api.ProgramResponse, len(content.Programs))
				for i, program := range content.Programs {
					programs[i] = api.ProgramResponse{
						ProgramID: program.ID.String(),
						Question:  &program.Question,
						Answer:    &program.Answer,
					}
				}
				return programs
			}(),
		}
	}
	return api.GetContents200JSONResponse(responses), nil
}

func (c impl) GetContentsByID(ctx context.Context, request api.GetContentsByIDRequestObject) (api.GetContentsByIDResponseObject, error) {
	_, err := service.ParseID(request.ID)
	if err != nil {
		return api.GetContentsByID400JSONResponse{Message: "not uuid"}, nil
	}
	content, err := c.getContent.Exec(ctx, request.ID)
	if err != nil {
		return nil, err
	}
	if content.IsEmpty() {
		return api.GetContentsByID404JSONResponse{Message: converter.ToPtr("not found")}, nil
	}
	return api.GetContentsByID200JSONResponse(api.ContentResponse{
		ContentID: content.ID.String(),
		Name:      content.Name,
		Programs: func() []api.ProgramResponse {
			programs := make([]api.ProgramResponse, len(content.Programs))
			for i, program := range content.Programs {
				programs[i] = api.ProgramResponse{
					ProgramID: program.ID.String(),
					Question:  &program.Question,
					Answer:    &program.Answer,
				}
			}
			return programs
		}(),
	}), nil
}
