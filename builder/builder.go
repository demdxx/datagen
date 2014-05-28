// @project datagen
// @copyright Dmitry Ponomarev <demdxx@gmail.com> 2014
//
// This work is licensed under the Creative Commons Attribution 4.0 International License.
// To view a copy of this license, visit http://creativecommons.org/licenses/by/4.0/.

package builder

import (
  "datagen/data"
  "datagen/generator"
)

type Builder struct {
  structInfo data.IContainer
  generator  generator.IGenerator
}

func MakeBuilder(sInfo data.IContainer, gen generator.IGenerator) *Builder {
  b := &Builder{
    structInfo: sInfo,
    generator:  gen,
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

func (self *Builder) Build() (string, error) {
  self.generator.SetGenerateStructure(self.structInfo)
  return self.generator.Generate()
}

func (self *Builder) GetGenerator() generator.IGenerator {
  self.generator.SetGenerateStructure(self.structInfo)
  return self.generator
}
