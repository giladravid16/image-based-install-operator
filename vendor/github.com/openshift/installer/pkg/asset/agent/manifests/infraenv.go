package manifests

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/coreos/stream-metadata-go/arch"
	"github.com/pkg/errors"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/yaml"

	aiv1beta1 "github.com/openshift/assisted-service/api/v1beta1"
	"github.com/openshift/installer/pkg/asset"
	"github.com/openshift/installer/pkg/asset/agent"
	"github.com/openshift/installer/pkg/asset/agent/agentconfig"
	"github.com/openshift/installer/pkg/asset/agent/joiner"
	"github.com/openshift/installer/pkg/asset/agent/workflow"
	"github.com/openshift/installer/pkg/types"
)

var (
	infraEnvFilename = filepath.Join(clusterManifestDir, "infraenv.yaml")
)

// InfraEnv generates the infraenv.yaml file.
type InfraEnv struct {
	File   *asset.File
	Config *aiv1beta1.InfraEnv
}

var _ asset.WritableAsset = (*InfraEnv)(nil)

// Name returns a human friendly name for the asset.
func (*InfraEnv) Name() string {
	return "InfraEnv Config"
}

// Dependencies returns all of the dependencies directly needed to generate
// the asset.
func (*InfraEnv) Dependencies() []asset.Asset {
	return []asset.Asset{
		&workflow.AgentWorkflow{},
		&joiner.ClusterInfo{},
		&agent.OptionalInstallConfig{},
		&agentconfig.AgentConfig{},
	}
}

// Generate generates the InfraEnv manifest.
func (i *InfraEnv) Generate(_ context.Context, dependencies asset.Parents) error {
	agentWorkflow := &workflow.AgentWorkflow{}
	clusterInfo := &joiner.ClusterInfo{}
	installConfig := &agent.OptionalInstallConfig{}
	agentConfig := &agentconfig.AgentConfig{}
	dependencies.Get(installConfig, agentConfig, agentWorkflow, clusterInfo)

	switch agentWorkflow.Workflow {
	case workflow.AgentWorkflowTypeInstall:
		if installConfig.Config != nil {
			err := i.generateManifest(installConfig.ClusterName(), installConfig.ClusterNamespace(), installConfig.Config.SSHKey, installConfig.Config.AdditionalTrustBundle, installConfig.Config.Proxy, string(installConfig.Config.ControlPlane.Architecture))
			if err != nil {
				return err
			}

			if agentConfig.Config != nil {
				i.Config.Spec.AdditionalNTPSources = agentConfig.Config.AdditionalNTPSources
			}

		}

	case workflow.AgentWorkflowTypeAddNodes:
		err := i.generateManifest(clusterInfo.ClusterName, clusterInfo.Namespace, clusterInfo.SSHKey, clusterInfo.UserCaBundle, clusterInfo.Proxy, clusterInfo.Architecture)
		if err != nil {
			return err
		}

	default:
		return fmt.Errorf("AgentWorkflowType value not supported: %s", agentWorkflow.Workflow)
	}

	return i.finish()
}

func (i *InfraEnv) generateManifest(clusterName, clusterNamespace, sshKey, additionalTrustBundle string, proxy *types.Proxy, architecture string) error {
	infraEnv := &aiv1beta1.InfraEnv{
		TypeMeta: metav1.TypeMeta{
			Kind:       "InfraEnv",
			APIVersion: aiv1beta1.GroupVersion.String(),
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      clusterName,
			Namespace: clusterNamespace,
		},
		Spec: aiv1beta1.InfraEnvSpec{
			ClusterRef: &aiv1beta1.ClusterReference{
				Name:      clusterName,
				Namespace: clusterNamespace,
			},
			SSHAuthorizedKey: strings.Trim(sshKey, "|\n\t"),
			PullSecretRef: &corev1.LocalObjectReference{
				Name: getPullSecretName(clusterName),
			},
			NMStateConfigLabelSelector: metav1.LabelSelector{
				MatchLabels: getNMStateConfigLabels(clusterName),
			},
		},
	}

	// Input values (amd64, arm64) must be converted to rpmArch because infraEnv.Spec.CpuArchitecture expects x86_64 or aarch64.
	if architecture != "" {
		infraEnv.Spec.CpuArchitecture = arch.RpmArch(architecture)
	}
	if additionalTrustBundle != "" {
		infraEnv.Spec.AdditionalTrustBundle = additionalTrustBundle
	}
	if proxy != nil {
		infraEnv.Spec.Proxy = getProxy(proxy)
	}

	i.Config = infraEnv

	return nil
}

// Files returns the files generated by the asset.
func (i *InfraEnv) Files() []*asset.File {
	if i.File != nil {
		return []*asset.File{i.File}
	}
	return []*asset.File{}
}

// Load returns infraenv asset from the disk.
func (i *InfraEnv) Load(f asset.FileFetcher) (bool, error) {

	file, err := f.FetchByName(infraEnvFilename)
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}
		return false, errors.Wrap(err, fmt.Sprintf("failed to load %s file", infraEnvFilename))
	}

	config := &aiv1beta1.InfraEnv{}
	if err := yaml.UnmarshalStrict(file.Data, config); err != nil {
		return false, errors.Wrapf(err, "failed to unmarshal %s", infraEnvFilename)
	}
	// If defined, convert to RpmArch amd64 -> x86_64 or arm64 -> aarch64
	if config.Spec.CpuArchitecture != "" {
		config.Spec.CpuArchitecture = arch.RpmArch(config.Spec.CpuArchitecture)
	}
	i.File, i.Config = file, config
	if err = i.finish(); err != nil {
		return false, err
	}

	return true, nil
}

func (i *InfraEnv) finish() error {

	if i.Config == nil {
		return errors.New("missing configuration or manifest file")
	}

	infraEnvData, err := yaml.Marshal(i.Config)
	if err != nil {
		return errors.Wrap(err, "failed to marshal agent installer infraEnv")
	}

	i.File = &asset.File{
		Filename: infraEnvFilename,
		Data:     infraEnvData,
	}

	// Throw an error if CpuArchitecture isn't x86_64, aarch64, ppc64le, s390x, or ""
	switch i.Config.Spec.CpuArchitecture {
	case arch.RpmArch(types.ArchitectureAMD64), arch.RpmArch(types.ArchitectureARM64), arch.RpmArch(types.ArchitecturePPC64LE), arch.RpmArch(types.ArchitectureS390X), "":
	default:
		return errors.Errorf("Config.Spec.CpuArchitecture %s is not supported ", i.Config.Spec.CpuArchitecture)
	}
	return nil
}