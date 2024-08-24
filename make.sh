#PATH=$PATH:~/go/bin
fyne build -tags debug -os linux -o biehdc.swipetabs.fakemsg # -release
#CGO_ENABLED=1 GOARCH=386 GOOS=windows CC=i686-w64-mingw32-gcc fyne package -release -os windows
ANDROID_NDK_HOME=~/code/android-ndk-r27 fyne package -os android/arm64
# Get it at https://github.com/android/ndk
