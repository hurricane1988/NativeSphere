package service

import (
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	nwv1 "k8s.io/api/networking/v1"
	"sort" // 自定义类型排序参考文档：https://segmentfault.com/a/1190000008062661
	"strings"
	"time"
)

// dataSelect 用于封装排序、过滤、分页的数据类型
type dataSelector struct {
	GenericDataList []DataCell
	DataSelectQuery *DataSelectQuery
}

// DataCell DataCell接口，用于各种资源list的类型转换，转换后可以使用dataSelector的自定义排序、过滤、分页方法
type DataCell interface {
	GetCreation() time.Time
	GetName() string
}

// DataSelectQuery 定义过滤和分页的属性，过滤:Name， 分页:Limit和Page
// Limit是单页的数据条数
// Page是第几页
type DataSelectQuery struct {
	Filter   *FilterQuery
	Paginate *PaginateQuery
}

// FilterQuery 过滤查询
type FilterQuery struct {
	Name string
}

// PaginateQuery 分页查询
type PaginateQuery struct {
	Limit int
	Page  int
}

// 定义deployment结构体
type deploymentCell appsv1.Deployment

// GetCreation 获取deployment的创建时间
func (d deploymentCell) GetCreation() time.Time {
	return d.CreationTimestamp.Time
}

// GetName 获取deployment的名称
func (d deploymentCell) GetName() string {
	return d.Name
}

//实现自定义结构的排序，需要重写Len、Swap、Less方法

// Len 方法用于获取数组长度
func (d *dataSelector) Len() int {
	return len(d.GenericDataList)
}

// Swap 方法用于数组中的元素在比较大小后的位置交换，可定义升序或降序
// i, j 为数据GenericDataList的下标
func (d *dataSelector) Swap(i, j int) {
	d.GenericDataList[i], d.GenericDataList[j] = d.GenericDataList[j], d.GenericDataList[i]
	// 常规做法
	//temp := d.GenericDataList[i]
	//d.GenericDataList[i] = d.GenericDataList[j]
	//d.GenericDataList[j] = temp
}

// Less 方法用于定义数组中元素排序的“大小”的比较方式
func (d *dataSelector) Less(i, j int) bool {
	a := d.GenericDataList[i].GetCreation()
	b := d.GenericDataList[j].GetCreation()
	return b.Before(a)
}

// Sort 重写以上3个方法用使用sort.Sort进行排序
func (d *dataSelector) Sort() *dataSelector {
	sort.Sort(d)
	return d
}

// Filter 方法用于过滤元素，比较元素的Name属性，若包含，再返回
func (d *dataSelector) Filter() *dataSelector {
	// 判断入参是否为空，若为空，则返回所有数据
	if d.DataSelectQuery.Filter.Name == "" {
		return d
	}
	// 若入参的传参不为空，则返回元素名中包含Name的所有元素
	// 声明一个新的数组,若Name包含,则把数据放进数组，返回出去
	var filtered []DataCell
	for _, value := range d.GenericDataList {
		// 定义是否匹配标签变量，默认匹配
		matches := true
		objName := value.GetName()
		if !strings.Contains(objName, d.DataSelectQuery.Filter.Name) {
			matches = false
			continue
		}
		if matches {
			filtered = append(filtered, value)
		}
	}
	d.GenericDataList = filtered
	return d
}

// Paginate 方法用于数据分页，根据limit和page的传参，取一定范围内的数据，返回
func (d *dataSelector) Paginate() *dataSelector {
	// 根据Limit和Page的入参，定义快捷变量
	limit := d.DataSelectQuery.Paginate.Limit
	page := d.DataSelectQuery.Paginate.Page
	// 鉴于参数的合法性
	if limit <= 0 || page <= 0 {
		return d
	}
	// 定义取值范围需要的startIndex和endIndex
	// 举例,有25个元素的数组，limit是10， page是3，startIndex是20，endIndex是30（endIndex是24）
	startIndex := limit * (page - 1)
	endIndex := limit*page - 1

	// 处理endIndex
	if endIndex > len(d.GenericDataList) {
		endIndex = len(d.GenericDataList) - 1
	}

	d.GenericDataList = d.GenericDataList[startIndex:endIndex]
	return d
}

// 定义podCell,重写GetCreation和GetName方法后，可以进行数据转换
// corev1.Pod -> podcell -> DataCell
// appsv1.Deployment -> deploycell -> DataCell
type podCell corev1.Pod

// GetCreation 重新DataCell接口的两个方法
func (p podCell) GetCreation() time.Time {
	return p.CreationTimestamp.Time
}

func (p podCell) GetName() string {
	return p.Name
}

/* daemonSet相关配置 */
type daemonSetCell appsv1.DaemonSet

func (d daemonSetCell) GetCreation() time.Time {
	return d.CreationTimestamp.Time
}

func (d daemonSetCell) GetName() string {
	return d.Name
}

/* statefulSet相关配置 */
type statefulSetCell appsv1.StatefulSet

func (s statefulSetCell) GetCreation() time.Time {
	return s.CreationTimestamp.Time
}

func (s statefulSetCell) GetName() string {
	return s.Name
}

/* service相关配置 */
type serviceCell corev1.Service

func (s serviceCell) GetCreation() time.Time {
	return s.CreationTimestamp.Time
}

func (s serviceCell) GetName() string {
	return s.Name
}

/* ingress相关配置 */
type ingressCell nwv1.Ingress

func (i ingressCell) GetCreation() time.Time {
	return i.CreationTimestamp.Time
}

func (i ingressCell) GetName() string {
	return i.Name
}

/* configMap相关配置 */
type configMapCell corev1.ConfigMap

func (c configMapCell) GetCreation() time.Time {
	return c.CreationTimestamp.Time
}

func (c configMapCell) GetName() string {
	return c.Name
}

/* secret相关配置 */
type secretCell corev1.Secret

func (s secretCell) GetCreation() time.Time {
	return s.CreationTimestamp.Time
}

func (s secretCell) GetName() string {
	return s.Name
}

/* node相关配置 */
type nodeCell corev1.Node

func (n nodeCell) GetCreation() time.Time {
	return n.CreationTimestamp.Time
}

func (n nodeCell) GetName() string {
	return n.Name
}

/* namespace相关配置 */
type namespaceCell corev1.Namespace

func (n namespaceCell) GetCreation() time.Time {
	return n.CreationTimestamp.Time
}

func (n namespaceCell) GetName() string {
	return n.Name
}

/* pv相关配置 */
type pvCell corev1.PersistentVolume

func (p pvCell) GetCreation() time.Time {
	return p.CreationTimestamp.Time
}

func (p pvCell) GetName() string {
	return p.Name
}

/* pvc相关配置 */
type pvcCell corev1.PersistentVolumeClaim

func (p pvcCell) GetCreation() time.Time {
	return p.CreationTimestamp.Time
}

func (p pvcCell) GetName() string {
	return p.Name
}
