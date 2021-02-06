package repository

import "github.com/TeamKitten/NoteNote/api/domain/model"

type NoteRepository interface {
	FindOne(id string) (note *model.Note, err error)
	Set(id string, note *model.Note) error
}
