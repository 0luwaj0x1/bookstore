package mytypes

import "strings"


type MyBuilder struct {
	strings.Builder
}

type StringUppercaser struct {
	strings.Builder
}

func (su StringUppercaser) ToUpper() string {
	return strings.ToUpper(su.String())
}

func (mb MyBuilder) Hello() string {
	return "Hello, Gophers!"
}


type MyInt int


func (input *MyInt) Double() {
	*input *= 2
}
