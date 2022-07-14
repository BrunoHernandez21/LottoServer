package gormdb

type Resultado struct {
	Id               uint32  `json:"id"`
	Comment_count    *int32  `json:"comment_count,omitempty"`
	Fechahoraapuesta *string `json:"fechahoraapuesta,omitempty"`
	Id_apuesta       *uint32 `json:"id_apuesta,omitempty"`
	Id_video         *uint32 `json:"id_video,omitempty"`
	Like_count       *uint32 `json:"like_count,omitempty"`
	View_count       *uint32 `json:"view_count,omitempty"`
}

func (product *Resultado) TableName() string {
	return "Resultado"
}
