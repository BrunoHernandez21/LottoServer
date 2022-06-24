package planes

import "lottomusic/src/models/gormdb"

type Set_User struct {
	MisApuestas  []gormdb.Apuesta_usuario `json:"mis_apuestas"`
	TipoApuestas []gormdb.Apuestas        `json:"tipo_apuestas"`
}
