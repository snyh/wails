package build

import (
	"github.com/leaanthony/wailsv2/v2/internal/logger"
	"github.com/leaanthony/wailsv2/v2/internal/project"
)

// Builder defines a builder that can build Wails applications
type Builder interface {
	SetProjectData(projectData *project.Project)
	BuildAssets(*Options) error
	BuildFrontend(*logger.Logger) error
	BuildRuntime(*Options) error
	CompileProject(*Options) error
	CleanUp()
}
