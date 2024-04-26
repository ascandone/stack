package internal

import (
	"reflect"

	"github.com/formancehq/stack/components/agent/internal/generated"
	sharedlogging "github.com/formancehq/stack/libs/go-libs/logging"
	"google.golang.org/protobuf/types/known/structpb"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/client-go/tools/cache"
)

func NewStackEventHandler(logger sharedlogging.Logger, membershipClient MembershipClient) cache.ResourceEventHandlerFuncs {
	sendStatus := func(stackName string, status *structpb.Struct) {
		if err := membershipClient.Send(&generated.Message{
			Message: &generated.Message_StatusChanged{
				StatusChanged: &generated.StatusChanged{
					ClusterName: stackName,
					Statuses:    status,
				},
			},
		}); err != nil {
			logger.Errorf("Unable to send stack status to server: %s", err)
		}
	}

	return cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			stack := obj.(*unstructured.Unstructured)
			status, err := getStatus(stack)
			if err != nil {
				logger.Errorf("Unable to generate message stack update: %s", err)
				return
			}

			logger.Infof("Stack '%s' added", stack.GetName())
			sendStatus(stack.GetName(), status)
		},
		UpdateFunc: func(oldObj, newObj interface{}) {
			oldStack := oldObj.(*unstructured.Unstructured)
			newStack := newObj.(*unstructured.Unstructured)

			oldStatus, err := getStatus(oldStack)
			if err != nil {
				logger.Errorf("Unable to get old stack status update: %s", err)
			}

			newStatus, err := getStatus(newStack)
			if err != nil {
				logger.Errorf("Unable to get new stack status update: %s", err)
				return
			}

			if reflect.DeepEqual(oldStatus, newStatus) {
				return
			}

			logger.Infof("Stack '%s' updated", newStack.GetName())
			sendStatus(newStack.GetName(), newStatus)
		},
		DeleteFunc: func(obj interface{}) {
			stack := obj.(*unstructured.Unstructured)
			logger.Infof("Stack '%s' deleted", stack.GetName())

			if err := membershipClient.Send(&generated.Message{
				Message: &generated.Message_StackDeleted{
					StackDeleted: &generated.DeletedStack{
						ClusterName: stack.GetName(),
					},
				},
			}); err != nil {
				logger.Errorf("Unable to send stack delete to server: %s", err)
			}
		},
	}
}
