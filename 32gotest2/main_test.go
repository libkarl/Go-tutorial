package main

import (
	
	"testing"

	"github.com/stretchr/testify/assert"
)
func TestFilterUnique(t *testing.T) {
	input := []Developer { 
		
		{Name: "Elliot"}, 
		{Name: "David"}, 
		{Name: "Alexander"},
		{Name: "Eva"}, 
		{Name: "Alan"}, 
	}

	expected := []string {
		 "Elliot",
		 "David",
		 "Alexander",
		 "Eva",
		 "Alan",
	}
	
	assert.Equal(t, expected, FilterUnique(input))
	// Equal funkce z testify, která otestuje jestli se dvě věci rovnají
	// Equal(t, 123,123) -> test by prošel

}

func TestNegativeFilterUnique(t *testing.T) {
	input := []Developer { 
		
		{Name: "Elliot"}, 
		{Name: "David"}, 
		{Name: "Alexander"},
		{Name: "Eva"}, 
		{Name: "Alan"}, 
	}

	expected := []string {
		 "Elliot",
		 "Eva",
		 "Alan",
	}
	// kontrolujeme, že se input a expected nerovnají 
	assert.NotEqual(t, expected, FilterUnique(input))
}