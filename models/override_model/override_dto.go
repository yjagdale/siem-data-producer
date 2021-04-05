package override_model

var OverrideValues map[string][]string

func connectDB() {
	if OverrideValues == nil {
		OverrideValues = make(map[string][]string)
	}
	return
}

func (overrides Override) Save() string {
	connectDB()
	if _, ok := OverrideValues[overrides.Key]; ok {
		OverrideValues[overrides.Key] = append(OverrideValues[overrides.Key], overrides.Values...)
	} else {
		OverrideValues[overrides.Key] = overrides.Values
	}

	return "Override Added Successfully"
}
