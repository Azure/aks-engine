// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

//go:generate mockgen -destination client_mock.go --package mock_kubernetes --source ../interfaces.go Client
//go:generate /usr/bin/env bash -c "cat ../../../scripts/copyright.txt client_mock.go > _client_mock.go && mv _client_mock.go client_mock.go"

package mock_kubernetes //nolint
