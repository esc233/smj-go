package autoConfig

func StartApp(configDir string, autoConfig AutoConfig) {
	LoadConfig(configDir)
	InitZeroLog(C.Log)
}
