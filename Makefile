
build: model
		#../../goadesign/goa/goagen/goagen bootstrap -d github.com/atclate/goa-exercise/design
		goagen bootstrap -d github.com/atclate/goa-exercise/design
		go build

model:
		#../../goadesign/goa/goagen/goagen --design=github.com/atclate/goa-exercise/design gen --pkg-path=github.com/atclate/gorma
		goagen --design=github.com/atclate/goa-exercise/design gen --pkg-path=github.com/goadesign/gorma

clean:
		rm -rf app swagger tool client models
		rm main.go source.go cache.go public.go
