package service

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/wonderivan/logger"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
	"strconv"
)

// Servicev1 定义Servicev1变量
var Servicev1 servicev1

// 定义service结构体
type servicev1 struct{}

type ServicesResp struct {
	Items []corev1.Service `json:"items"`
	Total int              `json:"total"`
}

// ServiceCreate 定义service创建结构体对象
type ServiceCreate struct {
	Name          string            `json:"name"`
	Namespace     string            `json:"namespace"`
	Type          string            `json:"type"`
	Protocol      string            `json:"protocol,omitempty"`
	ContainerPort int32             `json:"containerPort"`
	Port          int32             `json:"port"`
	NodePort      int32             `json:"nodePort"`
	Label         map[string]string `json:"label"`
}

// GetServices 获取service列表、支持过滤、排序和分页
func (s *servicev1) GetServices(filterName, namespace string, limit, page int) (servicesResp *ServicesResp, err error) {
	// 获取serviceList类型的service列表
	serviceList, err := K8s.ClientSet.CoreV1().Services(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		logger.Error(errors.New("获取Service列表失败,错误信息, " + err.Error()))
		return nil, errors.New("获取Service列表失败,错误信息, " + err.Error())
	}
	// 将serviceList中的service列表(Items),放进dataselector对象中，进行排序
	selectableData := &dataSelector{
		GenericDataList: s.toCells(serviceList.Items),
		DataSelectQuery: &DataSelectQuery{
			Filter: &FilterQuery{
				Name: filterName,
			},
			Paginate: &PaginateQuery{
				Limit: limit,
				Page:  page,
			},
		},
	}
	filtered := selectableData.Filter()
	total := len(filtered.GenericDataList)
	data := filtered.Sort().Paginate()

	// 将将[]DataCell类型的service列表转为v1.service列表
	services := s.fromCells(data.GenericDataList)
	return &ServicesResp{
		Items: services,
		Total: total,
	}, nil
}

// GetServiceDetail 获取service详情
func (s *servicev1) GetServiceDetail(serviceName, namespace string) (service *corev1.Service, err error) {
	service, err = K8s.ClientSet.CoreV1().Services(namespace).Get(context.TODO(), serviceName, metav1.GetOptions{})
	if err != nil {
		logger.Error(errors.New("获取Service " + serviceName + "详情失败,错误信息, " + err.Error()))
		return nil, errors.New("获取Service " + serviceName + "详情失败,错误信息, " + err.Error())
	}
	return service, nil
}

// CreateService 创建Service,接受ServiceCreate对象
func (s *servicev1) CreateService(data *ServiceCreate) (err error) {
	// 将data中的数据组装成corev1.Service对象
	service := &corev1.Service{
		// ObjectMeta中定义资源名称,命名空间以以及标签
		ObjectMeta: metav1.ObjectMeta{
			Name:      data.Name,
			Namespace: data.Namespace,
			Labels:    data.Label,
		},
		// Spec中定义类型、端口、选择器
		Spec: corev1.ServiceSpec{
			Type: corev1.ServiceType(data.Type),
			Ports: []corev1.ServicePort{
				{
					Name:     data.Protocol + "-" + strconv.Itoa(int(data.Port)),
					Port:     data.Port,
					Protocol: "TCP",
					TargetPort: intstr.IntOrString{
						Type:   0,
						IntVal: data.ContainerPort,
					},
				},
			},
			Selector: data.Label,
		},
	}
	// 默认ClusterIP,这里是判断NodePort,添加配置
	if data.NodePort != 0 && data.Type == "NodePort" {
		service.Spec.Ports[0].NodePort = data.NodePort
	}
	// 创建Service
	_, err = K8s.ClientSet.CoreV1().Services(data.Namespace).Create(context.TODO(), service, metav1.CreateOptions{})
	if err != nil {
		logger.Error(errors.New("创建Service失败,错误信息," + err.Error()))
		return errors.New("创建Service失败,错误信息," + err.Error())
	}
	return nil
}

// DeleteService 删除service
func (s *servicev1) DeleteService(serviceName, namespace string) (err error) {
	err = K8s.ClientSet.CoreV1().Services(namespace).Delete(context.TODO(), serviceName, metav1.DeleteOptions{})
	if err != nil {
		logger.Error(errors.New("删除Service " + serviceName + "失败,错误信息," + err.Error()))
		return errors.New("删除Service " + serviceName + "失败,错误信息," + err.Error())
	}
	return nil
}

// UpdateService 更新service
func (s *servicev1) UpdateService(namespace, content string) (err error) {
	var service = &corev1.Service{}
	err = json.Unmarshal([]byte(content), service)
	if err != nil {
		logger.Error(errors.New("反序列化失败, " + err.Error()))
		return errors.New("反序列化失败, " + err.Error())
	}
	_, err = K8s.ClientSet.CoreV1().Services(namespace).Update(context.TODO(), service, metav1.UpdateOptions{})
	if err != nil {
		logger.Error(errors.New("更新service失败, " + err.Error()))
		return errors.New("更新service失败, " + err.Error())
	}
	return nil
}

func (s *servicev1) toCells(std []corev1.Service) []DataCell {
	cells := make([]DataCell, len(std))
	for i := range std {
		cells[i] = serviceCell(std[i])
	}
	return cells
}

func (s *servicev1) fromCells(cells []DataCell) []corev1.Service {
	services := make([]corev1.Service, len(cells))
	for i := range cells {
		services[i] = corev1.Service(cells[i].(serviceCell))
	}
	return services
}
