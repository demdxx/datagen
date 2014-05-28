// @project datagen
// @copyright Dmitry Ponomarev <demdxx@gmail.com> 2014
//
// This work is licensed under the Creative Commons Attribution 4.0 International License.
// To view a copy of this license, visit http://creativecommons.org/licenses/by/4.0/.

package data

import (
  "math/rand"
)

type PageGenerator struct {
  Generator
  PageSize      uint
  ElementsCount uint
}

func MakePageGenerator(structInfo *Item, pageSize, elements uint) *Generator {
  return &PageGenerator{Generator{structInfo: structInfo}, PageSize: pageSize, ElementsCount: elements}
}

///////////////////////////////////////////////////////////////////////////////
/// Getters/Setters
///////////////////////////////////////////////////////////////////////////////

func (gen *PageGenerator) Count() uint {
  return gen.ElementsCount
}

func (gen *PageGenerator) PageCount() uint {
  var count uint = gen.ElementsCount / gen.PageSize
  if gen.ElementsCount%gen.PageSize > 0 {
    count++
  }
  return count
}

func (gen *PageGenerator) PageItemsCount(page uint) uint {
  pageCount := gen.PageCount()
  if page >= pageCount {
    return 0
  }
  if page == pageCount {
    return gen.ElementsCount % gen.PageSize
  }
  return gen.PageSize
}

func (gen *PageGenerator) UpdateCount(min, max uint) *PageGenerator {
  gen.ElementsCount = (uint)(rand.Intn(int(max)-int(min))) + min
  return gen
}

func (gen *PageGenerator) Get(item uint) interface{} {
  if item > gen.ElementsCount {
    return nil
  }
  return gen.Generator.Get(item)
}

///////////////////////////////////////////////////////////////////////////////
/// Generators
///////////////////////////////////////////////////////////////////////////////

func (gen *PageGenerator) Page(page uint) chan interface{} {
  pageSize := gen.PageItemsCount(page)
  out := make(chan interface{})
  go func() {
    startIndex := page * gen.PageSize
    for i := startIndex; i < startIndex+pageSize; i++ {
      out <- gen.Get(i)
    }
    close(out)
  }()
  return out
}

func (gen *PageGenerator) Set() chan interface{} {
  return gen.Generator.Set(gen.Count())
}
