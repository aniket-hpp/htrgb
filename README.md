# htrgb
A go module to convert a 6bit hex-color-string to (r, g, b).

# Install
```bash
go get github.com/aniket-hpp/htrgb
```

# Usage
```golang
    //returns (255, 160, 69, nil)
    r, g, b, err := htrgb.HexToRGB("#ffa045")
```