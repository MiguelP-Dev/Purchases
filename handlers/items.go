package handlers

import (
	"net/http"
	"strconv"
	"text/template"

	"github.com/MiguelP-Dev/Purchases/models"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func ToggleItemHandler(w http.ResponseWriter, r *http.Request) {
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

	// Solo permitir desmarcar (cambiar de comprado a no comprado)
	// Para marcar como comprado se debe usar MarkItemAsPurchasedHandler
	if item.IsPurchased {
		item.IsPurchased = false
		item.Price = 0.0
		item.Place = "supermercado"
	} else {
		// Si no está comprado, redirigir al formulario de precio
		http.Redirect(w, r, "/items/"+strconv.Itoa(itemID)+"/mark-purchased", http.StatusSeeOther)
		return
	}

	if err := db.Save(&item).Error; err != nil {
		http.Error(w, "Error al actualizar el item", http.StatusInternalServerError)
		return
	}

	// Redirigir de vuelta a la lista después de desmarcar
	http.Redirect(w, r, "/list/"+strconv.FormatUint(uint64(item.ShoppingListID), 10), http.StatusSeeOther)
}

// MarkItemAsPurchasedHandler marca un item como comprado con precio y lugar
func MarkItemAsPurchasedHandler(w http.ResponseWriter, r *http.Request) {
	db := r.Context().Value("db").(*gorm.DB)

	if r.Method == "GET" {
		vars := mux.Vars(r)
		itemID := vars["id"]

		// Mostrar formulario para ingresar precio y lugar
		tmpl, err := template.ParseFiles("templates/mark_purchased.html")
		if err != nil {
			http.Error(w, "Error al cargar la plantilla", http.StatusInternalServerError)
			return
		}

		tmpl.Execute(w, map[string]string{"item_id": itemID})
		return
	}

	if r.Method == "POST" {
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

		r.ParseForm()
		priceStr := r.FormValue("price")
		place := r.FormValue("place")

		if priceStr == "" {
			http.Error(w, "El precio es obligatorio", http.StatusBadRequest)
			return
		}

		price, err := strconv.ParseFloat(priceStr, 64)
		if err != nil {
			http.Error(w, "Precio inválido", http.StatusBadRequest)
			return
		}

		item.Price = price
		if place != "" {
			item.Place = place
		}
		item.IsPurchased = true

		if err := db.Save(&item).Error; err != nil {
			http.Error(w, "Error al actualizar el item", http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "/list/"+strconv.FormatUint(uint64(item.ShoppingListID), 10), http.StatusSeeOther)
	}
}

// EditItemHandler permite editar un item comprado (precio y lugar)
func EditItemHandler(w http.ResponseWriter, r *http.Request) {
	db := r.Context().Value("db").(*gorm.DB)

	if r.Method == "GET" {
		vars := mux.Vars(r)
		itemID := vars["id"]

		// Obtener el item para mostrar los valores actuales
		itemIDInt, err := strconv.Atoi(itemID)
		if err != nil {
			http.Error(w, "ID de item inválido", http.StatusBadRequest)
			return
		}

		var item models.ListItem
		if err := db.First(&item, itemIDInt).Error; err != nil {
			http.Error(w, "Item no encontrado", http.StatusNotFound)
			return
		}

		// Mostrar formulario para editar precio y lugar
		tmpl, err := template.ParseFiles("templates/edit_item.html")
		if err != nil {
			http.Error(w, "Error al cargar la plantilla", http.StatusInternalServerError)
			return
		}

		tmpl.Execute(w, item)
		return
	}

	if r.Method == "POST" {
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

		r.ParseForm()
		priceStr := r.FormValue("price")
		place := r.FormValue("place")

		if priceStr == "" {
			http.Error(w, "El precio es obligatorio", http.StatusBadRequest)
			return
		}

		price, err := strconv.ParseFloat(priceStr, 64)
		if err != nil {
			http.Error(w, "Precio inválido", http.StatusBadRequest)
			return
		}

		item.Price = price
		if place != "" {
			item.Place = place
		}

		if err := db.Save(&item).Error; err != nil {
			http.Error(w, "Error al actualizar el item", http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "/list/"+strconv.FormatUint(uint64(item.ShoppingListID), 10), http.StatusSeeOther)
	}
}
