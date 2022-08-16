package gormdb

import "time"

type Resultado struct {
	Id             uint32    `json:"id"`
	Hora_resultado time.Time `json:"fechahoraapuesta,omitempty"`
	Like_count     *uint32   `json:"like_count,omitempty"`
	Views_count    *uint32   `json:"view_count,omitempty"`
	Comments_count *uint32   `json:"comments_count,omitempty"`
	Dislikes_count *uint32   `json:"dislikes,omitempty"`
	Saved_count    *uint32   `json:"saved_count,omitempty"`
	Shared_count   *uint32   `json:"shared_count,omitempty"`
	Evento_id      *uint32   `json:"evento_id,omitempty"`
	Video_id       *uint32   `json:"video_id,omitempty"`
}

func (product *Resultado) TableName() string {
	return "resultado"
}
