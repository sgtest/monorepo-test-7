/*
Copyright 2015 The Kubernetes Authors.

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

package master

// These imports are the API groups the API server will support.
import (
	"fmt"

	"github.com/sourcegraph/monorepo-test-1/kubernetes-5/pkg/api"
	_ "github.com/sourcegraph/monorepo-test-1/kubernetes-5/pkg/api/install"
	_ "github.com/sourcegraph/monorepo-test-1/kubernetes-5/pkg/apis/apps/install"
	_ "github.com/sourcegraph/monorepo-test-1/kubernetes-5/pkg/apis/authentication/install"
	_ "github.com/sourcegraph/monorepo-test-1/kubernetes-5/pkg/apis/authorization/install"
	_ "github.com/sourcegraph/monorepo-test-1/kubernetes-5/pkg/apis/autoscaling/install"
	_ "github.com/sourcegraph/monorepo-test-1/kubernetes-5/pkg/apis/batch/install"
	_ "github.com/sourcegraph/monorepo-test-1/kubernetes-5/pkg/apis/certificates/install"
	_ "github.com/sourcegraph/monorepo-test-1/kubernetes-5/pkg/apis/componentconfig/install"
	_ "github.com/sourcegraph/monorepo-test-1/kubernetes-5/pkg/apis/extensions/install"
	_ "github.com/sourcegraph/monorepo-test-1/kubernetes-5/pkg/apis/imagepolicy/install"
	_ "github.com/sourcegraph/monorepo-test-1/kubernetes-5/pkg/apis/policy/install"
	_ "github.com/sourcegraph/monorepo-test-1/kubernetes-5/pkg/apis/rbac/install"
	_ "github.com/sourcegraph/monorepo-test-1/kubernetes-5/pkg/apis/settings/install"
	_ "github.com/sourcegraph/monorepo-test-1/kubernetes-5/pkg/apis/storage/install"
)

func init() {
	if missingVersions := api.Registry.ValidateEnvRequestedVersions(); len(missingVersions) != 0 {
		panic(fmt.Sprintf("KUBE_API_VERSIONS contains versions that are not installed: %q.", missingVersions))
	}
}