package main

import (
	"fmt"
	"log"
	"net/http"

	message "github.com/eduardomassami/message-receiver/handlers"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	messageHandler := message.NewMessageHandler()

	router.HandleFunc("/message/telegram", messageHandler.PostTelegram).
		Methods(http.MethodPost).
		Name("TelegramMessages")

	router.HandleFunc("/message/whatsapp", messageHandler.PostWhatsapp).
		Methods(http.MethodPost).
		Name("WhatsappMessages")

	// starting server
	// address := os.Getenv("SERVER_ADDRESS")
	// port := os.Getenv("SERVER_PORT")
	address := "127.0.0.1"
	port := "3001"
	log.Print(fmt.Sprintf("Starting server on %s:%s ...", address, port))
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", address, port), router))
}
