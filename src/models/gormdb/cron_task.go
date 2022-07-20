package gormdb

type Cron_task struct {
	Id         uint32  `json:"id"`
	Tarea_cron *string `json:"tarea_cron,omitempty"`
	Evento_id  *uint32 `json:"evento_id,omitempty"`
}

func (product *Cron_task) TableName() string {
	return "cron_task"
}
