package findfont

// #cgo pkg-config: fontconfig
// #include <stdlib.h>
// #include <fontconfig/fontconfig.h>
import "C"
import (
	"strings"
	"unsafe"

	"github.com/danielgatis/go-ptrloop/ptrloop"
	"github.com/pkg/errors"
)

type fontStyle byte

const (
	// FontRegular Regular style
	FontRegular fontStyle = 0x00
	// FontBold Bold Style
	FontBold fontStyle = 0x01
	// FontItalic Italic style
	FontItalic fontStyle = 0x02
)

// Find Returns a list of fonts through the fontconfig library
func Find(family string, style fontStyle) ([][]string, error) {
	var err error

	fontList := make([][]string, 0)
	format := C.CString("%{family};;%{style};;%{file}")

	family += strStyle(style)
	name := C.CString(family)

	pattern := C.FcNameParse((*C.FcChar8)(unsafe.Pointer(name)))
	C.FcConfigSubstitute(nil, pattern, C.FcMatchPattern)
	C.FcDefaultSubstitute(pattern)

	var t C.FcResult
	fontPatterns := C.FcFontSort(nil, pattern, C.FcTrue, nil, &t)

	fs := C.FcFontSetCreate()

	if fontPatterns == nil || fontPatterns.nfont == 0 {
		err = errors.New("No fonts installed on the system")
		goto Exit
	}

	ptrloop.Loop(
		unsafe.Pointer(fontPatterns.fonts),
		int(fontPatterns.nfont),
		func(ptr unsafe.Pointer, i int) bool {
			font := *(**C.FcPattern)(ptr)
			fontPattern := C.FcFontRenderPrepare(nil, pattern, font)

			if fontPattern != nil {
				C.FcFontSetAdd(fs, fontPattern)
			}

			return true
		},
	)

	ptrloop.Loop(
		unsafe.Pointer(fs.fonts),
		int(fs.nfont),
		func(ptr unsafe.Pointer, i int) bool {
			font := *(**C.FcPattern)(ptr)

			f := C.FcPatternFilter(font, nil)
			s := C.FcPatternFormat(f, (*C.FcChar8)(unsafe.Pointer(format)))
			r := C.GoString((*C.char)(unsafe.Pointer(s)))

			fontList = append(fontList, strings.Split(r, ";;"))
			C.FcPatternDestroy(f)

			return true
		},
	)

Exit:
	C.free(unsafe.Pointer(format))
	C.free(unsafe.Pointer(name))
	C.FcPatternDestroy(pattern)
	C.FcFontSetDestroy(fs)
	C.FcFontSetSortDestroy(fontPatterns)
	C.FcFini()

	return fontList, err
}

func strStyle(style fontStyle) string {
	switch style {
	case FontBold:
		return ":Bold"
	case FontItalic:
		return ":Italic"
	case FontBold | FontItalic:
		return ":Bold:Italic"
	default:
		return ":Regular"
	}
}
