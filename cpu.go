package main

type AddrMode uint8

const (
	_     AddrMode = iota
	ACC            // Accumulator
	ABS            // Absolute
	ABSX           // Absolute, X-indexed
	ABSY           // Absolute, Y-indexed
	IMM            // Immediate
	IMPL           // Implied
	IND            // Indirect
	INDX           // Indirect, X-indexed
	INDY           // Indirect, Y-indexed
	REL            // Relative
	ZERO           // Zeropage
	ZEROX          // Zeropage, X-indexed
	ZEROY          // Zeropage, Y-indexed
)

type Instruction struct {
	name     string
	addrMode AddrMode
	size     uint8
	cycle    uint8
}

var instructions = map[uint8]Instruction{
	0x00: {"BRK", IMM, 1, 7},
	0x01: {"ORA", INDX, 2, 6},
	0x05: {"ORA", ZERO, 2, 3},
	0x06: {"ASL", ZERO, 2, 5},
	0x08: {"PHP", IMPL, 1, 3},
	0x09: {"ORA", IMM, 2, 2},
	0x0a: {"ASL", ACC, 1, 2},
	0x0d: {"ORA", ABS, 3, 4},
	0x0e: {"ASL", ABS, 3, 6},

	0x10: {"BPL", REL, 2, 2}, // cycle:3 ?
	0x11: {"ORA", INDY, 2, 5},
	0x15: {"ORA", ZEROX, 2, 4},
	0x16: {"ASL", ZEROX, 2, 6},
	0x18: {"CLC", IMPL, 1, 2},
	0x19: {"ORA", ABSY, 3, 4},
	0x1d: {"ORA", ABSX, 3, 4},
	0x1e: {"ASL", ABSX, 3, 7},

	0x20: {"JSR", ABS, 3, 6},
	0x21: {"AND", INDX, 2, 6},
	0x24: {"BIT", ZERO, 2, 3},
	0x25: {"AND", ZERO, 2, 3},
	0x26: {"ROL", ZERO, 2, 5},
	0x28: {"PLP", IMPL, 1, 4},
	0x29: {"AND", IMM, 2, 2},
	0x2a: {"ROL", ACC, 1, 2},
	0x2c: {"BIT", ABS, 3, 4},
	0x2d: {"AND", ABS, 3, 4},
	0x2e: {"ROL", ABS, 3, 6},

	0x30: {"BMI", REL, 2, 2}, // cycle:3 ?
	0x31: {"AND", INDY, 2, 5},
	0x35: {"AND", ZEROX, 2, 4},
	0x36: {"ROL", ZEROX, 2, 6},
	0x38: {"SEC", IMPL, 1, 2},
	0x39: {"AND", ABSY, 3, 4},
	0x3d: {"AND", ABSX, 3, 4},
	0x3e: {"ROL", ABSX, 3, 7},

	0x40: {"RTI", IMPL, 1, 6},
	0x41: {"EOR", INDX, 2, 6},
	0x45: {"EOR", ZERO, 2, 3},
	0x46: {"LSR", ZERO, 2, 5},
	0x48: {"PHA", IMPL, 1, 3},
	0x49: {"EOR", IMM, 2, 2},
	0x4a: {"LSR", ACC, 1, 2},
	0x4c: {"JMP", ABS, 3, 3},
	0x4d: {"EOR", ABS, 3, 4},
	0x4e: {"LSR", ABS, 3, 6},

	0x50: {"BVC", REL, 2, 2}, // cycle:3 ?
	0x51: {"EOR", INDY, 2, 5},
	0x55: {"EOR", ZEROX, 2, 4},
	0x56: {"LSR", ZEROX, 2, 6},
	0x58: {"CLI", IMPL, 1, 2},
	0x59: {"EOR", ABSY, 3, 4},
	0x5d: {"EOR", ABSX, 3, 4},
	0x5e: {"LSR", ABSX, 3, 7},

	0x60: {"RTS", IMPL, 1, 6},
	0x61: {"ADC", INDX, 2, 6},
	0x65: {"ADC", ZERO, 2, 3},
	0x66: {"ROR", ZERO, 2, 5},
	0x68: {"PLA", IMPL, 1, 4},
	0x69: {"ADC", IMM, 2, 2},
	0x6a: {"ROR", ACC, 1, 2},
	0x6c: {"JMP", IND, 3, 5},
	0x6d: {"ADC", ABS, 3, 4},
	0x6e: {"ROR", ABSX, 3, 7},

	0x70: {"BVS", REL, 2, 2}, // cycle:3 ?
	0x71: {"ADC", INDY, 2, 5},
	0x75: {"ADC", ZEROX, 2, 4},
	0x76: {"ROR", ZEROX, 2, 6},
	0x78: {"SEI", IMPL, 1, 2},
	0x79: {"ADC", ABSY, 3, 4},
	0x7d: {"ADC", ABSX, 3, 4},
	0x7e: {"ROR", ABS, 3, 6},

	0x81: {"STA", INDX, 2, 6},
	0x84: {"STY", ZERO, 2, 3},
	0x85: {"STA", ZERO, 2, 3},
	0x86: {"STX", ZERO, 2, 3},
	0x88: {"DEY", IMPL, 1, 2},
	0x8a: {"TXA", IMPL, 1, 2},
	0x8c: {"STY", ABS, 3, 4},
	0x8d: {"STA", ABS, 3, 4},
	0x8e: {"STX", ABS, 3, 4},

	0x90: {"BCC", REL, 2, 2}, // cycle:3 ?
	0x91: {"STA", INDY, 2, 6},
	0x94: {"STY", ZEROX, 2, 4},
	0x95: {"STA", ZEROX, 2, 4},
	0x96: {"STX", ZEROY, 2, 4},
	0x98: {"TYA", IMPL, 1, 2},
	0x99: {"STA", ABSY, 3, 5},
	0x9a: {"TXS", IMPL, 1, 2},
	0x9d: {"STA", ABSX, 3, 5},

	0xa0: {"LDY", IMM, 2, 2},
	0xa1: {"LDA", INDX, 2, 6},
	0xa2: {"LDX", IMM, 2, 2},
	0xa4: {"LDY", ZERO, 2, 3},
	0xa5: {"LDA", ZERO, 2, 3},
	0xa6: {"LDX", ZERO, 2, 3},
	0xa8: {"TAY", IMPL, 1, 2},
	0xa9: {"LDA", IMM, 2, 2},
	0xaa: {"TAX", IMPL, 1, 2},
	0xac: {"LDY", ABS, 3, 4},
	0xad: {"LDA", ABS, 3, 4},
	0xae: {"LDX", ABS, 3, 4},

	0xb0: {"BCS", REL, 2, 2}, // cycle:3 ?
	0xb1: {"LDA", INDY, 2, 5},
	0xb4: {"LDY", ZEROX, 2, 4},
	0xb5: {"LDA", ZEROX, 2, 4},
	0xb6: {"LDX", ZEROY, 2, 4},
	0xb8: {"CLV", IMPL, 1, 2},
	0xb9: {"LDA", ABSY, 3, 4},
	0xba: {"TSX", IMPL, 1, 2},
	0xbc: {"LDY", ABSX, 3, 4},
	0xbd: {"LDA", ABSX, 3, 4},
	0xbe: {"LDX", ABSY, 3, 4},

	0xc0: {"CPY", IMM, 2, 2},
	0xc1: {"CMP", INDX, 2, 6},
	0xc4: {"CPY", ZERO, 2, 3},
	0xc5: {"CMP", ZERO, 2, 3},
	0xc6: {"DEC", ZERO, 2, 5},
	0xc8: {"INY", IMPL, 1, 2},
	0xc9: {"CMP", IMM, 2, 2},
	0xca: {"DEX", IMPL, 1, 2},
	0xcc: {"CPY", ABS, 3, 4},
	0xcd: {"CMP", ABS, 3, 4},
	0xce: {"DEC", ABS, 3, 6},

	0xd0: {"BNE", REL, 2, 2}, // cycle:3 ?
	0xd1: {"CMP", INDY, 2, 5},
	0xd5: {"CMP", ZEROX, 2, 4},
	0xd6: {"DEC", ZEROX, 2, 6},
	0xd8: {"CLD", IMPL, 1, 2},
	0xd9: {"CMP", ABSY, 3, 4},
	0xdd: {"CMP", ABSX, 3, 4},
	0xde: {"DEC", ABSX, 3, 7},

	0xe0: {"CPX", IMM, 2, 2},
	0xe1: {"SBC", INDX, 2, 6},
	0xe4: {"CPX", ZERO, 2, 3},
	0xe5: {"SBC", ZERO, 2, 3},
	0xe6: {"INC", ZERO, 2, 5},
	0xe8: {"INX", IMPL, 1, 2},
	0xe9: {"SBC", IMM, 2, 2},
	0xea: {"NOP", IMPL, 1, 2},
	0xec: {"CPX", ABS, 3, 4},
	0xed: {"SBC", ABS, 3, 4},
	0xee: {"INC", ABS, 3, 6},

	0xf0: {"BEQ", REL, 2, 2}, // cycle:3 ?
	0xf1: {"SBC", INDY, 2, 5},
	0xf5: {"SBC", ZEROX, 2, 4},
	0xf6: {"INC", ZEROX, 2, 6},
	0xf8: {"SED", IMPL, 1, 2},
	0xf9: {"SBC", ABSY, 3, 4},
	0xfd: {"SBC", ABSX, 3, 4},
	0xfe: {"INC", ABSX, 3, 7},
}
