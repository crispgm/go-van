package deploy

// Mode of a deployer
type Mode string

// Deploy Modes
const (
	ModeRSync Mode = "rsync"
)

// Deployer of a deploy method
type Deployer interface {
	Run(string, string, []string) ([]byte, error)
}

// NewDeployer create a specific deployer
func NewDeployer(deployMethod Mode) Deployer {
	switch deployMethod {
	case ModeRSync:
		return RSync{}
	}
	return nil
}
