package aws_test

import (
	"context"
	"testing"

	"github.com/Ilya-Guyduk/go-openinfra/pkg/provider"
	"github.com/Ilya-Guyduk/go-openinfra/pkg/provider/aws"
	"github.com/stretchr/testify/assert"
)

var conn_det = make(map[string]string)

func TestAWSProvider(t *testing.T) {
	awsProv := aws.New("https://api.aws.com", conn_det)
	ctx := context.Background()

	err := awsProv.Connect(ctx)
	assert.NoError(t, err, "Connect() должно работать без ошибок")

	resourceID, err := awsProv.CreateResource(ctx, provider.Resource{Type: "virtual_machine"})
	assert.NoError(t, err, "CreateResource() должно работать без ошибок")
	assert.NotEmpty(t, resourceID, "ID ресурса не должен быть пустым")

	res, err := awsProv.GetResource(ctx, resourceID)
	assert.NoError(t, err, "GetResource() должно работать без ошибок")
	assert.Equal(t, "virtual_machine", res.Type, "Тип ресурса должен быть 'virtual_machine'")

	err = awsProv.PerformAction(ctx, resourceID, "start")
	assert.NoError(t, err, "PerformAction() (start) должно работать без ошибок")

	err = awsProv.DeleteResource(ctx, resourceID)
	assert.NoError(t, err, "DeleteResource() должно работать без ошибок")

	err = awsProv.Disconnect(ctx)
	assert.NoError(t, err, "Disconnect() должно работать без ошибок")
}
