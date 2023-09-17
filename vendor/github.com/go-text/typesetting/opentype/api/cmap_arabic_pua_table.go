// SPDX-License-Identifier: Unlicense OR BSD-3-Clause

package api

// Code generated by typesettings-utils/generators/unicodedata/cmd/main.go DO NOT EDIT.

// Legacy Simplified Arabic encoding. Returns 0 if not found.
func arabicPUASimpMap(r rune) rune {
	switch {
	case 0x20 <= r && r <= 0x22:
		return [...]rune{0xf120, 0xf121, 0xf122}[r-0x20]
	case 0x25 == r:
		return 0xf125
	case 0x28 <= r && r <= 0x3b:
		return [...]rune{0xf128, 0xf129, 0xf12a, 0xf12b, 0xf15e, 0xf12d, 0xf12e, 0xf12f, 0xf1b0, 0xf1b1, 0xf1b2, 0xf1b3, 0xf1b4, 0xf1b5, 0xf1b6, 0xf1b7, 0xf1b8, 0xf1b9, 0xf13a, 0xf13b}[r-0x28]
	case 0x3d == r:
		return 0xf13d
	case 0x3f == r:
		return 0xf13f
	case 0x5b <= r && r <= 0x5d:
		return [...]rune{0xf15b, 0xf15c, 0xf15d}[r-0x5b]
	case 0xab == r:
		return 0xf123
	case 0xbb == r:
		return 0xf124
	case 0xd7 == r:
		return 0xf126
	case 0xf7 == r:
		return 0xf127
	case 0x60c == r:
		return 0xf12c
	case 0x61b == r:
		return 0xf13b
	case 0x61f == r:
		return 0xf13f
	case 0x621 <= r && r <= 0x65e:
		return [...]rune{0xf1ad, 0xf145, 0xf143, 0xf1bb, 0xf147, 0xf1ba, 0xf141, 0xf14a, 0xf1a9, 0xf14c, 0xf14e, 0xf151, 0xf154, 0xf157, 0xf158, 0xf159, 0xf15a, 0xf160, 0xf162, 0xf164, 0xf166, 0xf168, 0xf169, 0xf16a, 0xf16e, 0xf172, 0xf100, 0xf100, 0xf100, 0xf100, 0xf100, 0xf15f, 0xf175, 0xf178, 0xf17a, 0xf17c, 0xf17e, 0xf1e1, 0xf1a4, 0xf1a5, 0xf1ac, 0xf1a8, 0xf1c7, 0xf1c8, 0xf1cb, 0xf1c4, 0xf1c5, 0xf1ca, 0xf1c9, 0xf1c6, 0xf100, 0xf100, 0xf100, 0xf100, 0xf100, 0xf100, 0xf100, 0xf100, 0xf100, 0xf100, 0xf100, 0xf100}[r-0x621]
	case 0x660 <= r && r <= 0x669:
		return [...]rune{0xf130, 0xf131, 0xf132, 0xf133, 0xf134, 0xf135, 0xf136, 0xf137, 0xf138, 0xf139}[r-0x660]
	case 0x66b <= r && r <= 0x66c:
		return [...]rune{0xf15e, 0xf15e}[r-0x66b]
	case 0x200c <= r && r <= 0x200f:
		return [...]rune{0xf10c, 0xf10d, 0xf10e, 0xf10f}[r-0x200c]
	case 0x2018 <= r && r <= 0x2019:
		return [...]rune{0xf13c, 0xf13e}[r-0x2018]
	case 0xfe81 <= r && r <= 0xfefc:
		return [...]rune{0xf145, 0xf146, 0xf143, 0xf144, 0xf1bb, 0xf1bb, 0xf147, 0xf148, 0xf1ba, 0xf1af, 0xf1ae, 0xf1ae, 0xf141, 0xf142, 0xf14a, 0xf14a, 0xf149, 0xf149, 0xf1a9, 0xf1aa, 0xf14c, 0xf14c, 0xf14b, 0xf14b, 0xf14e, 0xf14e, 0xf14d, 0xf14d, 0xf151, 0xf150, 0xf14f, 0xf14f, 0xf154, 0xf153, 0xf152, 0xf152, 0xf157, 0xf156, 0xf155, 0xf155, 0xf158, 0xf158, 0xf159, 0xf159, 0xf15a, 0xf15a, 0xf160, 0xf160, 0xf162, 0xf162, 0xf161, 0xf161, 0xf164, 0xf164, 0xf163, 0xf163, 0xf166, 0xf166, 0xf165, 0xf165, 0xf168, 0xf168, 0xf167, 0xf167, 0xf169, 0xf169, 0xf169, 0xf169, 0xf16a, 0xf16a, 0xf16a, 0xf16a, 0xf16e, 0xf16d, 0xf16b, 0xf16c, 0xf172, 0xf171, 0xf16f, 0xf170, 0xf175, 0xf175, 0xf173, 0xf174, 0xf178, 0xf178, 0xf176, 0xf177, 0xf17a, 0xf17a, 0xf179, 0xf179, 0xf17c, 0xf17c, 0xf17b, 0xf17b, 0xf17e, 0xf17e, 0xf17d, 0xf17d, 0xf1e1, 0xf1e1, 0xf17f, 0xf17f, 0xf1a4, 0xf1a3, 0xf1a1, 0xf1a2, 0xf1a5, 0xf1a5, 0xf1ac, 0xf1ab, 0xf1a8, 0xf1a7, 0xf1a6, 0xf1a6, 0xf1c0, 0xf1c1, 0xf1be, 0xf1bf, 0xf1c2, 0xf1c3, 0xf1bd, 0xf1bc}[r-0xfe81]
	}
	return 0
}

// Legacy Traditional Arabic encoding. Returns 0 if not found.
func arabicPUATradMap(r rune) rune {
	switch {
	case 0x20 <= r && r <= 0x22:
		return [...]rune{0xf220, 0xf221, 0xf222}[r-0x20]
	case 0x25 == r:
		return 0xf225
	case 0x28 <= r && r <= 0x2f:
		return [...]rune{0xf228, 0xf229, 0xf22a, 0xf22b, 0xf25e, 0xf22d, 0xf22e, 0xf22f}[r-0x28]
	case 0x3a <= r && r <= 0x3b:
		return [...]rune{0xf23a, 0xf23b}[r-0x3a]
	case 0x3d == r:
		return 0xf23d
	case 0x3f == r:
		return 0xf23f
	case 0x5b == r:
		return 0xf25b
	case 0x5d == r:
		return 0xf25d
	case 0xab == r:
		return 0xf223
	case 0xbb == r:
		return 0xf224
	case 0xd7 == r:
		return 0xf226
	case 0xf7 == r:
		return 0xf227
	case 0x60c == r:
		return 0xf22c
	case 0x61b == r:
		return 0xf23b
	case 0x61f == r:
		return 0xf23f
	case 0x621 <= r && r <= 0x65e:
		return [...]rune{0xf2d5, 0xf245, 0xf243, 0xf2da, 0xf247, 0xf2d9, 0xf241, 0xf24c, 0xf2d1, 0xf250, 0xf254, 0xf258, 0xf260, 0xf264, 0xf265, 0xf267, 0xf269, 0xf26b, 0xf270, 0xf274, 0xf278, 0xf27e, 0xf2a2, 0xf2a3, 0xf2aa, 0xf2ae, 0xf200, 0xf200, 0xf200, 0xf200, 0xf200, 0xf25f, 0xf2b2, 0xf2b6, 0xf2ba, 0xf2be, 0xf2c2, 0xf2c6, 0xf2ca, 0xf2cb, 0xf2d4, 0xf2d0, 0xf2e7, 0xf2e8, 0xf2eb, 0xf2e4, 0xf2e5, 0xf2ea, 0xf2e9, 0xf2e6, 0xf200, 0xf200, 0xf200, 0xf200, 0xf200, 0xf200, 0xf200, 0xf200, 0xf200, 0xf200, 0xf200, 0xf200}[r-0x621]
	case 0x660 <= r && r <= 0x669:
		return [...]rune{0xf230, 0xf231, 0xf232, 0xf233, 0xf234, 0xf235, 0xf236, 0xf237, 0xf238, 0xf239}[r-0x660]
	case 0x66b <= r && r <= 0x66c:
		return [...]rune{0xf25e, 0xf25e}[r-0x66b]
	case 0x200c <= r && r <= 0x200f:
		return [...]rune{0xf20c, 0xf20d, 0xf20e, 0xf20f}[r-0x200c]
	case 0x201c <= r && r <= 0x201d:
		return [...]rune{0xf23c, 0xf23e}[r-0x201c]
	case 0xfc08 == r:
		return 0xf202
	case 0xfc0a == r:
		return 0xf21d
	case 0xfc0e == r:
		return 0xf203
	case 0xfc10 == r:
		return 0xf21e
	case 0xfc12 == r:
		return 0xf204
	case 0xfc32 == r:
		return 0xf29f
	case 0xfc3f <= r && r <= 0xfc42:
		return [...]rune{0xf212, 0xf213, 0xf214, 0xf205}[r-0xfc3f]
	case 0xfc44 == r:
		return 0xf21c
	case 0xfc4e == r:
		return 0xf206
	case 0xfc50 == r:
		return 0xf21f
	case 0xfc5e == r:
		return 0xf2ef
	case 0xfc60 <= r && r <= 0xfc62:
		return [...]rune{0xf2ec, 0xf2ed, 0xf2f0}[r-0xfc60]
	case 0xfc6a == r:
		return 0xf215
	case 0xfc6d == r:
		return 0xf292
	case 0xfc70 == r:
		return 0xf216
	case 0xfc73 == r:
		return 0xf293
	case 0xfc86 == r:
		return 0xf295
	case 0xfc91 == r:
		return 0xf217
	case 0xfc94 == r:
		return 0xf294
	case 0xfc9c <= r && r <= 0xfc9f:
		return [...]rune{0xf280, 0xf281, 0xf282, 0xf296}[r-0xfc9c]
	case 0xfca1 <= r && r <= 0xfca4:
		return [...]rune{0xf283, 0xf284, 0xf285, 0xf297}[r-0xfca1]
	case 0xfca8 == r:
		return 0xf29a
	case 0xfcaa == r:
		return 0xf29b
	case 0xfcac == r:
		return 0xf29c
	case 0xfcb0 == r:
		return 0xf218
	case 0xfcc9 <= r && r <= 0xfcd3:
		return [...]rune{0xf286, 0xf287, 0xf288, 0xf29d, 0xf21a, 0xf289, 0xf28a, 0xf28b, 0xf29e, 0xf28d, 0xf28e}[r-0xfcc9]
	case 0xfcd5 == r:
		return 0xf298
	case 0xfcda <= r && r <= 0xfcdd:
		return [...]rune{0xf28f, 0xf290, 0xf291, 0xf299}[r-0xfcda]
	case 0xfd30 == r:
		return 0xf219
	case 0xfd3e <= r && r <= 0xfd3f:
		return [...]rune{0xf27b, 0xf27d}[r-0xfd3e]
	case 0xfd88 == r:
		return 0xf210
	case 0xfe81 <= r && r <= 0xfefc:
		return [...]rune{0xf245, 0xf246, 0xf243, 0xf244, 0xf2da, 0xf2db, 0xf247, 0xf248, 0xf2d9, 0xf2d8, 0xf2d6, 0xf2d7, 0xf241, 0xf242, 0xf24c, 0xf24b, 0xf249, 0xf24a, 0xf2d1, 0xf2d2, 0xf250, 0xf24f, 0xf24d, 0xf24e, 0xf254, 0xf253, 0xf251, 0xf252, 0xf258, 0xf257, 0xf255, 0xf256, 0xf260, 0xf25c, 0xf259, 0xf25a, 0xf264, 0xf263, 0xf261, 0xf262, 0xf265, 0xf266, 0xf267, 0xf268, 0xf269, 0xf26a, 0xf26b, 0xf26c, 0xf270, 0xf26f, 0xf26d, 0xf26e, 0xf274, 0xf273, 0xf271, 0xf272, 0xf278, 0xf277, 0xf275, 0xf276, 0xf27e, 0xf27c, 0xf279, 0xf27a, 0xf2a2, 0xf2a1, 0xf27f, 0xf2f1, 0xf2a6, 0xf2a5, 0xf2a3, 0xf2a4, 0xf2aa, 0xf2a9, 0xf2a7, 0xf2a8, 0xf2ae, 0xf2ad, 0xf2ab, 0xf2ac, 0xf2b2, 0xf2b1, 0xf2af, 0xf2b0, 0xf2b6, 0xf2b5, 0xf2b3, 0xf2b4, 0xf2ba, 0xf2b9, 0xf2b7, 0xf2b8, 0xf2be, 0xf2bd, 0xf2bb, 0xf2bc, 0xf2c2, 0xf2c1, 0xf2bf, 0xf2c0, 0xf2c6, 0xf2c5, 0xf2c3, 0xf2c4, 0xf2ca, 0xf2c9, 0xf2c7, 0xf2c8, 0xf2cb, 0xf2cc, 0xf2d4, 0xf2d3, 0xf2d0, 0xf2cf, 0xf2cd, 0xf2ce, 0xf2e0, 0xf2e1, 0xf2de, 0xf2df, 0xf2e2, 0xf2e3, 0xf2dc, 0xf2dd}[r-0xfe81]
	}
	return 0
}
