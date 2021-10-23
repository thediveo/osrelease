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
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("assignment value", func() {

	It("unquotes", func() {
		cases := []struct {
			in  string
			out string
			ok  bool
		}{
			// GOOD
			{in: ``, out: ``, ok: true},
			{in: `aöc`, out: `aöc`, ok: true},
			{in: `"aöc"`, out: `aöc`, ok: true},
			{in: `'abc'`, out: `abc`, ok: true},
			{in: `"a\"\$\"c"`, out: `a"$"c`, ok: true},
			{in: `"a\"ö\"c"`, out: `a"ö"c`, ok: true},
			{in: `'a\"\$\"c'`, out: `a"$"c`, ok: true},
			{in: `'a\\\$b'`, out: `a\$b`, ok: true},
			// BAD
			{in: `"\abc`, ok: false},
			{in: `"abc`, ok: false},
			{in: `"a\\bc`, ok: false},
			{in: `"ab"c`, ok: false},
			{in: `"abc'`, ok: false},
			{in: `'abc`, ok: false},
			{in: `'abc"`, ok: false},
		}
		for _, testcase := range cases {
			out, err := Unquote(testcase.in)
			if testcase.ok {
				Expect(err).NotTo(HaveOccurred(), "case %q", testcase.in)
			} else {
				Expect(err).To(HaveOccurred(), "case %q", testcase.in)
				continue
			}
			Expect(out).To(Equal(testcase.out), "case %q", testcase.in)
		}
	})

	It("unescapes a character", func() {
		cases := []struct {
			in   string
			r    rune
			mb   bool
			tail string
			ok   bool
		}{
			// GOOD
			{in: `abc`, r: 'a', mb: false, tail: "bc", ok: true},
			{in: `öh`, r: 'ö', mb: true, tail: "h", ok: true},
			{in: `\\abc`, r: '\\', mb: false, tail: "abc", ok: true},
			{in: `\"abc`, r: '"', mb: false, tail: "abc", ok: true},
			{in: `\'abc`, r: '\'', mb: false, tail: "abc", ok: true},
			{in: `\$abc`, r: '$', mb: false, tail: "abc", ok: true},
			{in: "\\`abc", r: '`', mb: false, tail: "abc", ok: true},
			// BAD
			{in: `\`, ok: false},
			{in: `\n`, ok: false},
		}
		for _, testcase := range cases {
			r, mb, tail, err := unescapeChar(testcase.in)
			if testcase.ok {
				Expect(err).NotTo(HaveOccurred(), "case %q", testcase.in)
			} else {
				Expect(err).To(HaveOccurred(), "case %q", testcase.in)
				continue
			}
			Expect(r).To(Equal(testcase.r), "rune of case %q", testcase.in)
			Expect(mb).To(Equal(testcase.mb), "multibyte of case %q", testcase.in)
			Expect(tail).To(Equal(testcase.tail), "tail of case %q", testcase.in)
		}
	})

})
