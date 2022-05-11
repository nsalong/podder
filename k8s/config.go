package k8s

import (
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"os"
	"path/filepath"
)

func InitConfig(context, path string) (*rest.Config, error) {
	if path == "" {
		userHomeDir, err := os.UserHomeDir()
		if err != nil {
			return nil, err
		}

		path = filepath.Join(userHomeDir, ".kube", "config")
	}

	kubeConfig, err := buildConfigFromFlags(context, path)
	if err != nil {
		return nil, err
	} else {
		return kubeConfig, nil
	}
}

func buildConfigFromFlags(context, kubeconfigPath string) (*rest.Config, error) {
	return clientcmd.NewNonInteractiveDeferredLoadingClientConfig(
		&clientcmd.ClientConfigLoadingRules{ExplicitPath: kubeconfigPath},
		&clientcmd.ConfigOverrides{
			CurrentContext: context,
		}).ClientConfig()
}
