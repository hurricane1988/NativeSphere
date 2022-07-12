package service

import (
	"context"
	"errors"
	"github.com/wonderivan/logger"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// Namespace 定义namespace变量
var Namespace namespace

// 定义namespace空结构体
type namespace struct{}

// NamespaceResp 定义namnespaceResp结构体
type NamespaceResp struct {
	Items []corev1.Namespace `json:"items"`
	Total int                `json:"total"`
}

// NamespaceCreate 定义namespace结构体，用于创建namespace需要的参数属性的定义
type NamespaceCreate struct {
	Name       string            `json:"name"`
	Label      map[string]string `json:"label"`
	Annotation map[string]string `json:"annotation"`
}

// GetNamespaces 获取namespace列表、支持过滤、排序和分页
func (n *namespace) GetNamespaces(filterName string, limit, page int) (namespaceResp *NamespaceResp, err error) {
	// 获取namespaceList类型的namespace列表
	namespaceList, err := K8s.ClientSet.CoreV1().Namespaces().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		logger.Error(errors.New("获取namespace列表失败,错误信息" + err.Error()))
		return nil, errors.New("获取namespace列表失败,错误信息" + err.Error())
	}
	// 将namespaceList中的namespace列表（items）,放进dataselector对象中，进行排序
	selectableData := &dataSelector{
		GenericDataList: n.toCells(namespaceList.Items),
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

	// 将[]DataCell类型的namespace列表转为v1.namespace列表
	namespaces := n.fromCells(data.GenericDataList)

	return &NamespaceResp{
		Items: namespaces,
		Total: total,
	}, nil
}

// GetNamespaceDetail 获取namespace详情
func (n *namespace) GetNamespaceDetail(namespaceName string) (namespace *corev1.Namespace, err error) {
	namespace, err = K8s.ClientSet.CoreV1().Namespaces().Get(context.TODO(), namespaceName, metav1.GetOptions{})
	if err != nil {
		logger.Error(errors.New("获取namespace " + namespaceName + "详情失败，错误信息 " + err.Error()))
		return nil, errors.New("获取namespace" + namespaceName + "详情失败，错误信息 " + err.Error())
	}
	return namespace, nil
}

// DeleteNamespace 删除namespace
func (n *namespace) DeleteNamespace(namespaceName string) (err error) {
	err = K8s.ClientSet.CoreV1().Namespaces().Delete(context.TODO(), namespaceName, metav1.DeleteOptions{})
	if err != nil {
		logger.Error(errors.New("删除namespace " + namespaceName + "成功!"))
		return errors.New("删除namespace " + namespaceName + "成功!")
	}
	return nil
}

// CreateNamespace 创建namespace
func (n *namespace) CreateNamespace(data *NamespaceCreate) (err error) {
	// 将data中的数据组装成appsv1.Namespace对象
	namespace := &corev1.Namespace{
		// ObjectMeta中定义资源名、命名空间以及标签
		ObjectMeta: metav1.ObjectMeta{
			Name:        data.Name,
			Labels:      data.Label,
			Annotations: data.Annotation,
		},
	}
	// 调用sdk创建deployment
	_, err = K8s.ClientSet.CoreV1().Namespaces().Create(context.TODO(), namespace, metav1.CreateOptions{})
	if err != nil {
		logger.Error(errors.New("创建namespace " + namespace.Name + "成功!"))
		return errors.New("创建namespace " + namespace.Name + "成功!")
	}
	return nil
}

// toCells方法用于将pod类型数组，转换成DataCell类型数组
// 类型转换的方法，corev1.Namespace -> DataCell, DataCell -> corev1.Namespace
func (n *namespace) toCells(namespaces []corev1.Namespace) []DataCell {
	cells := make([]DataCell, len(namespaces))
	for i := range namespaces {
		cells[i] = namespaceCell(namespaces[i])
	}
	return cells
}

// fromCells方法用于将DataCell类型数组，转换成namespace类型数组
func (n *namespace) fromCells(cells []DataCell) []corev1.Namespace {
	namespaces := make([]corev1.Namespace, len(cells))
	for i := range cells {
		namespaces[i] = corev1.Namespace(cells[i].(namespaceCell))
	}
	return namespaces
}
