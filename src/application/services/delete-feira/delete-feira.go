package delete_feira

import (
	"api-unico/application/interfaces"
)

type deleteFeiraService struct {
	repository interfaces.FeiraRepository
}

func NewDeleteFeiraService(repository interfaces.FeiraRepository) deleteFeiraService {
	return deleteFeiraService{
		repository: repository,
	}
}

func (h *deleteFeiraService) Execute(id int64) (err error) {
	if err = h.repository.Delete(id); err != nil {
		return err
	}

	return nil
}
