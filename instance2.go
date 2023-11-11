package teste

type Singleton2 struct{}

var instance2 *Singleton

func GetInstance2() *Singleton {
	if instance == nil {
		instance = &Singleton{}
	}
	return instance
}
