package dg

import (
  "dg/data"
  "dg/generator"
)

type Builder struct {
  structInfo    data.IContainer
  generator     generator.IGenerator

  MaxItemCount  int
  MinItemCount  int
  PageSize      int
  PageCount     int
}

func MakeBuilder(sInfo data.IContainer, gen generator.IGenerator) *Builder {
  b := &Builder{
    structInfo: sInfo,
    generator: gen,
    MaxItemCount: 0,
    MinItemCount: 0,
    PageSize: 10,
    PageCount: 100,
  }
  return b
}

///////////////////////////////////////////////////////////////////////////////
/// Getters/Setters
///////////////////////////////////////////////////////////////////////////////

func (self *Builder) SetGenerator(gen generator.IGenerator) *Builder {
  self.generator = gen
  return self
}

func (self *Builder) SetGenerateStructure(sInfo data.IContainer) *Builder {
  self.structInfo = sInfo
  return self
}

func (self *Builder) GetMaxItemCount() int {
  return self.MaxItemCount
}

func (self *Builder) SetMaxItemCount(count int) *Builder {
  self.MaxItemCount = count
  return self
}

func (self *Builder) GetMinItemCount() int {
  return self.MinItemCount
}

func (self *Builder) SetMinItemCount(count int) *Builder {
  self.MinItemCount = count
  return self
}

func (self *Builder) Build() (string, error) {
  self.generator.SetGenerateStructure(self.structInfo)
  // self.generator.SetMinItemCount(self.GetMinItemCount())
  // self.generator.SetMaxItemCount(self.GetMaxItemCount())
  return self.generator.Generate()
}




