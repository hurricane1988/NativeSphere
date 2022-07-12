package service

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/wonderivan/logger"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
)

// DaemonSet 初始化daemonSet结构体
var DaemonSet daemonSet

// 定义daemonSet结构体
type daemonSet struct{}

// DaemonSetsResp 定义列表的返回内容，Items是deployment元素列表，Total为deployment元素数量
type DaemonSetsResp struct {
	Items []appsv1.DaemonSet `json:"items"`
	Total int                `json:"total"`
}

// DaemonSetCreate 定义DaemonSetCreate结构体，用于创建DaemonSet需要的参数属性的定义
type DaemonSetCreate struct {
	Name          string            `json:"name"`
	Namespace     string            `json:"namespace"`
	Image         string            `json:"image"`
	Label         map[string]string `json:"label"`
	Cpu           string            `json:"cpu"`
	Memory        string            `json:"memory"`
	ContainerPort int32             `json:"containerPort"`
	HealthCheck   bool              `json:"healthCheck"`
	HealthPath    string            `json:"healthPath"`
}

// DaemonSetNp 定义DeployNp类型，用于返回namespace中DaemonSet的数量
type DaemonSetNp struct {
	Namespace    string `json:"namespace"`
	DaemonSetNum int    `json:"daemonSet_num"`
}

// GetDaemonSets 获取DaemonSet列表，支持过滤、排序、分页
func (d *daemonSet) GetDaemonSets(filterName, namespace string, limit, page int) (deploymentsResp *DaemonSetsResp, err error) {
	// 获取deploymentList类型的deployment列表
	daemonSetList, err := K8s.ClientSet.AppsV1().DaemonSets(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		logger.Error(errors.New("获取namespace" + namespace + "中的pod失败,错误信息" + err.Error()))
		return nil, errors.New("获取namespace" + namespace + "中的pod失败,错误信息" + err.Error())
	}
	// 将deploymentList中的deployment列表(Items)，放进dataselector对象中，进行排序
	selectableData := &dataSelector{
		GenericDataList: d.toCells(daemonSetList.Items),
		DataSelectQuery: &DataSelectQuery{
			Filter: &FilterQuery{Name: filterName},
			Paginate: &PaginateQuery{
				limit,
				page,
			},
		},
	}
	filtered := selectableData.Filter()
	total := len(filtered.GenericDataList)
	data := filtered.Sort().Paginate()

	// 将[]DataCell类型的deployment列表转为appsv1.deployment列表
	daemonSets := d.fromCells(data.GenericDataList)

	return &DaemonSetsResp{
		Items: daemonSets,
		Total: total,
	}, nil
}

// GetDaemonSetDetail 获取daemonSet详情
func (d *daemonSet) GetDaemonSetDetail(daemonSetName, namespace string) (daemonSet *appsv1.DaemonSet, err error) {
	daemonSet, err = K8s.ClientSet.AppsV1().DaemonSets(namespace).Get(context.TODO(), daemonSetName, metav1.GetOptions{})
	if err != nil {
		logger.Error(errors.New("获取namespace" + namespace + "中的daemonSet详细信息失败,错误信息: " + err.Error()))
		return nil, errors.New("获取namespace" + namespace + "中的daemonSet详细信息失败,错误信息:" + err.Error())
	}
	return daemonSet, nil
}

// CreateDaemonSets 创建DaemonSets，接受DeployCreate对象
func (d *daemonSet) CreateDaemonSets(data *DaemonSetCreate) (err error) {
	// 将data中的数据组装成appsv1.Deployment对象
	daemonSet := &appsv1.DaemonSet{
		// ObjectMeta中定义资源名、命名空间以及标签
		ObjectMeta: metav1.ObjectMeta{
			Name:      data.Name,
			Namespace: data.Namespace,
			Labels:    data.Label,
		},
		// Spec中定义副本数、选择器、以及pod属性
		Spec: appsv1.DaemonSetSpec{
			Selector: &metav1.LabelSelector{
				MatchLabels: data.Label,
			},
			Template: corev1.PodTemplateSpec{
				// 定义pod名称和标签
				ObjectMeta: metav1.ObjectMeta{
					Name:   data.Name,
					Labels: data.Label,
				},
				// 定义容器名称，镜像和端口
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{
						{
							Name:  data.Name,
							Image: data.Image,
							Ports: []corev1.ContainerPort{
								{
									Name:          "http",
									Protocol:      corev1.ProtocolTCP,
									ContainerPort: 80,
								},
							},
						},
					},
				},
			},
		},
		// Status定义资源的运行状态，这里由于是新建，传入空的appsv1.DeployemntStatus{}对象即可
		Status: appsv1.DaemonSetStatus{},
	}
	// 判断是否打开健康检查功能，若打开，则定义ReadinessProbe和LivenessProbe
	if data.HealthCheck {
		// 设置第一个容器的ReadinessProbe，因为我们pod中只有一个容器，所以直接使用index 0即可
		// 若pod中有多个容器，则这里需要使用for循环去定义了
		daemonSet.Spec.Template.Spec.Containers[0].ReadinessProbe = &corev1.Probe{
			ProbeHandler: corev1.ProbeHandler{
				HTTPGet: &corev1.HTTPGetAction{
					Path: data.HealthPath,
					// intstr.IntOrString的作用是端口可以定义为整型，也可以定义为字符串
					// Type=0则表示表示该结构体实例内的数据为整型，转json时只使用IntVal的数据
					// Type=1则表示表示该结构体实例内的数据为字符串，转json时只使用StrVal的数据
					Port: intstr.IntOrString{
						Type:   0,
						IntVal: data.ContainerPort,
					},
				},
			},
			// 初始化等待时间
			InitialDelaySeconds: 15,
			// 超时时间
			TimeoutSeconds: 5,
			// 执行时间
			PeriodSeconds: 5,
		}
		daemonSet.Spec.Template.Spec.Containers[0].LivenessProbe = &corev1.Probe{
			ProbeHandler: corev1.ProbeHandler{
				HTTPGet: &corev1.HTTPGetAction{
					Path: data.HealthPath,
					Port: intstr.IntOrString{
						Type:   0,
						IntVal: data.ContainerPort,
					},
				},
			},
			InitialDelaySeconds: 15,
			TimeoutSeconds:      5,
			PeriodSeconds:       5,
		}
		// 定义容器的limit和request资源
		daemonSet.Spec.Template.Spec.Containers[0].Resources.Limits = map[corev1.ResourceName]resource.Quantity{
			corev1.ResourceCPU:    resource.MustParse(data.Cpu),
			corev1.ResourceMemory: resource.MustParse(data.Memory),
		}
		daemonSet.Spec.Template.Spec.Containers[0].Resources.Requests = map[corev1.ResourceName]resource.Quantity{
			corev1.ResourceCPU:    resource.MustParse(data.Cpu),
			corev1.ResourceMemory: resource.MustParse(data.Memory),
		}
	}
	// 调用sdk创建deployment
	_, err = K8s.ClientSet.AppsV1().DaemonSets(data.Namespace).Create(context.TODO(), daemonSet, metav1.CreateOptions{})
	if err != nil {
		logger.Error(errors.New("创建daemonSet" + daemonSet.Name + "失败,错误信息 " + err.Error()))
		return errors.New("创建daemonSet" + daemonSet.Name + "失败,错误信息 " + err.Error())
	}
	return nil
}

// DeleteDaemonSet 删除DaemonSet函数
func (d *daemonSet) DeleteDaemonSet(daemonSetName, namespace string) (err error) {
	err = K8s.ClientSet.AppsV1().DaemonSets(namespace).Delete(context.TODO(), daemonSetName, metav1.DeleteOptions{})
	if err != nil {
		logger.Error(errors.New("删除deployment " + daemonSetName + "失败，错误信息" + err.Error()))
		return errors.New("删除deployment " + daemonSetName + "失败，错误信息" + err.Error())
	}
	return nil
}

// UpdateDaemonSet 更新daemonSet
func (d *daemonSet) UpdateDaemonSet(namespace, content string) (err error) {
	var daemonSet = &appsv1.DaemonSet{}

	err = json.Unmarshal([]byte(content), daemonSet)
	if err != nil {
		logger.Error(errors.New("反序列化失败, " + err.Error()))
		return errors.New("反序列化失败, " + err.Error())
	}

	_, err = K8s.ClientSet.AppsV1().DaemonSets(namespace).Update(context.TODO(), daemonSet, metav1.UpdateOptions{})
	if err != nil {
		logger.Error(errors.New("更新DaemonSet失败, " + err.Error()))
		return errors.New("更新DaemonSet失败, " + err.Error())
	}
	return nil
}

// 类型转换
func (d *daemonSet) toCells(daemonSets []appsv1.DaemonSet) []DataCell {
	cells := make([]DataCell, len(daemonSets))
	for i := range daemonSets {
		cells[i] = daemonSetCell(daemonSets[i])
	}
	return cells
}

func (d *daemonSet) fromCells(cells []DataCell) []appsv1.DaemonSet {
	daemonSets := make([]appsv1.DaemonSet, len(cells))
	for i := range cells {
		daemonSets[i] = appsv1.DaemonSet(cells[i].(daemonSetCell))
	}

	return daemonSets
}
