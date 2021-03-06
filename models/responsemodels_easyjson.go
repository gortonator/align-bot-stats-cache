// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package models

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

func easyjson83f2289fDecodeGithubComRapidclockAlignBotStatsCacheModels(in *jlexer.Lexer, out *TotalProgramCost) {
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
		case "totalCost":
			out.TotalCost = string(in.String())
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
func easyjson83f2289fEncodeGithubComRapidclockAlignBotStatsCacheModels(out *jwriter.Writer, in TotalProgramCost) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"totalCost\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.TotalCost))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v TotalProgramCost) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson83f2289fEncodeGithubComRapidclockAlignBotStatsCacheModels(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v TotalProgramCost) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson83f2289fEncodeGithubComRapidclockAlignBotStatsCacheModels(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *TotalProgramCost) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson83f2289fDecodeGithubComRapidclockAlignBotStatsCacheModels(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *TotalProgramCost) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson83f2289fDecodeGithubComRapidclockAlignBotStatsCacheModels(l, v)
}
func easyjson83f2289fDecodeGithubComRapidclockAlignBotStatsCacheModels1(in *jlexer.Lexer, out *TopEmployers) {
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
		case "employers":
			if in.IsNull() {
				in.Skip()
				out.Employers = nil
			} else {
				in.Delim('[')
				if out.Employers == nil {
					if !in.IsDelim(']') {
						out.Employers = make([]string, 0, 4)
					} else {
						out.Employers = []string{}
					}
				} else {
					out.Employers = (out.Employers)[:0]
				}
				for !in.IsDelim(']') {
					var v1 string
					v1 = string(in.String())
					out.Employers = append(out.Employers, v1)
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
func easyjson83f2289fEncodeGithubComRapidclockAlignBotStatsCacheModels1(out *jwriter.Writer, in TopEmployers) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"employers\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		if in.Employers == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v2, v3 := range in.Employers {
				if v2 > 0 {
					out.RawByte(',')
				}
				out.String(string(v3))
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v TopEmployers) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson83f2289fEncodeGithubComRapidclockAlignBotStatsCacheModels1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v TopEmployers) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson83f2289fEncodeGithubComRapidclockAlignBotStatsCacheModels1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *TopEmployers) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson83f2289fDecodeGithubComRapidclockAlignBotStatsCacheModels1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *TopEmployers) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson83f2289fDecodeGithubComRapidclockAlignBotStatsCacheModels1(l, v)
}
func easyjson83f2289fDecodeGithubComRapidclockAlignBotStatsCacheModels2(in *jlexer.Lexer, out *TopBackgrounds) {
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
		case "backgrounds":
			if in.IsNull() {
				in.Skip()
				out.Backgrounds = nil
			} else {
				in.Delim('[')
				if out.Backgrounds == nil {
					if !in.IsDelim(']') {
						out.Backgrounds = make([]string, 0, 4)
					} else {
						out.Backgrounds = []string{}
					}
				} else {
					out.Backgrounds = (out.Backgrounds)[:0]
				}
				for !in.IsDelim(']') {
					var v4 string
					v4 = string(in.String())
					out.Backgrounds = append(out.Backgrounds, v4)
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
func easyjson83f2289fEncodeGithubComRapidclockAlignBotStatsCacheModels2(out *jwriter.Writer, in TopBackgrounds) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"backgrounds\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		if in.Backgrounds == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v5, v6 := range in.Backgrounds {
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
func (v TopBackgrounds) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson83f2289fEncodeGithubComRapidclockAlignBotStatsCacheModels2(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v TopBackgrounds) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson83f2289fEncodeGithubComRapidclockAlignBotStatsCacheModels2(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *TopBackgrounds) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson83f2289fDecodeGithubComRapidclockAlignBotStatsCacheModels2(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *TopBackgrounds) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson83f2289fDecodeGithubComRapidclockAlignBotStatsCacheModels2(l, v)
}
func easyjson83f2289fDecodeGithubComRapidclockAlignBotStatsCacheModels3(in *jlexer.Lexer, out *PerCreditProgramCost) {
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
		case "perCreditCost":
			out.PerCreditCost = string(in.String())
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
func easyjson83f2289fEncodeGithubComRapidclockAlignBotStatsCacheModels3(out *jwriter.Writer, in PerCreditProgramCost) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"perCreditCost\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.PerCreditCost))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v PerCreditProgramCost) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson83f2289fEncodeGithubComRapidclockAlignBotStatsCacheModels3(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v PerCreditProgramCost) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson83f2289fEncodeGithubComRapidclockAlignBotStatsCacheModels3(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *PerCreditProgramCost) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson83f2289fDecodeGithubComRapidclockAlignBotStatsCacheModels3(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *PerCreditProgramCost) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson83f2289fDecodeGithubComRapidclockAlignBotStatsCacheModels3(l, v)
}
func easyjson83f2289fDecodeGithubComRapidclockAlignBotStatsCacheModels4(in *jlexer.Lexer, out *GraduatesCount) {
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
		case "totalGraduates":
			out.TotalGraduates = string(in.String())
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
func easyjson83f2289fEncodeGithubComRapidclockAlignBotStatsCacheModels4(out *jwriter.Writer, in GraduatesCount) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"totalGraduates\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.TotalGraduates))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v GraduatesCount) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson83f2289fEncodeGithubComRapidclockAlignBotStatsCacheModels4(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GraduatesCount) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson83f2289fEncodeGithubComRapidclockAlignBotStatsCacheModels4(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GraduatesCount) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson83f2289fDecodeGithubComRapidclockAlignBotStatsCacheModels4(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GraduatesCount) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson83f2289fDecodeGithubComRapidclockAlignBotStatsCacheModels4(l, v)
}
func easyjson83f2289fDecodeGithubComRapidclockAlignBotStatsCacheModels5(in *jlexer.Lexer, out *GenderRatio) {
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
		case "male":
			out.Male = string(in.String())
		case "female":
			out.Female = string(in.String())
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
func easyjson83f2289fEncodeGithubComRapidclockAlignBotStatsCacheModels5(out *jwriter.Writer, in GenderRatio) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"male\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Male))
	}
	{
		const prefix string = ",\"female\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Female))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v GenderRatio) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson83f2289fEncodeGithubComRapidclockAlignBotStatsCacheModels5(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GenderRatio) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson83f2289fEncodeGithubComRapidclockAlignBotStatsCacheModels5(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GenderRatio) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson83f2289fDecodeGithubComRapidclockAlignBotStatsCacheModels5(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GenderRatio) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson83f2289fDecodeGithubComRapidclockAlignBotStatsCacheModels5(l, v)
}
func easyjson83f2289fDecodeGithubComRapidclockAlignBotStatsCacheModels6(in *jlexer.Lexer, out *EmploymentRate) {
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
		case "rate":
			out.Rate = string(in.String())
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
func easyjson83f2289fEncodeGithubComRapidclockAlignBotStatsCacheModels6(out *jwriter.Writer, in EmploymentRate) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"rate\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Rate))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v EmploymentRate) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson83f2289fEncodeGithubComRapidclockAlignBotStatsCacheModels6(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v EmploymentRate) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson83f2289fEncodeGithubComRapidclockAlignBotStatsCacheModels6(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *EmploymentRate) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson83f2289fDecodeGithubComRapidclockAlignBotStatsCacheModels6(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *EmploymentRate) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson83f2289fDecodeGithubComRapidclockAlignBotStatsCacheModels6(l, v)
}
func easyjson83f2289fDecodeGithubComRapidclockAlignBotStatsCacheModels7(in *jlexer.Lexer, out *Employer) {
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
		case "name":
			out.Name = string(in.String())
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
func easyjson83f2289fEncodeGithubComRapidclockAlignBotStatsCacheModels7(out *jwriter.Writer, in Employer) {
	out.RawByte('{')
	first := true
	_ = first
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
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Employer) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson83f2289fEncodeGithubComRapidclockAlignBotStatsCacheModels7(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Employer) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson83f2289fEncodeGithubComRapidclockAlignBotStatsCacheModels7(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Employer) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson83f2289fDecodeGithubComRapidclockAlignBotStatsCacheModels7(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Employer) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson83f2289fDecodeGithubComRapidclockAlignBotStatsCacheModels7(l, v)
}
func easyjson83f2289fDecodeGithubComRapidclockAlignBotStatsCacheModels8(in *jlexer.Lexer, out *DropOutRate) {
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
		case "rate":
			out.Rate = string(in.String())
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
func easyjson83f2289fEncodeGithubComRapidclockAlignBotStatsCacheModels8(out *jwriter.Writer, in DropOutRate) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"rate\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Rate))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v DropOutRate) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson83f2289fEncodeGithubComRapidclockAlignBotStatsCacheModels8(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v DropOutRate) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson83f2289fEncodeGithubComRapidclockAlignBotStatsCacheModels8(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *DropOutRate) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson83f2289fDecodeGithubComRapidclockAlignBotStatsCacheModels8(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *DropOutRate) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson83f2289fDecodeGithubComRapidclockAlignBotStatsCacheModels8(l, v)
}
func easyjson83f2289fDecodeGithubComRapidclockAlignBotStatsCacheModels9(in *jlexer.Lexer, out *Background) {
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
		case "name":
			out.Name = string(in.String())
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
func easyjson83f2289fEncodeGithubComRapidclockAlignBotStatsCacheModels9(out *jwriter.Writer, in Background) {
	out.RawByte('{')
	first := true
	_ = first
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
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Background) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson83f2289fEncodeGithubComRapidclockAlignBotStatsCacheModels9(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Background) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson83f2289fEncodeGithubComRapidclockAlignBotStatsCacheModels9(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Background) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson83f2289fDecodeGithubComRapidclockAlignBotStatsCacheModels9(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Background) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson83f2289fDecodeGithubComRapidclockAlignBotStatsCacheModels9(l, v)
}
func easyjson83f2289fDecodeGithubComRapidclockAlignBotStatsCacheModels10(in *jlexer.Lexer, out *AverageSalary) {
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
		case "avgSalary":
			out.AvgSalary = string(in.String())
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
func easyjson83f2289fEncodeGithubComRapidclockAlignBotStatsCacheModels10(out *jwriter.Writer, in AverageSalary) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"avgSalary\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.AvgSalary))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v AverageSalary) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson83f2289fEncodeGithubComRapidclockAlignBotStatsCacheModels10(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v AverageSalary) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson83f2289fEncodeGithubComRapidclockAlignBotStatsCacheModels10(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *AverageSalary) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson83f2289fDecodeGithubComRapidclockAlignBotStatsCacheModels10(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *AverageSalary) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson83f2289fDecodeGithubComRapidclockAlignBotStatsCacheModels10(l, v)
}
func easyjson83f2289fDecodeGithubComRapidclockAlignBotStatsCacheModels11(in *jlexer.Lexer, out *AcceptanceRate) {
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
		case "rate":
			out.Rate = string(in.String())
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
func easyjson83f2289fEncodeGithubComRapidclockAlignBotStatsCacheModels11(out *jwriter.Writer, in AcceptanceRate) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"rate\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Rate))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v AcceptanceRate) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson83f2289fEncodeGithubComRapidclockAlignBotStatsCacheModels11(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v AcceptanceRate) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson83f2289fEncodeGithubComRapidclockAlignBotStatsCacheModels11(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *AcceptanceRate) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson83f2289fDecodeGithubComRapidclockAlignBotStatsCacheModels11(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *AcceptanceRate) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson83f2289fDecodeGithubComRapidclockAlignBotStatsCacheModels11(l, v)
}
