package services

import (
	k "github.com/katsew/kuji"
)

type kujiService struct {
	k.Kuji
}

var kujiHandler *kujiService
