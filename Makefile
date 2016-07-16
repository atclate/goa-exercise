
build: model
		goagen bootstrap -d github.com/atclate/goa-exercise/design
		go build

model:
		goagen --design=github.com/atclate/goa-exercise/design gen --pkg-path=github.com/goadesign/gorma

clean:
		rm -rf app swagger tool client models
		rm main.go source.go cache.go public.go
