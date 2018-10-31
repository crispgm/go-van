package deploy

// Mode of a deployer
type Mode string

// Deploy Modes
const (
	ModeRSync Mode = "rsync"
)

// Deployer of a deploy method
type Deployer interface {
	Run(string, string) ([]byte, error)
}
