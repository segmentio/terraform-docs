// Package tfconfig is a helper library that does careful, shallow parsing of
// Terraform modules to provide access to high-level metadata while
// remaining broadly compatible with configurations targeting various
// different Terraform versions.
//
// This packge focuses on describing top-level objects only, and in particular
// does not attempt any sort of processing that would require access to plugins.
// Currently it allows callers to extract high-level information about
// variables, outputs, resource blocks, provider dependencies, and Terraform
// Core dependencies.
//
// This package only works at the level of single modules. A full configuration
// is a tree of potentially several modules, some of which may be references
// to remote packages. There are some basic helpers for traversing calls to
// modules at relative local paths, however.
//
// This package employs a "best effort" parsing strategy, producing as complete
// a result as possible even though the input may not be entirely valid. The
// intended use-case is high-level analysis and indexing of externally-facing
// module characteristics, as opposed to validating or even applying the module.
//
// Important Note: This is in-project fork of https://github.com/hashicorp/terraform-config-inspect
// and by no means should be used outside of this project. The reason for bringing
// the project over internally was that terraform-config-inspect is considered
// feature-complete and new enhancement pull requests are not accepted and we've
// decided to bring the project over internally and apply one change that we
// really depended on in terraform-docs.
package tfconfig
