package controller

import (
	"github.com/gin-gonic/gin"
)

// Router 实例化router结构体，可使用该对象点出首字母大写的方法(包外调用)
var Router router

// 创建router结构体
type router struct{}

// InitApiRouter 初始化路由规则，创建测试api接口
func (r *router) InitApiRouter(router *gin.Engine) {
	router.
		/* login登录路由 */
		POST("/api/v1/login", Login.Auth).
		// 使用jwt认证(需要在其他路由之前加载)
		//Use(middle.JWTAuth()).
		/* workflow工作流路由 */
		GET("/api/v1/k8s/workflows", Workflow.GetList).
		GET("/api/v1/k8s/workflow/detail", Workflow.GetById).
		POST("/api/v1/k8s/workflow/create", Workflow.Create).
		DELETE("/api/v1/k8s/workflow/del", Workflow.DelById).
		GET("/api/v1/k8s/testapi", TestApi.TestAPI).
		/* POD相关路由 */
		GET("/api/v1/k8s/pods", Pod.GetPods).
		GET("/api/v1/k8s/pod/detail", Pod.GetPodDetail).
		DELETE("/api/v1/k8s/pod/delete", Pod.DeletePod).
		PUT("/api/v1/k8s/pod/update", Pod.UpdatePod).
		GET("/api/v1/k8s/pod/container", Pod.GetPodContainer).
		GET("/api/v1/k8s/pod/log", Pod.GetPodLog).
		GET("/api/v1/k8s/pod/numnp", Pod.GetPodNumPerNp).
		/* Deployment相关路由 */
		GET("/api/v1/k8s/deployments", Deployment.GetDeployments).
		GET("/api/v1/k8s/deployment/rs", Deployment.GetDeployReplicaSets).
		GET("/api/v1/k8s/deployment/detail", Deployment.GetDeploymentDetail).
		PUT("/api/v1/k8s/deployment/scale", Deployment.ScaleDeployment).
		DELETE("/api/v1/k8s/deployment/del", Deployment.DeleteDeployment).
		PUT("/api/v1/k8s/deployment/restart", Deployment.RestartDeployment).
		PUT("/api/v1/k8s/deployment/update", Deployment.UpdateDeployment).
		GET("/api/v1/k8s/deployment/numnp", Deployment.GetDeployNumPerNp).
		POST("/api/v1/k8s/deployment/create", Deployment.CreateDeployment).
		/* DaemonSet相关路由 */
		GET("/api/v1/k8s/daemonsets", DaemonSet.GetDaemonSets).
		GET("/api/v1/k8s/daemonset/detail", DaemonSet.GetDaemonSetDetail).
		DELETE("/api/v1/k8s/daemonset/del", DaemonSet.DeleteDaemonSet).
		PUT("/api/v1/k8s/daemonset/update", DaemonSet.UpdateDaemonSet).
		/* statefulSet相关路由*/
		GET("/api/v1/k8s/statefulsets", StatefulSet.GetStatefulSets).
		GET("/api/v1/k8s/statefulset/detail", StatefulSet.GetStatefulSetDetail).
		DELETE("/api/v1/k8s/statefulset/del", StatefulSet.DeleteStatefulSet).
		PUT("/api/v1/k8s/statefulset/update", StatefulSet.UpdateStatefulSet).
		/* service相关路由 */
		GET("/api/v1/k8s/services", Servicev1.GetServices).
		GET("/api/v1/k8s/service/detail", Servicev1.GetServiceDetail).
		DELETE("/api/v1/k8s/service/del", Servicev1.DeleteService).
		PUT("/api/v1/k8s/service/update", Servicev1.UpdateService).
		POST("/api/v1/k8s/service/create", Servicev1.CreateService).
		/* ingress相关路由 */
		GET("/api/v1/k8s/ingresses", Ingress.GetIngresses).
		GET("/api/v1/k8s/ingress/detail", Ingress.GetIngressDetail).
		DELETE("/api/v1/k8s/ingress/del", Ingress.DeleteIngress).
		PUT("/api/v1/k8s/ingress/update", Ingress.UpdateIngress).
		POST("/api/v1/k8s/ingress/create", Ingress.CreateIngress).
		/* namespace相关 */
		GET("/api/v1/k8s/namespaces", Namespace.GetNamespaces).
		GET("/api/v1/k8s/namespace/detail", Namespace.GetNamespaceDetail).
		DELETE("/api/v1/k8s/namespace/del", Namespace.DeleteNamespace).
		POST("/api/v1/k8s/namespace/create", Namespace.CreateNamespace).
		/* PV相关路由 */
		GET("/api/v1/k8s/pvs", Pv.GetPvs).
		GET("/api/v1/k8s/pv/detail", Pv.GetPvDetail).
		DELETE("/api/v1/k8s/pv/del", Pv.DeletePv).
		/* PVC相关路由 */
		GET("/api/v1/k8s/pvcs", Pvc.GetPvcs).
		GET("/api/v1/k8s/pvc/detail", Pvc.GetPvcDetail).
		PUT("/api/v1/k8s/pvc/update", Pvc.UpdatePvc).
		DELETE("/api/v1/k8s/pvc/del", Pvc.DeletePvc).
		/* node相关路由 */
		GET("/api/v1/k8s/nodes", Node.GetNodes).
		GET("/api/v1/k8s/node/detail", Node.GetNodeDetail).
		/* configmap相关路由 */
		GET("/api/v1/k8s/configmaps", ConfigMap.GetConfigMaps).
		GET("/api/v1/k8s/configmap/detail", ConfigMap.GetConfigMapDetail).
		PUT("/api/v1/k8s/configmap/update", ConfigMap.UpdateConfigMap).
		DELETE("/api/v1/k8s/configmap/del", ConfigMap.DeleteConfigMap).
		/* Secret秘钥相关理由 */
		GET("/api/v1/k8s/secrets", Secret.GetSecrets).
		GET("/api/v1/k8s/secret/detail", Secret.GetSecretDetail).
		DELETE("/api/v1/k8s/secret/del", Secret.DeleteSecret).
		PUT("/api/v1/k8s/secret/update", Secret.UpdateSecret)
}
