package main

import (
	"errors"
	"fmt"

	"github.com/RangelReale/osin"
	"github.com/aoisensi/gonvey/controller"
	"github.com/jinzhu/gorm"
)

var OAuth *osin.Server

func InitAuth() {
	config := osin.NewServerConfig()
	storage := NewOAuthStorage(DB)
	OAuth = osin.NewServer(config, storage)
	controller.OAuth = OAuth
}

type OAuthStorage struct {
	osin.Storage
	db        gorm.DB
	authorize map[string]*osin.AuthorizeData
	access    map[string]*osin.AccessData
	refresh   map[string]string
}

func NewOAuthStorage(db gorm.DB) osin.Storage {
	s := &OAuthStorage{
		db:        db,
		authorize: make(map[string]*osin.AuthorizeData),
		access:    make(map[string]*osin.AccessData),
		refresh:   make(map[string]string),
	}
	return s
}

func (s *OAuthStorage) Clone() osin.Storage {
	return s
}

func (s *OAuthStorage) Close() {

}

func (s *OAuthStorage) GetClient(id string) (osin.Client, error) {
	if id == "gonvey" {
		return &osin.DefaultClient{
			Id:          "gonvey",
			Secret:      "aabbccdd",
			RedirectUri: "/auth",
		}, nil
	}
	return nil, errors.New("Client not found.")
} //fixme

func (s *OAuthStorage) SetClient(id string, client osin.Client) error {
	return errors.New("Not yet supported.")
}

func (s *OAuthStorage) SaveAuthorize(data *osin.AuthorizeData) error {
	fmt.Printf("SaveAuthorize: %s\n", data.Code)
	s.authorize[data.Code] = data
	return nil
}

func (s *OAuthStorage) LoadAuthorize(code string) (*osin.AuthorizeData, error) {
	if d, ok := s.authorize[code]; ok {
		return d, nil
	}
	return nil, errors.New("Authorize not found")
}

func (s *OAuthStorage) RemoveAuthorize(code string) error {
	delete(s.authorize, code)
	return nil
}

func (s *OAuthStorage) SaveAccess(data *osin.AccessData) error {
	s.access[data.AccessToken] = data
	if data.RefreshToken != "" {
		s.refresh[data.RefreshToken] = data.AccessToken
	}
	return nil
}

func (s *OAuthStorage) LoadAccess(code string) (*osin.AccessData, error) {
	if d, ok := s.access[code]; ok {
		return d, nil
	}
	return nil, errors.New("Access not found")
}

func (s *OAuthStorage) RemoveAccess(code string) error {
	delete(s.access, code)
	return nil
}

func (s *OAuthStorage) LoadRefresh(code string) (*osin.AccessData, error) {
	if d, ok := s.refresh[code]; ok {
		return s.LoadAccess(d)
	}
	return nil, errors.New("Refresh not found")
}

func (s *OAuthStorage) RemoveRefresh(code string) error {
	delete(s.refresh, code)
	return nil
}
