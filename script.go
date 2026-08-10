package sydercomponents

import (
	"context"
	"crypto/sha512"
	"encoding/base64"

	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

type ScriptRegistry struct {
	Hashes []string
}

func (r *ScriptRegistry) RegisterScript(scriptContent string) {
	hash := sha512.Sum512([]byte(scriptContent))
	encoded := base64.StdEncoding.EncodeToString(hash[:])
	r.Hashes = append(r.Hashes, "'sha512-"+encoded+"'")
}

type ctxKey string

const registryKey ctxKey = "script_registry"

func WithScriptRegistry(ctx context.Context, r *ScriptRegistry) context.Context {
	return context.WithValue(ctx, registryKey, r)
}

func Script(ctx context.Context, content string, opts ...Option) Component {
	if reg, ok := ctx.Value(registryKey).(*ScriptRegistry); ok {
		reg.RegisterScript(content)
	}

	opts = append([]Option{WithAttributeAppend(g.Raw(content))}, opts...)
	n := newComponent("", opts...)

	n.content = h.Script(n.manuallySetAttributes())
	return n
}
