// @project datagen
// @copyright Dmitry Ponomarev <demdxx@gmail.com> 2014
//
// This work is licensed under the Creative Commons Attribution 4.0 International License.
// To view a copy of this license, visit http://creativecommons.org/licenses/by/4.0/.

package data

import (
  "errors"
  "fmt"
)

type Dictionary struct {
  Item
  fields map[string]interface{}
}

func MakeDictionary() *Dictionary {
  m := &Dictionary{}
  m.fields = make(map[string]interface{})
  return m
}

///////////////////////////////////////////////////////////////////////////////
/// Getters/Setters
///////////////////////////////////////////////////////////////////////////////

func (dict *Dictionary) Fields() map[string]interface{} {
  return dict.fields
}

func (dict *Dictionary) Field(key string, item IItem) *Dictionary {
  dict.fields[key] = item
  return dict
}

func (dict *Dictionary) Dictionary(key string) *Dictionary {
  d := MakeDictionary()
  d.Item.parent = (IItem)(dict)
  dict.fields[key] = d
  return d
}

func (dict *Dictionary) List(key string) *List {
  list := MakeList()
  list.Item.parent = (IItem)(dict)
  dict.fields[key] = list
  return list
}

func (dict *Dictionary) Parent() IItem {
  return dict.Item.Parent()
}

func (dict *Dictionary) Generate() interface{} {
  m := make(map[string]interface{})
  for k, v := range dict.fields {
    if ie, ok := v.(IItem); ok {
      m[k] = ie.Generate()
    } else {
      m[k] = errors.New(fmt.Sprint("Invalid field type %s", k))
    }
  }
  return m
}

///////////////////////////////////////////////////////////////////////////////
/// Generators
///////////////////////////////////////////////////////////////////////////////

// An Infinity generator
func (dict *Dictionary) Generator() *Generator {
  return MakeGenerator((IItem)(dict))
}

func (dict *Dictionary) Pages(pageSize, elements uint) *PageGenerator {
  return MakePageGenerator((IItem)(dict), pageSize, elements)
}

func (dict *Dictionary) Set(count uint) chan interface{} {
  return dict.Generator().Set(count)
}

func (dict *Dictionary) SetRandom(min, max uint) chan interface{} {
  return dict.Generator().SetRandom(min, max)
}
