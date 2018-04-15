package join_ring

import "fmt"
import chord "../utils/node_defs"

// Gets sponsoring node ID to lookup
// Node object is the node that wants to join
func Join_ring(sponsoring_node_id int64, node *chord.Node){
	
    node.Predecessor = nil
    fmt.Printf("Node %d is joining the ring now", sponsoring_node_id)
    //The &node is the node that needs a successor
    //node.Successor = find_successor(sponsoring_node_id, &node)

}