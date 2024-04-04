package arm64

import (
	"fmt"
	"strings"
)

func (i *Instruction) annotate() (string, error) {
	var annotation strings.Builder
	output := false
	if i.readRegs != 0 {
		regs := make([]string, 0)
		annotation.WriteString("r:")
		if i.readRegs == RWREGS_ALL {
			annotation.WriteString("all")

		} else {
			for reg := 0; reg < 32; reg++ {
				if ((uint64(1) << reg) & i.readRegs) != 0 {
					regs = append(regs, fmt.Sprintf("%d", reg))
				}
			}
			if (i.readRegs & RWREGS_STATUS) != 0 {
				regs = append(regs, "nzcv")
			}
		}
		annotation.WriteString(strings.Join(regs, ","))
		output = true
	}
	if i.writeRegs != 0 {
		if output {
			annotation.WriteString(", ")
		}
		regs := make([]string, 0)
		annotation.WriteString("w:")
		if i.writeRegs == RWREGS_ALL {
			annotation.WriteString("all")

		} else {
			for reg := 0; reg < 32; reg++ {
				if ((uint64(1) << reg) & i.writeRegs) != 0 {
					regs = append(regs, fmt.Sprintf("%d", reg))
				}
			}
			if (i.writeRegs & RWREGS_STATUS) != 0 {
				regs = append(regs, "nzcv")
			}
		}
		annotation.WriteString(strings.Join(regs, ","))
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
