# go-rename

A simple CLI tool for renaming images to their SHA-1 checksum.

## Features

- Rename image files with extensions: `.jpeg`, `.jpg`, `.png`
- Outputs the original and new file names.

## Usage

```sh
go run . /path/to/your/image/directory
```

## Example

```
> go-rename go run . ~/Pictures/Wallpapers

Original      New
-----------   -----------
image_1.jpg   7a241a57d2475cad94811f78ed16bca6f4211c9d.jpg
image_2.png   0ce1279ff93f6aa62f1bb37c9bd8354038041e5c.png
```
