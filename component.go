package sydercomponents

import (
	"io"
	"strconv"
	"strings"

	g "maragu.dev/gomponents"
)

// ComponentKind describes what kind of [Component] it is.
// This is useful for testing to determine the [Component] kind.
// The [newComponent] constructor defaults to being [KindElement].
type ComponentKind string

const (
	// KindNone is used by [Empty] component.
	// It renders nothing and has no attributes and children.
	KindNone ComponentKind = ""

	// KindText is used by [Text] and [Raw] text components.
	// Its content is readable with [Components.GetText].
	KindText ComponentKind = "text"

	// KindNode is used by [Node] components wrapping gomponents.Node.
	// No getters work for this kind (except [GetKind]).
	KindNode ComponentKind = "node"

	// KindElement is used by all HTML components.
	// Its attributes can be accessed with Get* methods.
	KindElement ComponentKind = "element"
)

// Component represents a HTML element in the UI library.
// It stores attributes, nested children, kind and, in case of Text components, text,
// which are evaluated and converted into standard gomponents when Render is called.
type Component struct {
	children []Component
	content  g.Node
	attrs    map[string]string
	kind     ComponentKind
	text     string
}

// Option defines a functional configuration pattern.
// It is used to safely modify a Component's attributes, classes,
// or children before the component is rendered.
type Option func(c *Component)

// ChildrenAppend appends nested Component elements to the component.
func ChildrenAppend(value ...Component) Option {
	return func(c *Component) {
		c.children = append(c.children, value...)
	}
}

// Children overwrites nested Component elements of the component.
func Children(value ...Component) Option {
	return func(c *Component) {
		c.children = value
	}
}

// Attribute sets a custom HTML attribute.
func Attribute(name, value string) Option {
	return func(c *Component) {
		c.attrs[name] = value
	}
}

// ClassAppend safely adds a new CSS class to the component without overwriting existing classes.
func ClassAppend(value string) Option {
	return func(c *Component) {
		if value == "" {
			return
		}

		curClass := c.attrs["class"]
		trimmedValue := strings.Trim(value, " ")
		if curClass == "" {
			c.attrs["class"] = trimmedValue
		} else {
			c.attrs["class"] += " " + trimmedValue
		}
	}
}

// Class sets the exact CSS class string, overwrites any existing classes.
func Class(value string) Option {
	return func(c *Component) {
		c.attrs["class"] = value
	}
}

// ID sets the HTML id attribute.
func ID(value string) Option {
	return func(c *Component) {
		c.attrs["id"] = value
	}
}

// Hidden sets the HTML hidden attribute.
func Hidden() Option {
	return func(c *Component) {
		c.attrs["hidden"] = ""
	}
}

// Disabled sets the HTML disabled attribute.
func Disabled() Option {
	return func(c *Component) {
		c.attrs["disabled"] = ""
	}
}

// Required sets the HTML required attribute.
func Required() Option {
	return func(c *Component) {
		c.attrs["required"] = ""
	}
}

// Data sets the HTML data-* attribute.
func Data(key, value string) Option {
	return func(c *Component) {
		c.attrs["data-"+key] = value
	}
}

// DialogTrigger sets the required attributes to the trigger element.
func DialogTrigger(dialogID string) Option {
	return func(c *Component) {
		Attribute("commandfor", dialogID)(c)
		Attribute("command", "show-modal")(c)
	}
}

// DialogClose sets the required attributes to the closing element.
func DialogClose(dialogID string) Option {
	return func(c *Component) {
		Attribute("commandfor", dialogID)(c)
		Attribute("command", "close")(c)
	}
}

// PopoverTrigger sets the required attribute for non-modal dialog trigger element.
func PopoverTrigger(dialogID string) Option {
	return Attribute("popovertarget", dialogID)
}

// PopoverClose sets the required attributes for non-modal dialog closing element.
func PopoverClose(dialogID string) Option {
	return func(c *Component) {
		Attribute("popovertarget", dialogID)(c)
		Attribute("popovertargetaction", "hide")(c)
	}
}

// Aria sets the HTML aria-* attribute.
func Aria(key, value string) Option {
	return func(c *Component) {
		c.attrs["aria-"+key] = value
	}
}

// Style sets the HTML style attribute.
func Style(value string) Option {
	return func(c *Component) {
		c.attrs["style"] = value
	}
}

// Name sets the HTML name attribute.
func Name(value string) Option {
	return func(c *Component) {
		c.attrs["name"] = value
	}
}

// Type sets the HTML type attribute.
func Type(value string) Option {
	return func(c *Component) {
		c.attrs["type"] = value
	}
}

// Placeholder sets the HTML placeholder attribute.
func Placeholder(value string) Option {
	return func(c *Component) {
		c.attrs["placeholder"] = value
	}
}

// Value sets the HTML value attribute.
func Value(value string) Option {
	return func(c *Component) {
		c.attrs["value"] = value
	}
}

// Src sets the HTML src attribute.
func Src(value string) Option {
	return func(c *Component) {
		c.attrs["src"] = value
	}
}

// Alt sets the HTML alt attribute.
func Alt(value string) Option {
	return func(c *Component) {
		c.attrs["alt"] = value
	}
}

// Href sets the HTML href attribute.
func Href(value string) Option {
	return func(c *Component) {
		c.attrs["href"] = value
	}
}

// Target sets the HTML target attribute.
func Target(value string) Option {
	return func(c *Component) {
		c.attrs["target"] = value
	}
}

// TabIndex sets the HTML tabindex attribute.
func TabIndex(value int) Option {
	return func(c *Component) {
		c.attrs["tabindex"] = strconv.Itoa(value)
	}
}

// Render renders the underlying gomponent into the io.Writer.
func (c Component) Render(w io.Writer) error {
	if c.content == nil {
		return nil
	}
	return c.content.Render(w)
}

// newComponent creates a new component with the class attribute set and given options.
// The default kind is Element. Other kinds require manual Component instantiation.
func newComponent(cls string, opts ...Option) Component {
	attrs := map[string]string{}
	if cls != "" {
		attrs["class"] = cls
	}

	c := Component{
		attrs: attrs,
		kind:  KindElement,
	}
	for _, o := range opts {
		o(&c)
	}
	return c
}

// GetClass returns the applied class to the element.
func (c Component) GetClass() string {
	return c.attrs["class"]
}

// GetID returns the element id, but only if ID was called.
func (c Component) GetID() string {
	return c.attrs["id"]
}

// GetChildren returns the children nodes.
func (c Component) GetChildren() []Component {
	return c.children
}

// GetContent returns the rendered content node.
func (c Component) GetContent() g.Node {
	return c.content
}

// GetAttrs returns all attributes as a map.
func (c Component) GetAttrs() map[string]string {
	return c.attrs
}

// GetText returns the content of a raw/text node.
func (c Component) GetText() string {
	return c.text
}

// GetKind returns the kind of the component.
func (c Component) GetKind() ComponentKind {
	return c.kind
}

// GetData returns a data attribute value.
func (c Component) GetData(key string) (string, bool) {
	val, ok := c.attrs["data-"+key]
	return val, ok
}

// GetAria returns an aria attribute value.
func (c Component) GetAria(key string) (string, bool) {
	val, ok := c.attrs["aria-"+key]
	return val, ok
}

// GetStyle returns the style attribute value.
func (c Component) GetStyle() (string, bool) {
	val, ok := c.attrs["style"]
	return val, ok
}

// GetName returns the name attribute value.
func (c Component) GetName() (string, bool) {
	val, ok := c.attrs["name"]
	return val, ok
}

// GetType returns the type attribute value.
func (c Component) GetType() (string, bool) {
	val, ok := c.attrs["type"]
	return val, ok
}

// GetPlaceholder returns the placeholder attribute value.
func (c Component) GetPlaceholder() (string, bool) {
	val, ok := c.attrs["placeholder"]
	return val, ok
}

// GetValue returns the value attribute value.
func (c Component) GetValue() (string, bool) {
	val, ok := c.attrs["value"]
	return val, ok
}

// GetSrc returns the src attribute value.
func (c Component) GetSrc() (string, bool) {
	val, ok := c.attrs["src"]
	return val, ok
}

// GetAlt returns the alt attribute value.
func (c Component) GetAlt() (string, bool) {
	val, ok := c.attrs["alt"]
	return val, ok
}

// GetHref returns the href attribute value.
func (c Component) GetHref() (string, bool) {
	val, ok := c.attrs["href"]
	return val, ok
}

// GetTarget returns the target attribute value.
func (c Component) GetTarget() (string, bool) {
	val, ok := c.attrs["target"]
	return val, ok
}

// GetTabIndex returns the tabindex attribute value.
func (c Component) GetTabIndex() (int, bool) {
	val, ok := c.attrs["tabindex"]
	if !ok {
		return 0, false
	}
	i, err := strconv.Atoi(val)
	if err != nil {
		return 0, false
	}
	return i, true
}

// HasHidden returns true if the hidden attribute was set.
func (c Component) HasHidden() bool {
	_, ok := c.attrs["hidden"]
	return ok
}

// HasDisabled returns true if the disabled attribute was set.
func (c Component) HasDisabled() bool {
	_, ok := c.attrs["disabled"]
	return ok
}

// HasRequired returns true if the required attribute was set.
func (c Component) HasRequired() bool {
	_, ok := c.attrs["required"]
	return ok
}

// manuallySetAttributes is supposed to be called for every
// new component. It injects all manually set attributes, such as
// class (via the [Class] hook), ID (via the [ID] hook)
// children (via the [Children] hook) and others.
func (c Component) manuallySetAttributes() g.Group {
	nodes := make([]g.Node, 0, len(c.attrs)+1)
	for key, value := range c.attrs {
		nodes = append(nodes, g.Attr(key, value))
	}
	nodes = append(nodes, g.Map(c.children, func(child Component) g.Node { return child }))
	return nodes
}
