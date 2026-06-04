package distpacher

import (
	"encoding/binary"
)

var byteOrder binary.ByteOrder = binary.BigEndian

type PacketsIDS int32;

var(
	HandShake	PacketsIDS = -1
	Ping 		PacketsIDS = 0
	PlayPause 	PacketsIDS = 1
	SetVideo	PacketsIDS = 2
	SetTime 	PacketsIDS = 3
	VideoFinish PacketsIDS = 4
	VideoLoaded PacketsIDS = 5
)



