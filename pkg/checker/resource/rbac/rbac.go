package rbac

import (
	"github.com/aiyengar2/hull/pkg/checker"
	corev1 "k8s.io/api/core/v1"
	rbacv1 "k8s.io/api/rbac/v1"
)

func init() {
	if err := corev1.AddToScheme(checker.Scheme); err != nil {
		panic(err)
	}
	if err := rbacv1.AddToScheme(checker.Scheme); err != nil {
		panic(err)
	}
}

type ServiceAccounts []*corev1.ServiceAccount

type ClusterRoles []*rbacv1.ClusterRole
type ClusterRoleBindings []*rbacv1.ClusterRoleBinding

type Roles []*rbacv1.Role
type RoleBindings []*rbacv1.RoleBinding
