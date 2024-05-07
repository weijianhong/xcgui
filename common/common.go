package common

import (
	"unicode/utf16"
	"unsafe"
)

// Bytes2String byte切片到string.
//
//	@param b
//	@return string
func Bytes2String(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

func String2Bytes(s string) []byte {
	sh := (*stringHeader)(unsafe.Pointer(&s))
	bh := sliceHeader{
		Data: sh.Data,
		Len:  sh.Len,
		Cap:  sh.Len,
	}
	return *(*[]byte)(unsafe.Pointer(&bh))
}

// stringHeader is the runtime representation of a string.
// It cannot be used safely or portably and its representation may
// change in a later release.
// Moreover, the Data field is not sufficient to guarantee the data
// it references will not be garbage collected, so programs must keep
// a separate, correctly typed pointer to the underlying data.
type stringHeader struct {
	Data uintptr
	Len  int
}

// sliceHeader is the runtime representation of a slice.
// It cannot be used safely or portably and its representation may
// change in a later release.
// Moreover, the Data field is not sufficient to guarantee the data
// it references will not be garbage collected, so programs must keep
// a separate, correctly typed pointer to the underlying data.
type sliceHeader struct {
	Data uintptr
	Len  int
	Cap  int
}

func toUTF16Ptr(s string) (*uint16, int) {
	// 将字符串转换为 UTF-16 编码的 rune 切片
	utf16s := utf16.Encode([]rune(s))

	// 分配一个额外的 uint16 用于空字符（字符串终止符）
	utf16WithNull := make([]uint16, len(utf16s)+1)
	copy(utf16WithNull, utf16s)
	utf16WithNull[len(utf16s)] = 0 // 空字符作为字符串终止符

	// 返回指向第一个 UTF-16 字符的指针和长度（不包括空字符）
	return &utf16WithNull[0], len(utf16s)
}

// StrPtr 将string转换到uintptr.
//
//	@param s
//	@return uintptr
func StrPtr(s string) uintptr {
	if len(s) == 0 {
		return uintptr(0)
	}
	p, _ := toUTF16Ptr(s)

	return uintptr(unsafe.Pointer(p))
}

// UintPtrToString 将uintptr转换到string.
//
//	@param ptr
//	@return string
func UintPtrToString(ptr uintptr) string {
	if ptr == 0 {
		return ""
	}
	s := *(*[]uint16)(unsafe.Pointer(&ptr)) // uintptr转换到[]uint16
	for i := 0; i < len(s); i++ {
		if s[i] == 0 {
			(*sliceHeader)(unsafe.Pointer(&s)).Cap = i // 修改切片的cap
			s = s[0:i]
			break
		}
	}
	return string(utf16.Decode(s))
}

// Uint16SliceDataPtr 将uint16[0]指针转换到uintptr.
//
//	@param p
//	@return uintptr
func Uint16SliceDataPtr(p *[]uint16) uintptr {
	if len(*p) == 0 {
		return uintptr(0)
	}
	return uintptr(unsafe.Pointer(&(*p)[0]))
}

// BoolPtr 将bool转换到uintptr.
//
//	@param b
//	@return uintptr
func BoolPtr(b bool) uintptr {
	if b {
		return uintptr(1)
	}
	return uintptr(0)
}

// Float32Ptr 将float32转换到uintptr.
//
//	@param f
//	@return uintptr
func Float32Ptr(f float32) uintptr {
	return uintptr(*(*uint32)(unsafe.Pointer(&f)))
}

// UintPtrToFloat32 将uintptr转换到float32.
//
//	@param ptr
//	@return float32
func UintPtrToFloat32(ptr uintptr) float32 {
	if ptr == 0 {
		return 0
	}
	u := uint32(ptr)
	return *(*float32)(unsafe.Pointer(&u))
}

// ByteSliceDataPtr 将byte[0]指针转换到uintptr.
//
//	@param b
//	@return uintptr
func ByteSliceDataPtr(b *[]byte) uintptr {
	if len(*b) == 0 {
		return uintptr(0)
	}
	return uintptr(unsafe.Pointer(&(*b)[0]))
}

/* // byte指针到uintptr.
func bytePtr(p *byte) uintptr {
	return uintptr(unsafe.Pointer(p))
} */

/* func UintPtrToString2(r uintptr) string {
    p := (*uint16)(unsafe.Pointer(r))
    if p == nil {
        return ""
    }

    n, end, add := 0, unsafe.Pointer(p), unsafe.Sizeof(*p)
    for *(*uint16)(end) != 0 {
        end = unsafe.Add(end, add)
        n++
    }
    return string(utf16.Decode(unsafe.Slice(p, n)))
} */

/*// RuneToUint16Ptr 返回指向 UTF-8 字符串 r 的 UTF-16 编码的指针.
func RuneToUint16Ptr(r []rune) *uint16 {
	return &utf16.Encode(r)[0]
}*/

// StringToUint16Ptr 返回指向 UTF-8 字符串 s 的 UTF-16 编码的指针，与 syscall.UTF16PtrFromString 不同的是末尾没有添加终止 NUL。
func StringToUint16Ptr(s string) *uint16 {
	return &utf16.Encode([]rune(s))[0]
}

// Uint16SliceToStringSlice 按null字符分割, 把 []uint16 转换到 []string.
func Uint16SliceToStringSlice(s []uint16) []string {
	strSlice := make([]string, 0)
	start := 0
	for i := 0; i < len(s); i++ {
		if s[i] == 0 {
			strSlice = append(strSlice, string(utf16.Decode(s[start:i])))
			start = i + 1

			// 连续null字符, 所以跳出
			if s[start] == 0 {
				break
			}
		}
	}
	return strSlice
}
