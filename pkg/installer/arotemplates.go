package installer

// Copyright (c) Microsoft Corporation.
// Licensed under the Apache License 2.0.

import (
	"context"
	"os"
	"path/filepath"

	"github.com/openshift/installer/pkg/asset"
	"github.com/pkg/errors"
)

const (
	aroTemplatetDir = "manifests"
	rootPath        = "/opt/openshift"
)

// Custom ARO asset to add custom manifests to install graph in installer-wrapper similar to installer's manifests.Manifests
type AROTemplates struct {
	FileList []*asset.File
}

// ARO File Fetcher to read manifests
type aroFileFetcher struct {
	directory string
}

var (
	_ asset.WritableAsset = (*AROTemplates)(nil)
	_ asset.FileFetcher   = (*aroFileFetcher)(nil)
)

func (am *AROTemplates) Name() string {
	return "ARO Templates"
}

func (am *AROTemplates) Dependencies() []asset.Asset {
	return []asset.Asset{}
}

func (am *AROTemplates) Generate(ctx context.Context, dependencies asset.Parents) error {
	return nil
}

func (am *AROTemplates) Files() []*asset.File {
	return am.FileList
}

func (am *AROTemplates) Load(f asset.FileFetcher) (found bool, err error) {
	templateFileList, err := f.FetchByPattern(filepath.Join(aroManifestDir, "*.aro"))
	if err != nil {
		return false, errors.Wrap(err, "failed to load *.aro files")
	}
	templateFileList, err := f.FetchByPattern(filepath.Join(aroManifestDir, "*.aro"))
	if err != nil {
		return false, errors.Wrap(err, "failed to load *.aro files")
	}

	am.FileList = append(am.FileList, templateFileList...)
	am.FileList = append(am.FileList, templateFileList...)
	asset.SortFiles(am.FileList)

	return len(am.FileList) > 0, nil
}

func (f *aroFileFetcher) FetchByName(name string) (*asset.File, error) {
	data, err := os.ReadFile(filepath.Join(f.directory, name))
	if err != nil {
		return nil, err
	}
	return &asset.File{Filename: name, Data: data}, nil
}

func (f *aroFileFetcher) FetchByPattern(pattern string) (files []*asset.File, err error) {
	matches, err := filepath.Glob(filepath.Join(f.directory, pattern))
	if err != nil {
		return nil, err
	}

	files = make([]*asset.File, 0, len(matches))
	for _, path := range matches {
		data, err := os.ReadFile(path)
		if err != nil {
			return nil, err
		}
		files = append(files, &asset.File{
			Filename: path,
			Data:     data,
		})
	}

	return files, nil
}
