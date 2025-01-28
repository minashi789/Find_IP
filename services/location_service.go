package services

import (
	"WhatDownloadGo/repositories"
	"fmt"
	"os"
)

type LocationService interface {
	GetIPAndSaveToFile(filename string) error
}

type LocationServiceImpl struct {
	repository repositories.LocationRepository
}

func NewLocationService(repository repositories.LocationRepository) LocationService {
	return &LocationServiceImpl{
		repository: repository,
	}
}

func (s *LocationServiceImpl) GetIPAndSaveToFile(filename string) error {
	ip, err := s.repository.GetIP()
	if err != nil {
		return err
	}

	fmt.Println("Внешний IP:", ip)

	location, err := s.repository.GetLocation(ip)
	if err != nil {
		return fmt.Errorf("ошибка определения местоположения: %w", err)
	}

	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = fmt.Fprintf(file, "IP: %s, City: %s, Region: %s, Country: %s, Lat: %f, Lon: %f\n",
		ip, location.City, location.Region, location.Country, location.Lat, location.Lon)
	if err != nil {
		return err
	}

	fmt.Println("Информация записана в файл:", filename)
	return nil
}
