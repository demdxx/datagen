package generator

import (
  "dg/data"
  "strings"
  "errors"
)

type GoGenerator struct {
  Generator
}

func (self *GoGenerator) GenerateBody() (string, error) {
  return "", nil
}

func (self *GoGenerator) generateFieldString(item data.IItem) (string, error) {
  switch item.(type) {
  case data.IDictionary:
    return ".Dictionary()", nil
  case data.IList:
    return ".List()", nil
  // case data.TextGenerator:
  //   return "data.MakeTextGenerator()", nil
  // case data.NameGenerator:
  //   return "data.MakeNameGenerator()", nil
  // case data.IntegerGenerator:
  //   return "data.MakeIntegerGenerator()", nil
  }
  return "", errors.New("Undefined generation type")
}

func (self *GoGenerator) GenerateField(item data.IItem) (string, error) {
  switch item.(type) {
  case data.IDictionary:
    v, err := self.Generator.GenerateField(item)
    if nil != err { return "", err }
    return ".Dictionary()" + v, nil
  case data.IList:
    v, err := self.Generator.GenerateField(item)
    if nil != err { return "", err }
    return ".List()" + v, nil
  }

  v, err := self.GenerateField(item)
  if nil != err { return "", err }
  return strings.Replace(".Field({field})", "{field}", v, -1), nil
}

func (self *GoGenerator) GenerateKeyField(key string, item data.IItem) (string, error) {
  switch item.(type) {
  case data.IDictionary:
    v1, err1 := self.Generator.GenerateField(item)
    if nil != err1 { return "", err1 }
    return strings.Replace(".Dictionary(\"{key}\")", "{key}", key, -1) + v1, nil
  case data.IList:
    v2, err2 := self.Generator.GenerateField(item)
    if nil != err2 { return "", err2 }
    return strings.Replace(".List(\"{key}\")", "{key}", key, -1) + v2, nil
  }

  v, err := self.GenerateField(item)
  if nil != err { return "", err }
  s1 := strings.Replace(".Field(\"{key}\", {field})", "{key}", key, -1)
  return strings.Replace(s1, "{field}", v, -1), nil
}
