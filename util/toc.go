/*
   Copyright The starlight Authors.

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.

   file created by maverick in 2021
*/

package util

import (
	"fmt"
	"github.com/mc256/stargz-snapshotter/estargz"
	"github.com/opencontainers/go-digest"
)

type TraceableEntry struct {
	*estargz.TOCEntry

	Landmark int `json:"lm,omitempty"`

	// We need this otherwise source layer wont show in the toc json

	// Source starts from 1 not 0.
	// index 0 and -1 are reserved for special purpose.
	Source int `json:"s,omitempty"`

	// ConsolidatedSource starts from 1 not 0.
	// index 0 and -1 are reserved for special purpose.
	ConsolidatedSource int `json:"cs,omitempty"`

	Chunks      []*estargz.TOCEntry `json:"chunks,omitempty"`
	DeltaOffset *[]int64            `json:"df,omitempty"`

	UpdateMeta int `json:"md,omitempty"`
}

func GetRootNode() *TraceableEntry {
	return &TraceableEntry{
		TOCEntry: &estargz.TOCEntry{
			Name:    ".",
			Type:    "dir",
			Mode:    0755,
			NumLink: 1,
		},
		Chunks: nil,
	}
}

//  DeepCopy creates a deep copy of the object and clears the source layer identifier
//  You must assign a new source layer
func (t *TraceableEntry) DeepCopy() (d *TraceableEntry) {
	d = &TraceableEntry{
		TOCEntry: t.CopyEntry(),
		Landmark: t.Landmark,
		Source:   t.TOCEntry.GetSourceLayer(),
		Chunks:   t.Chunks,
	}
	return
}

//  ExtendEntry creates a deep copy of the t object and clears the source layer identifier
//  You must assign a new source layer
func ExtendEntry(t *estargz.TOCEntry) (d *TraceableEntry) {
	d = &TraceableEntry{
		TOCEntry: t.CopyEntry(),
		Landmark: 0,
		Source:   t.GetSourceLayer(),
		Chunks:   nil,
	}
	return
}

func (t *TraceableEntry) ShiftSource(offset int) {
	t.SetSourceLayer(t.TOCEntry.GetSourceLayer() + offset)
}

// SetSourceLayer sets the index of source layer. index should always starts from 1 if
// the entry comes from an actual layer.
func (t *TraceableEntry) SetSourceLayer(d int) {
	t.TOCEntry.SetSourceLayer(d)
	t.Source = d
}

func (t *TraceableEntry) GetSourceLayer() int {
	t.Source = t.TOCEntry.GetSourceLayer()
	return t.TOCEntry.GetSourceLayer()
}

// SetDeltaOffset sets the offset in the image body.
// If offset is zero, it means no changes were made to the file and the client will do nothing.
func (t *TraceableEntry) SetDeltaOffset(offsets *[]int64) {
	t.DeltaOffset = offsets
}

type TraceableBlobDigest struct {
	digest.Digest `json:"hash"`
	ImageName     string `json:"img"`
}

func (d TraceableBlobDigest) String() string {
	return fmt.Sprintf("%s-%s", d.ImageName, d.Digest.String())
}