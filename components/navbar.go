package components

import sc "github.com/geneton-ltd/syder_components"

// Navbar returns a HTML5 <div> element with simple styles
// acting as a container for navigation elements.
func Navbar(children ...sc.Node) sc.Node {
	children = append([]sc.Node{sc.Class("h-[50px] bg-gray-800 flex justify-between text-white")}, children...)
	return sc.Div(children...)
}

// NavLink returns a HTML5 <a> element, which is supposed to be put into [Navbar].
func NavLink(children ...sc.Node) sc.Node {
	children = append([]sc.Node{sc.Class("hover:bg-gray-600 flex items-center px-2")}, children...)
	return sc.A(children...)
}
