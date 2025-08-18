// Vikunja is a to-do list application to facilitate your life.
// Copyright 2018-present Vikunja and contributors. All rights reserved.
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

package migration

import (
	"bytes"
	"context"
	"net/http"
	"net/url"
	"os"
	"strings"
)

// DownloadFile downloads a file and returns its contents
func DownloadFile(urlStr string) (buf *bytes.Buffer, err error) {
	return DownloadFileWithHeaders(urlStr, nil)
}

// DownloadFileWithHeaders downloads a file and allows you to pass in headers
func DownloadFileWithHeaders(urlStr string, headers http.Header) (buf *bytes.Buffer, err error) {
	// Support local file URLs in tests: file:///abs/path/to/file
	if parsed, errp := url.Parse(urlStr); errp == nil {
		if parsed.Scheme == "file" {
			// On file URLs, parsed.Path is the absolute path
			b, err := os.ReadFile(parsed.Path)
			if err != nil {
				return nil, err
			}
			buf = bytes.NewBuffer(b)
			return buf, nil
		}
	}

	req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, urlStr, nil)
	if err != nil {
		return nil, err
	}

	for key, h := range headers {
		for _, hh := range h {
			req.Header.Add(key, hh)
		}
	}

	hc := http.Client{}
	resp, err := hc.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	buf = &bytes.Buffer{}
	_, err = buf.ReadFrom(resp.Body)

	return
}

// DoPost makes a form encoded post request
func DoPost(urlStr string, form url.Values) (resp *http.Response, err error) {
	return DoPostWithHeaders(urlStr, form, map[string]string{})
}

// DoPostWithHeaders does an api request and allows to pass in arbitrary headers
func DoPostWithHeaders(urlStr string, form url.Values, headers map[string]string) (resp *http.Response, err error) {
	req, err := http.NewRequestWithContext(context.Background(), http.MethodPost, urlStr, strings.NewReader(form.Encode()))
	if err != nil {
		return
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	for key, value := range headers {
		req.Header.Add(key, value)
	}

	hc := http.Client{}
	return hc.Do(req)
}
