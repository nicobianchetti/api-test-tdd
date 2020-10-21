package main

import (
	"database/sql"
	"errors"
)

type product struct {
	ID    int `json:"id"`
	Name  int `json:"name"`
	Price int `json:"price"`
}

func (p *product) getProduct(db *sql.DB) error {
	return errors.New("Not implemented")
}

func (p *product) updateProduct(db *sql.DB) error {
	return errors.New("Not implemented")
}

func (p *product) deleteProduct(db *sql.DB) error {
	return errors.New("Not implemented")
}

func (p *product) createProduct(db *sql.DB) error {
	return errors.New("Not implemented")
}

func getProducts(db *sql.DB, start, count int) ([]product, error) {
	return nil, errors.New("Not implemented")
}
