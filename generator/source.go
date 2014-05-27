package generator

type ISource interface {
  FirstnameArray(count int) []string
  LastnameArray(count int) []string
}

type Source interface {
  ISource
}
