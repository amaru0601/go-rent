// Code generated by entc, DO NOT EDIT.

package ent

import "context"

func (pr *Property) Owner(ctx context.Context) (*User, error) {
	result, err := pr.Edges.OwnerOrErr()
	if IsNotLoaded(err) {
		result, err = pr.QueryOwner().Only(ctx)
	}
	return result, MaskNotFound(err)
}

func (u *User) Properties(ctx context.Context) ([]*Property, error) {
	result, err := u.Edges.PropertiesOrErr()
	if IsNotLoaded(err) {
		result, err = u.QueryProperties().All(ctx)
	}
	return result, err
}