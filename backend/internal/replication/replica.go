package replication

import "dfs/internal/loadbalancer"

func SelectReplicaNodes(lb *loadbalancer.Balancer, primary string, replicationFactor int) []string {
	return lb.GetReplicaNodes(primary, replicationFactor)
}
