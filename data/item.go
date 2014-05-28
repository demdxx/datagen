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
  Parent() IItem
  Generate() interface{}
  Generator() *Generator
  Pages(pagecount, pagesize uint) *PageGenerator
  Set(elements uint) chan interface{}
  SetRandom(min, max uint) chan interface{}
}

type Item struct {
  parent IItem
}

///////////////////////////////////////////////////////////////////////////////
/// Getters/Setters
///////////////////////////////////////////////////////////////////////////////

func (item *Item) Parent() IItem {
  return item.parent
}

func (item *Item) Generate() interface{} {
  return errors.New("Empty field generator")
}

///////////////////////////////////////////////////////////////////////////////
/// Generators
///////////////////////////////////////////////////////////////////////////////

// An Infinity generator
func (item *Item) Generator() *Generator {
  return MakeGenerator(item)
}

func (item *Item) Pages(pageSize, elements uint) *PageGenerator {
  return MakePageGenerator(item, pageSize, elements)
}

func (item *Item) Set(count uint) chan interface{} {
  return item.Generator().Set(count)
}

func (item *Item) SetRandom(min, max uint) chan interface{} {
  return item.Generator().SetRandom(min, max)
}
