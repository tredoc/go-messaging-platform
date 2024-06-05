package command

type Command struct {
	CreateTemplate CreateTemplateHandler
	DeleteTemplate DeleteTemplateHandler
}

func NewCommand() Command {
	return Command{
		CreateTemplate: NewCreateTemplateHandler(),
		DeleteTemplate: NewDeleteTemplateHandler(),
	}
}
