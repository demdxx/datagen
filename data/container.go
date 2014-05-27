package data

type IContainer interface {
  IItem
  Parent() IContainer
  List() IList
  Dictionary() IDictionary
}

type IList interface {
  IContainer
  Field(item IItem) IList
  Fields() []IItem
}

type IDictionary interface {
  IContainer
  Field(key string, item IItem) IDictionary
  Fields() map[string]IItem
}
