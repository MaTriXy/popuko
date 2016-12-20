package main

import (
	"testing"
)

func TestOwnersFileToRepoInfo1(t *testing.T) {
	o := OwnersFile{
		EnableAutoMerge:      true,
		DeleteAfterAutoMerge: true,
	}

	ok, info := o.ToRepoInfo()
	if !ok {
		t.Fatal("should be success to convert from OwnersFile")
	}

	if !info.EnableAutoMerge {
		t.Fatal("ShouldMergeAutomatically: should be true")
	}

	if !info.DeleteAfterAutoMerge {
		t.Fatal("ShouldDeleteMerged: should be true")
	}
}

func TestOwnersFileToRepoInfo2(t *testing.T) {
	o := OwnersFile{
		EnableAutoMerge:      false,
		DeleteAfterAutoMerge: false,
	}

	ok, info := o.ToRepoInfo()
	if !ok {
		t.Fatal("should be success to convert from OwnersFile")
	}

	if info.EnableAutoMerge {
		t.Fatal("ShouldMergeAutomatically: should be false")
	}

	if info.DeleteAfterAutoMerge {
		t.Fatal("ShouldDeleteMerged: should be false")
	}
}