package service

import "URLshortener/repository"

const BASE_URL = "http://localhost:8080/"

type URLService struct {
	repository    repository.Repository
	codeGenerator CodeGenerator
}

func NewUrlService(repository repository.Repository, generator *CodeGenerator) *URLService {
	return &URLService{repository, *generator}
}

func (u *URLService) GetOriginalUrl(code string) (string, bool) {
	var originalUrl string = u.repository.GetUrl(code)
	return originalUrl, originalUrl != ""
}

func (u *URLService) AddNewUrlPair(originalUrl string) string {
	urlExists, code := u.repository.ContainsUrl(originalUrl)
	if urlExists {
		return BASE_URL + code
	} else {
		code = u.createNewCode()
		u.repository.InsertUrl(code, originalUrl)
		return BASE_URL + code
	}
}

func (u *URLService) createNewCode() string {
	var code string
	for {
		code = u.codeGenerator.GenerateCode()
		if u.codeExists(code) {
			continue
		} else {
			break
		}
	}
	return code
}

func (u *URLService) codeExists(code string) bool {
	return u.repository.GetUrl(code) != ""
}
