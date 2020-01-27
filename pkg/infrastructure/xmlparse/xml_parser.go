package xmlparse

import (
	"encoding/xml"
	"github.com/phodal/coca/pkg/infrastructure/container"
	"io"
	"strings"
)

type ElemType string

const (
	eleTpText ElemType = "text" // 静态文本节点
	eleTpNode ElemType = "XMLNode" // 节点子节点
)

type XMLNode struct {
	ID       string
	Name     string
	Attrs    map[string]xml.Attr
	Elements []element
}

type element struct {
	ElementType ElemType
	Val         interface{}
}

func ParseXML(r io.Reader) *XMLNode {
	parser := xml.NewDecoder(r)
	var root XMLNode

	st := container.NewStack()
	for {
		token, err := parser.Token()
		if err != nil {
			break
		}
		switch t := token.(type) {
		case xml.StartElement: //tag start
			name := t.Name.Local
			attr := t.Attr
			attrMap := make(map[string]xml.Attr)
			for _, val := range attr {
				attrMap[val.Name.Local] = val
			}
			node := XMLNode{
				Name:     name,
				Attrs:    attrMap,
				Elements: make([]element, 0),
			}
			for _, val := range attr {
				if val.Name.Local == "id" {
					node.ID = val.Value
				}
			}
			st.Push(node)

		case xml.EndElement: //tag end
			if st.Len() > 0 {
				//cur XMLNode
				n := st.Pop().(XMLNode)
				if st.Len() > 0 { //if the root XMLNode then append to element
					e := element{
						ElementType: eleTpNode,
						Val:         n,
					}

					pn := st.Pop().(XMLNode)
					els := pn.Elements
					els = append(els, e)
					pn.Elements = els
					st.Push(pn)
				} else { //else root = n
					root = n
				}
			}
		case xml.CharData: //tag content
			if st.Len() > 0 {
				n := st.Pop().(XMLNode)
				content := strings.TrimSpace(string(t))
				if content != "" {
					e := element{
						ElementType: eleTpText,
						Val:         content,
					}
					els := n.Elements
					els = append(els, e)
					n.Elements = els
				}

				st.Push(n)
			}

		case xml.Comment:
		case xml.ProcInst:
		case xml.Directive:
		default:
		}
	}

	if st.Len() != 0 {
		panic("Parse xml error, there is tag no close, please check your xml config!")
	}

	return &root
}
