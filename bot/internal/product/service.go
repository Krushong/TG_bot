package product

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"strings"
)

type Service struct {
}

func NewService() *Service {
	return &Service{}
}

func (s *Service) List() []Product {
	return allProduct
}

func (s *Service) Get(idx int) (*Product, error) {
	return &allProduct[idx], nil
}

func (s *Service) Delete(idx int) {
	allProduct = append(allProduct[:idx], allProduct[idx+1:]...)
}

func (s *Service) Add(newData string) {
	parsData := strings.TrimSpace(strings.Replace(newData, "/add", "", 1))
	allProduct = append(allProduct, Product{Title: parsData})
}

func (s *Service) Edit(idx int) {
	numberEditElement = idx
}

func (s *Service) Xo(idx int) {
	//numberEditElement = idx
}

func (s *Service) Edit_value(inputMessage *tgbotapi.Message) {
	allProduct[numberEditElement] = Product{
		Title: inputMessage.Text,
	}
}
