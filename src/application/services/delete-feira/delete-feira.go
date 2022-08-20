package delete_feira

import (
	"api-unico/application/interfaces"
	"api-unico/infra/logger"
)

type DeleteFeiraService struct {
	repository interfaces.FeiraRepository
}

func NewDeleteFeiraService(repository interfaces.FeiraRepository) DeleteFeiraService {
	return DeleteFeiraService{
		repository: repository,
	}
}

func (h *DeleteFeiraService) Execute(id int64) (err error) {
	if err = h.repository.Delete(id); err != nil {
		logger.LogError("DeleteFeiraService", err.Error())
		return err
	}

	return nil
}
