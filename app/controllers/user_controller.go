package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"go-crud/app/models"
	"go-crud/config"
	"go-crud/dto"
	"go-crud/libs"

	"github.com/go-chi/chi/v5"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// GET /users
func IndexUser(w http.ResponseWriter, r *http.Request) {
	var users []models.User
	config.DB.Find(&users)

	json.NewEncoder(w).Encode(users)
}

// GET /users/{id}
func ShowUser(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	var user models.User
	if err := config.DB.First(&user, id).Error; err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(user)
}

// PUT /users/{id}
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	var user models.User
	if err := config.DB.First(&user, id).Error; err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	var input struct {
		Name  string
		Email string
	}
	json.NewDecoder(r.Body).Decode(&input)

	user.Name = input.Name
	user.Email = input.Email

	config.DB.Save(&user)
	json.NewEncoder(w).Encode(user)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))

	config.DB.Delete(&models.User{}, id)
	w.WriteHeader(http.StatusNoContent)
}

func PaginatedUsers(w http.ResponseWriter, r *http.Request) {
	libs.TryCatch(w, func() {

		page, err := strconv.Atoi(r.URL.Query().Get("page"))
		if err != nil {
			page = 1
		}

		limit, err := strconv.Atoi(r.URL.Query().Get("limit"))
		if err != nil {
			limit = 10
		}

		if page <= 0 {
			page = 1
		}
		if limit <= 0 {
			limit = 10
		}

		offset := (page - 1) * limit

		var users []models.User
		var total int64

		if err := config.DB.Model(&models.User{}).Count(&total).Error; err != nil {
			panic(err)
		}

		if err := config.DB.
			Limit(limit).
			Offset(offset).
			Find(&users).Error; err != nil {
			panic(err)
		}

		totalPages := (total + int64(limit) - 1) / int64(limit)

		libs.ResponseAPI(
			w,
			users,
			map[string]any{
				"pagination": map[string]any{
					"current_page": page,
					"per_page":     limit,
					"total":        total,
					"total_page":   totalPages,
				},
			},
			"Data has been retrieved",
			0,
		)
	})
}

func CreateUserBaru() http.HandlerFunc {
	return libs.WithTransaction(func(
		w http.ResponseWriter,
		r *http.Request,
		tx *gorm.DB,
	) {

		var req dto.CreateUserRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			panic(err)
		}

		hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), 10)
		if err != nil {
			panic(err)
		}

		user := models.User{
			Name:     req.Name,
			Email:    req.Email,
			Password: string(hash),
		}

		if err := tx.Create(&user).Error; err != nil {
			panic(err)
		}

		libs.ResponseAPI(
			w,
			user,
			nil,
			"User successfully created",
			0,
		)
	})
}
