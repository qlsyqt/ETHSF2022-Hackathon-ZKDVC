package codec

import (
	"encoding/binary"
	ethcommon "github.com/ethereum/go-ethereum/common"
	core "github.com/iden3/go-iden3-core"
	"issuerserver/models/common"
	"issuerserver/pkg/dvc"
)

func EncodeDcpClaim(claim *core.Claim, dcp *common.DCP) error {
	indexSlotA := encodeUtf8String(dcp.DataCategory)
	subcategoryBytes := dvc.EncodeSubcategory(dcp.DataCategory, dcp.SubCategory)
	indexSlotB := encodeBytes(subcategoryBytes)

	valueSlotA := encodeBoundaries(&dcp.LowerBoundary, &dcp.UpperBoundary)
	valueSlotB := encodeEthAddress(dcp.HolderAddress)

	err := claim.SetIndexData(indexSlotA, indexSlotB)
	if err != nil {
		return err
	}
	err = claim.SetValueData(valueSlotA, valueSlotB)
	if err != nil {
		return err
	}
	return nil
}

func DecodeDcpClaim(claim *core.Claim) *common.DCP {
	indexSlots, valueSlots := claim.RawSlots()
	idxSlotA, idxSlotB := indexSlots[2], indexSlots[3]
	valSlotA, valSlotB := valueSlots[2], valueSlots[3]

	dataCategory := decodeUtf8String(idxSlotA)
	lb, ub := decodeBoundaries(valSlotA)
	dcp := &common.DCP{
		DataCategory:  dataCategory,
		SubCategory:   dvc.DecodeSubcategory(dataCategory, decodeBytes(idxSlotB)),
		LowerBoundary: *lb,
		UpperBoundary: *ub,
		HolderAddress: decodeEthAddress(valSlotB),
	}

	return dcp
}

// Convert string to bytes in slots(also make sure this it fits into slot which is 253bit)
func encodeUtf8String(str string) core.ElemBytes {
	bytes := []byte(str)
	return encodeBytes(bytes)
}

func decodeUtf8String(slot core.ElemBytes) string {
	return string(decodeBytes(slot))
}

func encodeBytes(bytes []byte) core.ElemBytes {
	slot := [32]byte{}

	slot[5] = uint8(len(bytes))
	copy(slot[6:], bytes)
	return slot
}

func decodeBytes(slot core.ElemBytes) []byte {
	length := slot[5]
	return slot[6 : 6+length]
}

// 1~4 bytes: Big endian of value
// 5th byte: stores include
func encodeBoundary(boundary *common.Boundary) [5]byte {
	b := [5]byte{}
	binary.BigEndian.PutUint32(b[:4], uint32(boundary.Value))
	if boundary.Include {
		b[4] = 1
	}
	return b
}

func decodeBoundary(slot []byte) *common.Boundary {
	r := binary.BigEndian.Uint32(slot[:4])
	val := int32(r)
	include := false
	if slot[4] == 1 {
		include = true
	}
	return &common.Boundary{
		Include: include,
		Value:   val,
	}
}

func encodeBoundaries(lb, ub *common.Boundary) core.ElemBytes {
	slot := [32]byte{}
	lbBytes := encodeBoundary(lb)
	ubBytes := encodeBoundary(ub)
	copy(slot[:5], lbBytes[:])
	copy(slot[5:10], ubBytes[:])
	return slot
}

func decodeBoundaries(slot core.ElemBytes) (*common.Boundary, *common.Boundary) {
	lbBytes := slot[:5]
	ubBytes := slot[5:10]

	return decodeBoundary(lbBytes), decodeBoundary(ubBytes)
}

// Encode 20 bytes of wallet
func encodeEthAddress(walletAddress string) core.ElemBytes {
	addressBytes := ethcommon.FromHex(walletAddress)
	slot := [32]byte{}
	copy(slot[5:], addressBytes)
	return slot
}

func decodeEthAddress(bytes core.ElemBytes) string {
	return ethcommon.BytesToAddress(bytes[5:25]).String()
}
