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
	"sort"
	"strings"
)

type ImageRef struct {
	ImageName string `json:"n"`
	ImageTag  string `json:"t"`
}

func (i ImageRef) String() string {
	return i.ImageName + ":" + i.ImageTag
}

type ByImageName []*ImageRef

func (b ByImageName) Len() int {
	return len(b)
}

func (b ByImageName) Less(i, j int) bool {
	if b[i].ImageName != b[j].ImageName {
		return b[i].ImageName < b[j].ImageName
	} else {
		return b[i].ImageTag < b[j].ImageTag
	}
}

func (b ByImageName) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

func (b ByImageName) String() string {
	var arr []string
	for _, v := range b {
		arr = append(arr, v.String())
	}
	return strings.Join(arr, ",")
}

func SortImageArrayString(images string) string {
	var arrRef []*ImageRef
	if images == "" {
		return ""
	}
	for _, v := range strings.Split(images, ",") {
		arr := strings.SplitN(v, ":", 2)
		arrRef = append(arrRef, &ImageRef{
			ImageName: arr[0],
			ImageTag:  arr[1],
		})
	}
	sort.Sort(ByImageName(arrRef))
	return ByImageName(arrRef).String()
}
