package dg;

func GeneratorTest() {
  d := MakeGenerationStructure()
        .Dictionary()
          .Dictionary("user")
            .Filed("sex", MakeBooleanGenerator().setTrueFalse('male', 'female'))
            .Field("name", MakeNameGenerator())
            .Field("birthday", MakeBirthdayGenerator("dd.MM.yyyy"))
            .Field("about", MakeTextGenerator().MaxWords(100).MinWords(25))
          .Parent()
            .Field("IP", MakeIPGenerator().IPV6(false))

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
  gb.SetGenerator(MakeObjectiveCGenerator())
    .SetGenerateStructure(d)
    .SetMaxItemsCount(100)
    .SetPageSize(10)
    .SetStoreElements(true)
    .Write(outputStream)
}
