package geecache

type PeerPicker interface {
	PeerPicker(key string) (peer PeerGetter, ok bool)
}

type PeerGetter interface {
	Get(group, key string) ([]byte, error)
}

func (p *HTTPPool) Set(peers ...string) {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.peers = NewMap(defaultReplicas, nil)
	p.peers.Add(peers...)
	p.httpGetters = make(map[string]*httpGetter, len(peers))
	for _, peer := range peers {
		p.httpGetters[peer] = &httpGetter{baseURL: peer + p.basePath}
	}
}

func (p *HTTPPool) PickPeer(key string) (PeerGetter, bool) {
	p.mu.Lock()
	defer p.mu.Unlock()
	if peer := p.peers.Get(key); peer != "" && peer != p.self {
		p.Log("Pick peer %s ", peer)
		return p.httpGetters[peer], true
	}
	return nil, false
}

//var _ PeerPicker = (*HTTPPool)(nil)
