package arm64

import (
	"fmt"
	"strings"
)

func GetRegString(regset uint64) string {
	regs := make([]string, 0)
	if regset == RWREGS_ALL {
		return "all"

	} else {
		for reg := 0; reg < 32; reg++ {
			if ((uint64(1) << reg) & regset) != 0 {
				regs = append(regs, fmt.Sprintf("%d", reg))
			}
		}
		if (regset & RWREGS_STATUS) != 0 {
			regs = append(regs, "nzcv")
		}
	}
	if len(regs) == 0 {
		return "none"
	}
	return strings.Join(regs, ",")
}

func (i *Instruction) annotate() (string, error) {
	var annotation strings.Builder
	output := false
	if i.readRegs != 0 {
		annotation.WriteString("r:" + GetRegString(i.readRegs))
		output = true
	}
	if i.writeRegs != 0 {
		if output {
			annotation.WriteString(", ")
		}
		annotation.WriteString("w:" + GetRegString(i.writeRegs))
		output = true
	}
	if i.branchType != BranchTypeNone {
		if output {
			annotation.WriteString(", ")
		}
		annotation.WriteString("b:")
		switch i.branchType {
		case BranchTypeCall:
			annotation.WriteString("call")
		case BranchTypeCond:
			annotation.WriteString("cond")
		case BranchTypeUncond:
			annotation.WriteString("uncond")
		case BranchTypeException:
			annotation.WriteString("exception")
		}
		output = true
	}
	if BranchType(i.branchTargetAddr) != 0 {
		if output {
			annotation.WriteString(", ")
		}
		annotation.WriteString(fmt.Sprintf("t:%#x", i.branchTargetAddr))
	}
	return annotation.String(), nil
}
