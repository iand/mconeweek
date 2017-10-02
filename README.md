Go port of the Minecraft One Week Challenge

Purely for my own amusement I ported the first day of the [Minecraft One Week
Challenge](https://github.com/Hopson97/MineCraft-One-Week-Challenge) to Go. This is just to the
point of a single grass textured quad that you can look at from various angles using the mouse and
the W,A,S,D keys. Use ESC to exit the window.

The code is a pretty close one-to-one port of the original C++ so it lacks some Go idioms. I've
tried to keep filenames aligned with the original C++ codebase but placed everything into a single
package for simplicity. There are some differences between `glfw` used here and `sfml` used in the
original, most noticeably in keyboard and mouse handling. There are Go bindings for `sfml` but I'm
more familiar with `glfw` so chose to work with that instead.

The original code in this repository is licensed under the terms of the Apache 2.0 license, the same as the
original MineCraft One Week Challenge. 

This was a fun afternoon project but I may continue to port the remaining days if I feel motivated
enough.

# Installation

Required Go packages are vendored for convenience but to build it you'll need some C libraries
dependent on your OS. See more details at [https://github.com/go-gl/gl](https://github.com/go-gl/gl)
and [https://github.com/go-gl/glfw](https://github.com/go-gl/glfw). On Debian-based Linux systems it
might be enough to install the `libgl1-mesa-dev` and `xorg-dev` packages along with `gcc`. I have
not tested on other operating systems.

Ensure you have a recent working Go installtion. Using go get should check out the code and build
the latest version so you can run it straight off:

    go get github.com/iand/mconeweek
    mconeweek

Alternatively you can check it out into your GOPATH manually and build it with:

    go install github.com/iand/mconeweek

