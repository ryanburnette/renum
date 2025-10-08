# renum

`renum` is a POSIX-compliant shell script that renames files in a directory to a
sequential format using a specified base name and numbered suffixes with leading
zeros (e.g., `foo-001.jpg`, `foo-002.jpg`). It supports sorting files by name or
modification date and includes options for dry runs and bypassing confirmation
prompts.

## Usage

```sh
renum [-order=name|date] [-y] [-d] file(s) new_file_name
```

- **file(s)**: Files to rename, specified directly or via glob patterns (e.g.,
  `*.jpg`). All files must be in the same directory.
- **new_file_name**: Base name for the renamed files.

### Options

- `-order=name|date`: Sort files by name (default) or modification date (oldest first).
- `-y`: Bypass the confirmation prompt for the entire job.
- `-d`: Perform a dry run, printing proposed changes without renaming files.

## Examples

```sh
# Rename all .jpg files in the current directory to photo-001.jpg, photo-002.jpg, etc.
./renum.sh *.jpg photo

# Dry run to preview renaming .png files, sorted by date
./renum.sh -order=date -d *.png image

# Rename .txt files without confirmation
./renum.sh -y *.txt document
```

## Features

- **Sequential Naming**: Adds numbered suffixes with leading zeros (`-n`, `-nn`,
  `-nnn`, etc.) based on the number of files.
- **Confirmation**: Prompts for confirmation before renaming unless `-y` is
  used.
- **Dry Run**: Shows planned actions without modifying files when `-d` is
  specified.
- **POSIX Compliance**: Works in any POSIX-compliant shell environment.

## Notes

- Ensure all files are in the same directory to avoid errors.
- Use with `genfiles.sh` to create test files for experimenting with `renum`.

