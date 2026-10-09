package sydercomponents

import (
	"fmt"
	"html"
	"io"
	"strconv"
)

// attribute is an internal [Node] of [KindAttr] kind.
// The name of the attribute can be accessed with the [attribute.Name] method.
// The value of the attribute can be accessed with the [attribute.Value] and [attribute.Attr] methods.
// Boolean attributes have an empty value and are rendered as 'key=""'.
// It has no children.
type attribute struct {
	name  string
	value string
}

func (a attribute) Kind() Kind       { return KindAttr }
func (a attribute) Name() string     { return a.name }
func (a attribute) Value() string    { return a.value }
func (a attribute) Children() []Node { return []Node{} }

// Attr returns the underlying attribute's value and true if the name matches.
// Otherwise, it returns an empty string and false.
func (a attribute) Attr(name string) (string, bool) {
	if a.name == name {
		return a.value, true
	}

	return "", false
}

// Render writes 'name="value"' into the provided [io.Writer] instance.
// In case of boolean attributes, 'name=""' is written into the writer.
func (a attribute) Render(w io.Writer) error {
	_, err := fmt.Fprintf(w, "%s=\"%s\"", html.EscapeString(a.name), html.EscapeString(a.value))
	return err
}

// Attribute sets a custom HTML attribute.
func Attr(name, value string) Node {
	return attribute{name: name, value: value}
}

// Class sets the CSS class. Can be used multiple times, their values will be joined with spaces.
func Class(value string) Node {
	return Attr("class", value)
}

// ID sets the HTML id attribute.
func ID(value string) Node {
	return Attr("id", value)
}

// Hidden sets the HTML hidden attribute.
func Hidden() Node {
	return Attr("hidden", "")
}

// Disabled sets the HTML disabled attribute.
func Disabled() Node {
	return Attr("disabled", "")
}

// Required sets the HTML required attribute.
func Required() Node {
	return Attr("required", "")
}

// Data sets the HTML data-* attribute.
func Data(key, value string) Node {
	return Attr("data-"+key, value)
}

// Aria sets the HTML aria-* attribute.
func Aria(key, value string) Node {
	return Attr("aria-"+key, value)
}

// Style sets the HTML style attribute.
func Style(value string) Node {
	return Attr("style", value)
}

// Name sets the HTML name attribute.
func Name(value string) Node {
	return Attr("name", value)
}

// Type sets the HTML type attribute.
func Type(value string) Node {
	return Attr("type", value)
}

// Placeholder sets the HTML placeholder attribute.
func Placeholder(value string) Node {
	return Attr("placeholder", value)
}

// Value sets the HTML value attribute.
func Value(value string) Node {
	return Attr("value", value)
}

// Src sets the HTML src attribute.
func Src(value string) Node {
	return Attr("src", value)
}

// Alt sets the HTML alt attribute.
func Alt(value string) Node {
	return Attr("alt", value)
}

// Href sets the HTML href attribute.
func Href(value string) Node {
	return Attr("href", value)
}

// Target sets the HTML target attribute.
func Target(value string) Node {
	return Attr("target", value)
}

// Lang sets the HTML lang attribute.
func Lang(value string) Node {
	return Attr("lang", value)
}

// Charset sets the HTML charset attribute.
func Charset(value string) Node {
	return Attr("charset", value)
}

// Rel sets the HTML rel attribute.
func Rel(value string) Node {
	return Attr("rel", value)
}

// Content sets the HTML content attribute.
func Content(value string) Node {
	return Attr("content", value)
}

// TabIndex sets the HTML tabindex attribute.
func TabIndex(value int) Node {
	return Attr("tabindex", strconv.Itoa(value))
}

// DialogTrigger sets the required attributes to the trigger element.
func DialogTrigger(dialogID string) Node {
	return Group(
		Attr("commandfor", dialogID),
		Attr("command", "show-modal"),
	)
}

// DialogClose sets the required attributes to the closing element.
func DialogClose(dialogID string) Node {
	return Group(
		Attr("commandfor", dialogID),
		Attr("command", "close"),
	)
}

// PopoverTrigger sets the required attribute for non-modal dialog trigger element.
func PopoverTrigger(dialogID string) Node {
	return Attr("popovertarget", dialogID)
}

// PopoverClose sets the required attributes for non-modal dialog closing element.
func PopoverClose(dialogID string) Node {
	return Group(
		Attr("popovertarget", dialogID),
		Attr("popovertargetaction", "hide"),
	)
}
