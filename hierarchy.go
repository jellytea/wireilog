// Copyright 2023-2024 JetERA Creative
// This Source Code Form is subject to the terms of the Mozilla Public License, v. 2.0
// that can be found in the LICENSE file and https://mozilla.org/MPL/2.0/.

package wireilog

const (
	_ = iota

	PASSIVE

	POWER

	INPUT

	PUSH_PULL
	OPEN_DRAIN

	TRI_STATE

	ANALOG_INPUT
	ANALOG_OUTPUT
)

type Pin struct {
	Kind int
	Name string
}

type Module struct {
	Nets [][][2]int // [0]: Module [1]: Pin
	Port []Pin
	Mods []*Module
}
