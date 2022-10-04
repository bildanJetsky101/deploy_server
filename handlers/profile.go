package handlers

import (
	"encoding/json"
	"net/http"
	"os"
	profiledto "server_wb/dto/profile"
	dto "server_wb/dto/result"
	"server_wb/models"
	"server_wb/repositories"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
)

type handlerProfile struct {
	ProfileRepository repositories.ProfileRepository
}

var path_file = os.Getenv("PATH_FILE")

func HandlerProfile(ProfileRepository repositories.ProfileRepository) *handlerProfile {
	return &handlerProfile{ProfileRepository}
}

func (h *handlerProfile) FindProfiles(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	profiles, err := h.ProfileRepository.FindProfiles()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
	}

	for i, p := range profiles {
		profiles[i].Image = path_file + p.Image
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: profiles}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerProfile) GetProfile(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	userInfo := r.Context().Value("userInfo").(jwt.MapClaims)
	userId := int(userInfo["id"].(float64))

	profile, err := h.ProfileRepository.GetProfile(int(userId))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	profile.Image = path_file + profile.Image

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: convertProfileResponse(profile)}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerProfile) CreateProfile(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	phone, _ := strconv.Atoi(r.FormValue("phone"))
	dataContex := r.Context().Value("dataFile")
	filename := dataContex.(string)

	userInfo := r.Context().Value("userInfo").(jwt.MapClaims)
	userId := int(userInfo["id"].(float64))

	request := profiledto.CreateProfile{
		Address: r.FormValue("address"),
		Gender:  r.FormValue("gender"),
		Phone:   phone,
		Image:   filename,
		UserID:  userId,
	}

	// if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
	// 	w.WriteHeader(http.StatusBadRequest)
	// 	response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
	// 	json.NewEncoder(w).Encode(response)
	// 	return
	// }

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	profile := models.Profile{
		Address: request.Address,
		Gender:  request.Gender,
		Phone:   request.Phone,
		Image:   request.Image,
		UserID:  request.UserID,
	}

	data, err := h.ProfileRepository.CreateProfile(profile)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: convertProfileResponse(data)}
	json.NewEncoder(w).Encode(response)
}

func convertProfileResponse(u models.Profile) profiledto.ProfileResponse {
	return profiledto.ProfileResponse{
		ID:      u.ID,
		Address: u.Address,
		Gender:  u.Gender,
		Phone:   u.Phone,
		Image:   u.Image,
		UserID:  u.UserID,
	}
}
