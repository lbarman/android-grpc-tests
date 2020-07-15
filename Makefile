prerequisites:
	go get golang.org/x/mobile/cmd/gomobile
	$(info "Install NDK https://developer.android.com/studio/projects/install-ndk#default-version")

go-service-binary:
	gomobile bind -o service-go/servicego.aar -target android github.com/lbarman/android-grpc-tests/service-go

bin-to-java:
	cp service-go/servicego.aar app-java/app/libs


all: go-service-binary bin-to-java