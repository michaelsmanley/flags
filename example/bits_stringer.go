// Generated by: main
// TypeWriter: stringer
// Directive: +gen on Bits

package example

import (
	"fmt"
)

const _Bits_name = "Bit01Bit02Bit03MutuallyExclusiveBit04Bit05Bit06Bit07Bit08Bit09Bit10Bit11Bit12Bit13Bit14Bit15Bit16Bit17Bit18Bit19Bit20Bit21Bit22Bit23Bit24Bit25Bit26Bit27Bit28Bit29Bit30Bit31AnyOddBit32AnyEven"

var _Bits_map = map[Bits]string{
	1:          _Bits_name[0:5],
	2:          _Bits_name[5:10],
	4:          _Bits_name[10:15],
	7:          _Bits_name[15:32],
	8:          _Bits_name[32:37],
	16:         _Bits_name[37:42],
	32:         _Bits_name[42:47],
	64:         _Bits_name[47:52],
	128:        _Bits_name[52:57],
	256:        _Bits_name[57:62],
	512:        _Bits_name[62:67],
	1024:       _Bits_name[67:72],
	2048:       _Bits_name[72:77],
	4096:       _Bits_name[77:82],
	8192:       _Bits_name[82:87],
	16384:      _Bits_name[87:92],
	32768:      _Bits_name[92:97],
	65536:      _Bits_name[97:102],
	131072:     _Bits_name[102:107],
	262144:     _Bits_name[107:112],
	524288:     _Bits_name[112:117],
	1048576:    _Bits_name[117:122],
	2097152:    _Bits_name[122:127],
	4194304:    _Bits_name[127:132],
	8388608:    _Bits_name[132:137],
	16777216:   _Bits_name[137:142],
	33554432:   _Bits_name[142:147],
	67108864:   _Bits_name[147:152],
	134217728:  _Bits_name[152:157],
	268435456:  _Bits_name[157:162],
	536870912:  _Bits_name[162:167],
	1073741824: _Bits_name[167:172],
	1431655765: _Bits_name[172:178],
	2147483648: _Bits_name[178:183],
	2863311530: _Bits_name[183:190],
}

func (i Bits) String() string {
	if str, ok := _Bits_map[i]; ok {
		return str
	}
	return fmt.Sprintf("Bits(%d)", i)
}