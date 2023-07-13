package main

import (
	"KidneySmartBackend/internal/config"
	"fmt"
)

//init router chi

func main() {

	//  init config
	cfg := config.MustLoad()
	fmt.Println(cfg)

}
