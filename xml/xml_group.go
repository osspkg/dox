package xml

type Group struct {
	ns    string
	tag   *_tag
	err   error
	index int64
}

func (v *Group) Tag(path, value string, attr ...any) {
	next := getTag(v.tag.tags, v.ns, path, v.index)
	next.value = value
	next.attr = attr
}
