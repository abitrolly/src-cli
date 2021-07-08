	"github.com/mitchellh/copystructure"
		BaseRepository: testRepo1.ID,
			BaseRef:        testRepo1.DefaultBranch.Name,
			BaseRev:        testRepo1.DefaultBranch.Target.OID,
			HeadRepository: testRepo1.ID,
		copy, err := copystructure.Copy(s)
		if err != nil {
			t.Fatalf("deep copying spec: %+v", err)
		}

		s = copy.(*batches.ChangesetSpec)
		Repository: testRepo1,
	taskWith := func(task *Task, f func(task *Task)) *Task {
		copy, err := copystructure.Copy(task)
		if err != nil {
			t.Fatalf("deep copying task: %+v", err)
		}

		task = copy.(*Task)
		f(task)
		return task
	featuresWithoutOptionalPublished := featuresAllEnabled()
	featuresWithoutOptionalPublished.AllowOptionalPublished = false

		name     string
		task     *Task
		features batches.FeatureFlags
		result   executionResult
			name:     "success",
			task:     defaultTask,
			features: featuresAllEnabled(),
			result:   defaultResult,
			features: featuresAllEnabled(),
			result:   defaultResult,
			features: featuresAllEnabled(),
			result:   defaultResult,
		{
			name: "publish in UI on a supported version",
			task: taskWith(defaultTask, func(task *Task) {
				task.Template.Published = nil
			}),
			features: featuresAllEnabled(),
			result:   defaultResult,
			want: []*batches.ChangesetSpec{
				specWith(defaultChangesetSpec, func(s *batches.ChangesetSpec) {
					s.Published = nil
				}),
			},
			wantErr: "",
		},
		{
			name: "publish in UI on an unsupported version",
			task: taskWith(defaultTask, func(task *Task) {
				task.Template.Published = nil
			}),
			features: featuresWithoutOptionalPublished,
			result:   defaultResult,
			want:     nil,
			wantErr:  errOptionalPublishedUnsupported.Error(),
		},
			have, err := createChangesetSpecs(tt.task, tt.result, tt.features)
func parsePublishedFieldString(t *testing.T, input string) *overridable.BoolOrString {
	return &result