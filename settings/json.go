/*  Copyright (c) 2013, Brian Hummer (brian@boggo.net)
All rights reserved.

Redistribution and use in source and binary forms, with or without
modification, are permitted provided that the following conditions are met:
    * Redistributions of source code must retain the above copyright
      notice, this list of conditions and the following disclaimer.
    * Redistributions in binary form must reproduce the above copyright
      notice, this list of conditions and the following disclaimer in the
      documentation and/or other materials provided with the distribution.
    * Neither the name of the boggo.net nor the
      names of its contributors may be used to endorse or promote products
      derived from this software without specific prior written permission.

THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS" AND
ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED
WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
DISCLAIMED. IN NO EVENT SHALL BRIAN HUMMER BE LIABLE FOR ANY
DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES
(INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES;
LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND
ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
(INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS
SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
*/

package settings

import (
	"encoding/json"
	"github.com/boggo/neat"
	"os"
)

type jsonSettings struct {
	path string
}

func NewJSON(path string) *jsonSettings {
	return &jsonSettings{path}
}

// Load the settings from a JSON file
func (x *jsonSettings) Load() (settings *neat.Settings, err error) {
	var f *os.File
	f, err = os.Open(x.path)
	if err != nil {
		return
	}
	defer f.Close()

	d := json.NewDecoder(f)
	settings = new(neat.Settings)
	err = d.Decode(settings)
	return
}

// Save the settings to a JSON file
func (x *jsonSettings) Save(settings *neat.Settings) (err error) {
	var f *os.File
	f, err = os.Create(x.path)
	if err != nil {
		return
	}
	defer f.Close()

	e := json.NewEncoder(f)
	err = e.Encode(settings)
	return
}
