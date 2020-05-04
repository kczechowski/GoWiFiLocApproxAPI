package main

import "github.com/kczechowski/GoWiFiLocApproxAPI/app"

func main() {
	a := &app.App{}
	a.Init()
	a.Run(":8080")
}
