package p503

// Contains values used by tests
import (
	"testing/quick"

	. "github.com/cloudflare/circl/dh/sidh/internal/common"
)

// Values omputed using Sage
var (
	// j = 3674553797500778604587777859668542828244523188705960771798425843588160903687122861541242595678107095655647237100722594066610650373491179241544334443939077738732728884873568393760629500307797547379838602108296735640313894560419*i + 3127495302417548295242630557836520229396092255080675419212556702820583041296798857582303163183558315662015469648040494128968509467224910895884358424271180055990446576645240058960358037224785786494172548090318531038910933793845
	expectedJ = Fp2{
		A: Fp{0x2c441d03b72e27c, 0xf2c6748151dbf84, 0x3a774f6191070e, 0xa7c6212c9c800ba6, 0x23921b5cf09abc27, 0x9e1baefbb3cd4265, 0x8cd6a289f12e10dc, 0x3fa364128cf87e},
		B: Fp{0xe7497ac2bf6b0596, 0x629ee01ad23bd039, 0x95ee11587a119fa7, 0x572fb28a24772269, 0x3c00410b6c71567e, 0xe681e83a345f8a34, 0x65d21b1d96bd2d52, 0x7889a47e58901},
	}

	// A = 8752234765512331234913716743014562460822083005386252003333602919474238975785850965349950219277942402920758585086620525443539725921333735154674119646075*i + 6339624979889725406021454983012408976766782818694212228554611573314701271183857175866122275755278397694585249002282183018114967373119429936587424396917
	curve_A = Fp2{
		A: Fp{0xd9816986a543095f, 0xa78cb1d7217bec21, 0x9595dc97b74ea70, 0x9120a1da6b42797d, 0x59ef9d903f74e47c, 0x4c58a4cdc45b6d0b, 0x816d5213aaf7ee6d, 0x3892fee6bb7343},
		B: Fp{0x28c5288acbedf11b, 0x2143a438c86f6c68, 0x7cb5c4ae9c4c8e34, 0xb478aea445eed48b, 0x24d5c175776db478, 0x234582f8676c0ebe, 0x56234267b625fb08, 0x2c6e58d84b1192}}

	// C = 10458464853790890798085664692909194316288127038910691163573355876336993883402795907795767791362493831987298578966325154262747805705783782806176495638177*i + 7770984753616185271325854825309278833018655051139367603077592443785629339985729818288672809062782315510526648882226172896710704020683893684611137718845
	curve_C = Fp2{
		A: Fp{0xe05948236f2f913b, 0xc45da9ad1219a255, 0x7a568972a32fc1d0, 0x30f00bdd7071c3b1, 0x3b761b8dac2c98bc, 0x760f21b2179737b6, 0x13217e6656a13476, 0x2606b798e685aa},
		B: Fp{0x1c0171f78820052e, 0x440b7f7087e57140, 0xe0510c07b31b0e96, 0xd0cf489b2ac4aea9, 0x4fb328f1c1fdf783, 0xb3b4912342951cb7, 0x70a4b64e81961c42, 0x33eed63cf07181}}
	// x(P) = 9720237205826983370867050298878715935679372786589878620121159082290288918688002583435964840822877971257659901481591644347943354235932355923042390796255*i + 634577413124118560098123299804750904956499531431297942628887930019161512075536652691244843248133437326050395005054997679717801535474938466995392156605
	affine_xP = Fp2{
		A: Fp{0xb606d954d407faf2, 0x58a1ef6cd213a203, 0x9823b55033e62f7b, 0x59cafc060d5e25a1, 0x529685f1753526fc, 0xc2eac3d219989c7d, 0xc5e30c75dfd343a0, 0x378285adc968a0},
		B: Fp{0x6670f36db977b9da, 0xa07e2fdda5e1a7f0, 0xf367a7a722aed87d, 0x6c269e06d595cd10, 0x8379aa6092d87700, 0x57276ce3557ee7ae, 0xac8107bfbcd28993, 0x3d6f98869617a7}}

	affine_xP2 = Fp2{
		A: Fp{0x4e1133c2b3855902, 0x875a775c67597fbb, 0xd17eb74254141abb, 0x1d5a464a4f3391f5, 0x24405c332811d007, 0x7e47e3eb489a7372, 0x65b130dfd9efe605, 0xfa69fac179803},
		B: Fp{0x329f5322e1be51ee, 0x9004dca8132ebd6f, 0x7cd87e447ca8a7b6, 0x10a6ec02c38ce69e, 0x8cef2ed7d112ac46, 0x5f385a9fc4b57cd7, 0x68a366354fe7a32e, 0x2223c1455486ac}}

	affine_xP4 = Fp2{
		A: Fp{0x4eb695d34b46be8f, 0xfb5e76c58585f2d2, 0xa41f8aafa6dbb531, 0x4db82f5db5cfd144, 0x14dab0e3200cbba0, 0x430381706a279f81, 0xdf6707a57161f81, 0x44740f17197c3},
		B: Fp{0xa2473705cdb6d4e9, 0xfa3cd67b9c15502c, 0xf0928166d0c5cee1, 0x6150aba0c874faaa, 0x6c0b18d6d92f9034, 0xcff71d340fc1e72e, 0x19a47027af917587, 0x25ed4bad443b8f}}

	affine_xP9 = Fp2{
		A: Fp{0x112da30e288217e0, 0x5b336d527320a5f7, 0xbbf4d9403b68e3c6, 0x55eccb31c40b359c, 0x8907129ab69b3203, 0x69cc8c750125a915, 0xa41a38e6f530c0e1, 0xbe68e23af1b8d},
		B: Fp{0x472c603765964213, 0xe4e64995b0769754, 0x4515583c74a6dd24, 0xff7c57f5818363a2, 0xbeaeb24662a92177, 0x8a54fa61fbf24c68, 0xa85542049eb45e12, 0x2b54caf655e285}}

	// Inputs for testing 3-point-ladder
	threePointLadderInputs = []ProjectivePoint{
		// x(P)
		{
			X: Fp2{
				A: Fp{0x43941FA9244C059E, 0xD1F337D076941189, 0x6B6A8B3A8763C96A, 0x6DF569708D6C9482, 0x487EE5707A52F4AA, 0xDE396F6E2559689E, 0xE5EE3895A8991469, 0x2B0946695790A8},
				B: Fp{0xAB552C0FDAED092E, 0x7DF895E43E7DCB1C, 0x35C700E761920C4B, 0xCC5807DD70DC117A, 0x0884039A5A8DB18A, 0xD04620B3D0738052, 0xA200835605138F10, 0x3FF2E59B2FDC6A}},
			Z: params.OneFp2,
		},
		// x(Q)
		{
			X: Fp2{
				A: Fp{0x77015826982BA1FD, 0x44024489673471E4, 0x1CAA2A5F4D5DA63B, 0xA183C07E50738C01, 0x8B97782D4E1A0DE6, 0x9B819522FBC38280, 0x0BDA46A937FB7B8A, 0x3B3614305914DF},
				B: Fp{0xBF0366E97B3168D9, 0xAA522AC3879CEF0F, 0x0AF5EC975BD035C8, 0x1F26FEE7BBAC165C, 0xA0EE6A637724A6AB, 0xFB52101E36BA3A38, 0xD29CF5E376E17376, 0x1374A50DF57071}},
			Z: params.OneFp2,
		},
		// x(P-Q)
		{
			X: Fp2{
				A: Fp{0xD99279BBD41EA559, 0x35CF18E72F578214, 0x90473B1DC77F73E8, 0xBFFEA930B25D7F66, 0xFD558EA177B900B2, 0x7CFAD273A782A23E, 0x6B1F610822E0F611, 0x26D2D2EF9619B5},
				B: Fp{0x534F83651CBCC75D, 0x591FB4757AED5D08, 0x0B04353D40BED542, 0x829A94703AAC9139, 0x0F9C2E6D7663EB5B, 0x5D2D0F90C283F746, 0x34C872AA12A7676E, 0x0ECDB605FBFA16}},
			Z: params.OneFp2,
		},
	}
	scalar3Pt = [...]uint8{0x9f, 0x3b, 0xe7, 0xf9, 0xf4, 0x7c, 0xe6, 0xce, 0x79, 0x3e, 0x3d, 0x9f, 0x9f, 0x3b, 0xe7, 0xf9, 0xf4, 0x7c, 0xe6, 0xce, 0x79, 0x3e, 0x3d, 0x9f}
)

var quickCheckConfig = &quick.Config{
	MaxCount: (1 << 15),
}
