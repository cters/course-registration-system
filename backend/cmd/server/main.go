package main

import (
	"github.com/QuanCters/backend/internal/initialize"
)

func main() {
	r := initialize.Run()
	r.Run(":8000")
}