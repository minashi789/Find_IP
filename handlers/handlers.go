package handlers

import (
	"WhatDownloadGo/services"
	"log"
)

type Handlers struct {
	service services.LocationService
}

func NewHandlers(service services.LocationService) *Handlers {
	return &Handlers{
		service: service,
	}
}

func (h *Handlers) Run() {
	filename := "local_ip.txt"
	err := h.service.GetIPAndSaveToFile(filename)
	if err != nil {
		log.Fatal(err)
	}
}
