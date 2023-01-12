package cmd

// Code generated by github.com/tinylib/msgp DO NOT EDIT.

import (
	"github.com/tinylib/msgp/msgp"
)

// DecodeMsg implements msgp.Decodable
func (z *SiteResyncStatus) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, err = dc.ReadMapHeader()
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		switch msgp.UnsafeString(field) {
		case "v":
			z.Version, err = dc.ReadInt()
			if err != nil {
				err = msgp.WrapError(err, "Version")
				return
			}
		case "ss":
			err = z.Status.DecodeMsg(dc)
			if err != nil {
				err = msgp.WrapError(err, "Status")
				return
			}
		case "did":
			z.DeplID, err = dc.ReadString()
			if err != nil {
				err = msgp.WrapError(err, "DeplID")
				return
			}
		case "bkts":
			var zb0002 uint32
			zb0002, err = dc.ReadMapHeader()
			if err != nil {
				err = msgp.WrapError(err, "BucketStatuses")
				return
			}
			if z.BucketStatuses == nil {
				z.BucketStatuses = make(map[string]ResyncStatusType, zb0002)
			} else if len(z.BucketStatuses) > 0 {
				for key := range z.BucketStatuses {
					delete(z.BucketStatuses, key)
				}
			}
			for zb0002 > 0 {
				zb0002--
				var za0001 string
				var za0002 ResyncStatusType
				za0001, err = dc.ReadString()
				if err != nil {
					err = msgp.WrapError(err, "BucketStatuses")
					return
				}
				err = za0002.DecodeMsg(dc)
				if err != nil {
					err = msgp.WrapError(err, "BucketStatuses", za0001)
					return
				}
				z.BucketStatuses[za0001] = za0002
			}
		case "tb":
			z.TotBuckets, err = dc.ReadInt()
			if err != nil {
				err = msgp.WrapError(err, "TotBuckets")
				return
			}
		case "cst":
			err = z.TargetReplicationResyncStatus.DecodeMsg(dc)
			if err != nil {
				err = msgp.WrapError(err, "TargetReplicationResyncStatus")
				return
			}
		default:
			err = dc.Skip()
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *SiteResyncStatus) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 6
	// write "v"
	err = en.Append(0x86, 0xa1, 0x76)
	if err != nil {
		return
	}
	err = en.WriteInt(z.Version)
	if err != nil {
		err = msgp.WrapError(err, "Version")
		return
	}
	// write "ss"
	err = en.Append(0xa2, 0x73, 0x73)
	if err != nil {
		return
	}
	err = z.Status.EncodeMsg(en)
	if err != nil {
		err = msgp.WrapError(err, "Status")
		return
	}
	// write "did"
	err = en.Append(0xa3, 0x64, 0x69, 0x64)
	if err != nil {
		return
	}
	err = en.WriteString(z.DeplID)
	if err != nil {
		err = msgp.WrapError(err, "DeplID")
		return
	}
	// write "bkts"
	err = en.Append(0xa4, 0x62, 0x6b, 0x74, 0x73)
	if err != nil {
		return
	}
	err = en.WriteMapHeader(uint32(len(z.BucketStatuses)))
	if err != nil {
		err = msgp.WrapError(err, "BucketStatuses")
		return
	}
	for za0001, za0002 := range z.BucketStatuses {
		err = en.WriteString(za0001)
		if err != nil {
			err = msgp.WrapError(err, "BucketStatuses")
			return
		}
		err = za0002.EncodeMsg(en)
		if err != nil {
			err = msgp.WrapError(err, "BucketStatuses", za0001)
			return
		}
	}
	// write "tb"
	err = en.Append(0xa2, 0x74, 0x62)
	if err != nil {
		return
	}
	err = en.WriteInt(z.TotBuckets)
	if err != nil {
		err = msgp.WrapError(err, "TotBuckets")
		return
	}
	// write "cst"
	err = en.Append(0xa3, 0x63, 0x73, 0x74)
	if err != nil {
		return
	}
	err = z.TargetReplicationResyncStatus.EncodeMsg(en)
	if err != nil {
		err = msgp.WrapError(err, "TargetReplicationResyncStatus")
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *SiteResyncStatus) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 6
	// string "v"
	o = append(o, 0x86, 0xa1, 0x76)
	o = msgp.AppendInt(o, z.Version)
	// string "ss"
	o = append(o, 0xa2, 0x73, 0x73)
	o, err = z.Status.MarshalMsg(o)
	if err != nil {
		err = msgp.WrapError(err, "Status")
		return
	}
	// string "did"
	o = append(o, 0xa3, 0x64, 0x69, 0x64)
	o = msgp.AppendString(o, z.DeplID)
	// string "bkts"
	o = append(o, 0xa4, 0x62, 0x6b, 0x74, 0x73)
	o = msgp.AppendMapHeader(o, uint32(len(z.BucketStatuses)))
	for za0001, za0002 := range z.BucketStatuses {
		o = msgp.AppendString(o, za0001)
		o, err = za0002.MarshalMsg(o)
		if err != nil {
			err = msgp.WrapError(err, "BucketStatuses", za0001)
			return
		}
	}
	// string "tb"
	o = append(o, 0xa2, 0x74, 0x62)
	o = msgp.AppendInt(o, z.TotBuckets)
	// string "cst"
	o = append(o, 0xa3, 0x63, 0x73, 0x74)
	o, err = z.TargetReplicationResyncStatus.MarshalMsg(o)
	if err != nil {
		err = msgp.WrapError(err, "TargetReplicationResyncStatus")
		return
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *SiteResyncStatus) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		switch msgp.UnsafeString(field) {
		case "v":
			z.Version, bts, err = msgp.ReadIntBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Version")
				return
			}
		case "ss":
			bts, err = z.Status.UnmarshalMsg(bts)
			if err != nil {
				err = msgp.WrapError(err, "Status")
				return
			}
		case "did":
			z.DeplID, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "DeplID")
				return
			}
		case "bkts":
			var zb0002 uint32
			zb0002, bts, err = msgp.ReadMapHeaderBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "BucketStatuses")
				return
			}
			if z.BucketStatuses == nil {
				z.BucketStatuses = make(map[string]ResyncStatusType, zb0002)
			} else if len(z.BucketStatuses) > 0 {
				for key := range z.BucketStatuses {
					delete(z.BucketStatuses, key)
				}
			}
			for zb0002 > 0 {
				var za0001 string
				var za0002 ResyncStatusType
				zb0002--
				za0001, bts, err = msgp.ReadStringBytes(bts)
				if err != nil {
					err = msgp.WrapError(err, "BucketStatuses")
					return
				}
				bts, err = za0002.UnmarshalMsg(bts)
				if err != nil {
					err = msgp.WrapError(err, "BucketStatuses", za0001)
					return
				}
				z.BucketStatuses[za0001] = za0002
			}
		case "tb":
			z.TotBuckets, bts, err = msgp.ReadIntBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "TotBuckets")
				return
			}
		case "cst":
			bts, err = z.TargetReplicationResyncStatus.UnmarshalMsg(bts)
			if err != nil {
				err = msgp.WrapError(err, "TargetReplicationResyncStatus")
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *SiteResyncStatus) Msgsize() (s int) {
	s = 1 + 2 + msgp.IntSize + 3 + z.Status.Msgsize() + 4 + msgp.StringPrefixSize + len(z.DeplID) + 5 + msgp.MapHeaderSize
	if z.BucketStatuses != nil {
		for za0001, za0002 := range z.BucketStatuses {
			_ = za0002
			s += msgp.StringPrefixSize + len(za0001) + za0002.Msgsize()
		}
	}
	s += 3 + msgp.IntSize + 4 + z.TargetReplicationResyncStatus.Msgsize()
	return
}