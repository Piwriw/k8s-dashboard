package pod

import (
	"context"
	"errors"
	"fmt"
	corev1 "k8s.io/api/core/v1"
	k8serror "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/watch"
	"k8sdashboar.com/global"
	pod_req "k8sdashboar.com/model/pod/request"
	pod_res "k8sdashboar.com/model/pod/response"
	"strings"
)

type PodService struct {
}

func (*PodService) DeletePod(namespace, name string) error {
	background := metav1.DeletePropagationBackground
	var gracePeriodSeconds int64 = 0
	return global.KubeConfigSet.CoreV1().Pods(namespace).Delete(context.TODO(), name,
		metav1.DeleteOptions{
			GracePeriodSeconds: &gracePeriodSeconds,
			PropagationPolicy:  &background,
		})
}
func (*PodService) GetPodList(namespace, keyword, nodeName string) (err error, _ []pod_res.PodListItem) {
	ctx := context.TODO()
	list, err := global.KubeConfigSet.CoreV1().Pods(namespace).List(ctx, metav1.ListOptions{})
	if err != nil {
		return
	}
	podList := make([]pod_res.PodListItem, 0)
	for _, item := range list.Items {
		if nodeName != "" && item.Spec.NodeName != nodeName {
			continue
		}
		if strings.Contains(item.Name, keyword) {
			podItem := podConvert.K8s2rResConvert.PodK8s2ItemRes(item)
			podList = append(podList, podItem)
		}

	}
	return err, podList
}
func (*PodService) GetPodDetail(namespace string, name string) (podReq pod_req.Pod, err error) {
	ctx := context.TODO()
	podApi := global.KubeConfigSet.CoreV1().Pods(namespace)
	k8sGetPod, err := podApi.Get(ctx, name, metav1.GetOptions{})
	if err != nil {
		errMsg := fmt.Sprintf("Pod[namespace=%s,name=%s]查询失败，detail:%s", namespace, name, err.Error())
		err = errors.New(errMsg)
		return
	}
	// k8sPod 转为 pod request
	podReq = podConvert.K8s2rReqConvert.PodK8s2Req(*k8sGetPod)
	return
}

// CreateOrUpdate 创建或者更新Pod
func (*PodService) CreateOrUpdate(podReq pod_req.Pod) (string, error) {
	k8sPod := podConvert.Req2K8sConvert.PodReq2K8s(podReq)
	ctx := context.TODO()
	podApi := global.KubeConfigSet.CoreV1().Pods(k8sPod.Namespace)
	if k8sGetPod, err := podApi.Get(ctx, k8sPod.Name, metav1.GetOptions{}); err != nil {
		k8sPodCopy := *k8sPod
		k8sPodCopy.Name = k8sPod.Name + "-validate"
		_, err := podApi.Create(ctx, &k8sPodCopy, metav1.CreateOptions{
			DryRun: []string{metav1.DryRunAll},
		})
		if err != nil {
			errMsg := fmt.Sprintf("Pod[namespace=%s,name=%s]更新失败，detail:%s", k8sPod.Namespace, k8sPod.Name, err.Error())
			return errMsg, err
		}
		// 强制的删除
		background := metav1.DeletePropagationBackground
		var gracePeriodSeconds int64 = 0
		err = podApi.Delete(ctx, k8sPod.Name, metav1.DeleteOptions{
			GracePeriodSeconds: &gracePeriodSeconds,
			PropagationPolicy:  &background,
		})
		if err != nil {
			errMsg := fmt.Sprintf("Pod[namespace=%s,name=%s]更新失败，detail:%s", k8sPod.Namespace, k8sPod.Name, err.Error())
			return errMsg, err
		}
		var labelsSelector []string
		for k, v := range k8sGetPod.Labels {
			labelsSelector = append(labelsSelector, fmt.Sprintf("%s=%s", k, v))
		}
		// 监听，等pod处于特瑞，terminating 删除完毕，建立pod
		// label 格式 app=test,app2=test2
		watcher, err := podApi.Watch(ctx, metav1.ListOptions{
			LabelSelector: strings.Join(labelsSelector, ","),
		})
		if err != nil {
			errMsg := fmt.Sprintf("Pod[namespace=%s,name=%s]更新失败，detail:%s", k8sPod.Namespace, k8sPod.Name, err.Error())
			return errMsg, err
		}
		for event := range watcher.ResultChan() {
			k8sPodChan := event.Object.(*corev1.Pod)
			// 查询k8s 是否已经删除，那就不要判断删除时间
			if _, err := podApi.Get(ctx, k8sPod.Name, metav1.GetOptions{}); k8serror.IsNotFound(err) {
				pod, err := global.KubeConfigSet.CoreV1().Pods("").Create(ctx, k8sPod, metav1.CreateOptions{})
				if err != nil {
					errMsg := fmt.Sprintf("Pod[namespace=%s,name=%s]更新失败，detail:%s", k8sPod.Namespace, k8sPod.Name, err.Error())
					return errMsg, err
				} else {
					successMsg := fmt.Sprintf("Pod[namespace=%s,name=%s]更新成功", pod.Namespace, pod.Name)
					return successMsg, nil
				}
			}

			switch event.Type {
			case watch.Deleted:
				if k8sPodChan.Name != k8sGetPod.Name {
					continue
				}
				// 重建
				pod, err := global.KubeConfigSet.CoreV1().Pods("").Create(ctx, k8sPod, metav1.CreateOptions{})
				if err != nil {
					errMsg := fmt.Sprintf("Pod[namespace=%s,name=%s]更新失败，detail:%s", k8sPod.Namespace, k8sPod.Name, err.Error())
					return errMsg, err
				} else {
					successMsg := fmt.Sprintf("Pod[namespace=%s,name=%s]更新成功", pod.Namespace, pod.Name)
					return successMsg, nil
				}
			}
		}
		return "", nil
	} else {
		pod, err := global.KubeConfigSet.CoreV1().Pods("").Create(ctx, k8sPod, metav1.CreateOptions{})
		if err != nil {
			errMsg := fmt.Sprintf("PodPod[namespace=%s,name=%s],创建失败，detail:%s", k8sPod.Namespace, k8sPod.Name, err.Error())
			return errMsg, err
		} else {
			successMsg := fmt.Sprintf("PodPod[namespace=%s,name=%s],创建成功", pod.Namespace, pod.Name)
			return successMsg, nil
		}
	}
	return "", nil
}
