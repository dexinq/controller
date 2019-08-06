package handler

import (
	"context"
	"encoding/json"

	"github.com/dexinq/utils/proto/controller"
	"github.com/dexinq/utils/proto/dataservice"
	"github.com/dexinq/utils/proto/global"

	"github.com/google/uuid"
	"github.com/micro/go-micro/broker"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/util/log"
)

const ControllerSyncTopic = "app.controller.sync"

type ControllerHandler struct {
	client client.Client
	broker broker.Broker
}

func NewControllerHandler(cl client.Client, broker broker.Broker) *ControllerHandler {
	return &ControllerHandler{client: cl, broker: broker}
}

func (ch *ControllerHandler) RegisterController(ctx context.Context, req *controller.ControllerObj, rsp *global.Response) error {
	ds := dataservice.NewControllerService("go.micro.srv.dataservice", ch.client)
	if _, err := ds.AddController(ctx, req); err != nil {
		log.Fatal(err.Error())
	}

	msg, err := json.Marshal(req)
	if err != nil {
		log.Fatalf(err.Error())
		return err
	}

	m := &broker.Message{
		Header: map[string]string{"id": uuid.New().String()},
		Body:   msg,
	}

	if err := broker.Publish(ControllerSyncTopic, m); err != nil {
		log.Fatalf(err.Error())
		return err
	}

	return nil
}

func (ch *ControllerHandler) UpdateController(ctx context.Context, req *controller.ControllerObj, rsp *global.Response) error {
	ds := dataservice.NewControllerService("go.micro.srv.dataservice", ch.client)
	if _, err := ds.UpdateController(ctx, req); err != nil {
		log.Fatalf(err.Error())
	}
	return nil
}
func (ch *ControllerHandler) GetAppObject(ctx context.Context, req *global.Empty, rsp *controller.ControllerList) error {
	ds := dataservice.NewControllerService("go.micro.srv.dataservice", ch.client)
	if _, err := ds.GetController(ctx, req); err != nil {
		log.Fatalf(err.Error())
	}
	return nil
}
