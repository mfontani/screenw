# screenw

A command to print out a list of characters showing you how wide your terminal
is.

Example for a 36 chars wide terminal:

    $ screenw
                                     =36
             1         2         3
    123456789012345678901234567890123456

... or an 80x terminal:

    $ screenw
                                                                                 =80
             1         2         3         4         5         6         7         8
    12345678901234567890123456789012345678901234567890123456789012345678901234567890

... or a ginormous 130x terminal:

    $ screenw
                                                                                                       1                          =130
             1         2         3         4         5         6         7         8         9         0         1         2         3
    1234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890

I've had a similar command in my local bins for a while, "but" it was
Perl-based, and I was looking for something a little bit more portable - hence
this.

# Copyright and License

`screenw` is Copyright (c) 2019, Marco Fontani MFONTANI@cpan.org

It is released under the MIT license - see the `LICENSE` file in this repository/directory.
