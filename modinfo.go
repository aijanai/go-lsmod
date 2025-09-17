package lsmod

// ModInfo contains a module info provided by /proc/modules
type ModInfo struct {
	Depends   []string    `json:"depends"`
	Mem       uint64      `json:"mem"`
	Instances uint64      `json:"instances"`
	Offset    uint64      `json:"offset"`
	Taineds   []modTained `json:"taints"`
	State     modState    `json:"state"`
}
