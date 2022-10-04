package inputs

import "lottomusic/src/models/gormdb"

type Get_Event_Video struct {
	Genero   *string         `json:"genero"`
	Video_id string          `json:"video_id"`
	Evento   *gormdb.Eventos `json:"evento"`
}
