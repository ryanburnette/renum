# [sequentialize](https://github.com/ryanburnette/go-sequentialize)

Sequentialize files in a directory based on a specified ordering and with a new
base file name.

## Usage

The `-order` flag accepts name or date for the resulting sequential ordering.

The `directory` argument is the directory to be sequentialized.

The `new_file_name` argument is the base name for sequentialized file names.

```shell
sequentialize [-order=name|date] [-d] directory new_file_name
```
