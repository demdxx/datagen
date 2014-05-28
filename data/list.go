// @project datagen
// @copyright Dmitry Ponomarev <demdxx@gmail.com> 2014
//
// This work is licensed under the Creative Commons Attribution 4.0 International License.
// To view a copy of this license, visit http://creativecommons.org/licenses/by/4.0/.

package data

type List struct {
  Item
  items []interface{}
}

func MakeList() *List {
  l := &List{}
  l.items = make([]interface{}, 0, 0)
  return l
}

///////////////////////////////////////////////////////////////////////////////
/// Base actions
///////////////////////////////////////////////////////////////////////////////

func (list *List) Fields() []interface{} {
  return list.items
}

func (list *List) Field(item interface{}) *List {
  list.items = append(list.items, item)
  return list
}

func (list *List) Dictionary() *Dictionary {
  dict := MakeDictionary()
  dict.Item.parent = (IItem)(list)
  list.items = append(list.items, dict)
  return dict
}

func (list *List) List() *List {
  l := &List{Item{parent: list}}
  l.Item.parent = (IItem)(list)
  list.items = append(list.items, l)
  return l
}

func (list *List) Parent() *Item {
  return list.Item.Parent()
}

func (list *List) Generate() interface{} {
  l := make([]interface{}, 0, len(list.items))
  for _, v := range list.items {
    if ie, ok := v.(IItem); ok {
      l = append(l, ie.Generate())
    } else {
      l = append(l, errors.New(fmt.Sprint("Invalid field type %s", k)))
    }
  }
  return l
}

///////////////////////////////////////////////////////////////////////////////
/// Generators
///////////////////////////////////////////////////////////////////////////////

// An Infinity generator
func (list *List) Generator() *Generator {
  return MakeGenerator((IItem)(List))
}

func (list *List) Pages(pageSize, elements uint) *PageGenerator {
  return MakePageGenerator((IItem)(List), pageSize, elements)
}

func (list *List) Set(count uint) chan interface{} {
  return List.Generator().Set(count)
}

func (list *List) SetRandom(min, max uint) chan interface{} {
  return dict.Generator().SetRandom(min, max)
}
