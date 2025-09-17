package lsmod

import (
	"fmt"
	"strconv"
	"strings"
)

type modTained uint32

// possibble tained flags of the module
const (
	TainedNone modTained = 0      // No tained flag
	TainedP    modTained = 1      // (P): A module with a non-GPL license has been loaded, this includes modules with no license. Set by modutils >= 2.4.9 and module-init-tools.
	TainedF    modTained = 2      // (F): A module was force loaded by insmod -f. Set by modutils >= 2.4.9 and module-init-tools.
	TainedS    modTained = 4      // (S): Unsafe SMP processors: SMP with CPUs not designed for SMP.
	TainedR    modTained = 8      // (R): A module was forcibly unloaded from the system by rmmod -f.
	TainedM    modTained = 16     // (M): A hardware machine check error occurred on the system.
	TainedB    modTained = 32     // (B): A bad page was discovered on the system.
	TainedU    modTained = 64     // (U): The user has asked that the system be marked "tainted". This could be because they are running software that directly modifies the hardware, or for other reasons.
	TainedD    modTained = 128    // (D): The system has died.
	TainedA    modTained = 256    // (A): The ACPI DSDT has been overridden with one supplied by the user instead of using the one provided by the hardware.
	TainedW    modTained = 512    // (W): A kernel warning has occurred.
	TainedC    modTained = 1024   // (C): A module from drivers/staging was loaded.
	TainedI    modTained = 2048   // (I): The system is working around a severe firmware bug.
	TainedO    modTained = 4096   // (O): An out-of-tree module has been loaded.
	TainedE    modTained = 8192   // (E): An unsigned module has been loaded in a kernel supporting module signature.
	TainedL    modTained = 16384  // (L): A soft lockup has previously occurred on the system.
	TainedK    modTained = 32768  // (K): The kernel has been live patched.
	TainedX    modTained = 65536  // (X): Auxiliary taint, defined and used by for distros.
	TainedT    modTained = 131072 // (T): The kernel was built with the struct randomization plugin.
)

func (t modTained) ToStringDescription() string {
	codeLookup := map[string]string{
		"1":      "proprietary module was loaded",
		"2":      "module was force loaded",
		"4":      "kernel running on an out of specification system",
		"8":      "module was force unloaded",
		"16":     "processor reported a Machine Check Exception (MCE)",
		"32":     "bad page referenced or some unexpected page flags",
		"64":     "taint requested by userspace application",
		"128":    "kernel died recently, i.e. there was an OOPS or BUG",
		"256":    "ACPI table overridden by user",
		"512":    "kernel issued warning",
		"1024":   "staging driver was loaded",
		"2048":   "workaround for bug in platform firmware applied",
		"4096":   "externally-built (“out-of-tree”) module was loaded",
		"8192":   "unsigned module was loaded",
		"16384":  "soft lockup occurred",
		"32768":  "kernel has been live patched",
		"65536":  "auxiliary taint, defined for and used by distros",
		"131072": "kernel was built with the struct randomization plugin",
		"262144": "an in-kernel test has been run",
		"524288": "userspace used a mutating debug operation in fwctl",
	}
	return codeLookup[strconv.Itoa(int(t))]
}

func parseTained(tg string) ([]modTained, error) { // nolint: gocyclo
	result := []modTained{}
	tTrim := strings.Trim(tg, "()")
	for i := 0; i < len(tTrim); i++ {
		t := string(tTrim[i])
		switch t {
		case "P":
			result = append(result, TainedP)
		case "F":
			result = append(result, TainedF)
		case "S":
			result = append(result, TainedS)
		case "R":
			result = append(result, TainedR)
		case "M":
			result = append(result, TainedM)
		case "B":
			result = append(result, TainedB)
		case "U":
			result = append(result, TainedU)
		case "D":
			result = append(result, TainedD)
		case "A":
			result = append(result, TainedA)
		case "W":
			result = append(result, TainedW)
		case "C":
			result = append(result, TainedC)
		case "I":
			result = append(result, TainedI)
		case "O":
			result = append(result, TainedO)
		case "E":
			result = append(result, TainedE)
		case "L":
			result = append(result, TainedL)
		case "K":
			result = append(result, TainedK)
		case "X":
			result = append(result, TainedX)
		case "T":
			result = append(result, TainedT)
		}
	}

	if len(result) == 0 {
		return nil, fmt.Errorf("unknown tained flag %q", tg)
	}

	return result, nil
}
