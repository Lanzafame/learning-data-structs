// Tree
package tree

type Node int

type Tree struct {
	Root   *Node
	LChild *Tree
	RChild *Tree
}

func (t *Tree) AddChild(c *Tree) {

}
