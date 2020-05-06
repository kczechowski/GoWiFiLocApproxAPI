package handlers

import (
	"fmt"
	"github.com/kczechowski/GoWiFiLocApproxAPI/app/container"
	"net/http"
)

func PostNetwork(container *container.Container, w http.ResponseWriter, r *http.Request) {
	fmt.Println("test")
}
