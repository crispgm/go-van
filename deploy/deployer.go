package deploy

// Mode of a deployer
type Mode string

// Path of a deployer
type Path string

// Deploy Modes
const (
	ModeRSync Mode = "rsync"
)

// Deployer of a deploy method
type Deployer interface {
	Run(string, string) error
}
