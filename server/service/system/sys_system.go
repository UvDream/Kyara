package system

type SysConfigService struct {
}

// GetSysConfig 获取系统信息
//func (sysConfigService *SysConfigService) GetSysConfig() (server *utils.Server, err error) {
//	var s utils.Server
//	s.Os = utils.InitOS()
//	if s.Cpu, err = utils.InitCPU(); err != nil {
//		global.Log.Error("获取CPU信息失败(utils.InitCPU())", zap.String("err", err.Error()))
//		return &s, err
//	}
//	if s.Ram, err = utils.InitRAM(); err != nil {
//		global.Log.Error("获取内存信息失败(utils.InitRAM())", zap.String("err", err.Error()))
//		return &s, err
//	}
//	if s.Disk, err = utils.InitDisk(); err != nil {
//		global.Log.Error("获取磁盘信息失败(utils.InitDisk())", zap.String("err", err.Error()))
//		return &s, err
//	}
//	return &s, nil
//}
