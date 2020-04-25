package gtimer

import "testing"

func Test_task_setState(t *testing.T) {
	type args struct {
		i int32
	}
	tests := []struct {
		name    string
		t       *task
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.t.setState(tt.args.i); (err != nil) != tt.wantErr {
				t.Errorf("task.setState() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
