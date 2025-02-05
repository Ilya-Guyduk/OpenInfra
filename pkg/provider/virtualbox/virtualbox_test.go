package virtualbox_test

import (
	"context"
	"testing"

	"github.com/Ilya-Guyduk/go-openinfra/pkg/provider"
	"github.com/Ilya-Guyduk/go-openinfra/pkg/provider/virtualbox"
	"github.com/stretchr/testify/assert"
)

var conn_det = make(map[string]string)

func TestVirtualBoxProvider(t *testing.T) {
	vb := virtualbox.New("localhost", conn_det)
	ctx := context.Background()

	err := vb.Connect(ctx)
	assert.NoError(t, err, "Connect() должно работать без ошибок")

	resourceID, err := vb.CreateResource(ctx, provider.Resource{Type: "virtual_machine"})
	assert.NoError(t, err, "CreateResource() должно работать без ошибок")
	assert.NotEmpty(t, resourceID, "ID ресурса не должен быть пустым")

	res, err := vb.GetResource(ctx, resourceID)
	assert.NoError(t, err, "GetResource() должно работать без ошибок")
	assert.Equal(t, "virtual_machine", res.Type, "Тип ресурса должен быть 'virtual_machine'")

	err = vb.PerformAction(ctx, resourceID, "start")
	assert.NoError(t, err, "PerformAction() (start) должно работать без ошибок")

	err = vb.DeleteResource(ctx, resourceID)
	assert.NoError(t, err, "DeleteResource() должно работать без ошибок")

	err = vb.Disconnect(ctx)
	assert.NoError(t, err, "Disconnect() должно работать без ошибок")
}
