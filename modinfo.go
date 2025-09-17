package lsmod

// ModInfo contains a module info provided by /proc/modules
type ModInfo struct {
	Depends   []string    `json:"depends"`
	Mem       uint64      `json:"memory"`
	Instances uint64      `json:-`
	Offset    uint64      `json:-`
	Taineds   []modTained `json:"taints"`
	State     modState    `json:"state"`
}
