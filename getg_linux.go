//go:build linux
// +build linux

package monkey

func getg() []byte {
	return []byte{
		// mov r12,QWORD PTR fs:0xfffffffffffffff8
		0x64, 0x4C, 0x8B, 0x24, 0x25, 0xF8, 0xFF, 0xFF, 0xFF,
	}
}
