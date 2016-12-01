package check_pod_status

import (
	"fmt"
	"os"

	"encoding/json"

	"strings"

	"github.com/appscode/searchlight/pkg/config"
	"github.com/appscode/searchlight/pkg/util"
	"github.com/spf13/cobra"
	kApi "k8s.io/kubernetes/pkg/api"
	"k8s.io/kubernetes/pkg/labels"
)

type request struct {
	host string
}

type objectInfo struct {
	Name      string `json:"name,omitempty"`
	Namespace string `json:"namespace,omitempty"`
	Status    string `json:"status,omitempty"`
}

type serviceOutput struct {
	Objects []*objectInfo `json:"objects,omitempty"`
	Message string        `json:"message,omitempty"`
}

func checkPodStatus(namespace, objectType, objectName string) {
	kubeClient, err := config.GetKubeClient()
	if err != nil {
		fmt.Fprintln(os.Stdout, util.State[3], err)
		os.Exit(3)
	}

	objectInfoList := make([]*objectInfo, 0)
	if objectType == config.TypePods {
		pod, err := kubeClient.Pods(namespace).Get(objectName)
		if err != nil {
			fmt.Fprintln(os.Stdout, util.State[3], err)
			os.Exit(3)
		}

		if !(pod.Status.Phase == kApi.PodSucceeded || pod.Status.Phase == kApi.PodRunning) {
			objectInfoList = append(objectInfoList, &objectInfo{Name: pod.Name, Status: string(pod.Status.Phase), Namespace: pod.Namespace})
		}
	} else {
		var labelSelector labels.Selector
		if objectType == "" {
			labelSelector = labels.Everything()
		} else {
			if labelSelector, err = kubeClient.GetLabels(namespace, objectType, objectName); err != nil {
				fmt.Fprintln(os.Stdout, util.State[3], err)
				os.Exit(3)
			}
		}

		podList, err := kubeClient.Pods(namespace).List(kApi.ListOptions{
			LabelSelector: labelSelector,
		})
		if err != nil {
			fmt.Fprintln(os.Stdout, util.State[3], err)
			os.Exit(3)
		}

		for _, pod := range podList.Items {
			if !(pod.Status.Phase == kApi.PodSucceeded || pod.Status.Phase == kApi.PodRunning) {
				objectInfoList = append(objectInfoList, &objectInfo{Name: pod.Name, Status: string(pod.Status.Phase), Namespace: pod.Namespace})
			}
		}
	}

	if len(objectInfoList) == 0 {
		fmt.Fprintln(os.Stdout, util.State[0], "All pods are Ready")
		os.Exit(0)
	} else {
		output := &serviceOutput{
			Objects: objectInfoList,
			Message: fmt.Sprintf("Found %d not running pods(s)", len(objectInfoList)),
		}
		outputByte, err := json.MarshalIndent(output, "", "  ")
		if err != nil {
			fmt.Fprintln(os.Stdout, util.State[3], err)
			os.Exit(3)
		}
		fmt.Fprintln(os.Stdout, util.State[0], string(outputByte))
		os.Exit(0)
	}
}

func NewCmd() *cobra.Command {
	var req request
	c := &cobra.Command{
		Use:     "pod_status",
		Short:   "Check Kubernetes Pod(s) status",
		Example: "",

		Run: func(cmd *cobra.Command, args []string) {
			util.EnsureFlagsSet(cmd, "host")

			parts := strings.Split(req.host, "@")
			if len(parts) != 2 {
				fmt.Fprintln(os.Stdout, util.State[3], "Invalid icinga host.name")
				os.Exit(3)
			}

			name := parts[0]
			namespace := parts[1]

			parts = strings.Split(name, ":")
			objectType := ""
			objectName := ""
			if len(parts) == 1 {
				if parts[0] != "pod_status" {
					objectType = config.TypePods
					objectName = parts[0]
				}
			} else if len(parts) == 2 {
				objectType = parts[0]
				objectName = parts[1]
			}

			checkPodStatus(namespace, objectType, objectName)
		},
	}
	c.Flags().StringVarP(&req.host, "host", "H", "", "Icinga host name")
	return c
}
