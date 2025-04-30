/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>

*/
package main

import (
	"expense-tracker/cmd"
	"expense-tracker/model"
	"expense-tracker/repository"
)

func main() {
	storage := repository.NewStorage[model.Expenses]("data.json")
	storage.Load(&model.Exps)

	cmd.Execute()
	storage.Save(model.Exps)
}
