# reference-addon-test-harness

This is a test harness meant for testing the reference addon in osde2e. It does the following:

* Tests for the existence of the relevant Addon CR and the reference-addon deployment. These should be present if the reference
  addon has been installed properly.
* Writes out a junit XML file with tests results to the /test-run-results directory as expected
  by the [https://github.com/openshift/osde2e](osde2e) test framework.
* Writes out an `addon-metadata.json` file which will also be consumed by the osde2e test framework.
