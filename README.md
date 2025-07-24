### Install

```bash
go install github.com/sss7526/gmd.git
```

### Help
```bash
gmd help
```

### Initialize a config file template
```bash
gmd init # Creates gmd-config.yaml template in current directory
```

### Generate markdown

NOTE: Edit gmd-config.yaml to define behavior
```bash
gmd # Just run the command in the project directory

# You can optionally specify pats to a config file and output directory
gmd --config=/path/to/gmd-config.yaml --ouput_dir=/path/to/ouput_dir
```