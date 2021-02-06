package infra

import (
	"encoding/json"
	"fmt"

	"github.com/TeamKitten/NoteNote/server/domain/model"
	"github.com/TeamKitten/NoteNote/server/domain/repository"
	"github.com/tidwall/buntdb"
)

type NoteRepository struct {
	DBHandler
}

func NewNoteRepository(dbHandler DBHandler) repository.NoteRepository {
	noteRepository := NoteRepository{dbHandler}
	return &noteRepository
}
func (noteRepo *NoteRepository) FindOne(id string) (note *model.Note, err error) {
	err = noteRepo.db.View(func(tx *buntdb.Tx) error {
		value, err := tx.Get(fmt.Sprintf("notes:%s", id))
		if err != nil {
			return err
		}
		var note *model.Note
		err = json.Unmarshal([]byte(value), &note)
		if err != nil {
			return err
		}
		return nil
	})
	return note, err
}
func (noteRepo *NoteRepository) Set(id string, note *model.Note) error {
	// t := time.Unix(1000000, 0)
	// entropy := ulid.Monotonic(rand.New(rand.NewSource(t.UnixNano())), 0)
	// newId := ulid.MustNew(ulid.Timestamp(t), entropy)
	noteBytes, err := json.Marshal(note)
	if err != nil {
		return err
	}

	return noteRepo.db.Update(func(tx *buntdb.Tx) error {
		_, _, err = tx.Set(fmt.Sprintf("notes:%s", id), string(noteBytes), nil)
		return err
	})
}
