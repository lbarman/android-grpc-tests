prerequisites:
	go get golang.org/x/mobile/cmd/gomobile
	$(info "Install NDK https://developer.android.com/studio/projects/install-ndk#default-version")

go-service-binary:
	gomobile bind -target android github.com/lbarman/android-grpc-tests/service-go