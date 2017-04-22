# golitch
produces a glitched or 'databent' version of an image, it's basically random, so for best results just run a few times and hope for something that looks pretty

## How to Use
On it's own, running the binary will look for a file called "in.jpg" and save it to "out.jpg", overwriting whatever is there if it exists already

You can also optionally give it an input and output file name, respectively, if the output is missing, it will assume "out.jpg"

eg:

```golitch.exe myimage.jpg result.jpg```

will read myimage.jpg, glitch it and save it to result.jpg

```golitch.exe anotherimage.jpg```

will read anotherimage.jpg, glitch it ands ave it to out.jpg

## Example Image

Before:
![Before Image](http://i.imgur.com/LvpvH5X.jpg)

After:
![After Image](http://i.imgur.com/xIgpr7v.jpg)