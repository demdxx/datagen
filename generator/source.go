// @project datagen
// @copyright Dmitry Ponomarev <demdxx@gmail.com> 2014
//
// This work is licensed under the Creative Commons Attribution 4.0 International License.
// To view a copy of this license, visit http://creativecommons.org/licenses/by/4.0/.

package generator

type ISource interface {
  FirstnameArray(count int) []string
  LastnameArray(count int) []string
}

type Source interface {
  ISource
}
