// Copyright (c) 2022 Gitpod GmbH. All rights reserved.
// Licensed under the GNU Affero General Public License (AGPL).
// See License-AGPL.txt in the project root for license information.

package db_test

import (
	"fmt"
	"github.com/gitpod-io/gitpod/usage/pkg/db"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestRead(t *testing.T) {
	conn := db.ConnectForTests(t)

	var ws db.Workspace
	conn.First(&ws)

	require.Contains(t, ws.CloneURL, "github.com")
}

func TestInsert(t *testing.T) {
	conn := db.ConnectForTests(t)

	var ws db.Workspace
	tx := conn.First(&ws)
	require.NoError(t, tx.Error)
	fmt.Println(ws)
}
