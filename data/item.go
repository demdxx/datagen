package data

import (
  "errors"
)

type IItem interface {
  Parent() *Item
  Generate() interface{}
}

type Item struct {
  parent *Item
}

func (item *Item) Parent() *Item {
  return item.parent
}

func (item *Item) Generate() interface{} {
  return errors.New("Empty field generator")
}
