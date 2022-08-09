# GoDFM
Simply go to the tags and download the .exe file (or compile it yourself by running `go build`).

Add it to your environment paths by going to settings and adding the directory of where the exe is located.

Call the exe by running `godfm`.

### Commands
- `-a` | `--add` | `--append` - takes a regex pattern and a directory. 

  Example:

  `godfm -a .*png C:\Users\User\Downloads\Images`

- `-s` | `--sort` - sorts your download folder depending on the pattern and folders given.

    Example:

    `godfm -s`

- `-p` | `--print` - print the regex pattern and the path it takes to.

    Example:

    `godfm -p`

- `-r` | `--remove` - removes the regex pattern given from the settings.

    Example:

    `godfm -r .*png`