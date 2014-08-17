package apps

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gexec"

	"github.com/cloudfoundry-community/staticfiles-buildpack-acceptance-tests/helpers"
	"github.com/cloudfoundry-incubator/cf-test-helpers/cf"
	"github.com/cloudfoundry-incubator/cf-test-helpers/generator"
)

var _ = Describe("Buildpacks", func() {
	var appName string

	BeforeEach(func() {
		appName = generator.RandomName()
	})

	AfterEach(func() {
		Expect(cf.Cf("delete", appName, "-f").Wait(DEFAULT_TIMEOUT)).To(Exit(0))
	})

	Describe("normal", func() {
		It("successfully stages and runs", func() {
			Expect(cf.Cf("push", appName, "-p", helpers.NewAssets().Normal, "-c", "node app.js").Wait(CF_PUSH_TIMEOUT)).To(Exit(0))

			Eventually(func() string {
				return helpers.CurlAppRoot(appName)
			}, DEFAULT_TIMEOUT).Should(ContainSubstring("Hello from a node app!"))
		})
	})

})
