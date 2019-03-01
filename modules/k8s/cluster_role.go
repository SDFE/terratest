package k8s

import (
	"testing"

	"github.com/stretchr/testify/require"
	rbacv1 "k8s.io/api/rbac/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// GetClusterRole returns a Kubernetes role resource with the given name. This will fail the test if there is an error.
func GetClusterRole(t *testing.T, options *KubectlOptions, roleName string) *rbacv1.ClusterRole {
	role, err := GetClusterRoleE(t, options, roleName)
	require.NoError(t, err)
	return role
}

// GetClusterRoleE returns a Kubernetes ClusterRole resource with the given name.
func GetClusterRoleE(t *testing.T, options *KubectlOptions, roleName string) (*rbacv1.ClusterRole, error) {
	clientset, err := GetKubernetesClientFromOptionsE(t, options)
	if err != nil {
		return nil, err
	}
	return clientset.RbacV1().ClusterRoles().Get(roleName, metav1.GetOptions{})
}
