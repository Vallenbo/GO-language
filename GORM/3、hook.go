package main

import (
	"fmt"
	"gorm.io/gorm"
)

func (t *Teacher) BeforeSave(db *gorm.DB) error {
	fmt.Println("hook BeforeSave")
	return nil
}

func (t *Teacher) AfterSave(db *gorm.DB) error {
	fmt.Println("hook AfterSave")
	return nil
}

func (t *Teacher) BeforeCreate(db *gorm.DB) error {
	fmt.Println("hook BeforeCreate")
	return nil
}

func (t *Teacher) AfterCreate(db *gorm.DB) error {
	fmt.Println("hook AfterCreate")
	return nil
}

func (t *Teacher) BeforeUpate(db *gorm.DB) error {
	fmt.Println("hook BeforeUpate")
	return nil
}
func (t *Teacher) AfterUpate(db *gorm.DB) error {
	fmt.Println("hook AfterUpate")
	return nil
}

func (t *Teacher) BeforeDelete(db *gorm.DB) error {
	fmt.Println("hook AfterUpate")
	return nil
}

func (t *Teacher) AfterDelete(db *gorm.DB) error {
	fmt.Println("hook AfterUpate")
	return nil
}

func main() {

}