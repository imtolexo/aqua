package aqua

import (
	"reflect"
	"strconv"
)

type Fixture struct {
	Prefix  string
	Root    string
	Url     string
	Version string
	Pretty  string
	Vendor  string
	Modules string
	Stub    string

	Cache string // cache store
	Ttl   string // cache ttl
	Auth  string // authentication method
	Log   bool
}

func NewFixtureFromTag(i interface{}, fieldName string) Fixture {
	out := Fixture{}
	field, _ := reflect.TypeOf(i).Elem().FieldByName(fieldName)
	tag := field.Tag

	var tmp string

	tmp = getTagValue(tag, "prefix", "pre")
	if tmp != "" {
		out.Prefix = tmp
	}

	tmp = getTagValue(tag, "root")
	if tmp != "" {
		out.Root = tmp
	}

	tmp = getTagValue(tag, "url")
	if tmp != "" {
		out.Url = tmp
	}

	tmp = getTagValue(tag, "version", "ver")
	if tmp != "" {
		out.Version = tmp
	}

	tmp = getTagValue(tag, "pretty", "pty")
	if tmp != "" {
		out.Pretty = tmp
	}

	tmp = getTagValue(tag, "vendor", "vnd")
	if tmp != "" {
		out.Vendor = tmp
	}

	tmp = getTagValue(tag, "modules", "mods")
	if tmp != "" {
		out.Modules = tmp
	}

	tmp = getTagValue(tag, "cache")
	if tmp != "" {
		out.Cache = tmp
	}

	tmp = getTagValue(tag, "ttl")
	if tmp != "" {
		out.Ttl = tmp
	}

	tmp = getTagValue(tag, "log")
	if tmp != "" {
		out.Log, _ = strconv.ParseBool(tmp)
	}

	tmp = getTagValue(tag, "stub")
	if tmp != "" {
		out.Stub = tmp
	}

	tmp = getTagValue(tag, "auth")
	if tmp != "" {
		out.Auth = tmp
	}

	return out
}

// Get the first non-empty matching tag value for given variations of a key
func getTagValue(tag reflect.StructTag, keys ...string) string {
	for _, key := range keys {
		if tag.Get(key) != "" {
			return tag.Get(key)
		}
	}
	return ""
}

func resolveInOrder(e ...Fixture) Fixture {
	out := Fixture{}
	empty := ""

	for _, ep := range e {
		if out.Prefix == empty && ep.Prefix != empty {
			out.Prefix = ep.Prefix
		}
		if out.Root == empty && ep.Root != empty {
			out.Root = ep.Root
		}
		if out.Url == empty && ep.Url != empty {
			out.Url = ep.Url
		}
		if out.Version == empty && ep.Version != empty {
			out.Version = ep.Version
		}
		if out.Pretty == empty && ep.Pretty != empty {
			out.Pretty = ep.Pretty
		}
		if out.Vendor == empty && ep.Vendor != empty {
			out.Vendor = ep.Vendor
		}
		if out.Modules == empty && ep.Modules != empty {
			out.Modules = ep.Modules
		}
		if out.Cache == empty && ep.Cache != empty {
			out.Cache = ep.Cache
		}
		if out.Ttl == empty && ep.Ttl != empty {
			out.Ttl = ep.Ttl
		}
		if out.Log == false && ep.Log == true {
			out.Log = ep.Log
		}
		if out.Stub == empty && ep.Stub != empty {
			out.Stub = ep.Stub
		}
		if out.Auth == empty && ep.Auth != empty {
			out.Auth = ep.Auth
		}
	}
	return out
}
