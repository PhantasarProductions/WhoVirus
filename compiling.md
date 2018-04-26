# Compiling notes

To compile you need the next tools:
- Git
- And the Go compiler


In order to compile you must install Tricky Units and the itch.io lzma library first.
Go to your ~/go/src folder or whatever folder you've placed your Go sources in and type:
~~~shell
git clone https://github.com/Tricky1975/trickyunits_go trickyunits
go get itchio/lzma
~~~

Assuming you put the contents of this repository in the ~/go/src/WhoVirus folder type:
~~~
go build WhoVirus
~~~


## Note to Windows users:

This document has been set up for usage in Linux and Mac or any other Unix based system supported by Go.
In Unix ~ means the home folder. Also Unix uses / for folder seperateion where Windows uses \.
If your Windows username is 'MyName' then you should substitute ~/go/src/ with C:\Users\MyName\
The commands described above should work, assuming the system path has access to both Go and Git.

## Version note:

This source REQUIRES Go 1.10.1 or later. On versions before that it may not work.
If it tells you the function "Round" doesn't exist, it means your version of Go is too old.


Once compiled there are no more dependencies needed, and the binary can be distributed in any way.
Linux users should be aware though that if the binary is stored on a Windows compatible device the 'x' attribute may not be present, making it officially impossible to run the game from there.
