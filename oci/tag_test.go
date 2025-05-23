/*
Copyright 2022 The Flux authors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package oci

import (
	"context"
	"fmt"
	"testing"

	"github.com/google/go-containerregistry/pkg/crane"
	"github.com/google/go-containerregistry/pkg/v1/random"
	. "github.com/onsi/gomega"
)

func Test_Tag(t *testing.T) {
	g := NewWithT(t)
	ctx := context.Background()
	c := NewClient(DefaultOptions())
	testRepo := "test-tag"
	url := fmt.Sprintf("%s/%s:v0.0.1", dockerReg, testRepo)
	img, err := random.Image(1024, 1)
	g.Expect(err).ToNot(HaveOccurred())
	err = crane.Push(img, url, c.options...)
	g.Expect(err).ToNot(HaveOccurred())

	_, err = c.Tag(ctx, url, "v0.0.2")
	g.Expect(err).ToNot(HaveOccurred())

	tags, err := crane.ListTags(fmt.Sprintf("%s/%s", dockerReg, testRepo))
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(len(tags)).To(BeEquivalentTo(2))
	g.Expect(tags).To(BeEquivalentTo([]string{"v0.0.1", "v0.0.2"}))
}
