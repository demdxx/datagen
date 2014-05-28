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

func (list *List) Dictionary(key string) *Dictionary {
  dict := &Dictionary{Item{parent: list}}
  list.fields[key] = dict
  return dict
}

func (list *List) List() *List {
  l := &List{Item{parent: list}}
  list.items = append(list.items, l)
  return list
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
