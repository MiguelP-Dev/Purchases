package main

import (
	"context"
	"log"
	"net/http"

	"github.com/MiguelP-Dev/Purchases/handlers"
	"github.com/MiguelP-Dev/Purchases/models"
	"github.com/gorilla/mux"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	// Configurar base de datos SQLite
	db, err := gorm.Open(sqlite.Open("sqlitefile.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Error al conectar con la base de datos:", err)
	}

	// Auto-migrar los modelos
	err = db.AutoMigrate(&models.ShoppingList{}, &models.ListItem{})
	if err != nil {
		log.Fatal("Error al migrar la base de datos:", err)
	}

	// Crear router
	router := mux.NewRouter()

	// Middleware para inyectar la base de datos en el contexto
	dbMiddleware := func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := context.WithValue(r.Context(), "db", db)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}

	// Aplicar middleware
	router.Use(dbMiddleware)

	// Servir favicon
	router.HandleFunc("/favicon.svg", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "templates/favicon.svg")
	})

	// Servir CSS
	router.HandleFunc("/styles.css", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/css")
		http.ServeFile(w, r, "templates/styles.css")
	})

	// Rutas principales
	router.HandleFunc("/", handlers.IndexHandler).Methods("GET")
	router.HandleFunc("/create", handlers.CreateListHandler).Methods("GET", "POST")

	// Rutas de listas
	router.HandleFunc("/list/{id}", handlers.ShowListHandler).Methods("GET")
	router.HandleFunc("/list/{id}/edit", handlers.EditListHandler).Methods("GET", "POST")
	router.HandleFunc("/list/{id}/delete", handlers.DeleteListHandler).Methods("POST")
	router.HandleFunc("/list/{id}/clone", handlers.CloneListHandler).Methods("POST")
	router.HandleFunc("/list/{id}/items", handlers.AddItemHandler).Methods("POST")

	// Rutas de items
	router.HandleFunc("/items/{id}/delete", handlers.DeleteItemHandler).Methods("POST")
	router.HandleFunc("/items/{id}/toggle", handlers.ToggleItemHandler).Methods("POST")
	router.HandleFunc("/items/{id}/mark-purchased", handlers.MarkItemAsPurchasedHandler).Methods("GET", "POST")
	router.HandleFunc("/items/{id}/edit", handlers.EditItemHandler).Methods("GET", "POST")

	// Configurar servidor
	port := ":8080"
	log.Printf("Servidor iniciado en http://localhost%s", port)
	log.Fatal(http.ListenAndServe(port, router))
}
