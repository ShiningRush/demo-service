package hellosrv

type HelloSrv interface {
	Hello()
}

var singleton HelloSrv

func Srv() HelloSrv {
	return singleton
}

func SetSrv(srv HelloSrv) {
	singleton = srv
}
