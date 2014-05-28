// @project datagen
// @copyright Dmitry Ponomarev <demdxx@gmail.com> 2014
//
// This work is licensed under the Creative Commons Attribution 4.0 International License.
// To view a copy of this license, visit http://creativecommons.org/licenses/by/4.0/.

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
