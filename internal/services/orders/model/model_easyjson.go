// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package model

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

func easyjsonC80ae7adDecodeCoffeeshopApiInternalServicesOrdersModel(in *jlexer.Lexer, out *OrdersStatusesResDelivery) {
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
		case "id":
			out.OrderID = uint64(in.Uint64())
		case "created_at":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.OrderCreatedAt).UnmarshalJSON(data))
			}
		case "status":
			out.OrderStatus = string(in.String())
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
func easyjsonC80ae7adEncodeCoffeeshopApiInternalServicesOrdersModel(out *jwriter.Writer, in OrdersStatusesResDelivery) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix[1:])
		out.Uint64(uint64(in.OrderID))
	}
	{
		const prefix string = ",\"created_at\":"
		out.RawString(prefix)
		out.Raw((in.OrderCreatedAt).MarshalJSON())
	}
	{
		const prefix string = ",\"status\":"
		out.RawString(prefix)
		out.String(string(in.OrderStatus))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v OrdersStatusesResDelivery) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC80ae7adEncodeCoffeeshopApiInternalServicesOrdersModel(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v OrdersStatusesResDelivery) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC80ae7adEncodeCoffeeshopApiInternalServicesOrdersModel(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *OrdersStatusesResDelivery) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC80ae7adDecodeCoffeeshopApiInternalServicesOrdersModel(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *OrdersStatusesResDelivery) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC80ae7adDecodeCoffeeshopApiInternalServicesOrdersModel(l, v)
}
func easyjsonC80ae7adDecodeCoffeeshopApiInternalServicesOrdersModel1(in *jlexer.Lexer, out *OrdersStatusesReqDelivery) {
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
func easyjsonC80ae7adEncodeCoffeeshopApiInternalServicesOrdersModel1(out *jwriter.Writer, in OrdersStatusesReqDelivery) {
	out.RawByte('{')
	first := true
	_ = first
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v OrdersStatusesReqDelivery) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC80ae7adEncodeCoffeeshopApiInternalServicesOrdersModel1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v OrdersStatusesReqDelivery) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC80ae7adEncodeCoffeeshopApiInternalServicesOrdersModel1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *OrdersStatusesReqDelivery) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC80ae7adDecodeCoffeeshopApiInternalServicesOrdersModel1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *OrdersStatusesReqDelivery) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC80ae7adDecodeCoffeeshopApiInternalServicesOrdersModel1(l, v)
}
func easyjsonC80ae7adDecodeCoffeeshopApiInternalServicesOrdersModel2(in *jlexer.Lexer, out *ListOrdersResDelivery) {
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
		case "orders":
			if in.IsNull() {
				in.Skip()
				out.Orders = nil
			} else {
				in.Delim('[')
				if out.Orders == nil {
					if !in.IsDelim(']') {
						out.Orders = make([]ListOrdersOrder, 0, 4)
					} else {
						out.Orders = []ListOrdersOrder{}
					}
				} else {
					out.Orders = (out.Orders)[:0]
				}
				for !in.IsDelim(']') {
					var v1 ListOrdersOrder
					easyjsonC80ae7adDecodeCoffeeshopApiInternalServicesOrdersModel3(in, &v1)
					out.Orders = append(out.Orders, v1)
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
func easyjsonC80ae7adEncodeCoffeeshopApiInternalServicesOrdersModel2(out *jwriter.Writer, in ListOrdersResDelivery) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"orders\":"
		out.RawString(prefix[1:])
		if in.Orders == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v2, v3 := range in.Orders {
				if v2 > 0 {
					out.RawByte(',')
				}
				easyjsonC80ae7adEncodeCoffeeshopApiInternalServicesOrdersModel3(out, v3)
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v ListOrdersResDelivery) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC80ae7adEncodeCoffeeshopApiInternalServicesOrdersModel2(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v ListOrdersResDelivery) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC80ae7adEncodeCoffeeshopApiInternalServicesOrdersModel2(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *ListOrdersResDelivery) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC80ae7adDecodeCoffeeshopApiInternalServicesOrdersModel2(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *ListOrdersResDelivery) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC80ae7adDecodeCoffeeshopApiInternalServicesOrdersModel2(l, v)
}
func easyjsonC80ae7adDecodeCoffeeshopApiInternalServicesOrdersModel3(in *jlexer.Lexer, out *ListOrdersOrder) {
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
		case "id":
			out.OrderID = uint64(in.Uint64())
		case "created_at":
			out.OrderCreatedAt = uint64(in.Uint64())
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
func easyjsonC80ae7adEncodeCoffeeshopApiInternalServicesOrdersModel3(out *jwriter.Writer, in ListOrdersOrder) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix[1:])
		out.Uint64(uint64(in.OrderID))
	}
	{
		const prefix string = ",\"created_at\":"
		out.RawString(prefix)
		out.Uint64(uint64(in.OrderCreatedAt))
	}
	out.RawByte('}')
}
func easyjsonC80ae7adDecodeCoffeeshopApiInternalServicesOrdersModel4(in *jlexer.Lexer, out *ListOrdersReqDelivery) {
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
		case "offset":
			out.Offset = uint64(in.Uint64())
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
func easyjsonC80ae7adEncodeCoffeeshopApiInternalServicesOrdersModel4(out *jwriter.Writer, in ListOrdersReqDelivery) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"offset\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Uint64(uint64(in.Offset))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v ListOrdersReqDelivery) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC80ae7adEncodeCoffeeshopApiInternalServicesOrdersModel4(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v ListOrdersReqDelivery) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC80ae7adEncodeCoffeeshopApiInternalServicesOrdersModel4(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *ListOrdersReqDelivery) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC80ae7adDecodeCoffeeshopApiInternalServicesOrdersModel4(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *ListOrdersReqDelivery) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC80ae7adDecodeCoffeeshopApiInternalServicesOrdersModel4(l, v)
}
func easyjsonC80ae7adDecodeCoffeeshopApiInternalServicesOrdersModel5(in *jlexer.Lexer, out *GetOrderInfoResDelivery) {
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
		case "order":
			easyjsonC80ae7adDecodeCoffeeshopApiInternalServicesOrdersModel6(in, &out.Order)
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
func easyjsonC80ae7adEncodeCoffeeshopApiInternalServicesOrdersModel5(out *jwriter.Writer, in GetOrderInfoResDelivery) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"order\":"
		out.RawString(prefix[1:])
		easyjsonC80ae7adEncodeCoffeeshopApiInternalServicesOrdersModel6(out, in.Order)
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v GetOrderInfoResDelivery) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC80ae7adEncodeCoffeeshopApiInternalServicesOrdersModel5(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetOrderInfoResDelivery) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC80ae7adEncodeCoffeeshopApiInternalServicesOrdersModel5(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetOrderInfoResDelivery) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC80ae7adDecodeCoffeeshopApiInternalServicesOrdersModel5(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetOrderInfoResDelivery) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC80ae7adDecodeCoffeeshopApiInternalServicesOrdersModel5(l, v)
}
func easyjsonC80ae7adDecodeCoffeeshopApiInternalServicesOrdersModel6(in *jlexer.Lexer, out *GetOrderInfoOrder) {
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
		case "id":
			out.OrderID = uint64(in.Uint64())
		case "created_at":
			out.CreatedAt = uint64(in.Uint64())
		case "status":
			out.Status = string(in.String())
		case "address":
			out.Address = string(in.String())
		case "items":
			if in.IsNull() {
				in.Skip()
				out.Items = nil
			} else {
				in.Delim('[')
				if out.Items == nil {
					if !in.IsDelim(']') {
						out.Items = make([]GetOrderInfoOrderItem, 0, 2)
					} else {
						out.Items = []GetOrderInfoOrderItem{}
					}
				} else {
					out.Items = (out.Items)[:0]
				}
				for !in.IsDelim(']') {
					var v4 GetOrderInfoOrderItem
					easyjsonC80ae7adDecodeCoffeeshopApiInternalServicesOrdersModel7(in, &v4)
					out.Items = append(out.Items, v4)
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
func easyjsonC80ae7adEncodeCoffeeshopApiInternalServicesOrdersModel6(out *jwriter.Writer, in GetOrderInfoOrder) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix[1:])
		out.Uint64(uint64(in.OrderID))
	}
	{
		const prefix string = ",\"created_at\":"
		out.RawString(prefix)
		out.Uint64(uint64(in.CreatedAt))
	}
	{
		const prefix string = ",\"status\":"
		out.RawString(prefix)
		out.String(string(in.Status))
	}
	{
		const prefix string = ",\"address\":"
		out.RawString(prefix)
		out.String(string(in.Address))
	}
	{
		const prefix string = ",\"items\":"
		out.RawString(prefix)
		if in.Items == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v5, v6 := range in.Items {
				if v5 > 0 {
					out.RawByte(',')
				}
				easyjsonC80ae7adEncodeCoffeeshopApiInternalServicesOrdersModel7(out, v6)
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}
func easyjsonC80ae7adDecodeCoffeeshopApiInternalServicesOrdersModel7(in *jlexer.Lexer, out *GetOrderInfoOrderItem) {
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
		case "coffee_id":
			out.CoffeeID = uint64(in.Uint64())
		case "topping":
			out.Topping = string(in.String())
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
func easyjsonC80ae7adEncodeCoffeeshopApiInternalServicesOrdersModel7(out *jwriter.Writer, in GetOrderInfoOrderItem) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"coffee_id\":"
		out.RawString(prefix[1:])
		out.Uint64(uint64(in.CoffeeID))
	}
	{
		const prefix string = ",\"topping\":"
		out.RawString(prefix)
		out.String(string(in.Topping))
	}
	out.RawByte('}')
}
func easyjsonC80ae7adDecodeCoffeeshopApiInternalServicesOrdersModel8(in *jlexer.Lexer, out *GetOrderInfoReqDelivery) {
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
		case "id":
			out.OrderID = uint64(in.Uint64())
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
func easyjsonC80ae7adEncodeCoffeeshopApiInternalServicesOrdersModel8(out *jwriter.Writer, in GetOrderInfoReqDelivery) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Uint64(uint64(in.OrderID))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v GetOrderInfoReqDelivery) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC80ae7adEncodeCoffeeshopApiInternalServicesOrdersModel8(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetOrderInfoReqDelivery) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC80ae7adEncodeCoffeeshopApiInternalServicesOrdersModel8(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetOrderInfoReqDelivery) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC80ae7adDecodeCoffeeshopApiInternalServicesOrdersModel8(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetOrderInfoReqDelivery) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC80ae7adDecodeCoffeeshopApiInternalServicesOrdersModel8(l, v)
}
func easyjsonC80ae7adDecodeCoffeeshopApiInternalServicesOrdersModel9(in *jlexer.Lexer, out *EmployeeCompleteOrderResDelivery) {
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
		case "id":
			out.OrderID = uint64(in.Uint64())
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
func easyjsonC80ae7adEncodeCoffeeshopApiInternalServicesOrdersModel9(out *jwriter.Writer, in EmployeeCompleteOrderResDelivery) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix[1:])
		out.Uint64(uint64(in.OrderID))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v EmployeeCompleteOrderResDelivery) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC80ae7adEncodeCoffeeshopApiInternalServicesOrdersModel9(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v EmployeeCompleteOrderResDelivery) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC80ae7adEncodeCoffeeshopApiInternalServicesOrdersModel9(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *EmployeeCompleteOrderResDelivery) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC80ae7adDecodeCoffeeshopApiInternalServicesOrdersModel9(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *EmployeeCompleteOrderResDelivery) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC80ae7adDecodeCoffeeshopApiInternalServicesOrdersModel9(l, v)
}
func easyjsonC80ae7adDecodeCoffeeshopApiInternalServicesOrdersModel10(in *jlexer.Lexer, out *EmployeeCompleteOrderReqDelivery) {
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
		case "id":
			out.OrderID = uint64(in.Uint64())
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
func easyjsonC80ae7adEncodeCoffeeshopApiInternalServicesOrdersModel10(out *jwriter.Writer, in EmployeeCompleteOrderReqDelivery) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix[1:])
		out.Uint64(uint64(in.OrderID))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v EmployeeCompleteOrderReqDelivery) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC80ae7adEncodeCoffeeshopApiInternalServicesOrdersModel10(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v EmployeeCompleteOrderReqDelivery) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC80ae7adEncodeCoffeeshopApiInternalServicesOrdersModel10(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *EmployeeCompleteOrderReqDelivery) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC80ae7adDecodeCoffeeshopApiInternalServicesOrdersModel10(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *EmployeeCompleteOrderReqDelivery) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC80ae7adDecodeCoffeeshopApiInternalServicesOrdersModel10(l, v)
}
func easyjsonC80ae7adDecodeCoffeeshopApiInternalServicesOrdersModel11(in *jlexer.Lexer, out *CreateOrderResDelivery) {
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
		case "id":
			out.OrderID = uint64(in.Uint64())
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
func easyjsonC80ae7adEncodeCoffeeshopApiInternalServicesOrdersModel11(out *jwriter.Writer, in CreateOrderResDelivery) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix[1:])
		out.Uint64(uint64(in.OrderID))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v CreateOrderResDelivery) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC80ae7adEncodeCoffeeshopApiInternalServicesOrdersModel11(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v CreateOrderResDelivery) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC80ae7adEncodeCoffeeshopApiInternalServicesOrdersModel11(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *CreateOrderResDelivery) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC80ae7adDecodeCoffeeshopApiInternalServicesOrdersModel11(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *CreateOrderResDelivery) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC80ae7adDecodeCoffeeshopApiInternalServicesOrdersModel11(l, v)
}
func easyjsonC80ae7adDecodeCoffeeshopApiInternalServicesOrdersModel12(in *jlexer.Lexer, out *CreateOrderReqDelivery) {
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
		case "address":
			out.Address = string(in.String())
		case "items":
			if in.IsNull() {
				in.Skip()
				out.Items = nil
			} else {
				in.Delim('[')
				if out.Items == nil {
					if !in.IsDelim(']') {
						out.Items = make([]CreateOrderOrderItem, 0, 4)
					} else {
						out.Items = []CreateOrderOrderItem{}
					}
				} else {
					out.Items = (out.Items)[:0]
				}
				for !in.IsDelim(']') {
					var v7 CreateOrderOrderItem
					(v7).UnmarshalEasyJSON(in)
					out.Items = append(out.Items, v7)
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
func easyjsonC80ae7adEncodeCoffeeshopApiInternalServicesOrdersModel12(out *jwriter.Writer, in CreateOrderReqDelivery) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"address\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Address))
	}
	{
		const prefix string = ",\"items\":"
		out.RawString(prefix)
		if in.Items == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v8, v9 := range in.Items {
				if v8 > 0 {
					out.RawByte(',')
				}
				(v9).MarshalEasyJSON(out)
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v CreateOrderReqDelivery) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC80ae7adEncodeCoffeeshopApiInternalServicesOrdersModel12(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v CreateOrderReqDelivery) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC80ae7adEncodeCoffeeshopApiInternalServicesOrdersModel12(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *CreateOrderReqDelivery) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC80ae7adDecodeCoffeeshopApiInternalServicesOrdersModel12(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *CreateOrderReqDelivery) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC80ae7adDecodeCoffeeshopApiInternalServicesOrdersModel12(l, v)
}
func easyjsonC80ae7adDecodeCoffeeshopApiInternalServicesOrdersModel13(in *jlexer.Lexer, out *CreateOrderOrderItem) {
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
		case "coffee_id":
			out.CoffeeID = uint64(in.Uint64())
		case "topping":
			if in.IsNull() {
				in.Skip()
				out.Topping = nil
			} else {
				if out.Topping == nil {
					out.Topping = new(string)
				}
				*out.Topping = string(in.String())
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
func easyjsonC80ae7adEncodeCoffeeshopApiInternalServicesOrdersModel13(out *jwriter.Writer, in CreateOrderOrderItem) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"coffee_id\":"
		out.RawString(prefix[1:])
		out.Uint64(uint64(in.CoffeeID))
	}
	{
		const prefix string = ",\"topping\":"
		out.RawString(prefix)
		if in.Topping == nil {
			out.RawString("null")
		} else {
			out.String(string(*in.Topping))
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v CreateOrderOrderItem) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC80ae7adEncodeCoffeeshopApiInternalServicesOrdersModel13(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v CreateOrderOrderItem) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC80ae7adEncodeCoffeeshopApiInternalServicesOrdersModel13(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *CreateOrderOrderItem) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC80ae7adDecodeCoffeeshopApiInternalServicesOrdersModel13(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *CreateOrderOrderItem) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC80ae7adDecodeCoffeeshopApiInternalServicesOrdersModel13(l, v)
}
func easyjsonC80ae7adDecodeCoffeeshopApiInternalServicesOrdersModel14(in *jlexer.Lexer, out *CancelOrderResDelivery) {
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
		case "id":
			out.OrderID = uint64(in.Uint64())
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
func easyjsonC80ae7adEncodeCoffeeshopApiInternalServicesOrdersModel14(out *jwriter.Writer, in CancelOrderResDelivery) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix[1:])
		out.Uint64(uint64(in.OrderID))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v CancelOrderResDelivery) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC80ae7adEncodeCoffeeshopApiInternalServicesOrdersModel14(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v CancelOrderResDelivery) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC80ae7adEncodeCoffeeshopApiInternalServicesOrdersModel14(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *CancelOrderResDelivery) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC80ae7adDecodeCoffeeshopApiInternalServicesOrdersModel14(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *CancelOrderResDelivery) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC80ae7adDecodeCoffeeshopApiInternalServicesOrdersModel14(l, v)
}
func easyjsonC80ae7adDecodeCoffeeshopApiInternalServicesOrdersModel15(in *jlexer.Lexer, out *CancelOrderReqDelivery) {
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
		case "id":
			out.OrderID = uint64(in.Uint64())
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
func easyjsonC80ae7adEncodeCoffeeshopApiInternalServicesOrdersModel15(out *jwriter.Writer, in CancelOrderReqDelivery) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Uint64(uint64(in.OrderID))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v CancelOrderReqDelivery) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC80ae7adEncodeCoffeeshopApiInternalServicesOrdersModel15(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v CancelOrderReqDelivery) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC80ae7adEncodeCoffeeshopApiInternalServicesOrdersModel15(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *CancelOrderReqDelivery) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC80ae7adDecodeCoffeeshopApiInternalServicesOrdersModel15(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *CancelOrderReqDelivery) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC80ae7adDecodeCoffeeshopApiInternalServicesOrdersModel15(l, v)
}
