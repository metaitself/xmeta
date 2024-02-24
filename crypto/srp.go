package crypto

import (
	"bytes"
	"math/big"
	"math/rand"
	"time"
)

type PasswordKdfAlgoModPow struct {
	Salt1 []byte
	Salt2 []byte
	G     int32
	P     []byte
}

type SRPUtil struct {
	*PasswordKdfAlgoModPow
	g      *big.Int
	gBytes []byte
	p      *big.Int
	k      *big.Int
}

func MakeSRPUtil(algo *PasswordKdfAlgoModPow) *SRPUtil {
	rand.Seed(time.Now().UnixNano())

	// TODO(@benqi): check algo
	srp := &SRPUtil{
		PasswordKdfAlgoModPow: algo,
		g:                     big.NewInt(int64(algo.G)),
		p:                     new(big.Int).SetBytes(algo.P),
	}

	srp.gBytes = getBigIntegerBytes(srp.g)
	kBytes := calcSHA256(algo.P, srp.gBytes)
	srp.k = new(big.Int).SetBytes(kBytes)

	return srp
}

func (m *SRPUtil) CheckNewSalt1(newSalt1 []byte) bool {
	if len(newSalt1) < 8 {
		return false
	}

	return bytes.Equal(m.Salt1, newSalt1[:8])
}

func (m *SRPUtil) CalcSRPB(vBytes []byte) ([]byte, []byte) {
	v := new(big.Int).SetBytes(vBytes)

	bNonce := RandomBytes(256)
	//fmt.Println(hex.EncodeToString(bNonce))
	b := new(big.Int).SetBytes(bNonce)
	//bBytes := getBigIntegerBytes(b)
	//fmt.Println(hex.EncodeToString(bBytes))
	// Host —> User: s,B=k*v+g**b(发送盐值，和公开的B，其中b是一个随机选取的值
	B := new(big.Int).Mod(new(big.Int).Add(new(big.Int).Mul(m.k, v), new(big.Int).Exp(m.g, b, m.p)), m.p)
	//B := new(big.Int).Add(new(big.Int).Mul(m.k, v), new(big.Int).Exp(m.g, b, m.p))
	BBytes := getBigIntegerBytes(B)

	return bNonce, BBytes
}

func (m *SRPUtil) CalcSRPB2(bNonce, vBytes []byte) []byte {
	v := new(big.Int).SetBytes(vBytes)

	// bNonce := RandomBytes(256)
	//fmt.Println(hex.EncodeToString(bNonce))
	b := new(big.Int).SetBytes(bNonce)
	//bBytes := getBigIntegerBytes(b)
	//fmt.Println(hex.EncodeToString(bBytes))
	// Host —> User: s,B=k*v+g**b(发送盐值，和公开的B，其中b是一个随机选取的值
	B := new(big.Int).Mod(new(big.Int).Add(new(big.Int).Mul(m.k, v), new(big.Int).Exp(m.g, b, m.p)), m.p)
	//B := new(big.Int).Add(new(big.Int).Mul(m.k, v), new(big.Int).Exp(m.g, b, m.p))

	BBytes := getBigIntegerBytes(B)

	return BBytes
}

func (m *SRPUtil) CalcM(newSalt1, vBytes, srpA, srpb, srpB []byte) []byte {
	v := new(big.Int).SetBytes(vBytes)

	A := new(big.Int).SetBytes(srpA)
	if A.Cmp(bigIntZero) <= 0 || A.Cmp(m.p) >= 0 {
		return nil
	}
	ABytes := getBigIntegerBytes(A)

	//// Host —> User: s,B=kv+gb(发送盐值，和公开的B，其中b是一个随机选取的值
	b := new(big.Int).SetBytes(srpb)

	// B := new(big.Int).SetBytes(srpB)
	// BBytes := getBigIntegerBytes(B)
	// BBytes := srpB // getBigIntegerBytes(B)

	//_ = BBytes

	uBytes := calcSHA256(ABytes, srpB)
	u := new(big.Int).SetBytes(uBytes)
	if u.Cmp(bigIntZero) == 0 {
		return nil
	}
	//fmt.Println("uBytes: " + hex.EncodeToString(uBytes))
	//fmt.Println("vBytes: " + hex.EncodeToString(vBytes))

	// Host: S=(A*v**u)**b(如果用户输入的口令是对的，那这里算出的S值和客户端算出的S值是一致的)
	S := new(big.Int).Exp(new(big.Int).Mod(new(big.Int).Mul(A, new(big.Int).Exp(v, u, m.p)), m.p), b, m.p)
	SBytes := getBigIntegerBytes(S)

	KBytes := calcSHA256(SBytes)

	//fmt.Println("pHash: " + hex.EncodeToString(m.P))
	//fmt.Println("gHash: " + hex.EncodeToString(m.gBytes))
	//fmt.Println("newSalt1: " + hex.EncodeToString(newSalt1))
	//fmt.Println("m.Salt2: " + hex.EncodeToString(m.Salt2))
	//fmt.Println("ABytes: " + hex.EncodeToString(ABytes))
	//fmt.Println("BBytes: " + hex.EncodeToString(srpB))
	//fmt.Println("KBytes: " + hex.EncodeToString(KBytes))

	pHash := calcSHA256(m.P)
	gHash := calcSHA256(m.gBytes)
	for i := 0; i < len(pHash); i++ {
		pHash[i] = gHash[i] ^ pHash[i]
	}

	return calcSHA256(pHash, calcSHA256(newSalt1), calcSHA256(m.Salt2), ABytes, srpB, KBytes)
}

func (m *SRPUtil) GetX(newSalt1, passwordBytes []byte) []byte {
	var xBytes []byte

	xBytes = calcSHA256(newSalt1, passwordBytes, newSalt1)
	xBytes = calcSHA256(m.Salt2, xBytes, m.Salt2)
	xBytes = calcPBKDF2(xBytes, newSalt1)

	return calcSHA256(m.Salt2, xBytes, m.Salt2)
}

func (m *SRPUtil) GetV(newSalt1, passwordBytes []byte) *big.Int {
	xBytes := m.GetX(newSalt1, passwordBytes)
	x := new(big.Int).SetBytes(xBytes)

	return new(big.Int).Exp(m.g, x, m.p)
}

func (m *SRPUtil) GetVBytes(newSalt1, passwordBytes []byte) []byte {
	return getBigIntegerBytes(m.GetV(newSalt1, passwordBytes))
}

func (m *SRPUtil) CalcClientM(newSalt1, xBytes, srpB []byte) ([]byte, []byte) {
	if len(xBytes) == 0 || len(srpB) == 0 || !isGoodPrime(m.P, int(m.G)) {
		// fmt.Println("check error")
		return nil, nil
	}

	x := new(big.Int).SetBytes(xBytes)

	aBytes := RandomBytes(256)
	//fmt.Println(hex.EncodeToString(aBytes))
	a := new(big.Int).SetBytes(aBytes)

	A := new(big.Int).Exp(m.g, a, m.p)
	ABytes := getBigIntegerBytes(A)

	B := new(big.Int).SetBytes(srpB)
	if B.Cmp(bigIntZero) <= 0 || B.Cmp(m.p) >= 0 {
		// fmt.Println("b error")
		return nil, nil
	}
	BBytes := getBigIntegerBytes(B)

	uBytes := calcSHA256(ABytes, BBytes)
	u := new(big.Int).SetBytes(uBytes)
	if u.Cmp(bigIntZero) == 0 {
		// fmt.Println("u error")
		return nil, nil
	}

	//// Host: S=(A*v**u)**b(如果用户输入的口令是对的，那这里算出的S值和客户端算出的S值是一致的)
	//S := new(big.Int).Exp(new(big.Int).Mod(new(big.Int).Mul(A, new(big.Int).Exp(v, u, m.p)), m.p), b, m.p)
	//SBytes := getBigIntegerBytes(S)

	// new(big.Int).Exp(v, u, m.p)
	// new(big.Int).Mul(A, new(big.Int).Exp(v, u, m.p))
	// new(big.Int).Exp(new(big.Int).Mul(A, new(big.Int).Exp(v, u, m.p)), b, m.p)
	// User: S=(B−k*g**x)**(a+u*x) (S就是用户计算的会话密钥生成值)
	BKgx := new(big.Int).Sub(B, new(big.Int).Mod(new(big.Int).Mul(m.k, new(big.Int).Exp(m.g, x, m.p)), m.p))
	if BKgx.Cmp(bigIntZero) < 0 {
		// fmt.Println("<0")
		BKgx = new(big.Int).Add(BKgx, m.p)
	}

	if !isGoodGaAndGb(BKgx, m.p) {
		// fmt.Println("isGoodGaAndGb error")
		return nil, nil
	}

	S := new(big.Int).Exp(BKgx, new(big.Int).Add(a, new(big.Int).Mul(u, x)), m.p)
	SBytes := getBigIntegerBytes(S)

	KBytes := calcSHA256(SBytes)

	pHash := calcSHA256(m.P)
	gHash := calcSHA256(m.gBytes)
	for i := 0; i < len(pHash); i++ {
		pHash[i] = gHash[i] ^ pHash[i]
	}

	return ABytes, calcSHA256(pHash, calcSHA256(newSalt1), calcSHA256(m.Salt2), ABytes, BBytes, KBytes)

	// result.M1 = Utilities.computeSHA256(p_hash, Utilities.computeSHA256(algo.salt1), Utilities.computeSHA256(algo.salt2), A_bytes, B_bytes, K_bytes);

}

func (m *SRPUtil) CalcClientM2(newSalt1, aBytes, ABytes, xBytes, srpB []byte) []byte {
	if len(xBytes) == 0 || len(srpB) == 0 || !isGoodPrime(m.P, int(m.G)) {
		// fmt.Println("check error")
		return nil
	}

	x := new(big.Int).SetBytes(xBytes)

	// aBytes := RandomBytes(256)
	// fmt.Println(hex.EncodeToString(aBytes))
	a := new(big.Int).SetBytes(aBytes)
	//
	//A := new(big.Int).Exp(m.g, a, m.p)
	//ABytes := A //getBigIntegerBytes(A)

	B := new(big.Int).SetBytes(srpB)
	if B.Cmp(bigIntZero) <= 0 || B.Cmp(m.p) >= 0 {
		// fmt.Println("b error")
		return nil
	}
	BBytes := getBigIntegerBytes(B)

	uBytes := calcSHA256(ABytes, BBytes)
	u := new(big.Int).SetBytes(uBytes)
	if u.Cmp(bigIntZero) == 0 {
		// fmt.Println("u error")
		return nil
	}

	// User: S=(B−k*g**x)**(a+u*x) (S就是用户计算的会话密钥生成值)
	BKgx := new(big.Int).Sub(B, new(big.Int).Mod(new(big.Int).Mul(m.k, new(big.Int).Exp(m.g, x, m.p)), m.p))
	if BKgx.Cmp(bigIntZero) < 0 {
		// fmt.Println("<0")
		BKgx = new(big.Int).Add(BKgx, m.p)
	}

	if !isGoodGaAndGb(BKgx, m.p) {
		// fmt.Println("isGoodGaAndGb error")
		return nil
	}

	//fmt.Println("uBytes: " + hex.EncodeToString(uBytes))
	S := new(big.Int).Exp(BKgx, new(big.Int).Add(a, new(big.Int).Mul(u, x)), m.p)
	SBytes := getBigIntegerBytes(S)
	KBytes := calcSHA256(SBytes)
	pHash := calcSHA256(m.P)
	gHash := calcSHA256(m.gBytes)

	for i := 0; i < len(pHash); i++ {
		pHash[i] = gHash[i] ^ pHash[i]
	}

	return calcSHA256(pHash, calcSHA256(newSalt1), calcSHA256(m.Salt2), ABytes, BBytes, KBytes)
}
