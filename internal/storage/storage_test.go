package storage

import "testing"

func TestInMemoryStorage_AddItem(t *testing.T) {
	type args struct {
		item Item
	}
	tests := []struct {
		name string
		i    *InMemoryStorage
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.i.AddItem(tt.args.item)
		})
	}
}
