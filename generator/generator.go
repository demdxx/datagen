package generator

import (
  "dg/data"
  "bytes"
  "strings"
  "errors"
)

type IGenerator interface {
  SetSource(source ISource) IGenerator
  SetGenerateStructure(sInfo data.IContainer) IGenerator
  Generate() (string, error)
  GenerateBody() (string, error)
  GenerateStructure() (string, error)
  GenerateField(item data.IItem) (string, error)
  GenerateKeyField(key string, item data.IItem) (string, error)
}

type Generator struct {
  IGenerator
  source        ISource
  structInfo    data.IContainer
}

func (self *Generator) SetSource(source ISource) IGenerator {
  self.source = source
  return self.IGenerator
}

func (self *Generator) SetGenerateStructure(sInfo data.IContainer) IGenerator {
  self.structInfo = sInfo
  return self
}

///////////////////////////////////////////////////////////////////////////////
/// Generating
///////////////////////////////////////////////////////////////////////////////

func (self *Generator) Generate() (string, error) {
  body, berr := self.GenerateBody()
  if nil != berr {
    return "", berr
  }

  st, serr := self.GenerateStructure()
  if nil != serr {
    return "", serr
  }

  return strings.Replace(body, "{structure}", st, -1), nil
}

func (self *Generator) GenerateBody() (string, error) {
  return "", nil
}

func (self *Generator) GenerateStructure() (string, error) {
  var buffer bytes.Buffer
  switch self.structInfo.(type) {
    case data.IDictionary:
      for k1, v1 := range self.structInfo.(data.IDictionary).Fields() {
        str1, err1 := self.GenerateKeyField(k1, v1)
        if nil != err1 {
          return "", err1
        }
        buffer.WriteString(str1)
      }
      return buffer.String(), nil
    case data.IList:
      for _, v2 := range self.structInfo.(data.IList).Fields() {
        str2, err2 := self.GenerateField(v2)
        if nil != err2 {
          return "", err2
        }
        buffer.WriteString(str2)
      }
      return buffer.String(), nil
  }
  return "", errors.New("Invalid structure container")
}

func (self *Generator) GenerateField(item data.IItem) (string, error) {
  var buffer bytes.Buffer
  switch item.(type) {
    case data.IDictionary:
      for k1, v1 := range self.structInfo.(data.IDictionary).Fields() {
        str1, err1 := self.GenerateKeyField(k1, v1)
        if nil != err1 {
          return "", err1
        }
        buffer.WriteString(str1)
      }
    case data.IList:
      for _, v1 := range self.structInfo.(data.IList).Fields() {
        str1, err1 := self.GenerateField(v1)
        if nil != err1 {
          return "", err1
        }
        buffer.WriteString(str1)
      }
  }
  return "", errors.New("Undefined generation type")
}

func (self *Generator) GenerateKeyField(key string, item data.IItem) (string, error) {
  return "", errors.New("Undefined generation type")
}


