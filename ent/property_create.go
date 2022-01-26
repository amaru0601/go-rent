// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/amaru0601/go-rent/ent/contract"
	"github.com/amaru0601/go-rent/ent/property"
	"github.com/amaru0601/go-rent/ent/user"
)

// PropertyCreate is the builder for creating a Property entity.
type PropertyCreate struct {
	config
	mutation *PropertyMutation
	hooks    []Hook
}

// SetClass sets the "class" field.
func (pc *PropertyCreate) SetClass(pr property.Class) *PropertyCreate {
	pc.mutation.SetClass(pr)
	return pc
}

// SetNillableClass sets the "class" field if the given value is not nil.
func (pc *PropertyCreate) SetNillableClass(pr *property.Class) *PropertyCreate {
	if pr != nil {
		pc.SetClass(*pr)
	}
	return pc
}

// SetAddress sets the "address" field.
func (pc *PropertyCreate) SetAddress(s string) *PropertyCreate {
	pc.mutation.SetAddress(s)
	return pc
}

// SetCity sets the "city" field.
func (pc *PropertyCreate) SetCity(s string) *PropertyCreate {
	pc.mutation.SetCity(s)
	return pc
}

// SetDescription sets the "description" field.
func (pc *PropertyCreate) SetDescription(s string) *PropertyCreate {
	pc.mutation.SetDescription(s)
	return pc
}

// SetDeleted sets the "deleted" field.
func (pc *PropertyCreate) SetDeleted(b bool) *PropertyCreate {
	pc.mutation.SetDeleted(b)
	return pc
}

// SetOwnerID sets the "owner" edge to the User entity by ID.
func (pc *PropertyCreate) SetOwnerID(id int) *PropertyCreate {
	pc.mutation.SetOwnerID(id)
	return pc
}

// SetNillableOwnerID sets the "owner" edge to the User entity by ID if the given value is not nil.
func (pc *PropertyCreate) SetNillableOwnerID(id *int) *PropertyCreate {
	if id != nil {
		pc = pc.SetOwnerID(*id)
	}
	return pc
}

// SetOwner sets the "owner" edge to the User entity.
func (pc *PropertyCreate) SetOwner(u *User) *PropertyCreate {
	return pc.SetOwnerID(u.ID)
}

// SetContractID sets the "contract" edge to the Contract entity by ID.
func (pc *PropertyCreate) SetContractID(id int) *PropertyCreate {
	pc.mutation.SetContractID(id)
	return pc
}

// SetNillableContractID sets the "contract" edge to the Contract entity by ID if the given value is not nil.
func (pc *PropertyCreate) SetNillableContractID(id *int) *PropertyCreate {
	if id != nil {
		pc = pc.SetContractID(*id)
	}
	return pc
}

// SetContract sets the "contract" edge to the Contract entity.
func (pc *PropertyCreate) SetContract(c *Contract) *PropertyCreate {
	return pc.SetContractID(c.ID)
}

// Mutation returns the PropertyMutation object of the builder.
func (pc *PropertyCreate) Mutation() *PropertyMutation {
	return pc.mutation
}

// Save creates the Property in the database.
func (pc *PropertyCreate) Save(ctx context.Context) (*Property, error) {
	var (
		err  error
		node *Property
	)
	pc.defaults()
	if len(pc.hooks) == 0 {
		if err = pc.check(); err != nil {
			return nil, err
		}
		node, err = pc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PropertyMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = pc.check(); err != nil {
				return nil, err
			}
			pc.mutation = mutation
			if node, err = pc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(pc.hooks) - 1; i >= 0; i-- {
			if pc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = pc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, pc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (pc *PropertyCreate) SaveX(ctx context.Context) *Property {
	v, err := pc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pc *PropertyCreate) Exec(ctx context.Context) error {
	_, err := pc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pc *PropertyCreate) ExecX(ctx context.Context) {
	if err := pc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pc *PropertyCreate) defaults() {
	if _, ok := pc.mutation.Class(); !ok {
		v := property.DefaultClass
		pc.mutation.SetClass(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pc *PropertyCreate) check() error {
	if _, ok := pc.mutation.Class(); !ok {
		return &ValidationError{Name: "class", err: errors.New(`ent: missing required field "class"`)}
	}
	if v, ok := pc.mutation.Class(); ok {
		if err := property.ClassValidator(v); err != nil {
			return &ValidationError{Name: "class", err: fmt.Errorf(`ent: validator failed for field "class": %w`, err)}
		}
	}
	if _, ok := pc.mutation.Address(); !ok {
		return &ValidationError{Name: "address", err: errors.New(`ent: missing required field "address"`)}
	}
	if _, ok := pc.mutation.City(); !ok {
		return &ValidationError{Name: "city", err: errors.New(`ent: missing required field "city"`)}
	}
	if _, ok := pc.mutation.Description(); !ok {
		return &ValidationError{Name: "description", err: errors.New(`ent: missing required field "description"`)}
	}
	if _, ok := pc.mutation.Deleted(); !ok {
		return &ValidationError{Name: "deleted", err: errors.New(`ent: missing required field "deleted"`)}
	}
	return nil
}

func (pc *PropertyCreate) sqlSave(ctx context.Context) (*Property, error) {
	_node, _spec := pc.createSpec()
	if err := sqlgraph.CreateNode(ctx, pc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (pc *PropertyCreate) createSpec() (*Property, *sqlgraph.CreateSpec) {
	var (
		_node = &Property{config: pc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: property.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: property.FieldID,
			},
		}
	)
	if value, ok := pc.mutation.Class(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: property.FieldClass,
		})
		_node.Class = value
	}
	if value, ok := pc.mutation.Address(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: property.FieldAddress,
		})
		_node.Address = value
	}
	if value, ok := pc.mutation.City(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: property.FieldCity,
		})
		_node.City = value
	}
	if value, ok := pc.mutation.Description(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: property.FieldDescription,
		})
		_node.Description = value
	}
	if value, ok := pc.mutation.Deleted(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: property.FieldDeleted,
		})
		_node.Deleted = value
	}
	if nodes := pc.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   property.OwnerTable,
			Columns: []string{property.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.user_properties = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pc.mutation.ContractIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   property.ContractTable,
			Columns: []string{property.ContractColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: contract.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// PropertyCreateBulk is the builder for creating many Property entities in bulk.
type PropertyCreateBulk struct {
	config
	builders []*PropertyCreate
}

// Save creates the Property entities in the database.
func (pcb *PropertyCreateBulk) Save(ctx context.Context) ([]*Property, error) {
	specs := make([]*sqlgraph.CreateSpec, len(pcb.builders))
	nodes := make([]*Property, len(pcb.builders))
	mutators := make([]Mutator, len(pcb.builders))
	for i := range pcb.builders {
		func(i int, root context.Context) {
			builder := pcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*PropertyMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, pcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, pcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{err.Error(), err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, pcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (pcb *PropertyCreateBulk) SaveX(ctx context.Context) []*Property {
	v, err := pcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pcb *PropertyCreateBulk) Exec(ctx context.Context) error {
	_, err := pcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pcb *PropertyCreateBulk) ExecX(ctx context.Context) {
	if err := pcb.Exec(ctx); err != nil {
		panic(err)
	}
}
