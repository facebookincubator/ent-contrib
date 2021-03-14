// Code generated by entc, DO NOT EDIT.

package messagewithfieldone

import (
	"entgo.io/contrib/entproto/internal/entprototest/ent/predicate"
	"entgo.io/ent/dialect/sql"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.MessageWithFieldOne {
	return predicate.MessageWithFieldOne(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.MessageWithFieldOne {
	return predicate.MessageWithFieldOne(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.MessageWithFieldOne {
	return predicate.MessageWithFieldOne(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.MessageWithFieldOne {
	return predicate.MessageWithFieldOne(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.MessageWithFieldOne {
	return predicate.MessageWithFieldOne(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.MessageWithFieldOne {
	return predicate.MessageWithFieldOne(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.MessageWithFieldOne {
	return predicate.MessageWithFieldOne(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.MessageWithFieldOne {
	return predicate.MessageWithFieldOne(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.MessageWithFieldOne {
	return predicate.MessageWithFieldOne(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// FieldOne applies equality check predicate on the "field_one" field. It's identical to FieldOneEQ.
func FieldOne(v int32) predicate.MessageWithFieldOne {
	return predicate.MessageWithFieldOne(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldFieldOne), v))
	})
}

// FieldOneEQ applies the EQ predicate on the "field_one" field.
func FieldOneEQ(v int32) predicate.MessageWithFieldOne {
	return predicate.MessageWithFieldOne(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldFieldOne), v))
	})
}

// FieldOneNEQ applies the NEQ predicate on the "field_one" field.
func FieldOneNEQ(v int32) predicate.MessageWithFieldOne {
	return predicate.MessageWithFieldOne(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldFieldOne), v))
	})
}

// FieldOneIn applies the In predicate on the "field_one" field.
func FieldOneIn(vs ...int32) predicate.MessageWithFieldOne {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.MessageWithFieldOne(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldFieldOne), v...))
	})
}

// FieldOneNotIn applies the NotIn predicate on the "field_one" field.
func FieldOneNotIn(vs ...int32) predicate.MessageWithFieldOne {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.MessageWithFieldOne(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldFieldOne), v...))
	})
}

// FieldOneGT applies the GT predicate on the "field_one" field.
func FieldOneGT(v int32) predicate.MessageWithFieldOne {
	return predicate.MessageWithFieldOne(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldFieldOne), v))
	})
}

// FieldOneGTE applies the GTE predicate on the "field_one" field.
func FieldOneGTE(v int32) predicate.MessageWithFieldOne {
	return predicate.MessageWithFieldOne(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldFieldOne), v))
	})
}

// FieldOneLT applies the LT predicate on the "field_one" field.
func FieldOneLT(v int32) predicate.MessageWithFieldOne {
	return predicate.MessageWithFieldOne(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldFieldOne), v))
	})
}

// FieldOneLTE applies the LTE predicate on the "field_one" field.
func FieldOneLTE(v int32) predicate.MessageWithFieldOne {
	return predicate.MessageWithFieldOne(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldFieldOne), v))
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.MessageWithFieldOne) predicate.MessageWithFieldOne {
	return predicate.MessageWithFieldOne(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.MessageWithFieldOne) predicate.MessageWithFieldOne {
	return predicate.MessageWithFieldOne(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.MessageWithFieldOne) predicate.MessageWithFieldOne {
	return predicate.MessageWithFieldOne(func(s *sql.Selector) {
		p(s.Not())
	})
}