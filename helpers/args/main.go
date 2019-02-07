package args

type StringArg struct {
	Short string
	Long string
	NoEqual bool
	Value string
}

type BoolArg struct {
	Short string
	Long string
	Value bool
}

func (a *StringArg) SetValue(value string) {
	a.Value = value
}

func (a *BoolArg) SetValue(value bool) {
	a.Value = value
}

func GenerateCommand(stringArgs map[int]StringArg, boolArgs  map[int]BoolArg) []string {
	
	var response []string

	for _, arg := range stringArgs {
		if arg.Value != "" && !arg.NoEqual {
			if len(arg.Long) > 1 {
				response = append(response, "--"+arg.Long+"=\""+arg.Value+"\"")
			}else {
				response = append(response, "-"+arg.Short+" \""+arg.Value+"\"")
			}
		}else if arg.Value != "" && arg.NoEqual {
			if len(arg.Long) > 1 {
				response = append(response, "--"+arg.Long+" \""+arg.Value+"\"")
			}else {
				response = append(response, "-"+arg.Short+" \""+arg.Value+"\"")
			}
		}
	}

	for _, arg := range boolArgs {
		if arg.Value {
			if len(arg.Long) > 1 {
				response = append(response, "--"+arg.Long)
			}else {
				response = append(response, "-"+arg.Short)
			}
		}
	}

	return response
}