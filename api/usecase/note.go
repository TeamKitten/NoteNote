package usecase

import (
	"github.com/TeamKitten/NoteNote/api/domain/model"
	"github.com/TeamKitten/NoteNote/api/domain/repository"
)

type NoteUsecase interface {
	Get(id string) (note *model.Note, err error)
	Set(id string, note *model.Note) error
}

type noteUsecase struct {
	noteRepo repository.NoteRepository
}

func NewNoteUsecase(noteRepo repository.NoteRepository) NoteUsecase {
	noteUsecase := noteUsecase{noteRepo}

	return &noteUsecase
}
func (usecase *noteUsecase) Get(id string) (note *model.Note, err error) {
	return usecase.noteRepo.FindOne(id)
}
func (usecase *noteUsecase) Set(id string, note *model.Note) error {
	return usecase.noteRepo.Set(id, note)
}
