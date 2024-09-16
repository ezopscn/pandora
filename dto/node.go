package dto

// 节点信息
type NodeStatus struct {
	NodeName string `json:"node_name"`
	NodeRole string `json:"node_role"`
	LiveTime string `json:"live_time"`
}
