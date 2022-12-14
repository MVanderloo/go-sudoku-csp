package csp

import (
	"Sudoku-CSP/util"
)

type Domain []int

/**
 * Variable constructor
 **/
func NewDomain(vals []int) Domain {
	if vals == nil {
		panic("Domain must have be non-nil")
	}
	return Domain(vals)
}

/**
 * Removes a value from the domain of the variable
 **/
func (d Domain) Remove(value int) Domain {
	d = util.RemoveOrdered(d, value)
	return d
}

/**
 * Adds a value from the domain of the variable
 **/
func (d Domain) Add(value int) Domain {
	d = append(d, value)
	return d
}

/**
 * Checks if a variable can take on a value
 **/
func (d Domain) Contains(value int) bool {
	return util.Contains(d, value)
}

/**
 * Checks if a variable's domain contains a value other than value
 **/
func (d Domain) ContainsOtherThan(value int) bool {
	for _, domain_value := range d {
		if domain_value != value {
			return true
		}
	}
	return false
}
