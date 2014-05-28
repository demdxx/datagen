// @project datagen
// @copyright Dmitry Ponomarev <demdxx@gmail.com> 2014
//
// This work is licensed under the Creative Commons Attribution 4.0 International License.
// To view a copy of this license, visit http://creativecommons.org/licenses/by/4.0/.

package data

type IGenerator interface {
  Count() int
  PageCount() int
  PageItemsCount(page int) int

  SetStructInfo(info IContainer) IGenerator
  GetMaxItemCount() int
  SetMaxItemCount(count int) IGenerator
  GetMinItemCount() int
  SetMinItemCount(count int) IGenerator

  Next() interface{}
  Get(i int) interface{}
  Generator() chan interface{}
  Page(page int) chan interface{}
}

type Generator struct {
  IGenerator

  structInfo IContainer

  maxItemCount int
  minItemCount int
  pageSize     int
  pageCount    int
}

///////////////////////////////////////////////////////////////////////////////
/// Getters/Setters
///////////////////////////////////////////////////////////////////////////////

func (self *Generator) Count() int {
  return self.pageSize * self.pageCount
}

func (self *Generator) PageCount() int {
  return self.pageCount
}

func (self *Generator) PageItemsCount(page int) int {
  return self.pageSize
}

func (self *Generator) SetStructInfo(info IContainer) IGenerator {
  self.structInfo = info
  return self.IGenerator
}

func (self *Generator) GetMaxItemCount() int {
  return self.maxItemCount
}

func (self *Generator) SetMaxItemCount(count int) IGenerator {
  self.maxItemCount = count
  return self.IGenerator
}

func (self *Generator) GetMinItemCount() int {
  return self.minItemCount
}

func (self *Generator) SetMinItemCount(count int) IGenerator {
  self.minItemCount = count
  return self.IGenerator
}

func (self *Generator) Next() interface{} {
  return self.structInfo.Generate()
}

func (self *Generator) Get(item int) interface{} {
  return self.Next()
}

///////////////////////////////////////////////////////////////////////////////
/// Generators
///////////////////////////////////////////////////////////////////////////////

func (self *Generator) Page(page int) chan interface{} {
  out := make(chan interface{})
  go func() {
    for i := page * self.pageSize; i < (1+page)*self.pageSize; i++ {
      out <- self.Get(i)
    }
    close(out)
  }()
  return out
}

func (self *Generator) Generator() chan interface{} {
  out := make(chan interface{})
  go func() {
    for i := 0; i < self.Count(); i++ {
      out <- self.Get(i)
    }
    close(out)
  }()
  return out
}
