package data

type Dictionary struct {
  IDictionary
  fields  map[string]IItem
}

func MakeDictionary() *Dictionary {
  m := &Dictionary{}
  m.fields = make(map[string]IItem)
  return m
}

func (self *Dictionary) Field(key string, item IItem) IDictionary {
  self.fields[key] = item
  return self.IDictionary
}

func (self *Dictionary) Fields() map[string]IItem {
  return self.fields
}

func (self *Dictionary) Generate() interface{} {
  m := make(map[string]interface{})
  for k, v := range self.fields {
    m[k] = v.Generate()
  }
  return m
}

