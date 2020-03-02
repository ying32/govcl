package main

func GetLazarusDir() (lazarusDir string, fpcExe string) {
	lazarusDir, fpcExe = readLazarusEnvFile("/etc")
	return
}

func GetBsdDir() (bsdDir string, userDir string) {
	return
}
