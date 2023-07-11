package storage

// Copyright (c) Microsoft Corporation.
// Licensed under the Apache License 2.0.

//go:generate rm -rf ../mocks/$GOPACKAGE
//go:generate go run ../../../vendor/github.com/golang/mock/mockgen -destination=../mocks/$GOPACKAGE/$GOPACKAGE.go github.com/openshift/ARO-Installer/pkg/util/$GOPACKAGE Manager
//go:generate go run ../../../vendor/golang.org/x/tools/cmd/goimports -local=github.com/openshift/ARO-Installer -e -w ../mocks/$GOPACKAGE/$GOPACKAGE.go
