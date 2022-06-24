package helpers

import "lottomusic/src/models/gormdb"

func RemoveDuplicate(intSlice []gormdb.Apuesta_usuario) []uint32 {
	keys := make(map[uint32]bool)
	list := []uint32{}

	for _, entry := range intSlice {
		if _, value := keys[*entry.Apuesta]; !value {
			keys[*entry.Apuesta] = true
			list = append(list, *entry.Apuesta)
		}
	}
	return list
}
