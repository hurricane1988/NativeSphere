package service

import (
	"context"
	"errors"
	"github.com/wonderivan/logger"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// Pv 定义PV变量
var Pv pv

// 定义pv结构提
type pv struct{}

// PvsResp 定义PvResp 结构体
type PvsResp struct {
	Items []corev1.PersistentVolume `json:"items"`
	Total int                       `json:"total"`
}

// GetPvs 获取pv列表、支持过滤、排序、分页
func (p *pv) GetPvs(filterName string, limit, page int) (PvResp *PvsResp, err error) {
	// 获取PVList类型的pv列表
	pvList, err := K8s.ClientSet.CoreV1().PersistentVolumes().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		logger.Error(errors.New("获取pv " + filterName + "失败，错误信息 " + err.Error()))
		return nil, errors.New("获取pv " + filterName + "失败，错误信息" + err.Error())
	}
	// 将PvList中的pv列表(items),放进dataselector对象中，进行排序
	selectableData := &dataSelector{
		GenericDataList: p.toCells(pvList.Items),
		DataSelectQuery: &DataSelectQuery{
			Filter: &FilterQuery{
				Name: filterName,
			},
			Paginate: &PaginateQuery{
				limit,
				page,
			},
		},
	}
	filtered := selectableData.Filter()
	total := len(filtered.GenericDataList)
	data := filtered.Sort().Paginate()

	// 将[]DataCell类型中的pv列表转为v1.pv列表
	pvs := p.fromCells(data.GenericDataList)

	return &PvsResp{
		Items: pvs,
		Total: total,
	}, nil
}

// GetPvDetail 获取pv详情
func (p *pv) GetPvDetail(pvName string) (pv *corev1.PersistentVolume, err error) {
	pv, err = K8s.ClientSet.CoreV1().PersistentVolumes().Get(context.TODO(), pvName, metav1.GetOptions{})
	if err != nil {
		logger.Error(errors.New("获取pv " + pvName + "详细失败，错误信息 " + err.Error()))
		return nil, errors.New("获取pv " + pvName + "详细失败，错误信息 " + err.Error())
	}
	return pv, nil
}

// DeletePv 删除pv
func (p *pv) DeletePv(pvName string) (err error) {
	err = K8s.ClientSet.CoreV1().PersistentVolumes().Delete(context.TODO(), pvName, metav1.DeleteOptions{})
	if err != nil {
		logger.Error(errors.New("删除pv " + pvName + "失败，错误信息 " + err.Error()))
		return errors.New("删除pv " + pvName + "失败，错误信息 " + err.Error())
	}
	return nil
}

func (p *pv) toCells(std []corev1.PersistentVolume) []DataCell {
	cells := make([]DataCell, len(std))
	for i := range std {
		cells[i] = pvCell(std[i])
	}
	return cells
}

func (p *pv) fromCells(cells []DataCell) []corev1.PersistentVolume {
	pvs := make([]corev1.PersistentVolume, len(cells))
	for i := range cells {
		pvs[i] = corev1.PersistentVolume(cells[i].(pvCell))
	}

	return pvs
}