### Go Wall

A simple program which sets random wallpaper from some images directory.

#### Dependencies

This program uses `feh` to set wallpapers.

#### Installation

-   Clone this repository and build it.

    ```bash
    git clone https://github.com/Awan/gowall.git
    cd gowall
    go build wallpaper.go
    mv wallpaper ~/.local/bin/
    ```
Or

-   Get the built one binary [here](https://github.com/Awan/gowall/releases/download/v1.0.0/wallpaper.binary).


#### Usage

Pretty simple.

```bash
wallpaper ~/Images 30
```

It will get a random image from `~/Images` directory and set it as wallpaper and 
will do the same thing after 30 seconds.

If you don't pass the time duration, it will just set a random wallpaper and 
exits.

You can put this line in your `~/.xinitrc` or `~/.xprofile` to automate the 
boring stuff of setting wallpapers.

```bash
...
~/.local/bin/wallpaper ~/Images 600 &
```
