// SPDX-License-Identifier: LGPL-3.0-or-later
// Copyright 2019 DNA Dev team
//
/*
 * Copyright (C) 2018 The ontology Authors
 * This file is part of The ontology library.
 *
 * The ontology is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Lesser General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * The ontology is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU Lesser General Public License for more details.
 *
 * You should have received a copy of the GNU Lesser General Public License
 * along with The ontology.  If not, see <http://www.gnu.org/licenses/>.
 */
package payload

import (
	"testing"

	"github.com/ontio/ontology/common"
	"github.com/stretchr/testify/assert"
)

func TestDeployCode_Serialize(t *testing.T) {
	ty, _ := VmTypeFromByte(1)
	deploy, err := NewDeployCode([]byte{1, 2, 3}, ty, "", "", "", "", "")
	assert.Nil(t, err)
	sink := common.NewZeroCopySink(nil)
	deploy.Serialization(sink)
	bs := sink.Bytes()
	var deploy2 DeployCode

	source := common.NewZeroCopySource(bs)
	deploy2.Deserialization(source)
	assert.Equal(t, &deploy2, deploy)

	source = common.NewZeroCopySource(bs[:len(bs)-1])
	err = deploy2.Deserialization(source)
	assert.NotNil(t, err)
}
