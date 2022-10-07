package interactive

func EventCommandsSection(cmdPrefix string, optionItems []OptionItem) Section {
	section := Section{
		Selects: Selects{
			ID: "",
			Items: []Select{
				{
					Name:    "Command",
					Command: cmdPrefix,
					OptionGroups: []OptionGroup{
						{
							Name:    "Commands",
							Options: optionItems,
						},
					},
				},
			},
		},
	}

	return section
}
