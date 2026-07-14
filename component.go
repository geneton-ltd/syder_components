package sydercomponents

import (
	"io"
	"strconv"

	g "maragu.dev/gomponents"
)

type Component struct {
	children   []Component
	attributes []g.Node // non-element attributes
	content    g.Node
	attrs      map[string]string
}

type Option func(c *Component)

func WithChildren(value ...Component) Option {
	return func(c *Component) {
		c.children = value
	}
}

func WithAttributes(value ...g.Node) Option {
	return func(c *Component) {
		c.attributes = value
	}
}

func WithClassAppend(value string) Option {
	return func(c *Component) {
		c.attrs["class"] += value
	}
}

func WithClass(value string) Option {
	return func(c *Component) {
		c.attrs["class"] = value
	}
}

func WithID(value string) Option {
	return func(c *Component) {
		c.attrs["id"] = value
	}
}

func WithHidden() Option {
	return func(c *Component) {
		c.attrs["hidden"] = ""
	}
}

func WithDisabled() Option {
	return func(c *Component) {
		c.attrs["disabled"] = ""
	}
}

func WithRequired() Option {
	return func(c *Component) {
		c.attrs["required"] = ""
	}
}

func WithData(key, value string) Option {
	return func(c *Component) {
		c.attrs["data-"+key] = value
	}
}

func WithAria(key, value string) Option {
	return func(c *Component) {
		c.attrs["aria-"+key] = value
	}
}

func WithStyle(value string) Option {
	return func(c *Component) {
		c.attrs["style"] = value
	}
}

func WithName(value string) Option {
	return func(c *Component) {
		c.attrs["name"] = value
	}
}

func WithType(value string) Option {
	return func(c *Component) {
		c.attrs["type"] = value
	}
}

func WithPlaceholder(value string) Option {
	return func(c *Component) {
		c.attrs["placeholder"] = value
	}
}

func WithValue(value string) Option {
	return func(c *Component) {
		c.attrs["value"] = value
	}
}

func WithSrc(value string) Option {
	return func(c *Component) {
		c.attrs["src"] = value
	}
}

func WithAlt(value string) Option {
	return func(c *Component) {
		c.attrs["alt"] = value
	}
}

func WithHref(value string) Option {
	return func(c *Component) {
		c.attrs["href"] = value
	}
}

func WithTarget(value string) Option {
	return func(c *Component) {
		c.attrs["target"] = value
	}
}

func WithTabIndex(value int) Option {
	return func(c *Component) {
		c.attrs["tabindex"] = strconv.Itoa(value)
	}
}

func (c Component) Render(w io.Writer) error {
	return c.content.Render(w)
}

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

// Class returns the applied class to the element
func (c Component) GetClass() string {
	return c.attrs["class"]
}

// ID returns the element id, but only if WithID was called
func (c Component) GetID() string {
	return c.attrs["id"]
}

// GetChildren returns the children nodes
func (c Component) GetChildren() []Component {
	return c.children
}

// GetAttributes returns the children nodes
func (c Component) GetAttributes() []g.Node {
	return c.attributes
}

// GetContent returns the rendered content node
func (c Component) GetContent() g.Node {
	return c.content
}

// GetAttrs returns all attributes as a map
func (c Component) GetAttrs() map[string]string {
	return c.attrs
}

// GetData returns a data attribute value
func (c Component) GetData(key string) (string, bool) {
	val, ok := c.attrs["data-"+key]
	return val, ok
}

// GetAria returns an aria attribute value
func (c Component) GetAria(key string) (string, bool) {
	val, ok := c.attrs["aria-"+key]
	return val, ok
}

// GetStyle returns the style attribute value
func (c Component) GetStyle() (string, bool) {
	val, ok := c.attrs["style"]
	return val, ok
}

// GetName returns the name attribute value
func (c Component) GetName() (string, bool) {
	val, ok := c.attrs["name"]
	return val, ok
}

// GetType returns the type attribute value
func (c Component) GetType() (string, bool) {
	val, ok := c.attrs["type"]
	return val, ok
}

// GetPlaceholder returns the placeholder attribute value
func (c Component) GetPlaceholder() (string, bool) {
	val, ok := c.attrs["placeholder"]
	return val, ok
}

// GetValue returns the value attribute value
func (c Component) GetValue() (string, bool) {
	val, ok := c.attrs["value"]
	return val, ok
}

// GetSrc returns the src attribute value
func (c Component) GetSrc() (string, bool) {
	val, ok := c.attrs["src"]
	return val, ok
}

// GetAlt returns the alt attribute value
func (c Component) GetAlt() (string, bool) {
	val, ok := c.attrs["alt"]
	return val, ok
}

// GetHref returns the href attribute value
func (c Component) GetHref() (string, bool) {
	val, ok := c.attrs["href"]
	return val, ok
}

// GetTarget returns the target attribute value
func (c Component) GetTarget() (string, bool) {
	val, ok := c.attrs["target"]
	return val, ok
}

// GetTabIndex returns the tabindex attribute value
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

// HasHidden returns true if the hidden attribute was set
func (c Component) HasHidden() bool {
	_, ok := c.attrs["hidden"]
	return ok
}

// HasDisabled returns true if the disabled attribute was set
func (c Component) HasDisabled() bool {
	_, ok := c.attrs["disabled"]
	return ok
}

// HasRequired returns true if the required attribute was set
func (c Component) HasRequired() bool {
	_, ok := c.attrs["required"]
	return ok
}

// manuallySetAttributes is supposed to be called for every
// new component. It injects all manually set attributes, such as
// class (via the WithClass hook), ID (via the WithID hook)
// children (via the WithChildren hook) and others
func (c Component) manuallySetAttributes() g.Group {
	nodes := make([]g.Node, 0, len(c.attrs)+1)
	for key, value := range c.attrs {
		nodes = append(nodes, g.Attr(key, value))
	}
	nodes = append(nodes, g.Group(c.attributes), g.Map(c.children, func(child Component) g.Node { return child }))
	return nodes
}
