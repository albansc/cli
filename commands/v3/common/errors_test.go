package common_test

import (
	"bytes"
	"text/template"

	. "code.cloudfoundry.org/cli/commands/v3/common"
	"code.cloudfoundry.org/cli/utils/ui"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
)

var _ = Describe("Translatable Errors", func() {
	translateFunc := func(s string, vars ...interface{}) string {
		formattedTemplate, err := template.New("test template").Parse(s)
		Expect(err).ToNot(HaveOccurred())
		formattedTemplate.Option("missingkey=error")

		var buffer bytes.Buffer
		if len(vars) > 0 {
			err = formattedTemplate.Execute(&buffer, vars[0])
			Expect(err).ToNot(HaveOccurred())

			return buffer.String()
		} else {
			return s
		}
	}

	DescribeTable("translates error",
		func(e error) {
			err, ok := e.(ui.TranslatableError)
			Expect(ok).To(BeTrue())
			err.Translate(translateFunc)
		},

		// Command prerequisite errors.
		Entry("NoAPISetError", NoAPISetError{}),
		Entry("NotLoggedInError", NotLoggedInError{}),
		Entry("NoTargetedOrgError", NoTargetedOrgError{}),
		Entry("NoTargetedSpaceError", NoTargetedSpaceError{}),

		// Cloud controller errors.
		Entry("APIRequestError", APIRequestError{}),
		Entry("InvalidSSLCertError", InvalidSSLCertError{}),
		Entry("APINotFoundError", APINotFoundError{}),

		// Actor errors.
		Entry("ApplicationNotFoundError", ApplicationNotFoundError{}),
		Entry("RunTaskError", RunTaskError{}),
	)
})
