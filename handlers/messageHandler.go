package message

import (
	"encoding/json"
	"fmt"
	"net/http"

	domain "github.com/eduardomassami/message-receiver/domain"
)

type MessageHandler struct {
}

func NewMessageHandler() *MessageHandler {
	return &MessageHandler{}
}

func (hdl *MessageHandler) PostTelegram(w http.ResponseWriter, r *http.Request) {
	fmt.Println(w)
	fmt.Println(r.Body)
	var request domain.TelegramPayload
	err := json.NewDecoder(r.Body).Decode(&request)

	if err != nil {
		writeResponse(w, http.StatusBadRequest, "error when decoding request")
		return
	} else {
		fmt.Println(request)
		// 	e := domain.ValidateStruct(request)
		// 	if e != nil {
		// 		writeResponse(w, http.StatusInternalServerError, "invalid body")
		// 		return
		// 	}
		// 	appError := hdl.marketsService.Post(request)
		// 	if appError != nil {
		// 		writeResponse(w, http.StatusInternalServerError, appError.Error())
		// 		return
		// 	}
		writeResponse(w, http.StatusAccepted, "received")
	}
}

func (hdl *MessageHandler) PostWhatsapp(w http.ResponseWriter, r *http.Request) {
	fmt.Println(w)
	fmt.Println(r.Body)
	var request domain.WhatsappPayload
	err := json.NewDecoder(r.Body).Decode(&request)

	if err != nil {
		writeResponse(w, http.StatusBadRequest, "error when decoding request")
		return
	} else {
		fmt.Println(request)
		// 	e := domain.ValidateStruct(request)
		// 	if e != nil {
		// 		writeResponse(w, http.StatusInternalServerError, "invalid body")
		// 		return
		// 	}
		// 	appError := hdl.marketsService.Post(request)
		// 	if appError != nil {
		// 		writeResponse(w, http.StatusInternalServerError, appError.Error())
		// 		return
		// 	}
		writeResponse(w, http.StatusAccepted, "received")
	}
}

func writeResponse(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		panic(err)
	}
}
