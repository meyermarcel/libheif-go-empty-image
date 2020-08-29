# ~~libheif golang API produces empty image~~ fixed

~~I read different image files from disk. After the conversion to HEIF by the amazing C library [libheif](https://github.com/strukturag/libheif) some of them were empty.~~

Version 1.8.0 of [libheif](https://github.com/strukturag/libheif) fixes this issue.

To reproduce this behaviour this project was created.

## Tested systems

This project should work on systems supporting Go and strukturag/libheif.

### Linux

Should work but it is not confirmed.

### macOS

Tested on macOS:

```
System Version: macOS 10.15.6 (19G2021)
go version go1.15.0 darwin/amd64
```

## What happens if the steps below are executed?

1. `sample-width<width>.png` image file is read

1. libheif Go API encodes image to to lossy heif with 75% quality

    For every encoding a new context is created. `main.go`:
    ```go
    ...
    ctx, err := heif.EncodeFromImage(imgNRGBA, ...)
    ...
    ```
    This is also suggested in https://github.com/strukturag/libheif/blob/master/examples/heif-test.go.

1. libheif Go API writes encoded image to file `sample-width<width>.heif`

1. Repeat steps from step 1. on with `width+1`

## Read and convert sample images

Read and convert sample images with pixel width from 446 to 486.

```
rm -f *.heif  &&  go run main.go 464 476  &&  ls -hl *.heif  |  awk '{print $5, $9}'

Output:
53K sample-width464.heif
54K sample-width465.heif
54K sample-width466.heif
54K sample-width467.heif
54K sample-width468.heif
54K sample-width469.heif
54K sample-width470.heif
54K sample-width471.heif
54K sample-width472.heif
55K sample-width473.heif
55K sample-width474.heif
55K sample-width475.heif
54K sample-width476.heif
```
