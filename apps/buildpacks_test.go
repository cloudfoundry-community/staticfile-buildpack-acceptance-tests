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
			Expect(cf.Cf("push", appName, "-p", helpers.NewAssets().Normal).Wait(CF_PUSH_TIMEOUT)).To(Exit(0))

			Eventually(func() string {
				return helpers.CurlAppRoot(appName)
			}, DEFAULT_TIMEOUT).Should(ContainSubstring("This is an example app for Cloud Foundry that is only static HTML/JS/CSS assets"))
		})
	})

	Describe("non_staticfile", func() {
		It("fails to stage", func() {
			Expect(cf.Cf("push", appName, "-p", helpers.NewAssets().NonStaticfile).Wait(CF_PUSH_TIMEOUT)).To(Exit(1))
		})
		// An app was not successfully detected by any available buildpack
	})

	Describe("alternate root", func() {
		It("successfully stages and runs", func() {
			Expect(cf.Cf("push", appName, "-p", helpers.NewAssets().AlternateRoot).Wait(CF_PUSH_TIMEOUT)).To(Exit(0))

			Eventually(func() string {
				return helpers.CurlAppRoot(appName)
			}, DEFAULT_TIMEOUT).Should(ContainSubstring("This index file comes from an alternate root"))
		})
	})

	Describe("basic auth", func() {
		It("successfully stages and runs", func() {
			Expect(cf.Cf("push", appName, "-p", helpers.NewAssets().BasicAuth).Wait(CF_PUSH_TIMEOUT)).To(Exit(0))

			Eventually(func() string {
				return helpers.CurlAppRoot(appName)
			}, DEFAULT_TIMEOUT).Should(ContainSubstring("401 Authorization Required"))

			Eventually(func() string {
				return helpers.CurlAppRootWithAuth(appName, "bob", "bob")
			}, DEFAULT_TIMEOUT).Should(ContainSubstring("This site is protected by basic auth."))
		})
	})

	Describe("directory index", func() {
		It("successfully stages and runs", func() {
			Expect(cf.Cf("push", appName, "-p", helpers.NewAssets().DirectoryIndex).Wait(CF_PUSH_TIMEOUT)).To(Exit(0))

			Eventually(func() string {
				return helpers.CurlAppRoot(appName)
			}, DEFAULT_TIMEOUT).Should(ContainSubstring("find-me.html"))
		})
	})
})
