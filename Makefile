
build:
		goagen bootstrap -d github.com/atclate/goa-exercise/design

model:
		goagen --design=github.com/atclate/goa-exercise/design gen --pkg-path=github.com/goadesign/gorma

clean:
		rm *.go
		rm -rf app swagger tool client
