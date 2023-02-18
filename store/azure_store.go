package store

import (
	"github.com/spctr-cc/backend-bugtrack/model"
)

type AzureStore struct {
}

func NewAzureStore() *AzureStore {
	return &AzureStore{}
}

func (as *AzureStore) Create(a *model.Azure) (err error) {
	return nil
}
