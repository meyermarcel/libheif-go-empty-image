# libheif golang API empty image

This project should reproduce an empty image by the Go API of the amazing C library [libheif](https://github.com/strukturag/libheif).

Shoud work on
```
System Version: macOS 10.15.3 (19D76)
go version go1.14 darwin/amd64
```

## What happens if the steps below are executed?

1. `sample.png` image file is read once

1. pure Go library [disintegration/imaging](https://github.com/disintegration/imaging) transforms image by resizing and cropping to different widths 

1. libheif Go API encodes image to heif

1. libheif Go API writes encoded image to file `sample-width<width>.heif`

1. Repeat steps from step 2. on

## Install library for transformation
```
go get -u github.com/disintegration/imaging
```
## Install libheif

### Install master libheif
```
brew install automake make pkg-config x265 libde265 libjpeg
```

```
go get -u github.com/strukturag/libheif/go/heif
```

```
cd $GOPATH/src/github.com/strukturag/libheif
./autogen.sh
./configure
make install
```

### Install latest libheif version

```
brew install automake make pkg-config x265 libde265 libjpeg libheif
```

## Create empty image

Create images with pixel width from 471 to 480.

```
rm -f *.heif && go run main.go 466 480 && ls -hl *.heif | awk '{print $5, $9}'

Output:
53K sample-width466.heif
54K sample-width467.heif
999B sample-width468.heif
54K sample-width469.heif
54K sample-width470.heif
54K sample-width471.heif
999B sample-width472.heif
54K sample-width473.heif
55K sample-width474.heif
55K sample-width475.heif
935B sample-width476.heif
54K sample-width477.heif
54K sample-width478.heif
55K sample-width479.heif
935B sample-width480.heif
```
