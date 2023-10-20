package main

import (
	"fmt"

	"github.com/golang/mock/gomock"
)

func main() {
	ctrl := gomock.NewController(nil)
	_ = ctrl
	fmt.Println("HW")
}
