package storage

import (
	"reflect"
	"testing"
)

func TestInMemoryStorage_AddItem(t *testing.T) {
	type args struct {
		item Item
		k    string
	}
	tests := []struct {
		name string

		args args
	}{
		{
			name: "success",
			args: args{
				item: Item{
					Name:     "pomodor",
					Sum:      10,
					Currency: "EUR",
				},
				k: "pomodor",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := NewMemoryStorage()
			a.AddItem(tt.args.item)
			f := a.List[tt.args.k]
			if !reflect.DeepEqual(f, tt.args.item) {
				t.Error("ne ravni")
			}
		})
	}
}

func TestInMemoryStorage_GetItem(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name string

		args args
		want Item
	}{
		{
			name: "succes",
			args: args{
				key: "xx",
			},
			want: Item{
				Name:     "og",
				Sum:      10,
				Currency: "EUR",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := NewMemoryStorage()

			a.List[tt.args.key] = tt.want

			if got := a.GetItem(tt.args.key); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InMemoryStorage.GetItem() = %v, want %v", got, tt.want)
			}
		})
	}
}
