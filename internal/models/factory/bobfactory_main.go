// Code generated by BobGen psql v0.26.1. DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package factory

type Factory struct {
	baseExampleMods ExampleModSlice
}

func New() *Factory {
	return &Factory{}
}

func (f *Factory) NewExample(mods ...ExampleMod) *ExampleTemplate {
	o := &ExampleTemplate{f: f}

	if f != nil {
		f.baseExampleMods.Apply(o)
	}

	ExampleModSlice(mods).Apply(o)

	return o
}

func (f *Factory) ClearBaseExampleMods() {
	f.baseExampleMods = nil
}

func (f *Factory) AddBaseExampleMod(mods ...ExampleMod) {
	f.baseExampleMods = append(f.baseExampleMods, mods...)
}
