package xml

import (
	"bytes"
	"fmt"
	"io"
	"sort"
	"strings"
	"sync/atomic"
)

type (
	XML struct {
		ns    map[string]string
		attr  []any
		tags  map[string]*_tag
		err   error
		index int64
	}

	_tag struct {
		ns    string
		tag   string
		value string
		attr  []any
		tags  map[string]*_tag
	}
)

func New() *XML {
	return &XML{
		ns:   make(map[string]string, 10),
		tags: make(map[string]*_tag, 10),
	}
}

func (v *XML) NS(ns ...string) {
	if len(ns)%2 == 1 {
		ns = append(ns, "")
	}
	for i := 0; i < len(ns); i += 2 {
		v.ns[ns[i]] = ns[i+1]
	}
}

func (v *XML) NSAttr(attr ...any) {
	if len(attr)%2 == 1 {
		attr = append(attr, "")
	}
	v.attr = append(v.attr, attr...)
}

func getTag(tags map[string]*_tag, ns, path string, index int64) *_tag {
	paths := strings.Split(strings.Trim(path, "/"), "/")
	var (
		next *_tag
		ok   bool
	)
	end := len(paths) - 1
	for i, tag := range paths {
		origTag := tag
		if i == end {
			tag = fmt.Sprintf("%s+%d", tag, index)
		}
		next, ok = tags[ns+":"+tag]
		if !ok {
			next = &_tag{
				ns:    ns,
				tag:   origTag,
				value: "",
				attr:  nil,
				tags:  make(map[string]*_tag, 10),
			}
			tags[ns+":"+tag] = next
		}
		tags = next.tags
	}
	return next
}

func (v *XML) Group(ns, path string, attr ...any) *Group {
	if _, ok := v.ns[ns]; !ok {
		v.err = fmt.Errorf("ns `%s` not foud for: %s", ns, path)
	}
	if v.err != nil {
		return &Group{err: v.err}
	}
	if len(attr)%2 == 1 {
		attr = append(attr, "")
	}
	index := atomic.AddInt64(&v.index, 1)
	next := getTag(v.tags, ns, path, index)
	next.attr = attr
	return &Group{
		ns:    ns,
		tag:   next,
		index: index,
	}
}

func (v *XML) Tag(ns, path, value string, attr ...any) {
	if v.err != nil {
		return
	}
	if len(attr)%2 == 1 {
		attr = append(attr, "")
	}
	if _, ok := v.ns[ns]; !ok {
		v.err = fmt.Errorf("ns `%s` not foud for: %s", ns, path)
		return
	}
	next := getTag(v.tags, ns, path, 0)
	next.value = value
	next.attr = attr
}

func (v *XML) Write(wo io.Writer) (err error) {
	if v.err != nil {
		return v.err
	}
	w := bytes.NewBuffer(make([]byte, 0, 1024))
	if _, err = fmt.Fprintln(w, "<?xml version=\"1.0\" encoding=\"UTF-8\" standalone=\"yes\"?>"); err != nil {
		return
	}
	keys := make([]string, 0, len(v.tags))
	for key := range v.tags {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	for _, key := range keys {
		if err = v.writeTag(w, v.tags[key], 0); err != nil {
			return
		}
	}
	_, err = w.WriteTo(wo)
	return
}

func (v *XML) writeTag(w io.Writer, tag *_tag, level int) (err error) {
	tab := strings.Repeat(" ", level*2)
	if _, err = fmt.Fprintf(w, "%s<%s:%s", tab, tag.ns, tag.tag); err != nil {
		return
	}
	if attrCount := len(tag.attr); attrCount > 0 {
		for i := 0; i < attrCount; i += 2 {
			if _, err = fmt.Fprintf(w, " %s:%s=\"%v\"", tag.ns, tag.attr[i], tag.attr[i+1]); err != nil {
				return
			}
		}
	}
	if level == 0 {
		keys := make([]string, 0, len(v.ns))
		for key := range v.ns {
			keys = append(keys, key)
		}
		sort.Strings(keys)
		for _, key := range keys {
			if _, err = fmt.Fprintf(w, " xmlns:%s=\"%v\"", key, v.ns[key]); err != nil {
				return
			}
		}
		if attrCount := len(v.attr); attrCount > 0 {
			for i := 0; i < attrCount; i += 2 {
				if _, err = fmt.Fprintf(w, " %s=\"%v\"", v.attr[i], v.attr[i+1]); err != nil {
					return
				}
			}
		}
	}
	space := ""
	if len(tag.attr) > 0 {
		space = " "
	}
	if len(tag.tags) == 0 && len(tag.value) == 0 {
		_, err = fmt.Fprintf(w, "%s/>\n", space)
		return
	}

	if len(tag.value) > 0 {
		_, err = fmt.Fprintf(w, "%s>%s</%s:%s>\n", space, tag.value, tag.ns, tag.tag)
	} else {
		if _, err = fmt.Fprintf(w, "%s>\n", space); err != nil {
			return
		}
		keys := make([]string, 0, len(tag.tags))
		for key := range tag.tags {
			keys = append(keys, key)
		}
		sort.Strings(keys)
		for _, key := range keys {
			if err = v.writeTag(w, tag.tags[key], level+1); err != nil {
				return
			}
		}
		_, err = fmt.Fprintf(w, "%s%s</%s:%s>\n", tab, tag.value, tag.ns, tag.tag)
	}
	return
}
