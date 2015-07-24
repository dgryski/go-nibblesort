/* Package nibblesort sorts nibbles in a uint64 */
/*

This is a port of  `jerome.c`, the winner from https://github.com/regehr/nibble-sort

For more information:

    http://blog.regehr.org/archives/1213
    http://www.hanshq.net/nibble_sort.html

This is MIT licensed, like the original.

*/
package nibblesort

var table [256]uint64
var offsets [256]uint64
var table3 [256 * 17]uint64

func init() {

	for i := uint64(0); i < 16; i++ {
		for j := uint64(0); j < 16; j++ {
			table[i*16+j] = (1 << (4 * i)) + (1 << (4 * j))
			offsets[i*16+j] = (i + j) << 8
			for k := uint64(0); k < 17; k++ {

				var v1 uint64
				if i+j+k < 16 {
					v1 = 0x1111111111111111 << (4 * (i + j + k))
				}

				var v2 uint64
				if j+k < 16 {
					v2 = 0x1111111111111111 << (4 * (j + k))
				}

				table3[i*16+j+k*256] = v1 + v2
			}
		}
	}
}

func Sort(word uint64) uint64 {

	if (word<<4 | word>>60) == word {
		return word
	}

	var counts uint64
	for i := uint64(0); i < 8; i++ {
		counts += table[(word>>(8*i))&0xff]
	}

	var output uint64
	var offset uint64
	for i := uint64(0); i < 8; i++ {
		output += table3[((counts>>(8*i))&0xff)+offset]
		offset += offsets[(counts>>(8*i))&0xff]
	}
	return output
}
