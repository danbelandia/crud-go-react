package handlers

import (
	"crud-stefanini/internal/core/domain"
	"crud-stefanini/internal/core/ports"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

type UserHandler struct {
	repo ports.UserRepository
}

func NewUserHandler(repo ports.UserRepository) *UserHandler {
	return &UserHandler{repo: repo}
}

func (h *UserHandler) HandleUsuarios(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.getUsers(w, r)
	case http.MethodPost:
		h.createUsers(w, r)
	case http.MethodPut:
		h.updateUsers(w, r)
	case http.MethodDelete:
		h.deleteUsers(w, r)
	default:
		http.Error(w, "Error de petición", http.StatusMethodNotAllowed)

	}
}

// Obtener usuario
func (h *UserHandler) getUsers(w http.ResponseWriter, r *http.Request) {
	path := strings.Trim(r.URL.Path, "/") // Quita "/" al inicio y al final
	parts := strings.Split(path, "/")     // Divide por partes
	//Si viene solo usuarios, lista todo
	if len(parts) == 1 && parts[0] == "usuarios" {
		limit := 5
		page := 1

		if lStr := r.URL.Query().Get("limit"); lStr != "" {
			if l, err := strconv.Atoi(lStr); err == nil && l > 0 {
				limit = l
			}
		}

		if pStr := r.URL.Query().Get("page"); pStr != "" {
			if p, err := strconv.Atoi(pStr); err == nil && p > 0 {
				page = p
			}
		}

		search := r.URL.Query().Get("search")
		offset := (page - 1) * limit

		users, err := h.repo.GetAll(limit, offset, search)
		if err != nil {
			http.Error(w, "Error al obtener los usuarios", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(users)
		return
	}
	//Si la division tiene 2 partes, convierte la 2da parte en numero y lo busca en BD
	if len(parts) == 2 && parts[0] == "usuarios" {
		id, err := strconv.Atoi(parts[1])
		if err != nil {
			http.Error(w, "el id debe ser un numero", http.StatusBadRequest)
			return
		}
		//Busca en la BD el ID
		user, err := h.repo.GetByID(id)
		if err != nil {
			http.Error(w, "Usuario no encontrado", http.StatusNotFound)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(user)
		return
	}

	http.Error(w, "Ruta no encontrada", http.StatusNotFound)

}

// Crear usuario
func (h *UserHandler) createUsers(w http.ResponseWriter, r *http.Request) {
	var u domain.User
	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		fmt.Println("ERROR DE SQLITE:", err)
		http.Error(w, "JSON invalido", http.StatusBadRequest)
		return
	}

	if err := h.repo.Create(&u); err != nil {
		http.Error(w, "Error al guardar en base de datos", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(u)
}

// Actualizar usuario
func (h *UserHandler) updateUsers(w http.ResponseWriter, r *http.Request) {
	path := strings.Trim(r.URL.Path, "/")
	parts := strings.Split(path, "/")

	if len(parts) != 2 || parts[0] != "usuarios" {
		http.Error(w, "Falta el id en la url", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(parts[1])
	if err != nil {
		http.Error(w, "El id debe ser un numero", http.StatusBadRequest)
		return
	}
	var u domain.User
	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		http.Error(w, "JSON invalido", http.StatusBadRequest)
		return
	}

	u.ID = id

	if err := h.repo.Update(&u); err != nil {
		if err.Error() == "No existe el usuario para actualizar" {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		http.Error(w, "Error interno", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(u)

}

// Eliminar usuario
func (h *UserHandler) deleteUsers(w http.ResponseWriter, r *http.Request) {
	path := strings.Trim(r.URL.Path, "/")
	parts := strings.Split(path, "/")

	if len(parts) != 2 || parts[0] != "usuarios" {
		http.Error(w, "Falta el id en la url", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(parts[1])
	if err != nil {
		http.Error(w, "El id debe ser un numero", http.StatusBadRequest)
		return
	}

	if err := h.repo.Delete(id); err != nil {
		if err.Error() == "No se encontro ningun usuario" {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		http.Error(w, "Error del servidor", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	message := map[string]string{"mensaje: ": "Usuario eliminado satisfactoriamente"}
	json.NewEncoder(w).Encode(message)

}
