package tests

import (
	"path/filepath"
	"testing"

	"github.com/mt-sre/reference-addon-test-harness/pkg/metadata"
	. "github.com/onsi/ginkgo"
	"github.com/onsi/ginkgo/reporters"
	. "github.com/onsi/gomega"
)

const (
	testResultsDirectory = "/test-run-results"
	jUnitOutputFilename  = "junit-reference-addon.xml"
	addonMetadataName    = "addon-metadata.json"
)

func TestAddonInstallation(t *testing.T) {
	RegisterFailHandler(Fail)
	jUnitReporter := reporters.NewJUnitReporter(filepath.Join(testResultsDirectory, jUnitOutputFilename))

	RunSpecsWithDefaultAndCustomReporters(t, "Reference Addon testharness", []Reporter{jUnitReporter})
	err := metadata.Instance.WriteToJSON(filepath.Join(testResultsDirectory, addonMetadataName))
	if err != nil {
		t.Errorf("error while writing metadata: %v", err)
	}
}
