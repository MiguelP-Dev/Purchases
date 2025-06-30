package handlers

import (
	"net/http"
	"strconv"
	"text/template"

	"github.com/MiguelP-Dev/Purchases/models"
	"github.com/MiguelP-Dev/Purchases/templates"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

// IndexHandler muestra la página principal con todas las listas
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	db := r.Context().Value("db").(*gorm.DB)

	var lists []models.ShoppingList
	if err := db.Preload("Items").Find(&lists).Error; err != nil {
		http.Error(w, "Error al cargar las listas", http.StatusInternalServerError)
		return
	}

	tmpl, err := template.New("index.html").Funcs(templates.Functions).ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, "Error al cargar la plantilla", http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, lists)
}

// CreateListHandler crea una nueva lista de compras
func CreateListHandler(w http.ResponseWriter, r *http.Request) {
	db := r.Context().Value("db").(*gorm.DB)

	if r.Method == "GET" {
		tmpl, err := template.ParseFiles("templates/create_list.html")
		if err != nil {
			http.Error(w, "Error al cargar la plantilla", http.StatusInternalServerError)
			return
		}
		tmpl.Execute(w, nil)
		return
	}

	if r.Method == "POST" {
		r.ParseForm()
		title := r.FormValue("title")

		if title == "" {
			http.Error(w, "El título es obligatorio", http.StatusBadRequest)
			return
		}

		list := models.ShoppingList{
			Title: title,
		}

		if err := db.Create(&list).Error; err != nil {
			http.Error(w, "Error al crear la lista", http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "/list/"+strconv.FormatUint(uint64(list.ID), 10), http.StatusSeeOther)
	}
}

// ShowListHandler muestra los detalles de una lista específica
func ShowListHandler(w http.ResponseWriter, r *http.Request) {
	db := r.Context().Value("db").(*gorm.DB)

	vars := mux.Vars(r)
	listID, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "ID de lista inválido", http.StatusBadRequest)
		return
	}

	var list models.ShoppingList
	if err := db.Preload("Items").First(&list, listID).Error; err != nil {
		http.Error(w, "Lista no encontrada", http.StatusNotFound)
		return
	}

	tmpl, err := template.New("list.html").Funcs(templates.Functions).ParseFiles("templates/list.html")
	if err != nil {
		http.Error(w, "Error al cargar la plantilla", http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, list)
}

// DeleteListHandler elimina una lista de compras
func DeleteListHandler(w http.ResponseWriter, r *http.Request) {
	db := r.Context().Value("db").(*gorm.DB)

	vars := mux.Vars(r)
	listID, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "ID de lista inválido", http.StatusBadRequest)
		return
	}

	if err := db.Delete(&models.ShoppingList{}, listID).Error; err != nil {
		http.Error(w, "Error al eliminar la lista", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

// CloneListHandler clona una lista existente
func CloneListHandler(w http.ResponseWriter, r *http.Request) {
	db := r.Context().Value("db").(*gorm.DB)

	vars := mux.Vars(r)
	listID, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "ID de lista inválido", http.StatusBadRequest)
		return
	}

	var originalList models.ShoppingList
	if err := db.Preload("Items").First(&originalList, listID).Error; err != nil {
		http.Error(w, "Lista original no encontrada", http.StatusNotFound)
		return
	}

	// Crear nueva lista con el mismo título + " (copia)"
	newList := models.ShoppingList{
		Title: originalList.Title + " (copia)",
	}

	if err := db.Create(&newList).Error; err != nil {
		http.Error(w, "Error al crear la lista clonada", http.StatusInternalServerError)
		return
	}

	// Clonar todos los items pero con IsPurchased = false
	for _, item := range originalList.Items {
		newItem := models.ListItem{
			ShoppingListID: newList.ID,
			ProductName:    item.ProductName,
			IsPurchased:    false,          // Siempre false en la copia
			Price:          0.0,            // Resetear precio
			Place:          "supermercado", // Resetear lugar
		}
		db.Create(&newItem)
	}

	http.Redirect(w, r, "/list/"+strconv.FormatUint(uint64(newList.ID), 10), http.StatusSeeOther)
}

// AddItemHandler agrega un nuevo item a una lista
func AddItemHandler(w http.ResponseWriter, r *http.Request) {
	db := r.Context().Value("db").(*gorm.DB)

	if r.Method != "POST" {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}

	vars := mux.Vars(r)
	listID, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "ID de lista inválido", http.StatusBadRequest)
		return
	}

	r.ParseForm()
	productName := r.FormValue("product_name")

	if productName == "" {
		http.Error(w, "El nombre del producto es obligatorio", http.StatusBadRequest)
		return
	}

	item := models.ListItem{
		ShoppingListID: uint(listID),
		ProductName:    productName,
		IsPurchased:    false,
		Price:          0.0,
		Place:          "supermercado",
	}

	if err := db.Create(&item).Error; err != nil {
		http.Error(w, "Error al agregar el item", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/list/"+strconv.Itoa(listID), http.StatusSeeOther)
}

// DeleteItemHandler elimina un item de una lista
func DeleteItemHandler(w http.ResponseWriter, r *http.Request) {
	db := r.Context().Value("db").(*gorm.DB)

	vars := mux.Vars(r)
	itemID, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "ID de item inválido", http.StatusBadRequest)
		return
	}

	var item models.ListItem
	if err := db.First(&item, itemID).Error; err != nil {
		http.Error(w, "Item no encontrado", http.StatusNotFound)
		return
	}

	listID := item.ShoppingListID

	if err := db.Delete(&item).Error; err != nil {
		http.Error(w, "Error al eliminar el item", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/list/"+strconv.FormatUint(uint64(listID), 10), http.StatusSeeOther)
}

// EditListHandler permite editar el título de una lista
func EditListHandler(w http.ResponseWriter, r *http.Request) {
	db := r.Context().Value("db").(*gorm.DB)

	if r.Method == "GET" {
		vars := mux.Vars(r)
		listID := vars["id"]

		// Obtener la lista para mostrar el título actual
		listIDInt, err := strconv.Atoi(listID)
		if err != nil {
			http.Error(w, "ID de lista inválido", http.StatusBadRequest)
			return
		}

		var list models.ShoppingList
		if err := db.First(&list, listIDInt).Error; err != nil {
			http.Error(w, "Lista no encontrada", http.StatusNotFound)
			return
		}

		// Mostrar formulario para editar título
		tmpl, err := template.ParseFiles("templates/edit_list.html")
		if err != nil {
			http.Error(w, "Error al cargar la plantilla", http.StatusInternalServerError)
			return
		}

		tmpl.Execute(w, list)
		return
	}

	if r.Method == "POST" {
		vars := mux.Vars(r)
		listID, err := strconv.Atoi(vars["id"])
		if err != nil {
			http.Error(w, "ID de lista inválido", http.StatusBadRequest)
			return
		}

		var list models.ShoppingList
		if err := db.First(&list, listID).Error; err != nil {
			http.Error(w, "Lista no encontrada", http.StatusNotFound)
			return
		}

		r.ParseForm()
		title := r.FormValue("title")

		if title == "" {
			http.Error(w, "El título es obligatorio", http.StatusBadRequest)
			return
		}

		list.Title = title

		if err := db.Save(&list).Error; err != nil {
			http.Error(w, "Error al actualizar la lista", http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "/list/"+strconv.FormatUint(uint64(list.ID), 10), http.StatusSeeOther)
	}
}
