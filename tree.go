package main

import (
	"log"
	"sort"

	"github.com/lxn/walk"
	"github.com/oriath-net/pogo/ggpk"
)

type ggpkTreeModel struct {
	walk.TreeModelBase
}

func (m *ggpkTreeModel) LazyPopulation() bool { return true }

func (m *ggpkTreeModel) RootCount() int {
	if app.ggpk == nil {
		return 0
	} else {
		return 1
	}
}

func (m *ggpkTreeModel) RootAt(index int) walk.TreeItem {
	n, err := app.ggpk.RootNode()
	if err != nil {
		log.Fatal("failed to get root node: ", err)
	}
	return &ggpkTreeItem{
		node: n,
	}
}

type ggpkTreeItem struct {
	parent   *ggpkTreeItem
	node     ggpk.AnyNode
	children []ggpk.AnyNode
}

func (ti *ggpkTreeItem) Text() string {
	if ti.parent == nil { // root node
		return app.ggpkFilename
	}
	return ti.node.Name()
}

func (ti *ggpkTreeItem) Parent() walk.TreeItem {
	if ti.parent == nil {
		return nil
	}
	return ti.parent
}

func (ti *ggpkTreeItem) loadChildren() {
	if ti.node.Type() != "PDIR" {
		ti.children = []ggpk.AnyNode{}
		return
	}

	children, err := ti.node.(*ggpk.DirectoryNode).Children()
	if err != nil {
		log.Println("failed to get children: ", err)
		return
	}

	sort.Slice(children, func(a, b int) bool {
		return children[a].Name() < children[b].Name()
	})

	ti.children = children
}

func (ti *ggpkTreeItem) ChildCount() int {
	if ti.children == nil {
		ti.loadChildren()
	}
	return len(ti.children)
}

func (ti *ggpkTreeItem) ChildAt(pos int) walk.TreeItem {
	if ti.children == nil {
		ti.loadChildren()
	}

	n := ti.children[pos]
	return &ggpkTreeItem{
		node:   n,
		parent: ti,
	}
}
