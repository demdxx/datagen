package main

import (
  "data"
)

func GeneratorTest() {
  d := data.MakeDictionary().
    Dictionary("user").
    Filed("sex", data.MakeBooleanGenerator().setTrueFalse("male", "female")).
    Field("name", data.MakeNameGenerator()).
    Field("birthday", data.MakeBirthdayGenerator("dd.MM.yyyy")).
    Field("about", data.MakeTextGenerator().MaxWords(100).MinWords(25)).
    Parent().
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

  gb := &GenBuilder
  gb.SetGenerator(generator.MakeObjectiveCGenerator()).
    SetGenerateStructure(d).
    SetMaxItemsCount(100).
    SetPageSize(10).
    SetStoreElements(true).
    Write(outputStream)
}
