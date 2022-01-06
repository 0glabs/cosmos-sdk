package types

import (
	fmt "fmt"
	runtime "github.com/cosmos/cosmos-proto/runtime"
	_ "github.com/gogo/protobuf/gogoproto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoiface "google.golang.org/protobuf/runtime/protoiface"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	io "io"
	reflect "reflect"
	sync "sync"
)

var (
	md_ConsensusParams           protoreflect.MessageDescriptor
	fd_ConsensusParams_block     protoreflect.FieldDescriptor
	fd_ConsensusParams_evidence  protoreflect.FieldDescriptor
	fd_ConsensusParams_validator protoreflect.FieldDescriptor
	fd_ConsensusParams_version   protoreflect.FieldDescriptor
)

func init() {
	file_tendermint_types_params_proto_init()
	md_ConsensusParams = File_tendermint_types_params_proto.Messages().ByName("ConsensusParams")
	fd_ConsensusParams_block = md_ConsensusParams.Fields().ByName("block")
	fd_ConsensusParams_evidence = md_ConsensusParams.Fields().ByName("evidence")
	fd_ConsensusParams_validator = md_ConsensusParams.Fields().ByName("validator")
	fd_ConsensusParams_version = md_ConsensusParams.Fields().ByName("version")
}

var _ protoreflect.Message = (*fastReflection_ConsensusParams)(nil)

type fastReflection_ConsensusParams ConsensusParams

func (x *ConsensusParams) ProtoReflect() protoreflect.Message {
	return (*fastReflection_ConsensusParams)(x)
}

func (x *ConsensusParams) slowProtoReflect() protoreflect.Message {
	mi := &file_tendermint_types_params_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_ConsensusParams_messageType fastReflection_ConsensusParams_messageType
var _ protoreflect.MessageType = fastReflection_ConsensusParams_messageType{}

type fastReflection_ConsensusParams_messageType struct{}

func (x fastReflection_ConsensusParams_messageType) Zero() protoreflect.Message {
	return (*fastReflection_ConsensusParams)(nil)
}
func (x fastReflection_ConsensusParams_messageType) New() protoreflect.Message {
	return new(fastReflection_ConsensusParams)
}
func (x fastReflection_ConsensusParams_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_ConsensusParams
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_ConsensusParams) Descriptor() protoreflect.MessageDescriptor {
	return md_ConsensusParams
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_ConsensusParams) Type() protoreflect.MessageType {
	return _fastReflection_ConsensusParams_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_ConsensusParams) New() protoreflect.Message {
	return new(fastReflection_ConsensusParams)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_ConsensusParams) Interface() protoreflect.ProtoMessage {
	return (*ConsensusParams)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_ConsensusParams) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.Block != nil {
		value := protoreflect.ValueOfMessage(x.Block.ProtoReflect())
		if !f(fd_ConsensusParams_block, value) {
			return
		}
	}
	if x.Evidence != nil {
		value := protoreflect.ValueOfMessage(x.Evidence.ProtoReflect())
		if !f(fd_ConsensusParams_evidence, value) {
			return
		}
	}
	if x.Validator != nil {
		value := protoreflect.ValueOfMessage(x.Validator.ProtoReflect())
		if !f(fd_ConsensusParams_validator, value) {
			return
		}
	}
	if x.Version != nil {
		value := protoreflect.ValueOfMessage(x.Version.ProtoReflect())
		if !f(fd_ConsensusParams_version, value) {
			return
		}
	}
}

// Has reports whether a field is populated.
//
// Some fields have the property of nullability where it is possible to
// distinguish between the default value of a field and whether the field
// was explicitly populated with the default value. Singular message fields,
// member fields of a oneof, and proto2 scalar fields are nullable. Such
// fields are populated only if explicitly set.
//
// In other cases (aside from the nullable cases above),
// a proto3 scalar field is populated if it contains a non-zero value, and
// a repeated field is populated if it is non-empty.
func (x *fastReflection_ConsensusParams) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "tendermint.types.ConsensusParams.block":
		return x.Block != nil
	case "tendermint.types.ConsensusParams.evidence":
		return x.Evidence != nil
	case "tendermint.types.ConsensusParams.validator":
		return x.Validator != nil
	case "tendermint.types.ConsensusParams.version":
		return x.Version != nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: tendermint.types.ConsensusParams"))
		}
		panic(fmt.Errorf("message tendermint.types.ConsensusParams does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_ConsensusParams) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "tendermint.types.ConsensusParams.block":
		x.Block = nil
	case "tendermint.types.ConsensusParams.evidence":
		x.Evidence = nil
	case "tendermint.types.ConsensusParams.validator":
		x.Validator = nil
	case "tendermint.types.ConsensusParams.version":
		x.Version = nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: tendermint.types.ConsensusParams"))
		}
		panic(fmt.Errorf("message tendermint.types.ConsensusParams does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_ConsensusParams) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "tendermint.types.ConsensusParams.block":
		value := x.Block
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	case "tendermint.types.ConsensusParams.evidence":
		value := x.Evidence
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	case "tendermint.types.ConsensusParams.validator":
		value := x.Validator
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	case "tendermint.types.ConsensusParams.version":
		value := x.Version
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: tendermint.types.ConsensusParams"))
		}
		panic(fmt.Errorf("message tendermint.types.ConsensusParams does not contain field %s", descriptor.FullName()))
	}
}

// Set stores the value for a field.
//
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType.
// When setting a composite type, it is unspecified whether the stored value
// aliases the source's memory in any way. If the composite value is an
// empty, read-only value, then it panics.
//
// Set is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_ConsensusParams) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "tendermint.types.ConsensusParams.block":
		x.Block = value.Message().Interface().(*BlockParams)
	case "tendermint.types.ConsensusParams.evidence":
		x.Evidence = value.Message().Interface().(*EvidenceParams)
	case "tendermint.types.ConsensusParams.validator":
		x.Validator = value.Message().Interface().(*ValidatorParams)
	case "tendermint.types.ConsensusParams.version":
		x.Version = value.Message().Interface().(*VersionParams)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: tendermint.types.ConsensusParams"))
		}
		panic(fmt.Errorf("message tendermint.types.ConsensusParams does not contain field %s", fd.FullName()))
	}
}

// Mutable returns a mutable reference to a composite type.
//
// If the field is unpopulated, it may allocate a composite value.
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType
// if not already stored.
// It panics if the field does not contain a composite type.
//
// Mutable is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_ConsensusParams) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "tendermint.types.ConsensusParams.block":
		if x.Block == nil {
			x.Block = new(BlockParams)
		}
		return protoreflect.ValueOfMessage(x.Block.ProtoReflect())
	case "tendermint.types.ConsensusParams.evidence":
		if x.Evidence == nil {
			x.Evidence = new(EvidenceParams)
		}
		return protoreflect.ValueOfMessage(x.Evidence.ProtoReflect())
	case "tendermint.types.ConsensusParams.validator":
		if x.Validator == nil {
			x.Validator = new(ValidatorParams)
		}
		return protoreflect.ValueOfMessage(x.Validator.ProtoReflect())
	case "tendermint.types.ConsensusParams.version":
		if x.Version == nil {
			x.Version = new(VersionParams)
		}
		return protoreflect.ValueOfMessage(x.Version.ProtoReflect())
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: tendermint.types.ConsensusParams"))
		}
		panic(fmt.Errorf("message tendermint.types.ConsensusParams does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_ConsensusParams) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "tendermint.types.ConsensusParams.block":
		m := new(BlockParams)
		return protoreflect.ValueOfMessage(m.ProtoReflect())
	case "tendermint.types.ConsensusParams.evidence":
		m := new(EvidenceParams)
		return protoreflect.ValueOfMessage(m.ProtoReflect())
	case "tendermint.types.ConsensusParams.validator":
		m := new(ValidatorParams)
		return protoreflect.ValueOfMessage(m.ProtoReflect())
	case "tendermint.types.ConsensusParams.version":
		m := new(VersionParams)
		return protoreflect.ValueOfMessage(m.ProtoReflect())
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: tendermint.types.ConsensusParams"))
		}
		panic(fmt.Errorf("message tendermint.types.ConsensusParams does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_ConsensusParams) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in tendermint.types.ConsensusParams", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_ConsensusParams) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_ConsensusParams) SetUnknown(fields protoreflect.RawFields) {
	x.unknownFields = fields
}

// IsValid reports whether the message is valid.
//
// An invalid message is an empty, read-only value.
//
// An invalid message often corresponds to a nil pointer of the concrete
// message type, but the details are implementation dependent.
// Validity is not part of the protobuf data model, and may not
// be preserved in marshaling or other operations.
func (x *fastReflection_ConsensusParams) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_ConsensusParams) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*ConsensusParams)
		if x == nil {
			return protoiface.SizeOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Size:              0,
			}
		}
		options := runtime.SizeInputToOptions(input)
		_ = options
		var n int
		var l int
		_ = l
		if x.Block != nil {
			l = options.Size(x.Block)
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.Evidence != nil {
			l = options.Size(x.Evidence)
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.Validator != nil {
			l = options.Size(x.Validator)
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.Version != nil {
			l = options.Size(x.Version)
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*ConsensusParams)
		if x == nil {
			return protoiface.MarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Buf:               input.Buf,
			}, nil
		}
		options := runtime.MarshalInputToOptions(input)
		_ = options
		size := options.Size(x)
		dAtA := make([]byte, size)
		i := len(dAtA)
		_ = i
		var l int
		_ = l
		if x.unknownFields != nil {
			i -= len(x.unknownFields)
			copy(dAtA[i:], x.unknownFields)
		}
		if x.Version != nil {
			encoded, err := options.Marshal(x.Version)
			if err != nil {
				return protoiface.MarshalOutput{
					NoUnkeyedLiterals: input.NoUnkeyedLiterals,
					Buf:               input.Buf,
				}, err
			}
			i -= len(encoded)
			copy(dAtA[i:], encoded)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(encoded)))
			i--
			dAtA[i] = 0x22
		}
		if x.Validator != nil {
			encoded, err := options.Marshal(x.Validator)
			if err != nil {
				return protoiface.MarshalOutput{
					NoUnkeyedLiterals: input.NoUnkeyedLiterals,
					Buf:               input.Buf,
				}, err
			}
			i -= len(encoded)
			copy(dAtA[i:], encoded)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(encoded)))
			i--
			dAtA[i] = 0x1a
		}
		if x.Evidence != nil {
			encoded, err := options.Marshal(x.Evidence)
			if err != nil {
				return protoiface.MarshalOutput{
					NoUnkeyedLiterals: input.NoUnkeyedLiterals,
					Buf:               input.Buf,
				}, err
			}
			i -= len(encoded)
			copy(dAtA[i:], encoded)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(encoded)))
			i--
			dAtA[i] = 0x12
		}
		if x.Block != nil {
			encoded, err := options.Marshal(x.Block)
			if err != nil {
				return protoiface.MarshalOutput{
					NoUnkeyedLiterals: input.NoUnkeyedLiterals,
					Buf:               input.Buf,
				}, err
			}
			i -= len(encoded)
			copy(dAtA[i:], encoded)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(encoded)))
			i--
			dAtA[i] = 0xa
		}
		if input.Buf != nil {
			input.Buf = append(input.Buf, dAtA...)
		} else {
			input.Buf = dAtA
		}
		return protoiface.MarshalOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Buf:               input.Buf,
		}, nil
	}
	unmarshal := func(input protoiface.UnmarshalInput) (protoiface.UnmarshalOutput, error) {
		x := input.Message.Interface().(*ConsensusParams)
		if x == nil {
			return protoiface.UnmarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Flags:             input.Flags,
			}, nil
		}
		options := runtime.UnmarshalInputToOptions(input)
		_ = options
		dAtA := input.Buf
		l := len(dAtA)
		iNdEx := 0
		for iNdEx < l {
			preIndex := iNdEx
			var wire uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
				}
				if iNdEx >= l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				wire |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			fieldNum := int32(wire >> 3)
			wireType := int(wire & 0x7)
			if wireType == 4 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: ConsensusParams: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: ConsensusParams: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Block", wireType)
				}
				var msglen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					msglen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if msglen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + msglen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if x.Block == nil {
					x.Block = &BlockParams{}
				}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.Block); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Evidence", wireType)
				}
				var msglen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					msglen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if msglen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + msglen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if x.Evidence == nil {
					x.Evidence = &EvidenceParams{}
				}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.Evidence); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			case 3:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Validator", wireType)
				}
				var msglen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					msglen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if msglen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + msglen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if x.Validator == nil {
					x.Validator = &ValidatorParams{}
				}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.Validator); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			case 4:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Version", wireType)
				}
				var msglen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					msglen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if msglen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + msglen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if x.Version == nil {
					x.Version = &VersionParams{}
				}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.Version); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			default:
				iNdEx = preIndex
				skippy, err := runtime.Skip(dAtA[iNdEx:])
				if err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				if (skippy < 0) || (iNdEx+skippy) < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if (iNdEx + skippy) > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if !options.DiscardUnknown {
					x.unknownFields = append(x.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
				}
				iNdEx += skippy
			}
		}

		if iNdEx > l {
			return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
		}
		return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, nil
	}
	return &protoiface.Methods{
		NoUnkeyedLiterals: struct{}{},
		Flags:             protoiface.SupportMarshalDeterministic | protoiface.SupportUnmarshalDiscardUnknown,
		Size:              size,
		Marshal:           marshal,
		Unmarshal:         unmarshal,
		Merge:             nil,
		CheckInitialized:  nil,
	}
}

var (
	md_BlockParams              protoreflect.MessageDescriptor
	fd_BlockParams_max_bytes    protoreflect.FieldDescriptor
	fd_BlockParams_max_gas      protoreflect.FieldDescriptor
	fd_BlockParams_time_iota_ms protoreflect.FieldDescriptor
)

func init() {
	file_tendermint_types_params_proto_init()
	md_BlockParams = File_tendermint_types_params_proto.Messages().ByName("BlockParams")
	fd_BlockParams_max_bytes = md_BlockParams.Fields().ByName("max_bytes")
	fd_BlockParams_max_gas = md_BlockParams.Fields().ByName("max_gas")
	fd_BlockParams_time_iota_ms = md_BlockParams.Fields().ByName("time_iota_ms")
}

var _ protoreflect.Message = (*fastReflection_BlockParams)(nil)

type fastReflection_BlockParams BlockParams

func (x *BlockParams) ProtoReflect() protoreflect.Message {
	return (*fastReflection_BlockParams)(x)
}

func (x *BlockParams) slowProtoReflect() protoreflect.Message {
	mi := &file_tendermint_types_params_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_BlockParams_messageType fastReflection_BlockParams_messageType
var _ protoreflect.MessageType = fastReflection_BlockParams_messageType{}

type fastReflection_BlockParams_messageType struct{}

func (x fastReflection_BlockParams_messageType) Zero() protoreflect.Message {
	return (*fastReflection_BlockParams)(nil)
}
func (x fastReflection_BlockParams_messageType) New() protoreflect.Message {
	return new(fastReflection_BlockParams)
}
func (x fastReflection_BlockParams_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_BlockParams
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_BlockParams) Descriptor() protoreflect.MessageDescriptor {
	return md_BlockParams
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_BlockParams) Type() protoreflect.MessageType {
	return _fastReflection_BlockParams_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_BlockParams) New() protoreflect.Message {
	return new(fastReflection_BlockParams)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_BlockParams) Interface() protoreflect.ProtoMessage {
	return (*BlockParams)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_BlockParams) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.MaxBytes != int64(0) {
		value := protoreflect.ValueOfInt64(x.MaxBytes)
		if !f(fd_BlockParams_max_bytes, value) {
			return
		}
	}
	if x.MaxGas != int64(0) {
		value := protoreflect.ValueOfInt64(x.MaxGas)
		if !f(fd_BlockParams_max_gas, value) {
			return
		}
	}
	if x.TimeIotaMs != int64(0) {
		value := protoreflect.ValueOfInt64(x.TimeIotaMs)
		if !f(fd_BlockParams_time_iota_ms, value) {
			return
		}
	}
}

// Has reports whether a field is populated.
//
// Some fields have the property of nullability where it is possible to
// distinguish between the default value of a field and whether the field
// was explicitly populated with the default value. Singular message fields,
// member fields of a oneof, and proto2 scalar fields are nullable. Such
// fields are populated only if explicitly set.
//
// In other cases (aside from the nullable cases above),
// a proto3 scalar field is populated if it contains a non-zero value, and
// a repeated field is populated if it is non-empty.
func (x *fastReflection_BlockParams) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "tendermint.types.BlockParams.max_bytes":
		return x.MaxBytes != int64(0)
	case "tendermint.types.BlockParams.max_gas":
		return x.MaxGas != int64(0)
	case "tendermint.types.BlockParams.time_iota_ms":
		return x.TimeIotaMs != int64(0)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: tendermint.types.BlockParams"))
		}
		panic(fmt.Errorf("message tendermint.types.BlockParams does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_BlockParams) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "tendermint.types.BlockParams.max_bytes":
		x.MaxBytes = int64(0)
	case "tendermint.types.BlockParams.max_gas":
		x.MaxGas = int64(0)
	case "tendermint.types.BlockParams.time_iota_ms":
		x.TimeIotaMs = int64(0)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: tendermint.types.BlockParams"))
		}
		panic(fmt.Errorf("message tendermint.types.BlockParams does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_BlockParams) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "tendermint.types.BlockParams.max_bytes":
		value := x.MaxBytes
		return protoreflect.ValueOfInt64(value)
	case "tendermint.types.BlockParams.max_gas":
		value := x.MaxGas
		return protoreflect.ValueOfInt64(value)
	case "tendermint.types.BlockParams.time_iota_ms":
		value := x.TimeIotaMs
		return protoreflect.ValueOfInt64(value)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: tendermint.types.BlockParams"))
		}
		panic(fmt.Errorf("message tendermint.types.BlockParams does not contain field %s", descriptor.FullName()))
	}
}

// Set stores the value for a field.
//
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType.
// When setting a composite type, it is unspecified whether the stored value
// aliases the source's memory in any way. If the composite value is an
// empty, read-only value, then it panics.
//
// Set is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_BlockParams) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "tendermint.types.BlockParams.max_bytes":
		x.MaxBytes = value.Int()
	case "tendermint.types.BlockParams.max_gas":
		x.MaxGas = value.Int()
	case "tendermint.types.BlockParams.time_iota_ms":
		x.TimeIotaMs = value.Int()
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: tendermint.types.BlockParams"))
		}
		panic(fmt.Errorf("message tendermint.types.BlockParams does not contain field %s", fd.FullName()))
	}
}

// Mutable returns a mutable reference to a composite type.
//
// If the field is unpopulated, it may allocate a composite value.
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType
// if not already stored.
// It panics if the field does not contain a composite type.
//
// Mutable is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_BlockParams) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "tendermint.types.BlockParams.max_bytes":
		panic(fmt.Errorf("field max_bytes of message tendermint.types.BlockParams is not mutable"))
	case "tendermint.types.BlockParams.max_gas":
		panic(fmt.Errorf("field max_gas of message tendermint.types.BlockParams is not mutable"))
	case "tendermint.types.BlockParams.time_iota_ms":
		panic(fmt.Errorf("field time_iota_ms of message tendermint.types.BlockParams is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: tendermint.types.BlockParams"))
		}
		panic(fmt.Errorf("message tendermint.types.BlockParams does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_BlockParams) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "tendermint.types.BlockParams.max_bytes":
		return protoreflect.ValueOfInt64(int64(0))
	case "tendermint.types.BlockParams.max_gas":
		return protoreflect.ValueOfInt64(int64(0))
	case "tendermint.types.BlockParams.time_iota_ms":
		return protoreflect.ValueOfInt64(int64(0))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: tendermint.types.BlockParams"))
		}
		panic(fmt.Errorf("message tendermint.types.BlockParams does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_BlockParams) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in tendermint.types.BlockParams", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_BlockParams) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_BlockParams) SetUnknown(fields protoreflect.RawFields) {
	x.unknownFields = fields
}

// IsValid reports whether the message is valid.
//
// An invalid message is an empty, read-only value.
//
// An invalid message often corresponds to a nil pointer of the concrete
// message type, but the details are implementation dependent.
// Validity is not part of the protobuf data model, and may not
// be preserved in marshaling or other operations.
func (x *fastReflection_BlockParams) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_BlockParams) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*BlockParams)
		if x == nil {
			return protoiface.SizeOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Size:              0,
			}
		}
		options := runtime.SizeInputToOptions(input)
		_ = options
		var n int
		var l int
		_ = l
		if x.MaxBytes != 0 {
			n += 1 + runtime.Sov(uint64(x.MaxBytes))
		}
		if x.MaxGas != 0 {
			n += 1 + runtime.Sov(uint64(x.MaxGas))
		}
		if x.TimeIotaMs != 0 {
			n += 1 + runtime.Sov(uint64(x.TimeIotaMs))
		}
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*BlockParams)
		if x == nil {
			return protoiface.MarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Buf:               input.Buf,
			}, nil
		}
		options := runtime.MarshalInputToOptions(input)
		_ = options
		size := options.Size(x)
		dAtA := make([]byte, size)
		i := len(dAtA)
		_ = i
		var l int
		_ = l
		if x.unknownFields != nil {
			i -= len(x.unknownFields)
			copy(dAtA[i:], x.unknownFields)
		}
		if x.TimeIotaMs != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.TimeIotaMs))
			i--
			dAtA[i] = 0x18
		}
		if x.MaxGas != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.MaxGas))
			i--
			dAtA[i] = 0x10
		}
		if x.MaxBytes != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.MaxBytes))
			i--
			dAtA[i] = 0x8
		}
		if input.Buf != nil {
			input.Buf = append(input.Buf, dAtA...)
		} else {
			input.Buf = dAtA
		}
		return protoiface.MarshalOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Buf:               input.Buf,
		}, nil
	}
	unmarshal := func(input protoiface.UnmarshalInput) (protoiface.UnmarshalOutput, error) {
		x := input.Message.Interface().(*BlockParams)
		if x == nil {
			return protoiface.UnmarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Flags:             input.Flags,
			}, nil
		}
		options := runtime.UnmarshalInputToOptions(input)
		_ = options
		dAtA := input.Buf
		l := len(dAtA)
		iNdEx := 0
		for iNdEx < l {
			preIndex := iNdEx
			var wire uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
				}
				if iNdEx >= l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				wire |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			fieldNum := int32(wire >> 3)
			wireType := int(wire & 0x7)
			if wireType == 4 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: BlockParams: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: BlockParams: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field MaxBytes", wireType)
				}
				x.MaxBytes = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.MaxBytes |= int64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
			case 2:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field MaxGas", wireType)
				}
				x.MaxGas = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.MaxGas |= int64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
			case 3:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field TimeIotaMs", wireType)
				}
				x.TimeIotaMs = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.TimeIotaMs |= int64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
			default:
				iNdEx = preIndex
				skippy, err := runtime.Skip(dAtA[iNdEx:])
				if err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				if (skippy < 0) || (iNdEx+skippy) < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if (iNdEx + skippy) > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if !options.DiscardUnknown {
					x.unknownFields = append(x.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
				}
				iNdEx += skippy
			}
		}

		if iNdEx > l {
			return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
		}
		return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, nil
	}
	return &protoiface.Methods{
		NoUnkeyedLiterals: struct{}{},
		Flags:             protoiface.SupportMarshalDeterministic | protoiface.SupportUnmarshalDiscardUnknown,
		Size:              size,
		Marshal:           marshal,
		Unmarshal:         unmarshal,
		Merge:             nil,
		CheckInitialized:  nil,
	}
}

var (
	md_EvidenceParams                    protoreflect.MessageDescriptor
	fd_EvidenceParams_max_age_num_blocks protoreflect.FieldDescriptor
	fd_EvidenceParams_max_age_duration   protoreflect.FieldDescriptor
	fd_EvidenceParams_max_bytes          protoreflect.FieldDescriptor
)

func init() {
	file_tendermint_types_params_proto_init()
	md_EvidenceParams = File_tendermint_types_params_proto.Messages().ByName("EvidenceParams")
	fd_EvidenceParams_max_age_num_blocks = md_EvidenceParams.Fields().ByName("max_age_num_blocks")
	fd_EvidenceParams_max_age_duration = md_EvidenceParams.Fields().ByName("max_age_duration")
	fd_EvidenceParams_max_bytes = md_EvidenceParams.Fields().ByName("max_bytes")
}

var _ protoreflect.Message = (*fastReflection_EvidenceParams)(nil)

type fastReflection_EvidenceParams EvidenceParams

func (x *EvidenceParams) ProtoReflect() protoreflect.Message {
	return (*fastReflection_EvidenceParams)(x)
}

func (x *EvidenceParams) slowProtoReflect() protoreflect.Message {
	mi := &file_tendermint_types_params_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_EvidenceParams_messageType fastReflection_EvidenceParams_messageType
var _ protoreflect.MessageType = fastReflection_EvidenceParams_messageType{}

type fastReflection_EvidenceParams_messageType struct{}

func (x fastReflection_EvidenceParams_messageType) Zero() protoreflect.Message {
	return (*fastReflection_EvidenceParams)(nil)
}
func (x fastReflection_EvidenceParams_messageType) New() protoreflect.Message {
	return new(fastReflection_EvidenceParams)
}
func (x fastReflection_EvidenceParams_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_EvidenceParams
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_EvidenceParams) Descriptor() protoreflect.MessageDescriptor {
	return md_EvidenceParams
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_EvidenceParams) Type() protoreflect.MessageType {
	return _fastReflection_EvidenceParams_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_EvidenceParams) New() protoreflect.Message {
	return new(fastReflection_EvidenceParams)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_EvidenceParams) Interface() protoreflect.ProtoMessage {
	return (*EvidenceParams)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_EvidenceParams) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.MaxAgeNumBlocks != int64(0) {
		value := protoreflect.ValueOfInt64(x.MaxAgeNumBlocks)
		if !f(fd_EvidenceParams_max_age_num_blocks, value) {
			return
		}
	}
	if x.MaxAgeDuration != nil {
		value := protoreflect.ValueOfMessage(x.MaxAgeDuration.ProtoReflect())
		if !f(fd_EvidenceParams_max_age_duration, value) {
			return
		}
	}
	if x.MaxBytes != int64(0) {
		value := protoreflect.ValueOfInt64(x.MaxBytes)
		if !f(fd_EvidenceParams_max_bytes, value) {
			return
		}
	}
}

// Has reports whether a field is populated.
//
// Some fields have the property of nullability where it is possible to
// distinguish between the default value of a field and whether the field
// was explicitly populated with the default value. Singular message fields,
// member fields of a oneof, and proto2 scalar fields are nullable. Such
// fields are populated only if explicitly set.
//
// In other cases (aside from the nullable cases above),
// a proto3 scalar field is populated if it contains a non-zero value, and
// a repeated field is populated if it is non-empty.
func (x *fastReflection_EvidenceParams) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "tendermint.types.EvidenceParams.max_age_num_blocks":
		return x.MaxAgeNumBlocks != int64(0)
	case "tendermint.types.EvidenceParams.max_age_duration":
		return x.MaxAgeDuration != nil
	case "tendermint.types.EvidenceParams.max_bytes":
		return x.MaxBytes != int64(0)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: tendermint.types.EvidenceParams"))
		}
		panic(fmt.Errorf("message tendermint.types.EvidenceParams does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_EvidenceParams) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "tendermint.types.EvidenceParams.max_age_num_blocks":
		x.MaxAgeNumBlocks = int64(0)
	case "tendermint.types.EvidenceParams.max_age_duration":
		x.MaxAgeDuration = nil
	case "tendermint.types.EvidenceParams.max_bytes":
		x.MaxBytes = int64(0)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: tendermint.types.EvidenceParams"))
		}
		panic(fmt.Errorf("message tendermint.types.EvidenceParams does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_EvidenceParams) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "tendermint.types.EvidenceParams.max_age_num_blocks":
		value := x.MaxAgeNumBlocks
		return protoreflect.ValueOfInt64(value)
	case "tendermint.types.EvidenceParams.max_age_duration":
		value := x.MaxAgeDuration
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	case "tendermint.types.EvidenceParams.max_bytes":
		value := x.MaxBytes
		return protoreflect.ValueOfInt64(value)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: tendermint.types.EvidenceParams"))
		}
		panic(fmt.Errorf("message tendermint.types.EvidenceParams does not contain field %s", descriptor.FullName()))
	}
}

// Set stores the value for a field.
//
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType.
// When setting a composite type, it is unspecified whether the stored value
// aliases the source's memory in any way. If the composite value is an
// empty, read-only value, then it panics.
//
// Set is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_EvidenceParams) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "tendermint.types.EvidenceParams.max_age_num_blocks":
		x.MaxAgeNumBlocks = value.Int()
	case "tendermint.types.EvidenceParams.max_age_duration":
		x.MaxAgeDuration = value.Message().Interface().(*durationpb.Duration)
	case "tendermint.types.EvidenceParams.max_bytes":
		x.MaxBytes = value.Int()
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: tendermint.types.EvidenceParams"))
		}
		panic(fmt.Errorf("message tendermint.types.EvidenceParams does not contain field %s", fd.FullName()))
	}
}

// Mutable returns a mutable reference to a composite type.
//
// If the field is unpopulated, it may allocate a composite value.
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType
// if not already stored.
// It panics if the field does not contain a composite type.
//
// Mutable is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_EvidenceParams) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "tendermint.types.EvidenceParams.max_age_duration":
		if x.MaxAgeDuration == nil {
			x.MaxAgeDuration = new(durationpb.Duration)
		}
		return protoreflect.ValueOfMessage(x.MaxAgeDuration.ProtoReflect())
	case "tendermint.types.EvidenceParams.max_age_num_blocks":
		panic(fmt.Errorf("field max_age_num_blocks of message tendermint.types.EvidenceParams is not mutable"))
	case "tendermint.types.EvidenceParams.max_bytes":
		panic(fmt.Errorf("field max_bytes of message tendermint.types.EvidenceParams is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: tendermint.types.EvidenceParams"))
		}
		panic(fmt.Errorf("message tendermint.types.EvidenceParams does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_EvidenceParams) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "tendermint.types.EvidenceParams.max_age_num_blocks":
		return protoreflect.ValueOfInt64(int64(0))
	case "tendermint.types.EvidenceParams.max_age_duration":
		m := new(durationpb.Duration)
		return protoreflect.ValueOfMessage(m.ProtoReflect())
	case "tendermint.types.EvidenceParams.max_bytes":
		return protoreflect.ValueOfInt64(int64(0))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: tendermint.types.EvidenceParams"))
		}
		panic(fmt.Errorf("message tendermint.types.EvidenceParams does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_EvidenceParams) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in tendermint.types.EvidenceParams", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_EvidenceParams) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_EvidenceParams) SetUnknown(fields protoreflect.RawFields) {
	x.unknownFields = fields
}

// IsValid reports whether the message is valid.
//
// An invalid message is an empty, read-only value.
//
// An invalid message often corresponds to a nil pointer of the concrete
// message type, but the details are implementation dependent.
// Validity is not part of the protobuf data model, and may not
// be preserved in marshaling or other operations.
func (x *fastReflection_EvidenceParams) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_EvidenceParams) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*EvidenceParams)
		if x == nil {
			return protoiface.SizeOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Size:              0,
			}
		}
		options := runtime.SizeInputToOptions(input)
		_ = options
		var n int
		var l int
		_ = l
		if x.MaxAgeNumBlocks != 0 {
			n += 1 + runtime.Sov(uint64(x.MaxAgeNumBlocks))
		}
		if x.MaxAgeDuration != nil {
			l = options.Size(x.MaxAgeDuration)
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.MaxBytes != 0 {
			n += 1 + runtime.Sov(uint64(x.MaxBytes))
		}
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*EvidenceParams)
		if x == nil {
			return protoiface.MarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Buf:               input.Buf,
			}, nil
		}
		options := runtime.MarshalInputToOptions(input)
		_ = options
		size := options.Size(x)
		dAtA := make([]byte, size)
		i := len(dAtA)
		_ = i
		var l int
		_ = l
		if x.unknownFields != nil {
			i -= len(x.unknownFields)
			copy(dAtA[i:], x.unknownFields)
		}
		if x.MaxBytes != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.MaxBytes))
			i--
			dAtA[i] = 0x18
		}
		if x.MaxAgeDuration != nil {
			encoded, err := options.Marshal(x.MaxAgeDuration)
			if err != nil {
				return protoiface.MarshalOutput{
					NoUnkeyedLiterals: input.NoUnkeyedLiterals,
					Buf:               input.Buf,
				}, err
			}
			i -= len(encoded)
			copy(dAtA[i:], encoded)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(encoded)))
			i--
			dAtA[i] = 0x12
		}
		if x.MaxAgeNumBlocks != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.MaxAgeNumBlocks))
			i--
			dAtA[i] = 0x8
		}
		if input.Buf != nil {
			input.Buf = append(input.Buf, dAtA...)
		} else {
			input.Buf = dAtA
		}
		return protoiface.MarshalOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Buf:               input.Buf,
		}, nil
	}
	unmarshal := func(input protoiface.UnmarshalInput) (protoiface.UnmarshalOutput, error) {
		x := input.Message.Interface().(*EvidenceParams)
		if x == nil {
			return protoiface.UnmarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Flags:             input.Flags,
			}, nil
		}
		options := runtime.UnmarshalInputToOptions(input)
		_ = options
		dAtA := input.Buf
		l := len(dAtA)
		iNdEx := 0
		for iNdEx < l {
			preIndex := iNdEx
			var wire uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
				}
				if iNdEx >= l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				wire |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			fieldNum := int32(wire >> 3)
			wireType := int(wire & 0x7)
			if wireType == 4 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: EvidenceParams: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: EvidenceParams: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field MaxAgeNumBlocks", wireType)
				}
				x.MaxAgeNumBlocks = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.MaxAgeNumBlocks |= int64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field MaxAgeDuration", wireType)
				}
				var msglen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					msglen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if msglen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + msglen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if x.MaxAgeDuration == nil {
					x.MaxAgeDuration = &durationpb.Duration{}
				}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.MaxAgeDuration); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			case 3:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field MaxBytes", wireType)
				}
				x.MaxBytes = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.MaxBytes |= int64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
			default:
				iNdEx = preIndex
				skippy, err := runtime.Skip(dAtA[iNdEx:])
				if err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				if (skippy < 0) || (iNdEx+skippy) < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if (iNdEx + skippy) > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if !options.DiscardUnknown {
					x.unknownFields = append(x.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
				}
				iNdEx += skippy
			}
		}

		if iNdEx > l {
			return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
		}
		return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, nil
	}
	return &protoiface.Methods{
		NoUnkeyedLiterals: struct{}{},
		Flags:             protoiface.SupportMarshalDeterministic | protoiface.SupportUnmarshalDiscardUnknown,
		Size:              size,
		Marshal:           marshal,
		Unmarshal:         unmarshal,
		Merge:             nil,
		CheckInitialized:  nil,
	}
}

var _ protoreflect.List = (*_ValidatorParams_1_list)(nil)

type _ValidatorParams_1_list struct {
	list *[]string
}

func (x *_ValidatorParams_1_list) Len() int {
	if x.list == nil {
		return 0
	}
	return len(*x.list)
}

func (x *_ValidatorParams_1_list) Get(i int) protoreflect.Value {
	return protoreflect.ValueOfString((*x.list)[i])
}

func (x *_ValidatorParams_1_list) Set(i int, value protoreflect.Value) {
	valueUnwrapped := value.String()
	concreteValue := valueUnwrapped
	(*x.list)[i] = concreteValue
}

func (x *_ValidatorParams_1_list) Append(value protoreflect.Value) {
	valueUnwrapped := value.String()
	concreteValue := valueUnwrapped
	*x.list = append(*x.list, concreteValue)
}

func (x *_ValidatorParams_1_list) AppendMutable() protoreflect.Value {
	panic(fmt.Errorf("AppendMutable can not be called on message ValidatorParams at list field PubKeyTypes as it is not of Message kind"))
}

func (x *_ValidatorParams_1_list) Truncate(n int) {
	*x.list = (*x.list)[:n]
}

func (x *_ValidatorParams_1_list) NewElement() protoreflect.Value {
	v := ""
	return protoreflect.ValueOfString(v)
}

func (x *_ValidatorParams_1_list) IsValid() bool {
	return x.list != nil
}

var (
	md_ValidatorParams               protoreflect.MessageDescriptor
	fd_ValidatorParams_pub_key_types protoreflect.FieldDescriptor
)

func init() {
	file_tendermint_types_params_proto_init()
	md_ValidatorParams = File_tendermint_types_params_proto.Messages().ByName("ValidatorParams")
	fd_ValidatorParams_pub_key_types = md_ValidatorParams.Fields().ByName("pub_key_types")
}

var _ protoreflect.Message = (*fastReflection_ValidatorParams)(nil)

type fastReflection_ValidatorParams ValidatorParams

func (x *ValidatorParams) ProtoReflect() protoreflect.Message {
	return (*fastReflection_ValidatorParams)(x)
}

func (x *ValidatorParams) slowProtoReflect() protoreflect.Message {
	mi := &file_tendermint_types_params_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_ValidatorParams_messageType fastReflection_ValidatorParams_messageType
var _ protoreflect.MessageType = fastReflection_ValidatorParams_messageType{}

type fastReflection_ValidatorParams_messageType struct{}

func (x fastReflection_ValidatorParams_messageType) Zero() protoreflect.Message {
	return (*fastReflection_ValidatorParams)(nil)
}
func (x fastReflection_ValidatorParams_messageType) New() protoreflect.Message {
	return new(fastReflection_ValidatorParams)
}
func (x fastReflection_ValidatorParams_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_ValidatorParams
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_ValidatorParams) Descriptor() protoreflect.MessageDescriptor {
	return md_ValidatorParams
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_ValidatorParams) Type() protoreflect.MessageType {
	return _fastReflection_ValidatorParams_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_ValidatorParams) New() protoreflect.Message {
	return new(fastReflection_ValidatorParams)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_ValidatorParams) Interface() protoreflect.ProtoMessage {
	return (*ValidatorParams)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_ValidatorParams) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if len(x.PubKeyTypes) != 0 {
		value := protoreflect.ValueOfList(&_ValidatorParams_1_list{list: &x.PubKeyTypes})
		if !f(fd_ValidatorParams_pub_key_types, value) {
			return
		}
	}
}

// Has reports whether a field is populated.
//
// Some fields have the property of nullability where it is possible to
// distinguish between the default value of a field and whether the field
// was explicitly populated with the default value. Singular message fields,
// member fields of a oneof, and proto2 scalar fields are nullable. Such
// fields are populated only if explicitly set.
//
// In other cases (aside from the nullable cases above),
// a proto3 scalar field is populated if it contains a non-zero value, and
// a repeated field is populated if it is non-empty.
func (x *fastReflection_ValidatorParams) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "tendermint.types.ValidatorParams.pub_key_types":
		return len(x.PubKeyTypes) != 0
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: tendermint.types.ValidatorParams"))
		}
		panic(fmt.Errorf("message tendermint.types.ValidatorParams does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_ValidatorParams) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "tendermint.types.ValidatorParams.pub_key_types":
		x.PubKeyTypes = nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: tendermint.types.ValidatorParams"))
		}
		panic(fmt.Errorf("message tendermint.types.ValidatorParams does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_ValidatorParams) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "tendermint.types.ValidatorParams.pub_key_types":
		if len(x.PubKeyTypes) == 0 {
			return protoreflect.ValueOfList(&_ValidatorParams_1_list{})
		}
		listValue := &_ValidatorParams_1_list{list: &x.PubKeyTypes}
		return protoreflect.ValueOfList(listValue)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: tendermint.types.ValidatorParams"))
		}
		panic(fmt.Errorf("message tendermint.types.ValidatorParams does not contain field %s", descriptor.FullName()))
	}
}

// Set stores the value for a field.
//
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType.
// When setting a composite type, it is unspecified whether the stored value
// aliases the source's memory in any way. If the composite value is an
// empty, read-only value, then it panics.
//
// Set is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_ValidatorParams) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "tendermint.types.ValidatorParams.pub_key_types":
		lv := value.List()
		clv := lv.(*_ValidatorParams_1_list)
		x.PubKeyTypes = *clv.list
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: tendermint.types.ValidatorParams"))
		}
		panic(fmt.Errorf("message tendermint.types.ValidatorParams does not contain field %s", fd.FullName()))
	}
}

// Mutable returns a mutable reference to a composite type.
//
// If the field is unpopulated, it may allocate a composite value.
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType
// if not already stored.
// It panics if the field does not contain a composite type.
//
// Mutable is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_ValidatorParams) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "tendermint.types.ValidatorParams.pub_key_types":
		if x.PubKeyTypes == nil {
			x.PubKeyTypes = []string{}
		}
		value := &_ValidatorParams_1_list{list: &x.PubKeyTypes}
		return protoreflect.ValueOfList(value)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: tendermint.types.ValidatorParams"))
		}
		panic(fmt.Errorf("message tendermint.types.ValidatorParams does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_ValidatorParams) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "tendermint.types.ValidatorParams.pub_key_types":
		list := []string{}
		return protoreflect.ValueOfList(&_ValidatorParams_1_list{list: &list})
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: tendermint.types.ValidatorParams"))
		}
		panic(fmt.Errorf("message tendermint.types.ValidatorParams does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_ValidatorParams) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in tendermint.types.ValidatorParams", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_ValidatorParams) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_ValidatorParams) SetUnknown(fields protoreflect.RawFields) {
	x.unknownFields = fields
}

// IsValid reports whether the message is valid.
//
// An invalid message is an empty, read-only value.
//
// An invalid message often corresponds to a nil pointer of the concrete
// message type, but the details are implementation dependent.
// Validity is not part of the protobuf data model, and may not
// be preserved in marshaling or other operations.
func (x *fastReflection_ValidatorParams) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_ValidatorParams) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*ValidatorParams)
		if x == nil {
			return protoiface.SizeOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Size:              0,
			}
		}
		options := runtime.SizeInputToOptions(input)
		_ = options
		var n int
		var l int
		_ = l
		if len(x.PubKeyTypes) > 0 {
			for _, s := range x.PubKeyTypes {
				l = len(s)
				n += 1 + l + runtime.Sov(uint64(l))
			}
		}
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*ValidatorParams)
		if x == nil {
			return protoiface.MarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Buf:               input.Buf,
			}, nil
		}
		options := runtime.MarshalInputToOptions(input)
		_ = options
		size := options.Size(x)
		dAtA := make([]byte, size)
		i := len(dAtA)
		_ = i
		var l int
		_ = l
		if x.unknownFields != nil {
			i -= len(x.unknownFields)
			copy(dAtA[i:], x.unknownFields)
		}
		if len(x.PubKeyTypes) > 0 {
			for iNdEx := len(x.PubKeyTypes) - 1; iNdEx >= 0; iNdEx-- {
				i -= len(x.PubKeyTypes[iNdEx])
				copy(dAtA[i:], x.PubKeyTypes[iNdEx])
				i = runtime.EncodeVarint(dAtA, i, uint64(len(x.PubKeyTypes[iNdEx])))
				i--
				dAtA[i] = 0xa
			}
		}
		if input.Buf != nil {
			input.Buf = append(input.Buf, dAtA...)
		} else {
			input.Buf = dAtA
		}
		return protoiface.MarshalOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Buf:               input.Buf,
		}, nil
	}
	unmarshal := func(input protoiface.UnmarshalInput) (protoiface.UnmarshalOutput, error) {
		x := input.Message.Interface().(*ValidatorParams)
		if x == nil {
			return protoiface.UnmarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Flags:             input.Flags,
			}, nil
		}
		options := runtime.UnmarshalInputToOptions(input)
		_ = options
		dAtA := input.Buf
		l := len(dAtA)
		iNdEx := 0
		for iNdEx < l {
			preIndex := iNdEx
			var wire uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
				}
				if iNdEx >= l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				wire |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			fieldNum := int32(wire >> 3)
			wireType := int(wire & 0x7)
			if wireType == 4 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: ValidatorParams: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: ValidatorParams: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field PubKeyTypes", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					stringLen |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				intStringLen := int(stringLen)
				if intStringLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.PubKeyTypes = append(x.PubKeyTypes, string(dAtA[iNdEx:postIndex]))
				iNdEx = postIndex
			default:
				iNdEx = preIndex
				skippy, err := runtime.Skip(dAtA[iNdEx:])
				if err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				if (skippy < 0) || (iNdEx+skippy) < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if (iNdEx + skippy) > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if !options.DiscardUnknown {
					x.unknownFields = append(x.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
				}
				iNdEx += skippy
			}
		}

		if iNdEx > l {
			return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
		}
		return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, nil
	}
	return &protoiface.Methods{
		NoUnkeyedLiterals: struct{}{},
		Flags:             protoiface.SupportMarshalDeterministic | protoiface.SupportUnmarshalDiscardUnknown,
		Size:              size,
		Marshal:           marshal,
		Unmarshal:         unmarshal,
		Merge:             nil,
		CheckInitialized:  nil,
	}
}

var (
	md_VersionParams             protoreflect.MessageDescriptor
	fd_VersionParams_app_version protoreflect.FieldDescriptor
)

func init() {
	file_tendermint_types_params_proto_init()
	md_VersionParams = File_tendermint_types_params_proto.Messages().ByName("VersionParams")
	fd_VersionParams_app_version = md_VersionParams.Fields().ByName("app_version")
}

var _ protoreflect.Message = (*fastReflection_VersionParams)(nil)

type fastReflection_VersionParams VersionParams

func (x *VersionParams) ProtoReflect() protoreflect.Message {
	return (*fastReflection_VersionParams)(x)
}

func (x *VersionParams) slowProtoReflect() protoreflect.Message {
	mi := &file_tendermint_types_params_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_VersionParams_messageType fastReflection_VersionParams_messageType
var _ protoreflect.MessageType = fastReflection_VersionParams_messageType{}

type fastReflection_VersionParams_messageType struct{}

func (x fastReflection_VersionParams_messageType) Zero() protoreflect.Message {
	return (*fastReflection_VersionParams)(nil)
}
func (x fastReflection_VersionParams_messageType) New() protoreflect.Message {
	return new(fastReflection_VersionParams)
}
func (x fastReflection_VersionParams_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_VersionParams
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_VersionParams) Descriptor() protoreflect.MessageDescriptor {
	return md_VersionParams
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_VersionParams) Type() protoreflect.MessageType {
	return _fastReflection_VersionParams_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_VersionParams) New() protoreflect.Message {
	return new(fastReflection_VersionParams)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_VersionParams) Interface() protoreflect.ProtoMessage {
	return (*VersionParams)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_VersionParams) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.AppVersion != uint64(0) {
		value := protoreflect.ValueOfUint64(x.AppVersion)
		if !f(fd_VersionParams_app_version, value) {
			return
		}
	}
}

// Has reports whether a field is populated.
//
// Some fields have the property of nullability where it is possible to
// distinguish between the default value of a field and whether the field
// was explicitly populated with the default value. Singular message fields,
// member fields of a oneof, and proto2 scalar fields are nullable. Such
// fields are populated only if explicitly set.
//
// In other cases (aside from the nullable cases above),
// a proto3 scalar field is populated if it contains a non-zero value, and
// a repeated field is populated if it is non-empty.
func (x *fastReflection_VersionParams) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "tendermint.types.VersionParams.app_version":
		return x.AppVersion != uint64(0)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: tendermint.types.VersionParams"))
		}
		panic(fmt.Errorf("message tendermint.types.VersionParams does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_VersionParams) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "tendermint.types.VersionParams.app_version":
		x.AppVersion = uint64(0)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: tendermint.types.VersionParams"))
		}
		panic(fmt.Errorf("message tendermint.types.VersionParams does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_VersionParams) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "tendermint.types.VersionParams.app_version":
		value := x.AppVersion
		return protoreflect.ValueOfUint64(value)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: tendermint.types.VersionParams"))
		}
		panic(fmt.Errorf("message tendermint.types.VersionParams does not contain field %s", descriptor.FullName()))
	}
}

// Set stores the value for a field.
//
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType.
// When setting a composite type, it is unspecified whether the stored value
// aliases the source's memory in any way. If the composite value is an
// empty, read-only value, then it panics.
//
// Set is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_VersionParams) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "tendermint.types.VersionParams.app_version":
		x.AppVersion = value.Uint()
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: tendermint.types.VersionParams"))
		}
		panic(fmt.Errorf("message tendermint.types.VersionParams does not contain field %s", fd.FullName()))
	}
}

// Mutable returns a mutable reference to a composite type.
//
// If the field is unpopulated, it may allocate a composite value.
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType
// if not already stored.
// It panics if the field does not contain a composite type.
//
// Mutable is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_VersionParams) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "tendermint.types.VersionParams.app_version":
		panic(fmt.Errorf("field app_version of message tendermint.types.VersionParams is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: tendermint.types.VersionParams"))
		}
		panic(fmt.Errorf("message tendermint.types.VersionParams does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_VersionParams) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "tendermint.types.VersionParams.app_version":
		return protoreflect.ValueOfUint64(uint64(0))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: tendermint.types.VersionParams"))
		}
		panic(fmt.Errorf("message tendermint.types.VersionParams does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_VersionParams) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in tendermint.types.VersionParams", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_VersionParams) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_VersionParams) SetUnknown(fields protoreflect.RawFields) {
	x.unknownFields = fields
}

// IsValid reports whether the message is valid.
//
// An invalid message is an empty, read-only value.
//
// An invalid message often corresponds to a nil pointer of the concrete
// message type, but the details are implementation dependent.
// Validity is not part of the protobuf data model, and may not
// be preserved in marshaling or other operations.
func (x *fastReflection_VersionParams) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_VersionParams) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*VersionParams)
		if x == nil {
			return protoiface.SizeOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Size:              0,
			}
		}
		options := runtime.SizeInputToOptions(input)
		_ = options
		var n int
		var l int
		_ = l
		if x.AppVersion != 0 {
			n += 1 + runtime.Sov(uint64(x.AppVersion))
		}
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*VersionParams)
		if x == nil {
			return protoiface.MarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Buf:               input.Buf,
			}, nil
		}
		options := runtime.MarshalInputToOptions(input)
		_ = options
		size := options.Size(x)
		dAtA := make([]byte, size)
		i := len(dAtA)
		_ = i
		var l int
		_ = l
		if x.unknownFields != nil {
			i -= len(x.unknownFields)
			copy(dAtA[i:], x.unknownFields)
		}
		if x.AppVersion != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.AppVersion))
			i--
			dAtA[i] = 0x8
		}
		if input.Buf != nil {
			input.Buf = append(input.Buf, dAtA...)
		} else {
			input.Buf = dAtA
		}
		return protoiface.MarshalOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Buf:               input.Buf,
		}, nil
	}
	unmarshal := func(input protoiface.UnmarshalInput) (protoiface.UnmarshalOutput, error) {
		x := input.Message.Interface().(*VersionParams)
		if x == nil {
			return protoiface.UnmarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Flags:             input.Flags,
			}, nil
		}
		options := runtime.UnmarshalInputToOptions(input)
		_ = options
		dAtA := input.Buf
		l := len(dAtA)
		iNdEx := 0
		for iNdEx < l {
			preIndex := iNdEx
			var wire uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
				}
				if iNdEx >= l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				wire |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			fieldNum := int32(wire >> 3)
			wireType := int(wire & 0x7)
			if wireType == 4 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: VersionParams: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: VersionParams: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field AppVersion", wireType)
				}
				x.AppVersion = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.AppVersion |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
			default:
				iNdEx = preIndex
				skippy, err := runtime.Skip(dAtA[iNdEx:])
				if err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				if (skippy < 0) || (iNdEx+skippy) < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if (iNdEx + skippy) > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if !options.DiscardUnknown {
					x.unknownFields = append(x.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
				}
				iNdEx += skippy
			}
		}

		if iNdEx > l {
			return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
		}
		return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, nil
	}
	return &protoiface.Methods{
		NoUnkeyedLiterals: struct{}{},
		Flags:             protoiface.SupportMarshalDeterministic | protoiface.SupportUnmarshalDiscardUnknown,
		Size:              size,
		Marshal:           marshal,
		Unmarshal:         unmarshal,
		Merge:             nil,
		CheckInitialized:  nil,
	}
}

var (
	md_HashedParams                 protoreflect.MessageDescriptor
	fd_HashedParams_block_max_bytes protoreflect.FieldDescriptor
	fd_HashedParams_block_max_gas   protoreflect.FieldDescriptor
)

func init() {
	file_tendermint_types_params_proto_init()
	md_HashedParams = File_tendermint_types_params_proto.Messages().ByName("HashedParams")
	fd_HashedParams_block_max_bytes = md_HashedParams.Fields().ByName("block_max_bytes")
	fd_HashedParams_block_max_gas = md_HashedParams.Fields().ByName("block_max_gas")
}

var _ protoreflect.Message = (*fastReflection_HashedParams)(nil)

type fastReflection_HashedParams HashedParams

func (x *HashedParams) ProtoReflect() protoreflect.Message {
	return (*fastReflection_HashedParams)(x)
}

func (x *HashedParams) slowProtoReflect() protoreflect.Message {
	mi := &file_tendermint_types_params_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_HashedParams_messageType fastReflection_HashedParams_messageType
var _ protoreflect.MessageType = fastReflection_HashedParams_messageType{}

type fastReflection_HashedParams_messageType struct{}

func (x fastReflection_HashedParams_messageType) Zero() protoreflect.Message {
	return (*fastReflection_HashedParams)(nil)
}
func (x fastReflection_HashedParams_messageType) New() protoreflect.Message {
	return new(fastReflection_HashedParams)
}
func (x fastReflection_HashedParams_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_HashedParams
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_HashedParams) Descriptor() protoreflect.MessageDescriptor {
	return md_HashedParams
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_HashedParams) Type() protoreflect.MessageType {
	return _fastReflection_HashedParams_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_HashedParams) New() protoreflect.Message {
	return new(fastReflection_HashedParams)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_HashedParams) Interface() protoreflect.ProtoMessage {
	return (*HashedParams)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_HashedParams) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.BlockMaxBytes != int64(0) {
		value := protoreflect.ValueOfInt64(x.BlockMaxBytes)
		if !f(fd_HashedParams_block_max_bytes, value) {
			return
		}
	}
	if x.BlockMaxGas != int64(0) {
		value := protoreflect.ValueOfInt64(x.BlockMaxGas)
		if !f(fd_HashedParams_block_max_gas, value) {
			return
		}
	}
}

// Has reports whether a field is populated.
//
// Some fields have the property of nullability where it is possible to
// distinguish between the default value of a field and whether the field
// was explicitly populated with the default value. Singular message fields,
// member fields of a oneof, and proto2 scalar fields are nullable. Such
// fields are populated only if explicitly set.
//
// In other cases (aside from the nullable cases above),
// a proto3 scalar field is populated if it contains a non-zero value, and
// a repeated field is populated if it is non-empty.
func (x *fastReflection_HashedParams) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "tendermint.types.HashedParams.block_max_bytes":
		return x.BlockMaxBytes != int64(0)
	case "tendermint.types.HashedParams.block_max_gas":
		return x.BlockMaxGas != int64(0)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: tendermint.types.HashedParams"))
		}
		panic(fmt.Errorf("message tendermint.types.HashedParams does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_HashedParams) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "tendermint.types.HashedParams.block_max_bytes":
		x.BlockMaxBytes = int64(0)
	case "tendermint.types.HashedParams.block_max_gas":
		x.BlockMaxGas = int64(0)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: tendermint.types.HashedParams"))
		}
		panic(fmt.Errorf("message tendermint.types.HashedParams does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_HashedParams) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "tendermint.types.HashedParams.block_max_bytes":
		value := x.BlockMaxBytes
		return protoreflect.ValueOfInt64(value)
	case "tendermint.types.HashedParams.block_max_gas":
		value := x.BlockMaxGas
		return protoreflect.ValueOfInt64(value)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: tendermint.types.HashedParams"))
		}
		panic(fmt.Errorf("message tendermint.types.HashedParams does not contain field %s", descriptor.FullName()))
	}
}

// Set stores the value for a field.
//
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType.
// When setting a composite type, it is unspecified whether the stored value
// aliases the source's memory in any way. If the composite value is an
// empty, read-only value, then it panics.
//
// Set is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_HashedParams) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "tendermint.types.HashedParams.block_max_bytes":
		x.BlockMaxBytes = value.Int()
	case "tendermint.types.HashedParams.block_max_gas":
		x.BlockMaxGas = value.Int()
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: tendermint.types.HashedParams"))
		}
		panic(fmt.Errorf("message tendermint.types.HashedParams does not contain field %s", fd.FullName()))
	}
}

// Mutable returns a mutable reference to a composite type.
//
// If the field is unpopulated, it may allocate a composite value.
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType
// if not already stored.
// It panics if the field does not contain a composite type.
//
// Mutable is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_HashedParams) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "tendermint.types.HashedParams.block_max_bytes":
		panic(fmt.Errorf("field block_max_bytes of message tendermint.types.HashedParams is not mutable"))
	case "tendermint.types.HashedParams.block_max_gas":
		panic(fmt.Errorf("field block_max_gas of message tendermint.types.HashedParams is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: tendermint.types.HashedParams"))
		}
		panic(fmt.Errorf("message tendermint.types.HashedParams does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_HashedParams) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "tendermint.types.HashedParams.block_max_bytes":
		return protoreflect.ValueOfInt64(int64(0))
	case "tendermint.types.HashedParams.block_max_gas":
		return protoreflect.ValueOfInt64(int64(0))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: tendermint.types.HashedParams"))
		}
		panic(fmt.Errorf("message tendermint.types.HashedParams does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_HashedParams) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in tendermint.types.HashedParams", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_HashedParams) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_HashedParams) SetUnknown(fields protoreflect.RawFields) {
	x.unknownFields = fields
}

// IsValid reports whether the message is valid.
//
// An invalid message is an empty, read-only value.
//
// An invalid message often corresponds to a nil pointer of the concrete
// message type, but the details are implementation dependent.
// Validity is not part of the protobuf data model, and may not
// be preserved in marshaling or other operations.
func (x *fastReflection_HashedParams) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_HashedParams) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*HashedParams)
		if x == nil {
			return protoiface.SizeOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Size:              0,
			}
		}
		options := runtime.SizeInputToOptions(input)
		_ = options
		var n int
		var l int
		_ = l
		if x.BlockMaxBytes != 0 {
			n += 1 + runtime.Sov(uint64(x.BlockMaxBytes))
		}
		if x.BlockMaxGas != 0 {
			n += 1 + runtime.Sov(uint64(x.BlockMaxGas))
		}
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*HashedParams)
		if x == nil {
			return protoiface.MarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Buf:               input.Buf,
			}, nil
		}
		options := runtime.MarshalInputToOptions(input)
		_ = options
		size := options.Size(x)
		dAtA := make([]byte, size)
		i := len(dAtA)
		_ = i
		var l int
		_ = l
		if x.unknownFields != nil {
			i -= len(x.unknownFields)
			copy(dAtA[i:], x.unknownFields)
		}
		if x.BlockMaxGas != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.BlockMaxGas))
			i--
			dAtA[i] = 0x10
		}
		if x.BlockMaxBytes != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.BlockMaxBytes))
			i--
			dAtA[i] = 0x8
		}
		if input.Buf != nil {
			input.Buf = append(input.Buf, dAtA...)
		} else {
			input.Buf = dAtA
		}
		return protoiface.MarshalOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Buf:               input.Buf,
		}, nil
	}
	unmarshal := func(input protoiface.UnmarshalInput) (protoiface.UnmarshalOutput, error) {
		x := input.Message.Interface().(*HashedParams)
		if x == nil {
			return protoiface.UnmarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Flags:             input.Flags,
			}, nil
		}
		options := runtime.UnmarshalInputToOptions(input)
		_ = options
		dAtA := input.Buf
		l := len(dAtA)
		iNdEx := 0
		for iNdEx < l {
			preIndex := iNdEx
			var wire uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
				}
				if iNdEx >= l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				wire |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			fieldNum := int32(wire >> 3)
			wireType := int(wire & 0x7)
			if wireType == 4 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: HashedParams: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: HashedParams: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field BlockMaxBytes", wireType)
				}
				x.BlockMaxBytes = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.BlockMaxBytes |= int64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
			case 2:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field BlockMaxGas", wireType)
				}
				x.BlockMaxGas = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.BlockMaxGas |= int64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
			default:
				iNdEx = preIndex
				skippy, err := runtime.Skip(dAtA[iNdEx:])
				if err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				if (skippy < 0) || (iNdEx+skippy) < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if (iNdEx + skippy) > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if !options.DiscardUnknown {
					x.unknownFields = append(x.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
				}
				iNdEx += skippy
			}
		}

		if iNdEx > l {
			return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
		}
		return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, nil
	}
	return &protoiface.Methods{
		NoUnkeyedLiterals: struct{}{},
		Flags:             protoiface.SupportMarshalDeterministic | protoiface.SupportUnmarshalDiscardUnknown,
		Size:              size,
		Marshal:           marshal,
		Unmarshal:         unmarshal,
		Merge:             nil,
		CheckInitialized:  nil,
	}
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.0
// 	protoc        (unknown)
// source: tendermint/types/params.proto

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// ConsensusParams contains consensus critical parameters that determine the
// validity of blocks.
type ConsensusParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Block     *BlockParams     `protobuf:"bytes,1,opt,name=block,proto3" json:"block,omitempty"`
	Evidence  *EvidenceParams  `protobuf:"bytes,2,opt,name=evidence,proto3" json:"evidence,omitempty"`
	Validator *ValidatorParams `protobuf:"bytes,3,opt,name=validator,proto3" json:"validator,omitempty"`
	Version   *VersionParams   `protobuf:"bytes,4,opt,name=version,proto3" json:"version,omitempty"`
}

func (x *ConsensusParams) Reset() {
	*x = ConsensusParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tendermint_types_params_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConsensusParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConsensusParams) ProtoMessage() {}

// Deprecated: Use ConsensusParams.ProtoReflect.Descriptor instead.
func (*ConsensusParams) Descriptor() ([]byte, []int) {
	return file_tendermint_types_params_proto_rawDescGZIP(), []int{0}
}

func (x *ConsensusParams) GetBlock() *BlockParams {
	if x != nil {
		return x.Block
	}
	return nil
}

func (x *ConsensusParams) GetEvidence() *EvidenceParams {
	if x != nil {
		return x.Evidence
	}
	return nil
}

func (x *ConsensusParams) GetValidator() *ValidatorParams {
	if x != nil {
		return x.Validator
	}
	return nil
}

func (x *ConsensusParams) GetVersion() *VersionParams {
	if x != nil {
		return x.Version
	}
	return nil
}

// BlockParams contains limits on the block size.
type BlockParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Max block size, in bytes.
	// Note: must be greater than 0
	MaxBytes int64 `protobuf:"varint,1,opt,name=max_bytes,json=maxBytes,proto3" json:"max_bytes,omitempty"`
	// Max gas per block.
	// Note: must be greater or equal to -1
	MaxGas int64 `protobuf:"varint,2,opt,name=max_gas,json=maxGas,proto3" json:"max_gas,omitempty"`
	// Minimum time increment between consecutive blocks (in milliseconds) If the
	// block header timestamp is ahead of the system clock, decrease this value.
	//
	// Not exposed to the application.
	TimeIotaMs int64 `protobuf:"varint,3,opt,name=time_iota_ms,json=timeIotaMs,proto3" json:"time_iota_ms,omitempty"`
}

func (x *BlockParams) Reset() {
	*x = BlockParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tendermint_types_params_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlockParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlockParams) ProtoMessage() {}

// Deprecated: Use BlockParams.ProtoReflect.Descriptor instead.
func (*BlockParams) Descriptor() ([]byte, []int) {
	return file_tendermint_types_params_proto_rawDescGZIP(), []int{1}
}

func (x *BlockParams) GetMaxBytes() int64 {
	if x != nil {
		return x.MaxBytes
	}
	return 0
}

func (x *BlockParams) GetMaxGas() int64 {
	if x != nil {
		return x.MaxGas
	}
	return 0
}

func (x *BlockParams) GetTimeIotaMs() int64 {
	if x != nil {
		return x.TimeIotaMs
	}
	return 0
}

// EvidenceParams determine how we handle evidence of malfeasance.
type EvidenceParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Max age of evidence, in blocks.
	//
	// The basic formula for calculating this is: MaxAgeDuration / {average block
	// time}.
	MaxAgeNumBlocks int64 `protobuf:"varint,1,opt,name=max_age_num_blocks,json=maxAgeNumBlocks,proto3" json:"max_age_num_blocks,omitempty"`
	// Max age of evidence, in time.
	//
	// It should correspond with an app's "unbonding period" or other similar
	// mechanism for handling [Nothing-At-Stake
	// attacks](https://github.com/ethereum/wiki/wiki/Proof-of-Stake-FAQ#what-is-the-nothing-at-stake-problem-and-how-can-it-be-fixed).
	MaxAgeDuration *durationpb.Duration `protobuf:"bytes,2,opt,name=max_age_duration,json=maxAgeDuration,proto3" json:"max_age_duration,omitempty"`
	// This sets the maximum size of total evidence in bytes that can be committed in a single block.
	// and should fall comfortably under the max block bytes.
	// Default is 1048576 or 1MB
	MaxBytes int64 `protobuf:"varint,3,opt,name=max_bytes,json=maxBytes,proto3" json:"max_bytes,omitempty"`
}

func (x *EvidenceParams) Reset() {
	*x = EvidenceParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tendermint_types_params_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EvidenceParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EvidenceParams) ProtoMessage() {}

// Deprecated: Use EvidenceParams.ProtoReflect.Descriptor instead.
func (*EvidenceParams) Descriptor() ([]byte, []int) {
	return file_tendermint_types_params_proto_rawDescGZIP(), []int{2}
}

func (x *EvidenceParams) GetMaxAgeNumBlocks() int64 {
	if x != nil {
		return x.MaxAgeNumBlocks
	}
	return 0
}

func (x *EvidenceParams) GetMaxAgeDuration() *durationpb.Duration {
	if x != nil {
		return x.MaxAgeDuration
	}
	return nil
}

func (x *EvidenceParams) GetMaxBytes() int64 {
	if x != nil {
		return x.MaxBytes
	}
	return 0
}

// ValidatorParams restrict the public key types validators can use.
// NOTE: uses ABCI pubkey naming, not Amino names.
type ValidatorParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PubKeyTypes []string `protobuf:"bytes,1,rep,name=pub_key_types,json=pubKeyTypes,proto3" json:"pub_key_types,omitempty"`
}

func (x *ValidatorParams) Reset() {
	*x = ValidatorParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tendermint_types_params_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ValidatorParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValidatorParams) ProtoMessage() {}

// Deprecated: Use ValidatorParams.ProtoReflect.Descriptor instead.
func (*ValidatorParams) Descriptor() ([]byte, []int) {
	return file_tendermint_types_params_proto_rawDescGZIP(), []int{3}
}

func (x *ValidatorParams) GetPubKeyTypes() []string {
	if x != nil {
		return x.PubKeyTypes
	}
	return nil
}

// VersionParams contains the ABCI application version.
type VersionParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AppVersion uint64 `protobuf:"varint,1,opt,name=app_version,json=appVersion,proto3" json:"app_version,omitempty"`
}

func (x *VersionParams) Reset() {
	*x = VersionParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tendermint_types_params_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VersionParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VersionParams) ProtoMessage() {}

// Deprecated: Use VersionParams.ProtoReflect.Descriptor instead.
func (*VersionParams) Descriptor() ([]byte, []int) {
	return file_tendermint_types_params_proto_rawDescGZIP(), []int{4}
}

func (x *VersionParams) GetAppVersion() uint64 {
	if x != nil {
		return x.AppVersion
	}
	return 0
}

// HashedParams is a subset of ConsensusParams.
//
// It is hashed into the Header.ConsensusHash.
type HashedParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BlockMaxBytes int64 `protobuf:"varint,1,opt,name=block_max_bytes,json=blockMaxBytes,proto3" json:"block_max_bytes,omitempty"`
	BlockMaxGas   int64 `protobuf:"varint,2,opt,name=block_max_gas,json=blockMaxGas,proto3" json:"block_max_gas,omitempty"`
}

func (x *HashedParams) Reset() {
	*x = HashedParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tendermint_types_params_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HashedParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HashedParams) ProtoMessage() {}

// Deprecated: Use HashedParams.ProtoReflect.Descriptor instead.
func (*HashedParams) Descriptor() ([]byte, []int) {
	return file_tendermint_types_params_proto_rawDescGZIP(), []int{5}
}

func (x *HashedParams) GetBlockMaxBytes() int64 {
	if x != nil {
		return x.BlockMaxBytes
	}
	return 0
}

func (x *HashedParams) GetBlockMaxGas() int64 {
	if x != nil {
		return x.BlockMaxGas
	}
	return 0
}

var File_tendermint_types_params_proto protoreflect.FileDescriptor

var file_tendermint_types_params_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x74, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x74, 0x2f, 0x74, 0x79, 0x70,
	0x65, 0x73, 0x2f, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x10, 0x74, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x74, 0x2e, 0x74, 0x79, 0x70, 0x65,
	0x73, 0x1a, 0x14, 0x67, 0x6f, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x67,
	0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x98, 0x02, 0x0a, 0x0f, 0x43, 0x6f, 0x6e, 0x73,
	0x65, 0x6e, 0x73, 0x75, 0x73, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x39, 0x0a, 0x05, 0x62,
	0x6c, 0x6f, 0x63, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x74, 0x65, 0x6e,
	0x64, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x74, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x42, 0x6c,
	0x6f, 0x63, 0x6b, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x42, 0x04, 0xc8, 0xde, 0x1f, 0x00, 0x52,
	0x05, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x42, 0x0a, 0x08, 0x65, 0x76, 0x69, 0x64, 0x65, 0x6e,
	0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x74, 0x65, 0x6e, 0x64, 0x65,
	0x72, 0x6d, 0x69, 0x6e, 0x74, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x45, 0x76, 0x69, 0x64,
	0x65, 0x6e, 0x63, 0x65, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x42, 0x04, 0xc8, 0xde, 0x1f, 0x00,
	0x52, 0x08, 0x65, 0x76, 0x69, 0x64, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x45, 0x0a, 0x09, 0x76, 0x61,
	0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e,
	0x74, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x74, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73,
	0x2e, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73,
	0x42, 0x04, 0xc8, 0xde, 0x1f, 0x00, 0x52, 0x09, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f,
	0x72, 0x12, 0x3f, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x74, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x74, 0x2e,
	0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x50, 0x61, 0x72,
	0x61, 0x6d, 0x73, 0x42, 0x04, 0xc8, 0xde, 0x1f, 0x00, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x22, 0x65, 0x0a, 0x0b, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x50, 0x61, 0x72, 0x61, 0x6d,
	0x73, 0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x61, 0x78, 0x5f, 0x62, 0x79, 0x74, 0x65, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x6d, 0x61, 0x78, 0x42, 0x79, 0x74, 0x65, 0x73, 0x12, 0x17,
	0x0a, 0x07, 0x6d, 0x61, 0x78, 0x5f, 0x67, 0x61, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x06, 0x6d, 0x61, 0x78, 0x47, 0x61, 0x73, 0x12, 0x20, 0x0a, 0x0c, 0x74, 0x69, 0x6d, 0x65, 0x5f,
	0x69, 0x6f, 0x74, 0x61, 0x5f, 0x6d, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x74,
	0x69, 0x6d, 0x65, 0x49, 0x6f, 0x74, 0x61, 0x4d, 0x73, 0x22, 0xa9, 0x01, 0x0a, 0x0e, 0x45, 0x76,
	0x69, 0x64, 0x65, 0x6e, 0x63, 0x65, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x2b, 0x0a, 0x12,
	0x6d, 0x61, 0x78, 0x5f, 0x61, 0x67, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x5f, 0x62, 0x6c, 0x6f, 0x63,
	0x6b, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0f, 0x6d, 0x61, 0x78, 0x41, 0x67, 0x65,
	0x4e, 0x75, 0x6d, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x12, 0x4d, 0x0a, 0x10, 0x6d, 0x61, 0x78,
	0x5f, 0x61, 0x67, 0x65, 0x5f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x08,
	0xc8, 0xde, 0x1f, 0x00, 0x98, 0xdf, 0x1f, 0x01, 0x52, 0x0e, 0x6d, 0x61, 0x78, 0x41, 0x67, 0x65,
	0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x61, 0x78, 0x5f,
	0x62, 0x79, 0x74, 0x65, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x6d, 0x61, 0x78,
	0x42, 0x79, 0x74, 0x65, 0x73, 0x22, 0x3f, 0x0a, 0x0f, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x6f, 0x72, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x22, 0x0a, 0x0d, 0x70, 0x75, 0x62, 0x5f,
	0x6b, 0x65, 0x79, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x0b, 0x70, 0x75, 0x62, 0x4b, 0x65, 0x79, 0x54, 0x79, 0x70, 0x65, 0x73, 0x3a, 0x08, 0xb8, 0xa0,
	0x1f, 0x01, 0xe8, 0xa0, 0x1f, 0x01, 0x22, 0x3a, 0x0a, 0x0d, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x61, 0x70, 0x70, 0x5f, 0x76,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x61, 0x70,
	0x70, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x3a, 0x08, 0xb8, 0xa0, 0x1f, 0x01, 0xe8, 0xa0,
	0x1f, 0x01, 0x22, 0x5a, 0x0a, 0x0c, 0x48, 0x61, 0x73, 0x68, 0x65, 0x64, 0x50, 0x61, 0x72, 0x61,
	0x6d, 0x73, 0x12, 0x26, 0x0a, 0x0f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x6d, 0x61, 0x78, 0x5f,
	0x62, 0x79, 0x74, 0x65, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x62, 0x6c, 0x6f,
	0x63, 0x6b, 0x4d, 0x61, 0x78, 0x42, 0x79, 0x74, 0x65, 0x73, 0x12, 0x22, 0x0a, 0x0d, 0x62, 0x6c,
	0x6f, 0x63, 0x6b, 0x5f, 0x6d, 0x61, 0x78, 0x5f, 0x67, 0x61, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x0b, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x4d, 0x61, 0x78, 0x47, 0x61, 0x73, 0x42, 0xbb,
	0x01, 0x0a, 0x14, 0x63, 0x6f, 0x6d, 0x2e, 0x74, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x6d, 0x69, 0x6e,
	0x74, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x42, 0x0b, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x31, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73,
	0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x74, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x6d,
	0x69, 0x6e, 0x74, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0xa2, 0x02, 0x03, 0x54, 0x54, 0x58, 0xaa,
	0x02, 0x10, 0x54, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x74, 0x2e, 0x54, 0x79, 0x70,
	0x65, 0x73, 0xca, 0x02, 0x10, 0x54, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x74, 0x5c,
	0x54, 0x79, 0x70, 0x65, 0x73, 0xe2, 0x02, 0x1c, 0x54, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x6d, 0x69,
	0x6e, 0x74, 0x5c, 0x54, 0x79, 0x70, 0x65, 0x73, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x11, 0x54, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x6d, 0x69, 0x6e,
	0x74, 0x3a, 0x3a, 0x54, 0x79, 0x70, 0x65, 0x73, 0xa8, 0xe2, 0x1e, 0x01, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tendermint_types_params_proto_rawDescOnce sync.Once
	file_tendermint_types_params_proto_rawDescData = file_tendermint_types_params_proto_rawDesc
)

func file_tendermint_types_params_proto_rawDescGZIP() []byte {
	file_tendermint_types_params_proto_rawDescOnce.Do(func() {
		file_tendermint_types_params_proto_rawDescData = protoimpl.X.CompressGZIP(file_tendermint_types_params_proto_rawDescData)
	})
	return file_tendermint_types_params_proto_rawDescData
}

var file_tendermint_types_params_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_tendermint_types_params_proto_goTypes = []interface{}{
	(*ConsensusParams)(nil),     // 0: tendermint.types.ConsensusParams
	(*BlockParams)(nil),         // 1: tendermint.types.BlockParams
	(*EvidenceParams)(nil),      // 2: tendermint.types.EvidenceParams
	(*ValidatorParams)(nil),     // 3: tendermint.types.ValidatorParams
	(*VersionParams)(nil),       // 4: tendermint.types.VersionParams
	(*HashedParams)(nil),        // 5: tendermint.types.HashedParams
	(*durationpb.Duration)(nil), // 6: google.protobuf.Duration
}
var file_tendermint_types_params_proto_depIdxs = []int32{
	1, // 0: tendermint.types.ConsensusParams.block:type_name -> tendermint.types.BlockParams
	2, // 1: tendermint.types.ConsensusParams.evidence:type_name -> tendermint.types.EvidenceParams
	3, // 2: tendermint.types.ConsensusParams.validator:type_name -> tendermint.types.ValidatorParams
	4, // 3: tendermint.types.ConsensusParams.version:type_name -> tendermint.types.VersionParams
	6, // 4: tendermint.types.EvidenceParams.max_age_duration:type_name -> google.protobuf.Duration
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_tendermint_types_params_proto_init() }
func file_tendermint_types_params_proto_init() {
	if File_tendermint_types_params_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_tendermint_types_params_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConsensusParams); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_tendermint_types_params_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BlockParams); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_tendermint_types_params_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EvidenceParams); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_tendermint_types_params_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ValidatorParams); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_tendermint_types_params_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VersionParams); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_tendermint_types_params_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HashedParams); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_tendermint_types_params_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_tendermint_types_params_proto_goTypes,
		DependencyIndexes: file_tendermint_types_params_proto_depIdxs,
		MessageInfos:      file_tendermint_types_params_proto_msgTypes,
	}.Build()
	File_tendermint_types_params_proto = out.File
	file_tendermint_types_params_proto_rawDesc = nil
	file_tendermint_types_params_proto_goTypes = nil
	file_tendermint_types_params_proto_depIdxs = nil
}
