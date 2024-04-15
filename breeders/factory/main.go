package main

import (
	"factory/products"
	"fmt"
	"log"
)

// Animal is type for abstract factory
type Animal interface {
	Says()
	LikesWater() bool
}

// create a concrete factory for dogs
type Dog struct {
}

// implement the abstract factory for dogs
func (d *Dog) Says() {
	fmt.Println("woof")
}
func (d *Dog) LikesWater() bool {
	return true
}

// create a concrete factory for cats
type Cat struct {
}

// implement the abstract factory for cats
func (c *Cat) Says() {
	fmt.Println("miau")
}
func (c *Cat) LikesWater() bool {
	return false
}

//

type AnimalFactory interface {
	New() Animal
}

type DogFactory struct {
}

func (df *DogFactory) New() Animal {
	return &Dog{}
}

type CatFactory struct {
}

func (cf *CatFactory) New() Animal {
	return &Cat{}
}
func main() {
	var factory products.Product
	pr := factory.New()
	pr.Name = "1"

	log.Println("my product was created at ", pr.CreatedAt.UTC())
	// --
	df := DogFactory{}
	cf := CatFactory{}
	d := df.New()
	c := cf.New()
	d.Says()
	c.Says()

}
