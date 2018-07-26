package main

import (
	"errors"
	"fmt"
	"hash/fnv"
	"io"
	"math"
	"sort"
)

const (
	replica_format = "R%d_%s"
	max_keys       = 1000000
)

var fnv_hash = fnv.New32()

type HashFunction (func(string) int)

type Server struct {
	Name string
	Data []interface{}
}

func (server *Server) String() string {
	return fmt.Sprintf("%s:%d", server.Name, len(server.Data))
}

type CHPool struct {
	Hasher            HashFunction
	server_keys       []int
	key_to_server_map map[int]*Server
}

func bad_hasher(str string) int {
	result := 0
	for _, c := range str {
		result = 31*result + int(c)
	}
	return result
}

func fnv_hasher(str string) int {
	io.WriteString(fnv_hash, str)
	return int(fnv_hash.Sum32())
}

func NewConsistentHashingPool(servers []string, replication_factor uint) (*CHPool, error) {
	if len(servers) == 0 {
		return nil, errors.New("List of servers cannot be empty")
	}
	if replication_factor == 0 {
		return nil, errors.New("Replication factor must be greater than 0")
	}
	chpool := &CHPool{
		Hasher:            fnv_hasher,
		key_to_server_map: make(map[int]*Server),
	}
	for _, server := range servers {
		chpool.AddServer(server, replication_factor)
	}
	return chpool, nil
}

func (pool *CHPool) PrintState() {
	fmt.Println(pool.key_to_server_map)
	avg := 0
	num_servers := len(pool.server_keys)
	for _, server := range pool.key_to_server_map {
		avg += len(server.Data) / num_servers
	}
	std_dev := 0
	for _, server := range pool.key_to_server_map {
		diff := len(server.Data) - avg
		std_dev += diff * diff
	}
	result := math.Sqrt(float64(std_dev / num_servers))
	fmt.Printf("Standard Deviation: %v\n", result)
}

func (pool *CHPool) AddServer(server string, replication_factor uint) {
	for i := 0; i < int(replication_factor); i++ {
		replica := fmt.Sprintf(replica_format, i, server)
		replica_key := pool.Hasher(replica)
		pool.server_keys = append(pool.server_keys, replica_key)
		pool.key_to_server_map[replica_key] = &Server{Name: server}
	}
	sort.Ints(pool.server_keys)
}

func (pool *CHPool) GetServer(key string) string {
	hash_code := pool.Hasher(key)
	num_servers := len(pool.server_keys)
	server_index := sort.Search(num_servers, func(i int) bool {
		return pool.server_keys[i] >= hash_code
	})
	if server_index == num_servers {
		server_index = 0
	}
	target_server_key := pool.server_keys[server_index]
	target_server := pool.key_to_server_map[target_server_key]
	target_server.Data = append(target_server.Data, key)
	return target_server.Name
}

func main() {
	servers := []string{"ip1:port1", "ip2:port2", "ip3:port3"}
	var replication_factor uint = 3
	if pool, err := NewConsistentHashingPool(servers, replication_factor); err == nil {
		for i := 0; i < max_keys; i++ {
			key := fmt.Sprintf("Key_%d", i)
			pool.GetServer(key)
		}
		pool.PrintState()
	} else {
		panic(err)
	}
}
