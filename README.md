# libheif golang API empty image

This project should reproduce an empty image by the Go API of the amazing C library [libheif](https://github.com/strukturag/libheif).

Shoud work on
```
System Version: macOS 10.15.3 (19D76)
go version go1.14 darwin/amd64
```

## What happens if the steps below are executed?

1. `sample.png` image (480x360px) file is read once

1. pure Go library [disintegration/imaging](https://github.com/disintegration/imaging) resizes image to different widths maintaining height

1. libheif Go API encodes image to to lossy heif with 75% quality

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

Create images with pixel width from 446 to 486.

```
rm -f *.heif && go run main.go 446 486 && ls -hl *.heif | awk '{print $5, $9}'

Output:
50K sample-width446.heif
50K sample-width447.heif
911B sample-width448.heif
50K sample-width449.heif
51K sample-width450.heif
51K sample-width451.heif
971B sample-width452.heif
52K sample-width453.heif
52K sample-width454.heif
51K sample-width455.heif
971B sample-width456.heif
52K sample-width457.heif
52K sample-width458.heif
52K sample-width459.heif
943B sample-width460.heif
52K sample-width461.heif
53K sample-width462.heif
52K sample-width463.heif
943B sample-width464.heif
53K sample-width465.heif
54K sample-width466.heif
53K sample-width467.heif
999B sample-width468.heif
54K sample-width469.heif
54K sample-width470.heif
54K sample-width471.heif
999B sample-width472.heif
54K sample-width473.heif
55K sample-width474.heif
54K sample-width475.heif
935B sample-width476.heif
54K sample-width477.heif
54K sample-width478.heif
54K sample-width479.heif
935B sample-width480.heif # original width of sample.png
54K sample-width481.heif
55K sample-width482.heif
55K sample-width483.heif
989B sample-width484.heif
55K sample-width485.heif
55K sample-width486.heif
```
