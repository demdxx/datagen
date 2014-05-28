// @project datagen
// @copyright Dmitry Ponomarev <demdxx@gmail.com> 2014
//
// This work is licensed under the Creative Commons Attribution 4.0 International License.
// To view a copy of this license, visit http://creativecommons.org/licenses/by/4.0/.

package main

import (
  "github.com/demdxx/datagen/data"
)

func main() {
  // TODO generation from json, xml or yuml

  structure := data.MakeDictionary().
    Dictionary("user").
    Filed("sex", data.MakeBooleanGenerator().setTrueFalse("male", "female")).
    Field("name", data.MakeNameGenerator()).
    Field("birthday", data.MakeBirthdayGenerator("dd.MM.yyyy")).
    Field("about", data.MakeTextGenerator().MaxWords(100).MinWords(25)).
    Parent().(*data.Dictionary).
    Field("IP", data.MakeIPGenerator().IPV6(false))

  for _, it := range structure.Pages(10 /* page count */, 10 /* elements in page */) {
    fmt.Println(it)
  }
}
