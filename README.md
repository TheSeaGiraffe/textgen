# `textgen`

Create files filled with dummy text for testing or other purposes. The text is just a
section of the classic *Lorem Ipsum* text that is shuffled around to form paragraphs.

## Basic usage

Calling `textgen` without any arguments creates a single file in the current directory.
You can specify the number of files to create as well as the number of paragraphs in the
file using the `-num-files` and `-num-paragraphs` flags respectively:

```bash
textgen -num-files 5 -num-paragraphs 10
```

This will create five files with 10 paragraphs each in the current directory. You can also
specify the output directory using the `-out` flag:

```bash
textgen -out dummy-files
```
