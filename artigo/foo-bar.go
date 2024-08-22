package main

type Foo struct{

}

type Bar struct{
	f Foo
}

func NewFoo() Foo{
 return Foo{}
}

// Bar vai ser instanciado e injetado pelo Colibri
func NewBar() Bar{
	// Foo não é instanciado e nem injetado pelo Colibri
	f := NewFoo()
	return Bar{f: f}
}