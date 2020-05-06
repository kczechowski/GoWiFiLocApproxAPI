package handlers

import (
	"fmt"
	"github.com/kczechowski/GoWiFiLocApproxAPI/app/container"
	"net/http"
)

func GetIndex(container *container.Container, w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "<a href=\"https://github.com/kczechowski/GoWiFiLocApproxAPI\">https://github.com/kczechowski/GoWiFiLocApproxAPI</a>")
}
