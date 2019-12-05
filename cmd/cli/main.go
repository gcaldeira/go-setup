package main

import (
	"bufio"
	"fmt"
	"github.com/gcaldeira/go-setup/domains"
	"github.com/gcaldeira/go-setup/domains/services"
	"os"
	"strings"
)

var greetingService domains.GreetingService

func main()  {

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Put your name:")
	name, _ := reader.ReadString('\n')
	name = strings.Replace(name, "\n", "", -1)

	greetingService = services.NewGreetingHappyService()

	g := greetingService.Hi(name)

	fmt.Println(g)
}