package lsmod

// ModInfo contains a module info provided by /proc/modules
type ModInfo struct {
	Depends   []string `json:"depends"`
	Mem       uint64   `json:"memory"`
	Instances uint64
	Offset    uint64
	Taineds   []modTained `json:"taineds"`
	State     modState    `json:"state"`
}
