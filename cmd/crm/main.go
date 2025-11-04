package main

import (
	"github.com/yumeei/go-tp/internal/app"
	"github.com/yumeei/go-tp/internal/storage"
)

func main() {
	var store = storage.NewMemoryStorage()
	app.Run(*store)
}