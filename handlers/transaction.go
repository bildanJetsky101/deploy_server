package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	dto "server_wb/dto/result"
	transactiondto "server_wb/dto/transaction"
	"server_wb/models"
	"server_wb/repositories"
	"strconv"

	// "math/rand"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
	"github.com/gorilla/mux"
	// "github.com/midtrans/midtrans-go"
	// "github.com/midtrans/midtrans-go/coreapi"
	// "github.com/midtrans/midtrans-go/snap"
)

type handlerTransaction struct {
	TransactionRepository repositories.TransactionRepository
}

func HandlerTransaction(TransactionRepository repositories.TransactionRepository) *handlerTransaction {
	return &handlerTransaction{TransactionRepository}
}

// var c = coreapi.Client{
// 	ServerKey: os.Getenv("SERVER_KEY"),
// 	ClientKey: os.Getenv("CLIENT_KEY"),
// }

func (h *handlerTransaction) FindTransactions(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	transactions, err := h.TransactionRepository.FindTransactions()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: transactions}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerTransaction) GetTransaction(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	// userInfo := r.Context().Value("userInfo").(jwt.MapClaims)
	// userId := int(userInfo["id"].(float64))

	cart, err := h.TransactionRepository.GetTransaction(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: convertTransactionResponse(cart)}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerTransaction) GetTransactions(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	//id, _ := strconv.Atoi(mux.Vars(r)["id"])
	userInfo := r.Context().Value("userInfo").(jwt.MapClaims)
	userId := int(userInfo["id"].(float64))

	transaction, err := h.TransactionRepository.GetTransactions(int(userId))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	fmt.Println(transaction)
	fmt.Println(userId)

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: transaction}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerTransaction) CreateTransaction(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	userInfo := r.Context().Value("userInfo").(jwt.MapClaims)
	userId := int(userInfo["id"].(float64))

	// cart, err := h.TransactionRepository.GetCartsTransaction(int(userId))
	// if err != nil {
	// 	w.WriteHeader(http.StatusBadRequest)
	// 	response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
	// 	json.NewEncoder(w).Encode(response)
	// 	return
	// }

	// fmt.Println(userId)
	// fmt.Println(cart)

	request := new(transactiondto.CreateTransaction)
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	// 	var TransIdIsMatch = false
	// 	var TransactionId int
	// 	for !TransIdIsMatch {
	// 		TransactionId = userId + rand.Intn(10000) - rand.Intn(100)
	// 		transactionData, _ := h.TransactionRepository.GetTransaction(TransactionId)
	// 		if transactionData.ID == 0 {
	// 			TransIdIsMatch = true
	// 		}
	// 	}

	// 	// 1. Initiate Snap client
	// var s = snap.Client{}
	// s.New(os.Getenv("SERVER_KEY"), midtrans.Sandbox)
	// // Use to midtrans.Production if you want Production Environment (accept real transaction).

	// // 2. Initiate Snap request param
	// req := &snap.Request{
	//   TransactionDetails: midtrans.TransactionDetails{
	//     OrderID:  strconv.Itoa(dataTransactions.ID),
	//     GrossAmt: int64(dataTransactions.Price),
	//   },
	//   CreditCard: &snap.CreditCardDetails{
	//     Secure: true,
	//   },
	//   CustomerDetail: &midtrans.CustomerDetails{
	//     FName: dataTransactions.Buyer.Name,
	//     Email: dataTransactions.Buyer.Email,
	//   },
	//   }

	// 	// 3. Execute request create Snap transaction to Midtrans Snap API
	// 	snapResp, _ := s.CreateTransaction(req)
	// 	w.WriteHeader(http.StatusOK)
	// 	response := dto.SuccessResult{Code: http.StatusOK, Data: snapResp}
	// 	json.NewEncoder(w).Encode(response)

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	// data form pattern submit to pattern entity db user
	transaction := models.Transaction{
		UserID: userId,
		Total:  request.Total,
		Status: request.Status,
		//Cart:   cart,
	}

	data, err := h.TransactionRepository.CreateTransaction(transaction)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: data}
	json.NewEncoder(w).Encode(response)
}

// func (h *handlerTransaction) Notification(w http.ResponseWriter, r *http.Request) {
// 	var notificationPayload map[string]interface{}

// 	err := json.NewDecoder(r.Body).Decode(&notificationPayload)
// 	if err != nil {
// 		w.WriteHeader(http.StatusBadRequest)
// 		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
// 		json.NewEncoder(w).Encode(response)
// 		return
// 	}

// 	transactionStatus := notificationPayload["transaction_status"].(string)
// 	fraudStatus := notificationPayload["fraud_status"].(string)
// 	orderId := notificationPayload["order_id"].(string)

// 	if transactionStatus == "capture" {
// 		if fraudStatus == "challenge" {
// 			// TODO set transaction status on your database to 'challenge'
// 			// e.g: 'Payment status challenged. Please take action on your Merchant Administration Portal
// 			h.TransactionRepository.UpdateTransaction("pending", orderId)
// 		} else if fraudStatus == "accept" {
// 			// TODO set transaction status on your database to 'success'
// 			h.TransactionRepository.UpdateTransaction("success", orderId)
// 		}
// 	} else if transactionStatus == "settlement" {
// 		// TODO set transaction status on your databaase to 'success'
// 		h.TransactionRepository.UpdateTransaction("success", orderId)
// 	} else if transactionStatus == "deny" {
// 		// TODO you can ignore 'deny', because most of the time it allows payment retries
// 		// and later can become success
// 		h.TransactionRepository.UpdateTransaction("failed", orderId)
// 	} else if transactionStatus == "cancel" || transactionStatus == "expire" {
// 		// TODO set transaction status on your databaase to 'failure'
// 		h.TransactionRepository.UpdateTransaction("failed", orderId)
// 	} else if transactionStatus == "pending" {
// 		// TODO set transaction status on your databaase to 'pending' / waiting payment
// 		h.TransactionRepository.UpdateTransaction("pending", orderId)
// 	}

// 	w.WriteHeader(http.StatusOK)
// }

func convertTransactionResponse(u models.Transaction) models.TransactionResponse {
	return models.TransactionResponse{
		ID:     u.ID,
		Status: u.Status,
		UserID: u.UserID,
		Total:  u.Total,
		User:   u.User,
	}
}
