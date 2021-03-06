package decisions

import (
	"bytes"
	"fmt"
	"image/color"
)

// Node includes an Action or Condition value
type Node struct {
	Color             color.RGBA
	ID                string
	NodeType          interface{}
	YesNode           *Node
	NoNode            *Node
	Metrics           map[Metric]float64
	MetricsAvgs       map[Metric]float64
	Uses              float64
	NumOrganismsUsing int
	Complexity        int
	UsedYes, UsedNo   bool
}

// IsAction returns true if Node's type is Action (false if Condition)
func (n *Node) IsAction() bool {
	return isAction(n.NodeType)
}

// IsCondition returns true if Node's type is Action (false if Condition)
func (n *Node) IsCondition() bool {
	return isCondition(n.NodeType)
}

// UpdateStats updates all Node Metrics according to a map of changes and
// increments number of Uses
func (n *Node) UpdateStats(metricsChange map[Metric]float64) {
	n.Uses++
	for key, change := range metricsChange {
		n.Metrics[key] += change
		uses := n.Uses
		n.MetricsAvgs[key] = (n.MetricsAvgs[key]*(uses-1.0) + change) / uses
	}
}

// UpdateNodeIDs sets a Node's ID to a hyphen-separated string listing its
// decision tree in serialized form.
//
// Recursively walks through Node tree updating ID for itself and all children.
func (n *Node) UpdateNodeIDs() string {
	var buffer bytes.Buffer
	nodeTypeString := fmt.Sprintf("%v", n.NodeType)
	buffer.WriteString(nodeTypeString)
	if !isAction(n.NodeType) {
		buffer.WriteString("-")
		buffer.WriteString(n.YesNode.UpdateNodeIDs())
		buffer.WriteString("-")
		buffer.WriteString(n.NoNode.UpdateNodeIDs())
	}
	n.ID = buffer.String()
	return n.ID
}

// UpdateNumOrganismsUsing updates the current number of organisms using this
// node (by +1 or -1), recursively calling for all sub-trees
func (n *Node) UpdateNumOrganismsUsing(change int) {
	n.NumOrganismsUsing += change
	if !isAction(n.NodeType) {
		n.YesNode.UpdateNumOrganismsUsing(change)
		n.NoNode.UpdateNumOrganismsUsing(change)
	}
}

// TreeFromAction creates a simple Node object from an Action type
func TreeFromAction(action Action) Node {
	node := Node{
		NodeType: action,
		Uses:     0,
		YesNode:  nil,
		NoNode:   nil,
		UsedYes:  false,
		UsedNo:   false,
	}
	node.Metrics = InitializeMetricsMap()
	node.MetricsAvgs = InitializeMetricsMap()
	node.UpdateNodeIDs()
	return node
}
