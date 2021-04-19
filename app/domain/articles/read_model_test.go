package articles

import (
	"github.com/HETIC-MT-P2021/CQRSES_GROUP3/shared/core/es"
	"github.com/HETIC-MT-P2021/CQRSES_GROUP3/shared/models"
	uuid "github.com/satori/go.uuid"
	"reflect"
	"testing"
	"time"
)

func TestReadModel_ProjectNewReadModel(t *testing.T) {
	type fields struct {
		AggregateID  string
		FinalPayload interface{}
		CreatedAt    time.Time
	}
	type args struct {
		eventList []*es.Event
	}
	var events []*es.Event
	for i := 0; i < 5; i++ {
		events = append(events, &es.Event{
			AggregateID: "",
			Typology:    "",
			Payload:     map[string]interface{}{},
			CreatedAt:   time.Time{},
			Index:       0,
		})
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    models.Article
		wantErr bool
	}{
		{
			name: "TestReadModelProjectionSuccess",
			fields: fields{
				AggregateID:  uuid.NewV4().String(),
				FinalPayload: reflect.Interface,
				CreatedAt:    time.Time{},
			},
			args: args{
				eventList: events,
			},
			want:    models.Article{},
			wantErr: false,
		},
		{
			name: "TestReadModelProjectionWithEmptyEvents",
			fields: fields{
				AggregateID:  uuid.NewV4().String(),
				FinalPayload: reflect.Interface,
				CreatedAt:    time.Time{},
			},
			args: args{
				eventList: []*es.Event{},
			},
			want:    models.Article{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &ReadModel{
				AggregateID:  tt.fields.AggregateID,
				FinalPayload: tt.fields.FinalPayload,
				CreatedAt:    tt.fields.CreatedAt,
			}
			if _, err, _ := r.ProjectNewReadModel(tt.args.eventList); (err != nil) != tt.wantErr {
				t.Errorf("ProjectNewReadModel() got1 = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
