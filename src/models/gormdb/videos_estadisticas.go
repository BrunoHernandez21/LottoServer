package gormdb

import "time"

type VideosEstadisticas struct {
	Id             uint32    `json:"id"`
	Video_id       uint32    `json:"video_id"`
	Fecha          time.Time `json:"fecha,omitempty"`
	Like_count     *uint32   `json:"like_count,omitempty"`
	Views_count    *uint32   `json:"views_count,omitempty"`
	Comments_count *uint32   `json:"comments_count,omitempty"`
	Dislikes_count *uint32   `json:"dislikes_count,omitempty"`
	Saved_count    *uint32   `json:"saved_count,omitempty"`
	Shared_count   *uint32   `json:"shared_count,omitempty"`
}

func (product *VideosEstadisticas) TableName() string {
	return "videos_estadisticas"
}
