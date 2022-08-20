package delete_feira

import (
	"api-unico/application/interfaces"
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
		return err
	}

	return nil
}
