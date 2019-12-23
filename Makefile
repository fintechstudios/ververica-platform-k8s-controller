# Image URL to use all building/pushing image targets
VERSION?=v0.6.0
IMG?=fintechstudios/ververica-platform-k8s-operator
PKG=github.com/fintechstudios.com/ververica-platform-k8s-operator
VERSION_PKG=main
BUILD=$(shell date -u +'%Y-%m-%dT%H:%M:%SZ')
# Produce CRDs that work back to Kubernetes 1.11 (no version conversion)
CRD_OPTIONS?="crd:trivialVersions=true"

LD_FLAGS="-X $(VERSION_PKG).operatorVersion='$(VERSION)' -X $(VERSION_PKG).gitCommit='$(GIT_COMMIT)' -X $(VERSION_PKG).buildDate='$(BUILD)'"

TEST_CLUSTER_NAME=ververica-platform-k8s-operator-cluster

# Get the currently used golang install path (in GOPATH/bin, unless GOBIN is set)
ifeq (,$(shell go env GOBIN))
GOBIN=$(shell go env GOPATH)/bin
else
GOBIN=$(shell go env GOBIN)
endif

all: manager

# find or download controller-gen
.PHONY: controller-gen
controller-gen:
ifeq (, $(shell which controller-gen))
	@{ \
	set -e ;\
	CONTROLLER_GEN_TMP_DIR=$$(mktemp -d) ;\
	cd $$CONTROLLER_GEN_TMP_DIR ;\
	go mod init tmp ;\
	go get sigs.k8s.io/controller-tools/cmd/controller-gen@v0.2.4 ;\
	rm -rf $$CONTROLLER_GEN_TMP_DIR ;\
	}
CONTROLLER_GEN=$(GOBIN)/controller-gen
else
CONTROLLER_GEN=$(shell which controller-gen)
endif

# find or download kustomize
.PHONY: kustomize
kustomize:
ifeq (, $(shell which kustomize))
	@{ \
	set -e ;\
	KUSTOMIZE_TMP_DIR=$$(mktemp -d) ;\
	cd $$KUSTOMIZE_TMP_DIR ;\
	go mod init tmp ;\
	go get sigs.k8s.io/kustomize/kustomize/v3@v3.3.0 ;\
	rm -rf $$KUSTOMIZE_TMP_DIR ;\
	}
KUSTOMIZE=$(GOBIN)/kustomize
else
KUSTOMIZE=$(shell which kustomize)
endif

# Run tests
.PHONY: test
test: generate manifests
	go test -ldflags $(LD_FLAGS) ./api/... ./controllers/... ./ -coverprofile cover.out

# Build manager binary
.PHONY: manager
manager: generate
	go build $(ARGS) -ldflags $(LD_FLAGS) -o bin/manager main.go

# Run against the configured Kubernetes cluster in ~/.kube/config
.PHONY: run
run: generate
	go run -ldflags $(LD_FLAGS) ./main.go

# Install CRDs into a cluster
.PHONY: install
install: manifests
	$(KUSTOMIZE) build config/crd | kubectl apply -f -

# Uninstall CRDs from a cluster
uninstall: manifests
	$(KUSTOMIZE) build config/crd | kubectl delete -f -

# Deploy controller in the configured Kubernetes cluster in ~/.kube/config
.PHONY: deploy
deploy: manifests
	cd config/manager && $(KUSTOMIZE) edit set image controller=${IMG}
	$(KUSTOMIZE) build config/default | kubectl apply -f -

# Generate manifests e.g. CRD, RBAC etc.
.PHONY: manifests
manifests: controller-gen
	$(CONTROLLER_GEN) $(CRD_OPTIONS) rbac:roleName=manager-role webhook paths="./..." output:crd:artifacts:config=config/crd/bases

# Run gofmt against non-generated code
.PHONY: fmt
fmt:
	gofmt -s -w ./api ./controllers *.go

.PHONY: lint
lint:
	golangci-lint run --timeout=120s --verbose

# Generate code
.PHONY: generate
generate: controller-gen
	$(CONTROLLER_GEN) object:headerFile=./hack/boilerplate.go.txt paths=./api/...

# Patch the latest image version into the default kustomize image patch
.PHONY: patch-image
patch-image:
	sed -i'' -e 's@image: .*@image: '"$(IMG):$(VERSION)"'@' ./config/default/manager_image_patch.yaml

# Build the k8s resources for deployment
kustomize-build: patch-image
	$(KUSTOMIZE) build config/default > resources.yaml

# Update the Swagger Client API
.PHONY: swagger-gen
swagger-gen:
	./hack/update-app-manager-swagger-codegen.sh \
	 && ./hack/update-platform-swagger-codegen.sh

# Create the test cluster using kind
# install local path storage as defult storage class (see: https://github.com/kubernetes-sigs/kind/issues/118#issuecomment-475134086)
.PHONY: test-cluster-create
test-cluster-create:
	kind create cluster --name $(TEST_CLUSTER_NAME) \
		&& sleep 10 \
		&& kubectl apply -f https://raw.githubusercontent.com/rancher/local-path-provisioner/master/deploy/local-path-storage.yaml \
		&& kubectl patch storageclass standard -p '{"metadata": {"annotations":{"storageclass.kubernetes.io/is-default-class":"false", "storageclass.beta.kubernetes.io/is-default-class":"false"}}}' \
		&& kubectl patch storageclass local-path -p '{"metadata": {"annotations":{"storageclass.kubernetes.io/is-default-class":"true", "storageclass.beta.kubernetes.io/is-default-class":"true"}}}'

# Delete the test cluster using kind
.PHONY: test-cluster-delete
test-cluster-delete:
	kind delete cluster --name $(TEST_CLUSTER_NAME)

