package caravan

// ShowConf prints conf
func ShowConf(conf *Conf) {
	if conf == nil {
		return
	}
	PrintNotice("=>", "debug:", conf.Debug)
	PrintNotice("=>", "once:", conf.Once)
	PrintNotice("=>", "src:", conf.Source)
	PrintNotice("=>", "dst:", conf.Destination)
	PrintNotice("=>", "deploy_mode:", conf.Mode)
	PrintNotice("=>", "incremental:", conf.Incremental)
	PrintNotice("=>", "exclude:", conf.Exclude)
}
