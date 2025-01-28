package repositories

import (
	"WhatDownloadGo/models"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type LocationRepository interface {
	GetIP() (string, error)
	GetLocation(ip string) (*models.Location, error)
}

type HTTPLocationRepository struct{}

func (r *HTTPLocationRepository) GetIP() (string, error) {
	url := "https://api64.ipify.org?format=json"
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("ошибка: получен код состояния %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var result models.Response
	err = json.Unmarshal(body, &result)
	if err != nil {
		return "", err
	}
	return result.IP, nil
}

func (r *HTTPLocationRepository) GetLocation(ip string) (*models.Location, error) {
	url := fmt.Sprintf("http://ip-api.com/json/%s", ip)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("код ошибки: %d", resp.StatusCode)
	}

	var location models.Location
	err = json.NewDecoder(resp.Body).Decode(&location)
	if err != nil {
		return nil, err
	}

	if location.Status != "success" {
		return nil, fmt.Errorf("ошибка API: %s", location.Message)
	}

	return &location, nil
}
