package main

import (
	"github.com/go-playground/validator/v10"
	"log"
)

var valiate *validator.Validate

func init() {
	valiate = validator.New()
	valiate.SetTagName("v")
}

func outRes(tag string, err *error) {
	log.Println("------------------start" + tag + "---------------")
	log.Println(*err)
	log.Println("------------------end" + tag + "---------------")
	err = nil
}