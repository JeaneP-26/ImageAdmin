package storage

import (
	"fmt"
    "IMAGEADMIN/models"
	
type Storage struct {
	usuarios    map[int]models.Usuario
	imagenes    map[int]models.Imagen
	categorias  map[int]models.Categoria
	contUsers   int
	contImages  int
	contCats    int
}
 
func NuevoStorage() *Storage {
	return &Storage{
		usuarios:   make(map[int]models.Usuario),
		imagenes:   make(map[int]models.Imagen),
		categorias: make(map[int]models.Categoria),
	}
}