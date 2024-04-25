// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cidrtypes

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

var (
	_ basetypes.StringTypable = (*IPv4PrefixType)(nil)
)

// IPv4PrefixType is an attribute type that represents a valid IPv4 CIDR string (RFC 4632). No semantic equality
// logic is defined for IPv4PrefixType, so it will follow Terraform's data-consistency rules for strings, which must match byte-for-byte.
type IPv4PrefixType struct {
	basetypes.StringType
}

// String returns a human readable string of the type name.
func (t IPv4PrefixType) String() string {
	return "cidrtypes.IPv4PrefixType"
}

// ValueType returns the Value type.
func (t IPv4PrefixType) ValueType(ctx context.Context) attr.Value {
	return IPv4Prefix{}
}

// Equal returns true if the given type is equivalent.
func (t IPv4PrefixType) Equal(o attr.Type) bool {
	other, ok := o.(IPv4PrefixType)

	if !ok {
		return false
	}

	return t.StringType.Equal(other.StringType)
}

// ValueFromString returns a StringValuable type given a StringValue.
func (t IPv4PrefixType) ValueFromString(ctx context.Context, in basetypes.StringValue) (basetypes.StringValuable, diag.Diagnostics) {
	return IPv4Prefix{
		StringValue: in,
	}, nil
}

// ValueFromTerraform returns a Value given a tftypes.Value.  This is meant to convert the tftypes.Value into a more convenient Go type
// for the provider to consume the data with.
func (t IPv4PrefixType) ValueFromTerraform(ctx context.Context, in tftypes.Value) (attr.Value, error) {
	attrValue, err := t.StringType.ValueFromTerraform(ctx, in)

	if err != nil {
		return nil, err
	}

	stringValue, ok := attrValue.(basetypes.StringValue)

	if !ok {
		return nil, fmt.Errorf("unexpected value type of %T", attrValue)
	}

	stringValuable, diags := t.ValueFromString(ctx, stringValue)

	if diags.HasError() {
		return nil, fmt.Errorf("unexpected error converting StringValue to StringValuable: %v", diags)
	}

	return stringValuable, nil
}
