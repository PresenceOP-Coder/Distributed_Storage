package loadbalancer

import (
	"fmt"
	"sync"
)

type Node struct {
	URL   string
	Alive bool
	Load  int
}

type Balancer struct {
	mu    sync.Mutex
	nodes []Node
}

func New(urls []string) *Balancer {
	nodes := make([]Node, 0, len(urls))
	for _, u := range urls {
		nodes = append(nodes, Node{URL: u, Alive: true})
	}
	return &Balancer{nodes: nodes}
}

func (b *Balancer) GetBestNode() (string, error) {
	b.mu.Lock()
	defer b.mu.Unlock()

	bestIdx := -1
	for i, n := range b.nodes {
		if !n.Alive {
			continue
		}
		if bestIdx == -1 || n.Load < b.nodes[bestIdx].Load {
			bestIdx = i
		}
	}
	if bestIdx == -1 {
		return "", fmt.Errorf("no alive nodes available")
	}
	return b.nodes[bestIdx].URL, nil
}

func (b *Balancer) TrackLoad(nodeURL string, delta int) {
	b.mu.Lock()
	defer b.mu.Unlock()

	for i := range b.nodes {
		if b.nodes[i].URL == nodeURL {
			b.nodes[i].Load += delta
			if b.nodes[i].Load < 0 {
				b.nodes[i].Load = 0
			}
			return
		}
	}
}

func (b *Balancer) GetReplicaNodes(startURL string, replicationFactor int) []string {
	b.mu.Lock()
	defer b.mu.Unlock()

	if len(b.nodes) == 0 || replicationFactor <= 0 {
		return nil
	}

	start := 0
	for i, n := range b.nodes {
		if n.URL == startURL {
			start = i
			break
		}
	}

	selected := make([]string, 0, replicationFactor)
	for i := 0; i < len(b.nodes) && len(selected) < replicationFactor; i++ {
		node := b.nodes[(start+i)%len(b.nodes)]
		if node.Alive {
			selected = append(selected, node.URL)
		}
	}
	return selected
}

func (b *Balancer) NodeURLs() []string {
	b.mu.Lock()
	defer b.mu.Unlock()

	urls := make([]string, 0, len(b.nodes))
	for _, n := range b.nodes {
		urls = append(urls, n.URL)
	}
	return urls
}

func (b *Balancer) MarkNodeAlive(nodeURL string, alive bool) {
	b.mu.Lock()
	defer b.mu.Unlock()

	for i := range b.nodes {
		if b.nodes[i].URL == nodeURL {
			b.nodes[i].Alive = alive
			return
		}
	}
}
