package services

import (
	"fmt"
	"github.com/gcaldeira/go-setup/domains"
)

//===========================================================
type GreetingHappyService struct {

}

func (s GreetingHappyService) Hi(h string) string {

	return fmt.Sprintf("Hi %s, I'm so %s friendly! (•◡•) /", h, s.Humor())
}

func (s GreetingHappyService) Humor() string {
	return "happy"
}

func NewGreetingHappyService() domains.GreetingService {

	return GreetingHappyService{}
}