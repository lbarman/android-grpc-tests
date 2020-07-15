# GoMobile

## Prerequisites

- GoMobile `go get golang.org/x/mobile/cmd/gomobile`
- Android NDK https://developer.android.com/studio/projects/install-ndk#default-version

Note: installing the gomobile is painful and I had to try multiple times. 
In the end, I installed Android Studio from sources in `/opt/`.
I had `ANDROID_HOME` pointing to `~/Android/Sdk`. This was added in my `.bash_rc`. Then, in Android Studio, using the SDK manager, I downloaded `NDK (side by side)` and `cmake` and other binutils tools.

I made sure all licenses were accepted and that the SDK was up-to-date with:
`~/Android/Sdk/tools/bin ❯ ./sdkmanager --update`
and:
`~/Android/Sdk/tools/bin ❯ ./sdkmanager --licenses`

Finally, the magic trick: I needed to to do the following:
`~/Android/Sdk ❯ ln -s ndk/21.3.6528147 ndk-bundle`

## Compile Go service as an Android lib

`gomobile bind -o service-go/servicego.aar -target android github.com/lbarman/android-grpc-tests/service-go`

To be made in the appropriate folder in `GOPATH` as always.

This creates an `.aar` file that can be imported by Java.

Put it in `JAVA_APP/app/libs`

## Importing the .aar in Java

Project-level `build.grade`:
```
allprojects {
  repositories {
    flatDir {
        dirs 'libs'
    }
  }
}
```

App-level `build.gradle`:
```
dependencies{
  implementation(name: 'servicego', ext: 'aar')
}
```