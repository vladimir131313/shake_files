# shake_files
The program rename files. It adds a random prefix. It's can be used for shake files in old mp3 players which doesn't support random functions.

Build:

go build shake_files.go 
go build shake_files2.go

shake_files - just rename files with random prefix.
shake_files2 - better version. Firstly it shake list files and then added sequens number in name. 

Help:

./shake_files2 -h
Usage of ./shake_files2:
  -d	Rename dirs too
  -p string
    	path to directory with files (default ".")
  -r	Delete prefix in filenames if it exist
  -w	Rename files (Default: Just print result without rename files)
MacBook-Pro-Vladimir:shake_files vladimir$ ./shake_files2 testdir/


Example:

Dry run:
./shake_files2 -p testdir/
testdir/f.mp3                                      --->       testdir/0RND-f.mp3
testdir/e.mp3                                      --->       testdir/1RND-e.mp3
testdir/g.mp3                                      --->       testdir/2RND-g.mp3
testdir/b.mp3                                      --->       testdir/3RND-b.mp3
testdir/d.mp3                                      --->       testdir/4RND-d.mp3
testdir/c.mp3                                      --->       testdir/5RND-c.mp3
testdir/a.mp3                                      --->       testdir/6RND-a.mp3


Rename:
./shake_files2 -w -p testdir/
MacBook-Pro-Vladimir:shake_files vladimir$ ls -l testdir/
total 0
-rw-r--r--  1 vladimir  staff  0 Apr 18 19:26 0RND-e.mp3
-rw-r--r--  1 vladimir  staff  0 Apr 18 19:26 1RND-c.mp3
-rw-r--r--  1 vladimir  staff  0 Apr 18 19:26 2RND-d.mp3
-rw-r--r--  1 vladimir  staff  0 Apr 18 19:26 3RND-a.mp3
-rw-r--r--  1 vladimir  staff  0 Apr 18 19:27 4RND-f.mp3
-rw-r--r--  1 vladimir  staff  0 Apr 18 19:26 5RND-b.mp3
-rw-r--r--  1 vladimir  staff  0 Apr 18 19:27 6RND-g.mp3

