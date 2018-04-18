package Landscape

import (
  "github.com/go-gl/gl/v4.1-core/gl"
  "github.com/go-gl/mathgl/mgl32"
)

/*
======================
Node
======================
*/

type Node struct {
  parent *Node
  children []*Node
  patch *Patch
  is_leaf bool

  world_transform, local_transform *mgl32.Mat4
}

func (node *Node) Init(parent *Node, patch *Patch) *Node {
  node.parent = parent
  node.patch = patch

  return node
}

func (node *Node) add_child(child *Node) {
  child.parent = node
  node.children = append(node.children, child)
}

func (node *Node) update(elapsed float32) {
  for _, child := range node.children {
    child.update(elapsed)
  }
}

func (node *Node) render() {
  if node.is_leaf {
    node.patch.render()
  } else {
    for _, child := range node.children {
      child.render()
    }
  }
}

/*
======================
Quadtree
======================
*/

type Quadtree struct {
  nodes_amt int
  shader uint32

  root_node *Node
}

func (quad *Quadtree) Init(shader uint32) *Quadtree {
  quad.shader = shader
  quad.nodes_amt = 8
  for i := float32(0); i < float32(quad.nodes_amt); i++ {
    for j := float32(0); j < float32(quad.nodes_amt); j++ {
      new_patch := new(Patch).Init(mgl32.Vec2{i/float32(quad.nodes_amt), j/float32(quad.nodes_amt)}, 0, mgl32.Vec2{i, j}, quad.shader)
      quad.root_node = new(Node).Init(nil, new_patch)
    }
  }

  return quad
}

func (quad *Quadtree) Update(elapsed float32) {
  quad.root_node.update(elapsed)
}

func (quad *Quadtree) Render() {
  gl.UseProgram(quad.shader)
  quad.root_node.render()
}
