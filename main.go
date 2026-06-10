package main

import (
	"fmt"
	"image-system/config"
	"image-system/db"
	"image-system/handlers"
	"net/http"
)

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, X-User-ID")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func main() {
	cfg := config.Obtener()
	base := db.Conectar(cfg.ArchivoBD)

	imagenHandler := handlers.NuevoImagenHandler(base, cfg)
	albumHandler := handlers.NuevoAlbumHandler(base, cfg)

	rutas := http.NewServeMux()

	rutas.HandleFunc("/images/upload", imagenHandler.SubirImagen)
	rutas.HandleFunc("/images", imagenHandler.ListarImagenes)
	rutas.HandleFunc("/images/editar", imagenHandler.EditarDescripcion)
	rutas.HandleFunc("/images/eliminar", imagenHandler.EliminarImagen)
	rutas.HandleFunc("/images/buscar", imagenHandler.BuscarImagenes)
	rutas.HandleFunc("/albums", albumHandler.ListarAlbums)
	rutas.HandleFunc("/albums/crear", albumHandler.CrearAlbum)

	fmt.Println("Servidor corriendo en", cfg.PuertoGo)

	if err := http.ListenAndServe(cfg.PuertoGo, corsMiddleware(rutas)); err != nil {
		fmt.Println("Error al iniciar el servidor:", err)
	}
}
