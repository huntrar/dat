package dat

//// DO NOT EDIT, auto-generated: godo builder-boilerplate

// Interpolate interpolates this builders sql.
func (b *CallBuilder) Interpolate() (string, []interface{}, error) {
	return interpolate(b)
}

// IsInterpolated determines if this builder will interpolate when
// Interpolate() is called.
func (b *CallBuilder) IsInterpolated() bool {
	return b.isInterpolated
}

// SetIsInterpolated sets whether this builder should interpolate.
func (b *CallBuilder) SetIsInterpolated(enable bool) *CallBuilder {
	b.isInterpolated = enable
	return b
}

// CanJSON determines if a builder can output JSON.
func (b *CallBuilder) CanJSON() bool {

	return false

}

// Interpolate interpolates this builders sql.
func (b *DeleteBuilder) Interpolate() (string, []interface{}, error) {
	return interpolate(b)
}

// IsInterpolated determines if this builder will interpolate when
// Interpolate() is called.
func (b *DeleteBuilder) IsInterpolated() bool {
	return b.isInterpolated
}

// SetIsInterpolated sets whether this builder should interpolate.
func (b *DeleteBuilder) SetIsInterpolated(enable bool) *DeleteBuilder {
	b.isInterpolated = enable
	return b
}

// CanJSON determines if a builder can output JSON.
func (b *DeleteBuilder) CanJSON() bool {

	return false

}

// Interpolate interpolates this builders sql.
func (b *InsectBuilder) Interpolate() (string, []interface{}, error) {
	return interpolate(b)
}

// IsInterpolated determines if this builder will interpolate when
// Interpolate() is called.
func (b *InsectBuilder) IsInterpolated() bool {
	return b.isInterpolated
}

// SetIsInterpolated sets whether this builder should interpolate.
func (b *InsectBuilder) SetIsInterpolated(enable bool) *InsectBuilder {
	b.isInterpolated = enable
	return b
}

// CanJSON determines if a builder can output JSON.
func (b *InsectBuilder) CanJSON() bool {

	return false

}

// Interpolate interpolates this builders sql.
func (b *InsertBuilder) Interpolate() (string, []interface{}, error) {
	return interpolate(b)
}

// IsInterpolated determines if this builder will interpolate when
// Interpolate() is called.
func (b *InsertBuilder) IsInterpolated() bool {
	return b.isInterpolated
}

// SetIsInterpolated sets whether this builder should interpolate.
func (b *InsertBuilder) SetIsInterpolated(enable bool) *InsertBuilder {
	b.isInterpolated = enable
	return b
}

// CanJSON determines if a builder can output JSON.
func (b *InsertBuilder) CanJSON() bool {

	return false

}

// Interpolate interpolates this builders sql.
func (b *JSQLBuilder) Interpolate() (string, []interface{}, error) {
	return interpolate(b)
}

// IsInterpolated determines if this builder will interpolate when
// Interpolate() is called.
func (b *JSQLBuilder) IsInterpolated() bool {
	return b.isInterpolated
}

// SetIsInterpolated sets whether this builder should interpolate.
func (b *JSQLBuilder) SetIsInterpolated(enable bool) *JSQLBuilder {
	b.isInterpolated = enable
	return b
}

// CanJSON determines if a builder can output JSON.
func (b *JSQLBuilder) CanJSON() bool {

	return true

}

// Interpolate interpolates this builders sql.
func (b *RawBuilder) Interpolate() (string, []interface{}, error) {
	return interpolate(b)
}

// IsInterpolated determines if this builder will interpolate when
// Interpolate() is called.
func (b *RawBuilder) IsInterpolated() bool {
	return b.isInterpolated
}

// SetIsInterpolated sets whether this builder should interpolate.
func (b *RawBuilder) SetIsInterpolated(enable bool) *RawBuilder {
	b.isInterpolated = enable
	return b
}

// CanJSON determines if a builder can output JSON.
func (b *RawBuilder) CanJSON() bool {

	return false

}

// Interpolate interpolates this builders sql.
func (b *SelectBuilder) Interpolate() (string, []interface{}, error) {
	return interpolate(b)
}

// IsInterpolated determines if this builder will interpolate when
// Interpolate() is called.
func (b *SelectBuilder) IsInterpolated() bool {
	return b.isInterpolated
}

// SetIsInterpolated sets whether this builder should interpolate.
func (b *SelectBuilder) SetIsInterpolated(enable bool) *SelectBuilder {
	b.isInterpolated = enable
	return b
}

// CanJSON determines if a builder can output JSON.
func (b *SelectBuilder) CanJSON() bool {

	return false

}

// Interpolate interpolates this builders sql.
func (b *SelectDocBuilder) Interpolate() (string, []interface{}, error) {
	return interpolate(b)
}

// IsInterpolated determines if this builder will interpolate when
// Interpolate() is called.
func (b *SelectDocBuilder) IsInterpolated() bool {
	return b.isInterpolated
}

// SetIsInterpolated sets whether this builder should interpolate.
func (b *SelectDocBuilder) SetIsInterpolated(enable bool) *SelectDocBuilder {
	b.isInterpolated = enable
	return b
}

// CanJSON determines if a builder can output JSON.
func (b *SelectDocBuilder) CanJSON() bool {

	return true

}

// Interpolate interpolates this builders sql.
func (b *UpdateBuilder) Interpolate() (string, []interface{}, error) {
	return interpolate(b)
}

// IsInterpolated determines if this builder will interpolate when
// Interpolate() is called.
func (b *UpdateBuilder) IsInterpolated() bool {
	return b.isInterpolated
}

// SetIsInterpolated sets whether this builder should interpolate.
func (b *UpdateBuilder) SetIsInterpolated(enable bool) *UpdateBuilder {
	b.isInterpolated = enable
	return b
}

// CanJSON determines if a builder can output JSON.
func (b *UpdateBuilder) CanJSON() bool {

	return false

}

// Interpolate interpolates this builders sql.
func (b *UpsertBuilder) Interpolate() (string, []interface{}, error) {
	return interpolate(b)
}

// IsInterpolated determines if this builder will interpolate when
// Interpolate() is called.
func (b *UpsertBuilder) IsInterpolated() bool {
	return b.isInterpolated
}

// SetIsInterpolated sets whether this builder should interpolate.
func (b *UpsertBuilder) SetIsInterpolated(enable bool) *UpsertBuilder {
	b.isInterpolated = enable
	return b
}

// CanJSON determines if a builder can output JSON.
func (b *UpsertBuilder) CanJSON() bool {

	return false

}
