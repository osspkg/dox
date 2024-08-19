package xml_test

import (
	"bytes"
	"fmt"
	"testing"

	"go.osspkg.com/dox/xml"
)

func TestUnit_Write(t *testing.T) {
	x := xml.New()
	x.NS("a", "aaa/aaa", "b", "bbb/bbb")
	x.NSAttr("b:ignore", "a")
	x.Tag("a", "root/aaa/bbb", "", "a1", 1, "a2", "aaa")
	x.Tag("a", "root/aaa/ccc", "hello", "c1", 5, "c2", "ccc")
	x.Tag("a", "root/aaa/ccc/ddd", "", "d1", 3, "d2", "ddd")
	x.Group("a", "root/www").Tag("vvv", "", "vvv1", 1)
	x.Group("a", "root/www").Tag("vvv", "", "vvv2", 1)
	x.Group("a", "root/www", "attr", 1).Tag("vvv", "", "vvv2", 1)
	var b bytes.Buffer
	if err := x.Write(&b); err != nil {
		t.Error(err)
	}

	want := `<?xml version="1.0" encoding="UTF-8" standalone="yes"?>
<a:root xmlns:a="aaa/aaa" xmlns:b="bbb/bbb" b:ignore="a">
  <a:aaa>
    <a:bbb a:a1="1" a:a2="aaa" />
    <a:ccc>
      <a:ddd a:d1="3" a:d2="ddd" />
    </a:ccc>
    <a:ccc a:c1="5" a:c2="ccc" >hello</a:ccc>
  </a:aaa>
  <a:www>
    <a:vvv a:vvv1="1" />
  </a:www>
  <a:www>
    <a:vvv a:vvv2="1" />
  </a:www>
  <a:www a:attr="1" >
    <a:vvv a:vvv2="1" />
  </a:www>
</a:root>
`

	if b.String() != want {
		fmt.Println(b.String())
		t.FailNow()
	}
}
