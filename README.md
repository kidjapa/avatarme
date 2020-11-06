# AvatarMe

## Skills developed

* Hashing
* image manipulation

## Assigment

Given a personal information such as an email address, IP address, or a
public key, the program you will write needs to generate a unique avatar.
Imagine that you are building a new application and you want all of your
users to have a default and unique avatar. The package you will write
will allow the generation of such avatars. GitHub recently used such an
approach and generates an identicon for all new users who don't have a
gravatar account attached.

## Result

- Create a avatar with a cli commands:

```
create a png avatarme picture

Usage:
  avatarme [flags]

Flags:
  -B, --border uint        Define the border distance (pixel) in image size. (default 30)
  -C, --columns uint       Define the block columns quantity. (default 4)
  -d, --dir string         Define the directory need to save the avatar. If not passed, the binary directory will be used (default ".")
  -f, --file-name string   Define the image name for save avatar image file (default "image")
  -H, --height uint        Define the output image height size. (default 1024)
  -h, --help               help for avatarme
  -L, --lines uint         Define the block lines quantity. (default 4)
  -s, --str string         Define the string for create avatar hash (default "testavatar")
  -W, --width uint         Define the output image width size. (default 1024)
```

![](_img/image.png)

## Resources

* http://golang.org/pkg/crypto/
* http://golang.org/pkg/image/
* http://en.wikipedia.org/wiki/Identicon
* http://haacked.com/archive/2007/01/22/Identicons_as_Visual_Fingerprints.aspx/
