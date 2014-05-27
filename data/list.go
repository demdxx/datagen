package data

type List struct {
  IList
  items   []IItem
}

func MakeList() *List {
  l := &List{}
  l.items = make([]IItem, 0, 0)
  return l
}

func (self *List) Field(item IItem) IList {
  self.items = append(self.items, item)
  return self
}

func (self *List) Fields() []IItem {
  return self.items
}

func (self *List) Generate() interface{} {
  l := make([]interface{}, 0, len(self.items))
  for _, v := range self.items {
    l = append(l, v.Generate())
  }
  return l
}
