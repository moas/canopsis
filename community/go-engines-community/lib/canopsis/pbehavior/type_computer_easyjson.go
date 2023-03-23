// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package pbehavior

import (
	json "encoding/json"
	pattern "git.canopsis.net/canopsis/canopsis-community/community/go-engines-community/lib/canopsis/pattern"
	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjson950e241aDecodeGitCanopsisNetCanopsisCanopsisCommunityCommunityGoEnginesCommunityLibCanopsisPbehavior(in *jlexer.Lexer, out *Types) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "T":
			if in.IsNull() {
				in.Skip()
			} else {
				in.Delim('{')
				out.T = make(map[string]Type)
				for !in.IsDelim('}') {
					key := string(in.String())
					in.WantColon()
					var v1 Type
					easyjson950e241aDecodeGitCanopsisNetCanopsisCanopsisCommunityCommunityGoEnginesCommunityLibCanopsisPbehavior1(in, &v1)
					(out.T)[key] = v1
					in.WantComma()
				}
				in.Delim('}')
			}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson950e241aEncodeGitCanopsisNetCanopsisCanopsisCommunityCommunityGoEnginesCommunityLibCanopsisPbehavior(out *jwriter.Writer, in Types) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"T\":"
		out.RawString(prefix[1:])
		if in.T == nil && (out.Flags&jwriter.NilMapAsEmpty) == 0 {
			out.RawString(`null`)
		} else {
			out.RawByte('{')
			v2First := true
			for v2Name, v2Value := range in.T {
				if v2First {
					v2First = false
				} else {
					out.RawByte(',')
				}
				out.String(string(v2Name))
				out.RawByte(':')
				easyjson950e241aEncodeGitCanopsisNetCanopsisCanopsisCommunityCommunityGoEnginesCommunityLibCanopsisPbehavior1(out, v2Value)
			}
			out.RawByte('}')
		}
	}
	out.RawByte('}')
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Types) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson950e241aEncodeGitCanopsisNetCanopsisCanopsisCommunityCommunityGoEnginesCommunityLibCanopsisPbehavior(w, v)
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Types) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson950e241aDecodeGitCanopsisNetCanopsisCanopsisCommunityCommunityGoEnginesCommunityLibCanopsisPbehavior(l, v)
}
func easyjson950e241aDecodeGitCanopsisNetCanopsisCanopsisCommunityCommunityGoEnginesCommunityLibCanopsisPbehavior1(in *jlexer.Lexer, out *Type) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "_id":
			out.ID = string(in.String())
		case "name":
			out.Name = string(in.String())
		case "description":
			out.Description = string(in.String())
		case "type":
			out.Type = string(in.String())
		case "priority":
			out.Priority = int(in.Int())
		case "icon_name":
			out.IconName = string(in.String())
		case "color":
			out.Color = string(in.String())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson950e241aEncodeGitCanopsisNetCanopsisCanopsisCommunityCommunityGoEnginesCommunityLibCanopsisPbehavior1(out *jwriter.Writer, in Type) {
	out.RawByte('{')
	first := true
	_ = first
	if in.ID != "" {
		const prefix string = ",\"_id\":"
		first = false
		out.RawString(prefix[1:])
		out.String(string(in.ID))
	}
	{
		const prefix string = ",\"name\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Name))
	}
	{
		const prefix string = ",\"description\":"
		out.RawString(prefix)
		out.String(string(in.Description))
	}
	{
		const prefix string = ",\"type\":"
		out.RawString(prefix)
		out.String(string(in.Type))
	}
	{
		const prefix string = ",\"priority\":"
		out.RawString(prefix)
		out.Int(int(in.Priority))
	}
	{
		const prefix string = ",\"icon_name\":"
		out.RawString(prefix)
		out.String(string(in.IconName))
	}
	{
		const prefix string = ",\"color\":"
		out.RawString(prefix)
		out.String(string(in.Color))
	}
	out.RawByte('}')
}
func easyjson950e241aDecodeGitCanopsisNetCanopsisCanopsisCommunityCommunityGoEnginesCommunityLibCanopsisPbehavior2(in *jlexer.Lexer, out *ComputedPbehavior) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "n":
			out.Name = string(in.String())
		case "rn":
			out.ReasonName = string(in.String())
		case "r":
			out.ReasonID = string(in.String())
		case "f":
			out.Filter = string(in.String())
		case "t":
			if in.IsNull() {
				in.Skip()
				out.Types = nil
			} else {
				in.Delim('[')
				if out.Types == nil {
					if !in.IsDelim(']') {
						out.Types = make([]ComputedType, 0, 0)
					} else {
						out.Types = []ComputedType{}
					}
				} else {
					out.Types = (out.Types)[:0]
				}
				for !in.IsDelim(']') {
					var v3 ComputedType
					easyjson950e241aDecodeGitCanopsisNetCanopsisCanopsisCommunityCommunityGoEnginesCommunityLibCanopsisPbehavior3(in, &v3)
					out.Types = append(out.Types, v3)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "c":
			out.Created = int64(in.Int64())
		case "p":
			if in.IsNull() {
				in.Skip()
				out.Pattern = nil
			} else {
				in.Delim('[')
				if out.Pattern == nil {
					if !in.IsDelim(']') {
						out.Pattern = make(pattern.Entity, 0, 2)
					} else {
						out.Pattern = pattern.Entity{}
					}
				} else {
					out.Pattern = (out.Pattern)[:0]
				}
				for !in.IsDelim(']') {
					var v4 []pattern.FieldCondition
					if in.IsNull() {
						in.Skip()
						v4 = nil
					} else {
						in.Delim('[')
						if v4 == nil {
							if !in.IsDelim(']') {
								v4 = make([]pattern.FieldCondition, 0, 0)
							} else {
								v4 = []pattern.FieldCondition{}
							}
						} else {
							v4 = (v4)[:0]
						}
						for !in.IsDelim(']') {
							var v5 pattern.FieldCondition
							easyjson950e241aDecodeGitCanopsisNetCanopsisCanopsisCommunityCommunityGoEnginesCommunityLibCanopsisPattern(in, &v5)
							v4 = append(v4, v5)
							in.WantComma()
						}
						in.Delim(']')
					}
					out.Pattern = append(out.Pattern, v4)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "q":
			if in.IsNull() {
				in.Skip()
			} else {
				in.Delim('{')
				if !in.IsDelim('}') {
					out.OldMongoQuery = make(map[string]interface{})
				} else {
					out.OldMongoQuery = nil
				}
				for !in.IsDelim('}') {
					key := string(in.String())
					in.WantColon()
					var v6 interface{}
					if m, ok := v6.(easyjson.Unmarshaler); ok {
						m.UnmarshalEasyJSON(in)
					} else if m, ok := v6.(json.Unmarshaler); ok {
						_ = m.UnmarshalJSON(in.Raw())
					} else {
						v6 = in.Interface()
					}
					(out.OldMongoQuery)[key] = v6
					in.WantComma()
				}
				in.Delim('}')
			}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson950e241aEncodeGitCanopsisNetCanopsisCanopsisCommunityCommunityGoEnginesCommunityLibCanopsisPbehavior2(out *jwriter.Writer, in ComputedPbehavior) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"n\":"
		out.RawString(prefix[1:])
		out.String(string(in.Name))
	}
	{
		const prefix string = ",\"rn\":"
		out.RawString(prefix)
		out.String(string(in.ReasonName))
	}
	{
		const prefix string = ",\"r\":"
		out.RawString(prefix)
		out.String(string(in.ReasonID))
	}
	{
		const prefix string = ",\"f\":"
		out.RawString(prefix)
		out.String(string(in.Filter))
	}
	{
		const prefix string = ",\"t\":"
		out.RawString(prefix)
		if in.Types == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v7, v8 := range in.Types {
				if v7 > 0 {
					out.RawByte(',')
				}
				easyjson950e241aEncodeGitCanopsisNetCanopsisCanopsisCommunityCommunityGoEnginesCommunityLibCanopsisPbehavior3(out, v8)
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"c\":"
		out.RawString(prefix)
		out.Int64(int64(in.Created))
	}
	if len(in.Pattern) != 0 {
		const prefix string = ",\"p\":"
		out.RawString(prefix)
		{
			out.RawByte('[')
			for v9, v10 := range in.Pattern {
				if v9 > 0 {
					out.RawByte(',')
				}
				if v10 == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
					out.RawString("null")
				} else {
					out.RawByte('[')
					for v11, v12 := range v10 {
						if v11 > 0 {
							out.RawByte(',')
						}
						easyjson950e241aEncodeGitCanopsisNetCanopsisCanopsisCommunityCommunityGoEnginesCommunityLibCanopsisPattern(out, v12)
					}
					out.RawByte(']')
				}
			}
			out.RawByte(']')
		}
	}
	if len(in.OldMongoQuery) != 0 {
		const prefix string = ",\"q\":"
		out.RawString(prefix)
		{
			out.RawByte('{')
			v13First := true
			for v13Name, v13Value := range in.OldMongoQuery {
				if v13First {
					v13First = false
				} else {
					out.RawByte(',')
				}
				out.String(string(v13Name))
				out.RawByte(':')
				if m, ok := v13Value.(easyjson.Marshaler); ok {
					m.MarshalEasyJSON(out)
				} else if m, ok := v13Value.(json.Marshaler); ok {
					out.Raw(m.MarshalJSON())
				} else {
					out.Raw(json.Marshal(v13Value))
				}
			}
			out.RawByte('}')
		}
	}
	out.RawByte('}')
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v ComputedPbehavior) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson950e241aEncodeGitCanopsisNetCanopsisCanopsisCommunityCommunityGoEnginesCommunityLibCanopsisPbehavior2(w, v)
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *ComputedPbehavior) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson950e241aDecodeGitCanopsisNetCanopsisCanopsisCommunityCommunityGoEnginesCommunityLibCanopsisPbehavior2(l, v)
}
func easyjson950e241aDecodeGitCanopsisNetCanopsisCanopsisCommunityCommunityGoEnginesCommunityLibCanopsisPattern(in *jlexer.Lexer, out *pattern.FieldCondition) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "field":
			out.Field = string(in.String())
		case "field_type":
			out.FieldType = string(in.String())
		case "cond":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.Condition).UnmarshalJSON(data))
			}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson950e241aEncodeGitCanopsisNetCanopsisCanopsisCommunityCommunityGoEnginesCommunityLibCanopsisPattern(out *jwriter.Writer, in pattern.FieldCondition) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"field\":"
		out.RawString(prefix[1:])
		out.String(string(in.Field))
	}
	if in.FieldType != "" {
		const prefix string = ",\"field_type\":"
		out.RawString(prefix)
		out.String(string(in.FieldType))
	}
	{
		const prefix string = ",\"cond\":"
		out.RawString(prefix)
		easyjson950e241aEncodeGitCanopsisNetCanopsisCanopsisCommunityCommunityGoEnginesCommunityLibCanopsisPattern1(out, in.Condition)
	}
	out.RawByte('}')
}
func easyjson950e241aDecodeGitCanopsisNetCanopsisCanopsisCommunityCommunityGoEnginesCommunityLibCanopsisPattern1(in *jlexer.Lexer, out *pattern.Condition) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "type":
			out.Type = string(in.String())
		case "value":
			if m, ok := out.Value.(easyjson.Unmarshaler); ok {
				m.UnmarshalEasyJSON(in)
			} else if m, ok := out.Value.(json.Unmarshaler); ok {
				_ = m.UnmarshalJSON(in.Raw())
			} else {
				out.Value = in.Interface()
			}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson950e241aEncodeGitCanopsisNetCanopsisCanopsisCommunityCommunityGoEnginesCommunityLibCanopsisPattern1(out *jwriter.Writer, in pattern.Condition) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"type\":"
		out.RawString(prefix[1:])
		out.String(string(in.Type))
	}
	{
		const prefix string = ",\"value\":"
		out.RawString(prefix)
		if m, ok := in.Value.(easyjson.Marshaler); ok {
			m.MarshalEasyJSON(out)
		} else if m, ok := in.Value.(json.Marshaler); ok {
			out.Raw(m.MarshalJSON())
		} else {
			out.Raw(json.Marshal(in.Value))
		}
	}
	out.RawByte('}')
}
func easyjson950e241aDecodeGitCanopsisNetCanopsisCanopsisCommunityCommunityGoEnginesCommunityLibCanopsisPbehavior3(in *jlexer.Lexer, out *ComputedType) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "t":
			out.ID = string(in.String())
		case "s":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.Span).UnmarshalJSON(data))
			}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson950e241aEncodeGitCanopsisNetCanopsisCanopsisCommunityCommunityGoEnginesCommunityLibCanopsisPbehavior3(out *jwriter.Writer, in ComputedType) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"t\":"
		out.RawString(prefix[1:])
		out.String(string(in.ID))
	}
	{
		const prefix string = ",\"s\":"
		out.RawString(prefix)
		out.Raw((in.Span).MarshalJSON())
	}
	out.RawByte('}')
}