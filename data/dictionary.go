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
/// Base actions
///////////////////////////////////////////////////////////////////////////////

func (dict *Dictionary) Fields() map[string]interface{} {
  return dict.fields
}

func (dict *Dictionary) Field(key string, item IItem) *Dictionary {
  dict.fields[key] = item
  return dict
}

func (dict *Dictionary) Dictionary(key string) *Dictionary {
  d := &Dictionary{Item{parent: dict}}
  dict.fields[key] = d
  return d
}

func (dict *Dictionary) List(key string) *List {
  list := &List{Item{parent: dict}}
  dict.fields[key] = list
  return list
}

func (dict *Dictionary) Parent() *Item {
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
