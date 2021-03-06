/*
 * Copyright 2018 It-chain
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package mem_test

import (
	"testing"

	"github.com/it-chain/engine/consensus/pbft"
	"github.com/it-chain/engine/consensus/pbft/infra/mem"
	"github.com/magiconair/properties/assert"
)

func TestParliamentRepository_Load_Save(t *testing.T) {
	p := mem.NewParliamentRepository()
	p.Save(pbft.Parliament{
		Leader: pbft.Leader{LeaderId: "123"},
	})

	parliament := p.Load()

	assert.Equal(t, parliament.Leader, pbft.Leader{LeaderId: "123"})
}
