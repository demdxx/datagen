// @project datagen
// @copyright Dmitry Ponomarev <demdxx@gmail.com> 2014
//
// This work is licensed under the Creative Commons Attribution 4.0 International License.
// To view a copy of this license, visit http://creativecommons.org/licenses/by/4.0/.

package data

type Generator struct {
  structInfo *Item
}

func MakeGenerator(structInfo *Item) *Generator {
  return &Generator{structInfo: structInfo}
}

///////////////////////////////////////////////////////////////////////////////
/// Getters/Setters
///////////////////////////////////////////////////////////////////////////////

func (gen *Generator) SetStructInfo(info *Item) *Generator {
  gen.structInfo = info
  return gen.IGenerator
}

func (gen *Generator) Next() interface{} {
  return gen.structInfo.Generate()
}

func (gen *Generator) Get(item uint) interface{} {
  return gen.Next()
}

///////////////////////////////////////////////////////////////////////////////
/// Generators
///////////////////////////////////////////////////////////////////////////////

func (gen *Generator) Set(count uint) chan interface{} {
  out := make(chan interface{})
  go func() {
    for i := 0; i < count; i++ {
      out <- gen.Get(i)
    }
    close(out)
  }()
  return out
}

func (gen *Generator) SetRandom(min, max uint) chan interface{} {
  return gen.Set((uint)(rand.Intn(int(max)-int(min))) + min)
}
