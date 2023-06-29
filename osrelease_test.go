// Copyright 2021 Harald Albrecht.
//
// Licensed under the Apache License, Version 2.0 (the "License"); you may not
// use this file except in compliance with the License. You may obtain a copy
// of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
// License for the specific language governing permissions and limitations
// under the License.

package osrelease

import (
	"errors"
	"io"
	"os"
	"strings"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	. "github.com/thediveo/success"
)

var _ = Describe("os-release", func() {

	It("finds system os-release information", func() {
		_, err1 := os.Stat("/etc/os-release")
		_, err2 := os.Stat("/usr/lib/os-release")
		if os.IsNotExist(err1) && os.IsNotExist(err2) {
			Skip("this system does not use os-release")
		}
		vars := New()
		Expect(vars).NotTo(BeNil())
		Expect(vars).NotTo(BeEmpty())
		Expect(vars).To(HaveKey("ID"))
	})

	It("finds os-release information", func() {
		vars := new("./test/etc")
		Expect(vars).To(HaveLen(1))
		Expect(vars).To(HaveKeyWithValue("OSRELEASE", "/etc/os-release"))

		vars = new("./test/usr")
		Expect(vars).To(HaveLen(1))
		Expect(vars).To(HaveKeyWithValue("OSRELEASE", "/usr/lib/os-release"))
	})

	It("parses assignments", func() {
		cases := []struct {
			in       string
			variable string
			value    string
		}{
			// GOOD
			{in: ``, variable: "", value: ""},
			{in: `# comment`, variable: "", value: ""},
			{in: `A=B`, variable: "A", value: "B"},
			{in: `A=""`, variable: "A", value: ""},
			{in: `A=`, variable: "A", value: ""},
			{in: `A="A\$BC"`, variable: "A", value: "A$BC"},
			{in: `A=A\$BC`, variable: "A", value: `A\$BC`}, // sic!
			// BAD
			{in: ` #`, variable: "", value: ""},
			{in: `FOO="BAR`, variable: "", value: ""},
		}
		for _, testcase := range cases {
			variable, value := assignment(testcase.in)
			Expect(variable).To(Equal(testcase.variable), "variable of case %q", testcase.in)
			Expect(value).To(Equal(testcase.value), "variable of case %q", testcase.in)
		}
	})

	It("reads assignments from a reader", func() {
		vars := Successful(assignmentsFromReader(strings.NewReader(`
# This is a test.
FOO=Bar
BAR="Baz"`)))
		Expect(vars).NotTo(BeNil())
		Expect(vars).To(HaveLen(2))
		Expect(vars).To(And(
			HaveKeyWithValue("FOO", "Bar"),
			HaveKeyWithValue("BAR", "Baz"),
		))
	})

	It("returns nil map when reader fails", func() {
		Expect(
			assignmentsFromReader(ErrReader(errors.New("DOH!"))),
		).Error().To(HaveOccurred())
	})

})

// ErrReader returns an io.Reader that returns 0, err from all Read calls. As
// ErrReader doesn't exist in older Go versions, we are bringing it here in
// ourselves.
func ErrReader(err error) io.Reader {
	return &errReader{err: err}
}

type errReader struct {
	err error
}

// Read always return 0, err for all calls.
func (r *errReader) Read(p []byte) (int, error) {
	return 0, r.err
}
