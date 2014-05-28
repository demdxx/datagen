// @project datagen
// @copyright Dmitry Ponomarev <demdxx@gmail.com> 2014
//
// This work is licensed under the Creative Commons Attribution 4.0 International License.
// To view a copy of this license, visit http://creativecommons.org/licenses/by/4.0/.

package main

import (
  "bytes"
  "datagen/data"
  "datagen/generator"
  "fmt"
  "testing"
)

var structure *Item

func init() {
  structure = data.MakeDictionary().
    Dictionary("user").
    Filed("sex", data.MakeBooleanGenerator().setTrueFalse("male", "female")).
    Field("name", data.MakeNameGenerator()).
    Field("birthday", data.MakeBirthdayGenerator("dd.MM.yyyy")).
    Field("about", data.MakeTextGenerator().MaxWords(100).MinWords(25)).
    Parent().(*data.Dictionary).
    Field("IP", data.MakeIPGenerator().IPV6(false))

  /* Output

     [
       {
         'IP': '192.168.0.1',
         'user': {
           'sex': 'male',
           'name': 'Mike Handson',
           'birthday': '1.11.2062',
           'about': 'text about Mike Handson'
         }
       }
     ]

  */
}

func GeneratorTest(t *testing.T) {
  for _, it := range structure.Pages(10 /* page count */, 10 /* elements in page */) {
    fmt.Println(it)
  }
}

func FormatTest(t *testing.T) {
  outputStream := bytes.NewBuffer(nil)
  if err := structure.Set(10 /* Elements in set */).Write(outputStream); nil != err {
    t.Error(err)
  }
}

func BuilderTest(t *testing.T) {
  outputStream := bytes.NewBuffer(nil)
  gb := MakeBuilder(d, generator.MakeObjectiveCGenerator())
  if err := gb.Write(outputStream); nil != err {
    t.Error(err)
  }
}
