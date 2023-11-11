# picker-go
a go implementation of both file and folder picking using [go-common-file-dialog](https://github.com/harryjph/go-common-file-dialog) <br>
this utility returns the path of a folder or file. e.g. `C:\Windows\System32\notepad.exe` or `C:\Windows\System32\`

```
Usage of picker.exe:
  -ext string
        File extension filter (used in 'file' mode)
  -mode string
        Mode of operation: 'file' or 'folder'

Example:
  picker.exe -mode file -ext .txt
  picker.exe -mode folder
```

built upon [filepicker](https://github.com/Atlas-OS/utilities/tree/main#filepicker) by the atlasos team.
