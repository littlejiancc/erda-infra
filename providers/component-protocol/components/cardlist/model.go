// Copyright (c) 2021 Terminus, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cardlist

import (
	model "github.com/erda-project/erda-infra/providers/component-protocol/components/cardlist/models"
	"github.com/erda-project/erda-infra/providers/component-protocol/cptype"
)

type (
	// Data card std data
	Data struct {
		Total        int    `json:"total"`
		Title        string `json:"title"`
		TitleSummary string `json:"titleSummary"`
		Cards        []Card `json:"cards,omitempty"`
	}
	// Card .
	Card struct {
		ID       string           `json:"id"`
		ImgURL   string           `json:"imgURL"`
		Icon     string           `json:"icon"`
		Title    string           `json:"title"`
		Labels   []Label          `json:"labels"`
		Star     bool             `json:"star"`
		TextMeta []model.TextMeta `json:"textMeta"`
		cptype.Extra
	}
	// Label .
	Label struct {
		Label string `json:"label"`
	}
)