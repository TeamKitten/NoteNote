package handler

import (
	"encoding/json"

	"github.com/TeamKitten/NoteNote/api/domain/model"
	"github.com/TeamKitten/NoteNote/api/usecase"
	"github.com/gin-gonic/gin"
	"github.com/tidwall/buntdb"
)

type NoteHandler struct {
	noteUsecase usecase.NoteUsecase
}

func NewNoteHandler(noteUsecase usecase.NoteUsecase) NoteHandler {
	noteHandler := NoteHandler{noteUsecase}
	return noteHandler
}

func (noteHandler *NoteHandler) Set(c *gin.Context) {
	id := c.Param("id")
	buf := make([]byte, 2048)
	_, err := c.Request.Body.Read(buf)
	if err != nil {
		c.Status(400)
		return
	}
	var note model.Note
	json.Unmarshal(buf, &note)
	err = noteHandler.noteUsecase.Set(id, &note)
	if err != nil {
		c.Status(500)
		return
	}
	c.Status(200)
}

func (noteHandler *NoteHandler) Get(c *gin.Context) {
	note, err := noteHandler.noteUsecase.Get(c.Param("id"))
	if err != nil {
		if err == buntdb.ErrNotFound {
			c.Status(404)
			return
		}
		c.Status(500)
		return
	}
	c.JSON(200, note)
}
