package injector

import (
	"github.com/TeamKitten/NoteNote/server/domain/repository"
	"github.com/TeamKitten/NoteNote/server/handler"
	"github.com/TeamKitten/NoteNote/server/infra"
	"github.com/TeamKitten/NoteNote/server/usecase"
)

func InjectDB() infra.DBHandler {
	dbHandler := infra.NewDBHandler()
	return *dbHandler
}

func InjectNoteRepository() repository.NoteRepository {
	dbHandler := InjectDB()
	return infra.NewNoteRepository(dbHandler)
}

func InjectNoteUsecase() usecase.NoteUsecase {
	noteRepo := InjectNoteRepository()
	return usecase.NewNoteUsecase(noteRepo)
}

func InjectNoteHandler() handler.NoteHandler {
	noteUsecase := InjectNoteUsecase()
	return handler.NewNoteHandler(noteUsecase)
}
