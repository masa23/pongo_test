package main

import (
	"github.com/flosch/pongo2"
	"github.com/k0kubun/pp"
)

func main() {
	var template = pongo2.Must(pongo2.FromFile("template.j2"))
	str := []string{"hoge", "fuga"}
	out, err := template.Execute(pongo2.Context{"hoge": str})
	if err != nil {
		panic(err)
	}

	pp.Println(out)
}
