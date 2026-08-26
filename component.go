package sydercomponents

import (
	"io"
	"strconv"
	"strings"

	g "maragu.dev/gomponents"
)

// Component represents a HTML element in the UI library.
// It stores attributes, nested children, and raw nodes, which are
// evaluated and converted into standard gomponents when Render is called.
type Component struct {
	children []Component
	nodes    []g.Node // non-component nodes
	content  g.Node
	attrs    map[string]string
}

// Option defines a functional configuration pattern.
// It is used to safely modify a Component's attributes, classes,
// or children before the component is rendered.
type Option func(c *Component)

// WithChildrenAppend appends nested Component elements to the component.
func WithChildrenAppend(value ...Component) Option {
	return func(c *Component) {
		c.children = append(c.children, value...)
	}
}

// WithChildren overwrites nested Component elements of the component.
func WithChildren(value ...Component) Option {
	return func(c *Component) {
		c.children = value
	}
}

// WithNodeAppend appends nested gomponents nodes to the component.
func WithNodeAppend(value g.Node) Option {
	return func(c *Component) {
		c.nodes = append(c.nodes, value)
	}
}

// WithNodes overwrites nested gomponents nodes of the component.
func WithNodes(value ...g.Node) Option {
	return func(c *Component) {
		c.nodes = value
	}
}

// WithAttribute sets a custom HTML attribute..
func WithAttribute(name, value string) Option {
	return func(c *Component) {
		c.attrs[name] = value
	}
}

// WithClassAppend safely adds a new CSS class to the component without overwriting existing classes.
func WithClassAppend(value string) Option {
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

// WithClass sets the exact CSS class string, overwrites any existing classes.
func WithClass(value string) Option {
	return func(c *Component) {
		c.attrs["class"] = value
	}
}

// WithID sets the HTML id attribute.
func WithID(value string) Option {
	return func(c *Component) {
		c.attrs["id"] = value
	}
}

// WithHidden sets the HTML hidden attribute.
func WithHidden() Option {
	return func(c *Component) {
		c.attrs["hidden"] = ""
	}
}

// WithDisabled sets the HTML disabled attribute.
func WithDisabled() Option {
	return func(c *Component) {
		c.attrs["disabled"] = ""
	}
}

// WithRequired sets the HTML required attribute.
func WithRequired() Option {
	return func(c *Component) {
		c.attrs["required"] = ""
	}
}

// WithData sets the HTML data-* attribute.
func WithData(key, value string) Option {
	return func(c *Component) {
		c.attrs["data-"+key] = value
	}
}

// WithDialogTrigger sets the required data attribute to the trigger element.
func WithDialogTrigger(id string) Option {
	return WithData("modal", id)
}

// WithAria sets the HTML aria-* attribute.
func WithAria(key, value string) Option {
	return func(c *Component) {
		c.attrs["aria-"+key] = value
	}
}

// WithStyle sets the HTML style attribute.
func WithStyle(value string) Option {
	return func(c *Component) {
		c.attrs["style"] = value
	}
}

// WithName sets the HTML name attribute.
func WithName(value string) Option {
	return func(c *Component) {
		c.attrs["name"] = value
	}
}

// WithType sets the HTML type attribute.
func WithType(value string) Option {
	return func(c *Component) {
		c.attrs["type"] = value
	}
}

// WithPlaceholder sets the HTML placeholder attribute.
func WithPlaceholder(value string) Option {
	return func(c *Component) {
		c.attrs["placeholder"] = value
	}
}

// WithValue sets the HTML value attribute.
func WithValue(value string) Option {
	return func(c *Component) {
		c.attrs["value"] = value
	}
}

// WithSrc sets the HTML src attribute.
func WithSrc(value string) Option {
	return func(c *Component) {
		c.attrs["src"] = value
	}
}

// WithAlt sets the HTML alt attribute.
func WithAlt(value string) Option {
	return func(c *Component) {
		c.attrs["alt"] = value
	}
}

// WithHref sets the HTML href attribute.
func WithHref(value string) Option {
	return func(c *Component) {
		c.attrs["href"] = value
	}
}

// WithTarget sets the HTML target attribute.
func WithTarget(value string) Option {
	return func(c *Component) {
		c.attrs["target"] = value
	}
}

// WithTabIndex sets the HTML tabindex attribute.
func WithTabIndex(value int) Option {
	return func(c *Component) {
		c.attrs["tabindex"] = strconv.Itoa(value)
	}
}

// Render renders the underlying gomponent into the io.Writer.
func (c Component) Render(w io.Writer) error {
	return c.content.Render(w)
}

// newComponent creates a new component with the class attribute set and given options.
func newComponent(cls string, opts ...Option) Component {
	c := Component{
		attrs: map[string]string{
			"class": cls,
		},
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

// GetID returns the element id, but only if WithID was called.
func (c Component) GetID() string {
	return c.attrs["id"]
}

// GetChildren returns the children nodes.
func (c Component) GetChildren() []Component {
	return c.children
}

// GetNodes returns the children gomponents nodes.
func (c Component) GetNodes() []g.Node {
	return c.nodes
}

// GetContent returns the rendered content node.
func (c Component) GetContent() g.Node {
	return c.content
}

// GetAttrs returns all attributes as a map.
func (c Component) GetAttrs() map[string]string {
	return c.attrs
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
// class (via the WithClass hook), ID (via the WithID hook)
// children (via the WithChildren hook) and others.
func (c Component) manuallySetAttributes() g.Group {
	nodes := make([]g.Node, 0, len(c.attrs)+1)
	for key, value := range c.attrs {
		nodes = append(nodes, g.Attr(key, value))
	}
	nodes = append(nodes, g.Group(c.nodes), g.Map(c.children, func(child Component) g.Node { return child }))
	return nodes
}
