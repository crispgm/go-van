package deploy

// Mode of a deployer
type Mode string

// Path of a deployer
type Path string

// Deploy Modes
const (
	RSync Mode = "rsync"
)

// Deployer of a deploy method
type Deployer interface {
	Run(Path, Path) int
}
