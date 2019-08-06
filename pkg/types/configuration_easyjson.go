// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package types

import (
	json "encoding/json"
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

func easyjsonA0a30a08DecodeGithubComQordobacodeCliV2PkgTypes(in *jlexer.Lexer, out *SourceConfig) {
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
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "Files":
			if in.IsNull() {
				in.Skip()
				out.Files = nil
			} else {
				in.Delim('[')
				if out.Files == nil {
					if !in.IsDelim(']') {
						out.Files = make([]string, 0, 4)
					} else {
						out.Files = []string{}
					}
				} else {
					out.Files = (out.Files)[:0]
				}
				for !in.IsDelim(']') {
					var v1 string
					v1 = string(in.String())
					out.Files = append(out.Files, v1)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "Folders":
			if in.IsNull() {
				in.Skip()
				out.Folders = nil
			} else {
				in.Delim('[')
				if out.Folders == nil {
					if !in.IsDelim(']') {
						out.Folders = make([]string, 0, 4)
					} else {
						out.Folders = []string{}
					}
				} else {
					out.Folders = (out.Folders)[:0]
				}
				for !in.IsDelim(']') {
					var v2 string
					v2 = string(in.String())
					out.Folders = append(out.Folders, v2)
					in.WantComma()
				}
				in.Delim(']')
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
func easyjsonA0a30a08EncodeGithubComQordobacodeCliV2PkgTypes(out *jwriter.Writer, in SourceConfig) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"Files\":"
		out.RawString(prefix[1:])
		if in.Files == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v3, v4 := range in.Files {
				if v3 > 0 {
					out.RawByte(',')
				}
				out.String(string(v4))
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"Folders\":"
		out.RawString(prefix)
		if in.Folders == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v5, v6 := range in.Folders {
				if v5 > 0 {
					out.RawByte(',')
				}
				out.String(string(v6))
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v SourceConfig) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonA0a30a08EncodeGithubComQordobacodeCliV2PkgTypes(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v SourceConfig) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonA0a30a08EncodeGithubComQordobacodeCliV2PkgTypes(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *SourceConfig) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonA0a30a08DecodeGithubComQordobacodeCliV2PkgTypes(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *SourceConfig) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonA0a30a08DecodeGithubComQordobacodeCliV2PkgTypes(l, v)
}
func easyjsonA0a30a08DecodeGithubComQordobacodeCliV2PkgTypes1(in *jlexer.Lexer, out *QordobaConfig) {
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
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "AccessToken":
			out.AccessToken = string(in.String())
		case "OrganizationID":
			out.OrganizationID = int64(in.Int64())
		case "WorkspaceID":
			out.WorkspaceID = int64(in.Int64())
		case "ProjectID":
			out.ProjectID = int64(in.Int64())
		case "AudienceMap":
			if in.IsNull() {
				in.Skip()
			} else {
				in.Delim('{')
				if !in.IsDelim('}') {
					out.AudienceMap = make(map[string]string)
				} else {
					out.AudienceMap = nil
				}
				for !in.IsDelim('}') {
					key := string(in.String())
					in.WantColon()
					var v7 string
					v7 = string(in.String())
					(out.AudienceMap)[key] = v7
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
func easyjsonA0a30a08EncodeGithubComQordobacodeCliV2PkgTypes1(out *jwriter.Writer, in QordobaConfig) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"AccessToken\":"
		out.RawString(prefix[1:])
		out.String(string(in.AccessToken))
	}
	{
		const prefix string = ",\"OrganizationID\":"
		out.RawString(prefix)
		out.Int64(int64(in.OrganizationID))
	}
	{
		const prefix string = ",\"WorkspaceID\":"
		out.RawString(prefix)
		out.Int64(int64(in.WorkspaceID))
	}
	{
		const prefix string = ",\"ProjectID\":"
		out.RawString(prefix)
		out.Int64(int64(in.ProjectID))
	}
	{
		const prefix string = ",\"AudienceMap\":"
		out.RawString(prefix)
		if in.AudienceMap == nil && (out.Flags&jwriter.NilMapAsEmpty) == 0 {
			out.RawString(`null`)
		} else {
			out.RawByte('{')
			v8First := true
			for v8Name, v8Value := range in.AudienceMap {
				if v8First {
					v8First = false
				} else {
					out.RawByte(',')
				}
				out.String(string(v8Name))
				out.RawByte(':')
				out.String(string(v8Value))
			}
			out.RawByte('}')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v QordobaConfig) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonA0a30a08EncodeGithubComQordobacodeCliV2PkgTypes1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v QordobaConfig) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonA0a30a08EncodeGithubComQordobacodeCliV2PkgTypes1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *QordobaConfig) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonA0a30a08DecodeGithubComQordobacodeCliV2PkgTypes1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *QordobaConfig) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonA0a30a08DecodeGithubComQordobacodeCliV2PkgTypes1(l, v)
}
func easyjsonA0a30a08DecodeGithubComQordobacodeCliV2PkgTypes2(in *jlexer.Lexer, out *PushConfig) {
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
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "Sources":
			(out.Sources).UnmarshalEasyJSON(in)
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
func easyjsonA0a30a08EncodeGithubComQordobacodeCliV2PkgTypes2(out *jwriter.Writer, in PushConfig) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"Sources\":"
		out.RawString(prefix[1:])
		(in.Sources).MarshalEasyJSON(out)
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v PushConfig) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonA0a30a08EncodeGithubComQordobacodeCliV2PkgTypes2(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v PushConfig) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonA0a30a08EncodeGithubComQordobacodeCliV2PkgTypes2(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *PushConfig) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonA0a30a08DecodeGithubComQordobacodeCliV2PkgTypes2(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *PushConfig) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonA0a30a08DecodeGithubComQordobacodeCliV2PkgTypes2(l, v)
}
func easyjsonA0a30a08DecodeGithubComQordobacodeCliV2PkgTypes3(in *jlexer.Lexer, out *DownloadConfig) {
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
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "Targets":
			if in.IsNull() {
				in.Skip()
				out.Targets = nil
			} else {
				in.Delim('[')
				if out.Targets == nil {
					if !in.IsDelim(']') {
						out.Targets = make([]string, 0, 4)
					} else {
						out.Targets = []string{}
					}
				} else {
					out.Targets = (out.Targets)[:0]
				}
				for !in.IsDelim(']') {
					var v9 string
					v9 = string(in.String())
					out.Targets = append(out.Targets, v9)
					in.WantComma()
				}
				in.Delim(']')
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
func easyjsonA0a30a08EncodeGithubComQordobacodeCliV2PkgTypes3(out *jwriter.Writer, in DownloadConfig) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"Targets\":"
		out.RawString(prefix[1:])
		if in.Targets == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v10, v11 := range in.Targets {
				if v10 > 0 {
					out.RawByte(',')
				}
				out.String(string(v11))
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v DownloadConfig) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonA0a30a08EncodeGithubComQordobacodeCliV2PkgTypes3(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v DownloadConfig) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonA0a30a08EncodeGithubComQordobacodeCliV2PkgTypes3(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *DownloadConfig) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonA0a30a08DecodeGithubComQordobacodeCliV2PkgTypes3(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *DownloadConfig) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonA0a30a08DecodeGithubComQordobacodeCliV2PkgTypes3(l, v)
}
func easyjsonA0a30a08DecodeGithubComQordobacodeCliV2PkgTypes4(in *jlexer.Lexer, out *Config) {
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
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "Qordoba":
			(out.Qordoba).UnmarshalEasyJSON(in)
		case "Push":
			(out.Push).UnmarshalEasyJSON(in)
		case "Download":
			(out.Download).UnmarshalEasyJSON(in)
		case "Blacklist":
			(out.Blacklist).UnmarshalEasyJSON(in)
		case "BaseURL":
			out.BaseURL = string(in.String())
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
func easyjsonA0a30a08EncodeGithubComQordobacodeCliV2PkgTypes4(out *jwriter.Writer, in Config) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"Qordoba\":"
		out.RawString(prefix[1:])
		(in.Qordoba).MarshalEasyJSON(out)
	}
	{
		const prefix string = ",\"Push\":"
		out.RawString(prefix)
		(in.Push).MarshalEasyJSON(out)
	}
	{
		const prefix string = ",\"Download\":"
		out.RawString(prefix)
		(in.Download).MarshalEasyJSON(out)
	}
	{
		const prefix string = ",\"Blacklist\":"
		out.RawString(prefix)
		(in.Blacklist).MarshalEasyJSON(out)
	}
	{
		const prefix string = ",\"BaseURL\":"
		out.RawString(prefix)
		out.String(string(in.BaseURL))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Config) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonA0a30a08EncodeGithubComQordobacodeCliV2PkgTypes4(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Config) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonA0a30a08EncodeGithubComQordobacodeCliV2PkgTypes4(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Config) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonA0a30a08DecodeGithubComQordobacodeCliV2PkgTypes4(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Config) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonA0a30a08DecodeGithubComQordobacodeCliV2PkgTypes4(l, v)
}
func easyjsonA0a30a08DecodeGithubComQordobacodeCliV2PkgTypes5(in *jlexer.Lexer, out *BlacklistConfig) {
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
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "Sources":
			if in.IsNull() {
				in.Skip()
				out.Sources = nil
			} else {
				in.Delim('[')
				if out.Sources == nil {
					if !in.IsDelim(']') {
						out.Sources = make([]string, 0, 4)
					} else {
						out.Sources = []string{}
					}
				} else {
					out.Sources = (out.Sources)[:0]
				}
				for !in.IsDelim(']') {
					var v12 string
					v12 = string(in.String())
					out.Sources = append(out.Sources, v12)
					in.WantComma()
				}
				in.Delim(']')
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
func easyjsonA0a30a08EncodeGithubComQordobacodeCliV2PkgTypes5(out *jwriter.Writer, in BlacklistConfig) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"Sources\":"
		out.RawString(prefix[1:])
		if in.Sources == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v13, v14 := range in.Sources {
				if v13 > 0 {
					out.RawByte(',')
				}
				out.String(string(v14))
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v BlacklistConfig) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonA0a30a08EncodeGithubComQordobacodeCliV2PkgTypes5(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v BlacklistConfig) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonA0a30a08EncodeGithubComQordobacodeCliV2PkgTypes5(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *BlacklistConfig) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonA0a30a08DecodeGithubComQordobacodeCliV2PkgTypes5(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *BlacklistConfig) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonA0a30a08DecodeGithubComQordobacodeCliV2PkgTypes5(l, v)
}